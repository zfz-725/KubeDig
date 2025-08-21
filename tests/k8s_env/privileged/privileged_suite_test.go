// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 Authors of KubeDig

package privileged

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSmoke(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Privileged Suite")
}
