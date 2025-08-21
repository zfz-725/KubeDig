#!/bin/bash
# SPDX-License-Identifier: Apache-2.0
# Copyright 2021 Authors of KubeDig

# remove old images
docker images | grep ubuntu-w-utils | awk '{print $3}' | xargs -I {} docker rmi -f {} 2> /dev/null

# create new images
docker build --tag kubedig/ubuntu-w-utils:0.1 --tag kubedig/ubuntu-w-utils:latest .

# push new images
docker push kubedig/ubuntu-w-utils:0.1
docker push kubedig/ubuntu-w-utils:latest
