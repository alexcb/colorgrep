## First generate a token

go to https://github.com/settings/tokens

## Pick a new release tag

    export RELEASE_TAG=v1.2.3

## Then deploy

    earthly --secret GITHUB_TOKEN=ghp_mytoken --build-arg RELEASE_TAG --push +release
