// SPDX-License-Identifier: Apache-2.0
// Copyright 2025 Authors of KubeDig

// Package buildinfo is responsible for providing kubedig build info.
package buildinfo

import (
	kg "github.com/zfz-725/KubeDig/KubeDig/log"
)

// GitSummary represents build-time info for git commit,tag
var GitSummary string

// GitBranch represents build-time info for git branch
var GitBranch string

// BuildDate represents build-time info for build date
var BuildDate string

// Provides KubeDig build time info
func PrintBuildDetails() {
	if GitSummary == "" {
		return
	}
	kg.Printf("BUILD-INFO: version: %v, branch: %v, date: %v",
		GitSummary, GitBranch, BuildDate)
}
