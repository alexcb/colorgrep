# colorgrep

Reads from stdin and highlights any matched patterns

## Building

First download earthly.

Then run:

    earthly +all

builds are written to `build/ubuntu/colorgrep` and `build/alpine/colorgrep`

## Example use

    build/alpine/colorgrep pattern1 pattern2

## Futurework

- support a `--color=<color-name>` flag which will force the highlighted pattern to be a specific color
