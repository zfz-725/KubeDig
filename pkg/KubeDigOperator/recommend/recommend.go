// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 Authors of KubeDig

package recommend

import (
	"embed"
)

//go:embed *.yaml
var CRDFs embed.FS
