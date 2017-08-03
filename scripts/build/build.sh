#!/bin/bash

if [ -z ${VERSION} ];
then
    echo "ERROR: Please set environment variable VERSION"
    echo ""
    echo "Ideally this script should be run via 'make build-all' command."
    echo "If you want to run this outside of make, please ensure that appropriate"
    echo "version information is set in VERSION environment variable."
    exit 1
fi

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
BASEDIR="${SCRIPTDIR}/../.."

. ${BASEDIR}/scripts/build/common.sh
. ${BASEDIR}/scripts/build/functions.sh

mkdir -p ${BUILD_DIR} ${DIST_DIR}

PLATFORM_BUILD_DIR=$(get_build_directory ${OS_VERSION} amd64)
build_binary ${OS_VERSION} amd64 ${PLATFORM_BUILD_DIR}
copy_wrapper_script ${PLATFORM_BUILD_DIR}
