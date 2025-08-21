#!/bin/bash
# SPDX-License-Identifier: Apache-2.0
# Copyright 2021 Authors of KubeDig

# remove docker
sudo dnf remove docker-ce docker-ce-cli containerd.io

# remove any assoicated files
sudo rm -rf /var/lib/docker
sudo rm -rf /var/lib/containerd
