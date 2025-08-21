// SPDX-License-Identifier: Apache-2.0
// Copyright 2022 Authors of KubeDig

package configmap_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestKubedigConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KubedigConfig Suite")
}
