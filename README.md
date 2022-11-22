# colorgrep

Reads from stdin and highlights any matched patterns

## Building

First download earthly.

Then run:

    earthly +all

builds are written to `build/ubuntu/colorgrep` and `build/alpine/colorgrep`

## Example use

    cat fruit.txt | colorgrep -c yellow berry -c green -i '^ap.*e$'

Which will highlight matches:

    ![screenshot](screenshot.png)
