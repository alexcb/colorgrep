# colorgrep

Reads from stdin and highlights any matched patterns

## Usage

    colorgrep [options] <pattern> [[-i] [-v] <pattern> [...]]
      -i          Case insensitive matching
      -w          word boundary matching
      -h, --help  display this help

## Example

    cat fruit.txt | colorgrep -c yellow berry -c green -i '^ap.*e$'

Which will highlight matches:

![screenshot](screenshot.png)

## Building

First download [earthly](https://github.com/earthly/earthly).

Then run:

    earthly +all

builds are written to `build/ubuntu/colorgrep` and `build/alpine/colorgrep`

