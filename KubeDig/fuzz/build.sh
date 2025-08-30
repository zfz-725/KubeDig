# SPDX-License-Identifier: Apache-2.0
# Copyright 2024 Authors of KubeDig
printf "package transform\nimport _ \"github.com/AdamKorcz/go-118-fuzz-build/testing\"\n" > $SRC/KubeDig/KubeDig/register.go
go mod tidy
compile_native_go_fuzzer github.com/zfz-725/KubeDig/KubeDig/core FuzzContainerPolicy FuzzContainerPolicy
compile_native_go_fuzzer github.com/zfz-725/KubeDig/KubeDig/core FuzzHostPolicy FuzzHostPolicy

