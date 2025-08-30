// SPDX-License-Identifier: Apache-2.0
// Copyright 2022 Authors of KubeDig

package main

import (
	"context"

	"github.com/zfz-725/KubeDig/KubeDig/types"
)

type handler interface {
	listContainers(ctx context.Context) ([]types.Container, error)
	close() error
}
