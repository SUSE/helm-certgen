#!/bin/sh

get_build_directory(){
    local BUILDOS=$1
    local BUILDARCH=$2
    echo "${BUILD_DIR}/${BUILDOS}-${BUILDARCH}/certgen"
}

build_binary(){
    local BUILDOS=$1
    local BUILDARCH=$2
    local P_BUILD_DIR=$3
    mkdir -p ${P_BUILD_DIR}

    echo "Building helm-certgen binary for ${BUILDOS}-${BUILDARCH} @ ${P_BUILD_DIR}"
    GOOS=${BUILDOS} GOARCH=${BUILDARCH} go build -ldflags '-s' -o ${P_BUILD_DIR}/certgen ${BASEDIR}/main.go
}

copy_wrapper_script(){
    local P_BUILD_DIR=$1
    cp -r ${BASEDIR}/LICENSE ${P_BUILD_DIR}/
    cp -r ${BASEDIR}/plugin/* ${P_BUILD_DIR}/
    sed "s|0.1.0|${VERSION}|g" ${BASEDIR}/plugin/plugin.yaml > ${P_BUILD_DIR}/plugin.yaml
}

create_distributable(){
    local BUILDOS=$1
    local BUILDARCH=$2
    local P_BUILD_DIR=$3

    if [ "${BUILDOS}" == "windows" ];
    then
        zip -jr ${DIST_DIR}/certgen-${BUILDOS}-${BUILDARCH}-${VERSION_DIR}.zip ${P_BUILD_DIR}/*
    else
        tar -cvzf ${DIST_DIR}/certgen-${BUILDOS}-${BUILDARCH}-${VERSION_DIR}.tgz -C ${P_BUILD_DIR} .
    fi
}