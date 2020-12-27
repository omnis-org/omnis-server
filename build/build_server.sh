#!/bin/bash


VERSION="0.2"
TIME=$(date +"%d-%m-%y")
WINDOWS="windows"

sed -i "/BuildVersion string/c\    BuildVersion string = \"$VERSION\"" ../internal/version/version.go
sed -i "/BuildDate string/c\    BuildDate string = \"$TIME\"" ../internal/version/version.go



if [[ "$WINDOWS" == $1 ]]; then
    echo "Build for windows"
    GOOS=windows go build ../cmd/omnis-server
else
    echo "Build for linux"
    go build ../cmd/omnis-server
fi