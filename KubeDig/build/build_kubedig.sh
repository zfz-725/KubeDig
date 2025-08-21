#!/bin/bash
# SPDX-License-Identifier: Apache-2.0
# Copyright 2021 Authors of KubeDig

[[ "$REPO" == "" ]] && REPO="kubedig/kubedig"

UBIREPO="kubedig/kubedig-ubi"

realpath() {
    CURR=$PWD

    cd "$(dirname "$0")"
    LINK=$(readlink "$(basename "$0")")

    while [ "$LINK" ]; do
        cd "$(dirname "$LINK")"
        LINK=$(readlink "$(basename "$1")")
    done

    REALPATH="$PWD/$(basename "$1")"
    echo "$REALPATH"

    cd $CURR
}

ARMOR_HOME=`dirname $(realpath "$0")`/..
cd $ARMOR_HOME/build
pwd

VERSION=latest

# check version
if [ ! -z $1 ]; then
    VERSION=$1
fi

export DOCKER_BUILDKIT=1

# remove old images
docker images | grep kubedig | awk '{print $3}' | xargs -I {} docker rmi -f {} 2> /dev/null
echo "[INFO] Removed existing $REPO images"

# set LABEL
unset LABEL
[[ "$GITHUB_SHA" != "" ]] && LABEL="--label github_sha=$GITHUB_SHA"

# set the $IS_COVERAGE env var to 'true' to build the kubedig-test image for coverage calculation
if [[ "$IS_COVERAGE" == "true" ]]; then
    REPO="kubedig/kubedig-test"
    
    # build a kubedig-test image
    DTAG="-t $REPO:$VERSION"
    echo "[INFO] Building $DTAG"
    cd $ARMOR_HOME/..; docker build $DTAG -f Dockerfile --target kubedig-test . $LABEL

    if [ $? != 0 ]; then
        echo "[FAILED] Failed to build $REPO:$VERSION"
        exit 1
    fi
    echo "[PASSED] Built $REPO:$VERSION"
    
    # build a kubedig-test-init image
    DTAGINI="-t $REPO-init:$VERSION"
    echo "[INFO] Building $DTAGINI"
    cd $ARMOR_HOME/..; docker build $DTAGINI -f Dockerfile.init --build-arg VERSION=$VERSION --target kubedig-init . $LABEL

    if [ $? != 0 ]; then
        echo "[FAILED] Failed to build $REPO-init:$VERSION"
        exit 1
    fi
    echo "[PASSED] Built $REPO-init:$VERSION"

    # build kubedig-ubi-test image
    DTAGUBITEST="-t $UBIREPO-test:$VERSION"
    echo "[INFO] Building $DTAGUBITEST"
    cd $ARMOR_HOME/..; docker build $DTAGUBITEST -f Dockerfile --target kubedig-ubi-test . $LABEL

    if [ $? != 0 ]; then
        echo "[FAILED] Failed to build $DTAGUBITEST:$VERSION"
        exit 1
    fi
    echo "[PASSED] Built $DTAGUBITEST:$VERSION"
    
    exit 0
fi

# build a kubedig image
DTAG="-t $REPO:$VERSION"
echo "[INFO] Building $DTAG"
cd $ARMOR_HOME/..; docker build $DTAG -f Dockerfile --target kubedig . $LABEL

if [ $? != 0 ]; then
    echo "[FAILED] Failed to build $REPO:$VERSION"
    exit 1
fi
echo "[PASSED] Built $REPO:$VERSION"

# build a kubedig-init image
DTAGINI="-t $REPO-init:$VERSION"
echo "[INFO] Building $DTAGINI"
cd $ARMOR_HOME/..; docker build $DTAGINI -f Dockerfile.init --build-arg VERSION=$VERSION --target kubedig-init . $LABEL

if [ $? != 0 ]; then
    echo "[FAILED] Failed to build $REPO-init:$VERSION"
    exit 1
fi
echo "[PASSED] Built $REPO-init:$VERSION"

# build a kubedig-ubi image
DTAGUBI="-t $UBIREPO:$VERSION"
echo "[INFO] Building $UBIREPO"
cd $ARMOR_HOME/..; docker build $DTAGUBI -f Dockerfile --build-arg VERSION=$VERSION --target kubedig-ubi . $LABEL

if [ $? != 0 ]; then
    echo "[FAILED] Failed to build $DTAGUBI:$VERSION"
    exit 1
fi
echo "[PASSED] Built $DTAGUBI:$VERSION"

exit 0
