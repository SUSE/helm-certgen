#!/bin/sh

set -e

workdir=.cover
profile="$workdir/cover.out"
mode=count

# clean up the artifacts for coverage.
rm -rf $workdir

generate_cover_data() {
    mkdir -p "$workdir"

    for pkg in "$@"; do
        echo $pkg
        f="$workdir/$(echo $pkg | tr / -).cover"
        echo $f
        godep go test -covermode="$mode" -coverprofile="$f" "$pkg"
    done

    echo "mode: $mode" >"$profile"
    grep -h -v "^mode:" "$workdir"/*.cover >>"$profile"
}

generate_cover_data $(go list ./... | grep -v vendor)
go tool cover -func="$profile"
go tool cover -html="$profile"
