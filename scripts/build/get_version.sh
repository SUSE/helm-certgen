#!/bin/sh
set -u

export BRANCH=$(git name-rev --name-only HEAD)
export BRANCH_TAG=$(echo $BRANCH | tr "/" "-" | cut -c1-5)

build_commit_hash=$(git rev-parse --short HEAD)
build_time=$(date -u +%Y%m%d%H%M%S)

if [ "${BRANCH}" = "master" ]; then
    export VERSION=$(git describe --tags --long)
else
    LATEST_TAG_VERSION=$(git describe --tags --long | cut -d "." -f 1,2,3)
    export VERSION="${LATEST_TAG_VERSION}+${BRANCH}.${build_time}"
fi
echo ${VERSION}
