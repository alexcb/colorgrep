package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"golang.org/x/tools/godoc/util"
)

type pattern struct {
	re     *regexp.Regexp
	negate bool
	color  string
}

func die(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}

func main() {
	progName := "colorgrep"
	if len(os.Args) > 0 {
		progName = os.Args[0]
	}

	var patterns []*pattern
	negate := false
	insensitive := false
	ignoreDashes := false
	wordBoundary := false
	showHelp := false
	colorNext := false
	var color string
	for _, arg := range os.Args[1:] {
		if len(arg) == 0 {
			die("empty args are not supported\n")
		}
		if colorNext {
			color = arg
			colorNext = false
			continue
		}
		if arg[0] == '-' && !ignoreDashes {
			for _, short := range arg[1:] {
				switch short {
				case 'h':
					showHelp = true
				case 'i':
					if insensitive {
						die("two -i's in a row not supported\n")
					}
					insensitive = true
				case 'c':
					colorNext = true
				case 'v':
					if negate {
						die("two -v's in a row not supported\n")
					}
					negate = true
				case 'w':
					if wordBoundary {
						die("two -w's in a row not supported\n")
					}
					wordBoundary = true
				case 'e':
					ignoreDashes = true
				default:
					die("Error: -%c not recognized\n", short)
				}
			}
			continue
		}
		if wordBoundary {
			arg = "\\b" + arg + "\\b"
		}
		if insensitive {
			arg = "(?i)" + arg
		}
		r, err := regexp.Compile(arg)
		if err != nil {
			die("%s is not a valid regexp: %s\n", arg, err.Error())
		}
		patterns = append(patterns, &pattern{
			re:     r,
			negate: negate,
			color:  color,
		})
		negate = false
		insensitive = false
		wordBoundary = false
		ignoreDashes = false
		color = ""
	}

	if showHelp {
		fmt.Printf(
			"%s [options] <pattern> [[-i] [-v] [-c <color>] <pattern> [...]]\n"+
				"  -i          case insensitive matching\n"+
				"  -w          word boundary matching\n"+
				"  -c <color>  color to highlight match\n"+
				"  -h, --help  display this help\n", progName)
		os.Exit(0)
	}

	colors := map[string]string{
		"red":         "\033[0;31m",
		"green":       "\033[0;32m",
		"orange":      "\033[0;33m",
		"blue":        "\033[0;34m",
		"purple":      "\033[0;35m",
		"cyan":        "\033[0;36m",
		"lightgray":   "\033[0;37m",
		"darkgray":    "\033[1;30m",
		"lightred":    "\033[1;31m",
		"lightgreen":  "\033[1;32m",
		"yellow":      "\033[1;33m",
		"lightblue":   "\033[1;34m",
		"lightpurple": "\033[1;35m",
		"lightcyan":   "\033[1;36m",
	}
	nc := "\033[0m"

	defaultColors := []string{
		colors["red"],
		colors["green"],
		colors["orange"],
		colors["blue"],
		colors["purple"],
		colors["cyan"],
		colors["yellow"],
		colors["lightred"],
		colors["lightgreen"],
		colors["lightblue"],
		colors["lightpurple"],
		colors["lightcyan"],
	}

	getColor := func(i int) string {
		n := len(defaultColors)
		return defaultColors[i%n]
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l := scanner.Text()
		for i, pat := range patterns {
			col := getColor(i)
			if pat.color != "" {
				var ok bool
				col, ok = colors[pat.color]
				if !ok {
					die("color %s not found", pat.color)
				}
			}
			l = pat.re.ReplaceAllString(l, col+"$0"+nc)
		}
		fmt.Println(l)
	}

	if err := scanner.Err(); err != nil {
		die("%s\n", err)
	}
}

func grepFile(path string, patterns []*pattern) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if !util.IsText(data) {
		return nil
	}

	lines := strings.Split(string(data), "\n")
	numLines := len(lines)
	numRegexps := len(patterns)
	matches := make([]bool, numLines*numRegexps)
	for i, line := range lines {
		for j, pat := range patterns {
			k := i*numRegexps + j
			match := pat.re.Match([]byte(line))
			matches[k] = match != pat.negate // xor
		}
	}

	// TODO make use of context size; for now just print the file if all patterns are found.

	var numUniqRegFound int
	regFound := make([]bool, numRegexps)
	for i := 0; i < numLines; i++ {
		for j := 0; j < numRegexps; j++ {
			k := i*numRegexps + j
			if matches[k] && !regFound[j] {
				regFound[j] = true
				numUniqRegFound++
			}
		}
	}
	if numUniqRegFound == numRegexps {
		fmt.Println(path)
	}

	return nil
}
