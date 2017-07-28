#!/bin/bash

OS_VERSION=$(go env GOOS)

BUILD_DIR="${BASEDIR}/build"
DIST_DIR="${BASEDIR}/dist"

VERSION_DIR=$(echo $VERSION | tr "." "-" | tr "+" "-")