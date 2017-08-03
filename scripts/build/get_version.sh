#!/bin/bash -x

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
BASEDIR="${SCRIPTDIR}/../.."

if [ -z ${VERSION} ];
then
    BRANCH=$(git name-rev --name-only HEAD | tr "/" "-")
    VERSION_FROM_FILE=$(cat ${BASEDIR}/version)
    LAST_COMMIT=$(git rev-parse --short HEAD)
    LAST_TIMESTAMP=$(git log -1 --pretty=format:%ct)
    export VERSION="${VERSION_FROM_FILE}-${LAST_TIMESTAMP}.${LAST_COMMIT}+${BRANCH}"
fi
echo ${VERSION}
