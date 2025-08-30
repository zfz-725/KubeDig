// SPDX-License-Identifier: Apache-2.0
// Copyright 2021 Authors of KubeDig

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"testing"
)

var clusterPtr, gRPCPtr, logPathPtr *string
var enableKubeDigPolicyPtr, enableKubeDigHostPolicyPtr, enableKubeDigVMPtr, coverageTestPtr, enableK8sEnv, tlsEnabled *bool
var defaultFilePosturePtr, defaultCapabilitiesPosturePtr, defaultNetworkPosturePtr, hostDefaultCapabilitiesPosturePtr, hostDefaultNetworkPosturePtr, hostDefaultFilePosturePtr, procFsMountPtr *string

func init() {
	// options (string)
	clusterPtr = flag.String("cluster", "default", "cluster name")

	// options (string)
	gRPCPtr = flag.String("gRPC", "32767", "gRPC port number")
	logPathPtr = flag.String("logPath", "none", "log file path")

	// options (string)
	defaultFilePosturePtr = flag.String("defaultFilePosture", "block", "configuring default enforcement action in global file context {allow|audit|block}")
	defaultNetworkPosturePtr = flag.String("defaultNetworkPosture", "block", "configuring default enforcement action in global network context {allow|audit|block}")
	defaultCapabilitiesPosturePtr = flag.String("defaultCapabilitiesPosture", "block", "configuring default enforcement action in global capability context {allow|audit|block}")

	hostDefaultFilePosturePtr = flag.String("hostDefaultFilePosture", "block", "configuring default enforcement action in global file context {allow|audit|block}")
	hostDefaultNetworkPosturePtr = flag.String("hostDefaultNetworkPosture", "block", "configuring default enforcement action in global network context {allow|audit|block}")
	hostDefaultCapabilitiesPosturePtr = flag.String("hostDefaultCapabilitiesPosture", "block", "configuring default enforcement action in global capability context {allow|audit|block}")

	procFsMountPtr = flag.String("procfsMount", "/proc", "Path to the BPF filesystem to use for storing maps")

	// options (boolean)
	enableKubeDigPolicyPtr = flag.Bool("enableKubeDigPolicy", true, "enabling KubeDigPolicy")
	enableKubeDigHostPolicyPtr = flag.Bool("enableKubeDigHostPolicy", true, "enabling KubeDigHostPolicy")
	enableKubeDigVMPtr = flag.Bool("enableKubeDigVm", false, "enabling KubeDigVM")

	enableK8sEnv = flag.Bool("k8s", true, "is k8s env?")
	tlsEnabled = flag.Bool("tlsEnabled", false, "enable tls for secure connection?")

	// options (boolean)
	coverageTestPtr = flag.Bool("coverageTest", false, "enabling CoverageTest")

}

// TestMain - test to drive external testing coverage
func TestMain(t *testing.T) {
	// Reset Test Flags before executing main
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	os.Args = []string{
		fmt.Sprintf("-cluster=%s", *clusterPtr),
		fmt.Sprintf("-gRPC=%s", *gRPCPtr),
		fmt.Sprintf("-logPath=%s", *logPathPtr),
		fmt.Sprintf("-defaultFilePosture=%s", *defaultFilePosturePtr),
		fmt.Sprintf("-defaultNetworkPosture=%s", *defaultNetworkPosturePtr),
		fmt.Sprintf("-defaultCapabilitiesPosture=%s", *defaultCapabilitiesPosturePtr),
		fmt.Sprintf("-hostDefaultFilePosture=%s", *hostDefaultFilePosturePtr),
		fmt.Sprintf("-hostDefaultNetworkPosture=%s", *hostDefaultNetworkPosturePtr),
		fmt.Sprintf("-hostDefaultCapabilitiesPosture=%s", *hostDefaultCapabilitiesPosturePtr),
		fmt.Sprintf("-k8s=%s", strconv.FormatBool(*enableK8sEnv)),
		fmt.Sprintf("-enableKubeDigPolicy=%s", strconv.FormatBool(*enableKubeDigPolicyPtr)),
		fmt.Sprintf("-enableKubeDigHostPolicy=%s", strconv.FormatBool(*enableKubeDigHostPolicyPtr)),
		fmt.Sprintf("-coverageTest=%s", strconv.FormatBool(*coverageTestPtr)),
		fmt.Sprintf("-tlsEnabled=%s", strconv.FormatBool(*tlsEnabled)),
		fmt.Sprintf("-procfsMount=%s", *procFsMountPtr),
	}

	t.Log("[INFO] Executed KubeDig")
	main()
	t.Log("[INFO] Terminated KubeDig")
}
