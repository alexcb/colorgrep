# colorgrep

Reads from stdin and highlights any matched patterns

## Usage

    colorgrep [options] <pattern> [[-i] [-w] [-c <color>] <pattern> [...]]
      -i           case insensitive matching
      -w           word boundary matching
      -c <color>   color to highlight match
      -e <pattern> use pattern (useful for patterns starting with a hyphen)
      -h, --help   display this help


## Example

    cat fruit.txt | colorgrep -c yellow berry -c green -i '^ap.*e$'

Which will highlight matches:

![screenshot](screenshot.png)

## Building

First download [earthly](https://github.com/earthly/earthly).

Then run:

    earthly +all

builds are written to `build/ubuntu/colorgrep` and `build/alpine/colorgrep`

