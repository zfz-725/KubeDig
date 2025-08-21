# SPDX-License-Identifier: Apache-2.0
# Copyright 2021 Authors of KubeDig
#!/usr/bin/env bash

set -e

if [ ! -e "/sys/kernel/btf/vmlinux" ]; then
    # compile BPF programs
    make -C /opt/kubedig/BPF/
fi

# update karmor SELinux module if BPFLSM is not present 
lsm_file="/sys/kernel/security/lsm"
bpf="bpf"
if ! grep -q "$bpf" "$lsm_file"; then
    if [ -x "$(command -v semanage)" ]; then
        # old karmor SELinux module
        /opt/kubedig/templates/uninstall.sh

        # new karmor SELinux module
        /opt/kubedig/templates/install.sh

    fi
fi

# start kubedig.service
/bin/systemctl daemon-reload
/bin/systemctl start kubedig.service
