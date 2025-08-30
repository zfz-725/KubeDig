#!/bin/sh
# SPDX-License-Identifier: Apache-2.0
# Copyright 2021 Authors of KubeDig

cd /KubeDig/BPF
make clean

if [[ -n "$KRNDIR" ]]; then
    make KRNDIR=$KRNDIR
else
    make
fi

cp *.bpf.o /opt/kubedig/BPF/
