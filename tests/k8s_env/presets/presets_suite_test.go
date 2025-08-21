// SPDX-License-Identifier: Apache-2.0
// Copyright 2022 Authors of KubeDig

package presets_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPresets(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Presets Suite")
}
