FROM alpine:3.12

deps:
    ARG DISTRO
    IF [ "$DISTRO" = "alpine" ]
        FROM golang:1.16-alpine3.14
    ELSE IF [ "$DISTRO" = "ubuntu" ]
        FROM golang:1.16-bullseye
    ELSE
        RUN --no-cache echo "$DISTRO not supported" && false
    END
    WORKDIR /code
    COPY go.mod go.sum ./
    RUN go mod download
    # Output these back in case go mod download changes them.
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

build:
    ARG DISTRO="alpine"
    FROM +deps --DISTRO="$DISTRO"
    COPY --dir cmd .
    RUN go build -o build/colorgrep cmd/main.go
    RUN test -n "$DISTRO"
    SAVE ARTIFACT build/colorgrep /go-example AS LOCAL build/$DISTRO/colorgrep

all:
    BUILD +build --DISTRO=alpine --DISTRO=ubuntu
