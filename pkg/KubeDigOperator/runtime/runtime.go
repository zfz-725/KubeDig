// SPDX-License-Identifier: Apache-2.0
// Copyright 2022 Authors of KubeDig

package runtime

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zfz-725/KubeDig/KubeDig/log"
	"github.com/zfz-725/KubeDig/pkg/KubeDigOperator/common"
	"go.uber.org/zap"
)

func DetectNRI(pathPrefix string) (string, error) {
	for _, path := range common.ContainerRuntimeSocketMap["nri"] {
		if _, err := os.Stat(filepath.Clean(pathPrefix + path)); err == nil || os.IsPermission(err) {
			return path, nil
		} else {
			log.Warnf("%s", err)
		}
	}
	return "NA", fmt.Errorf("NRI not available")
}

func DetectRuntimeViaMap(pathPrefix string, k8sRuntime string, log zap.SugaredLogger) (string, string, string) {
	log.Infof("Checking for %s socket\n", k8sRuntime)
	if k8sRuntime != "" {
		for _, path := range common.ContainerRuntimeSocketMap[k8sRuntime] {
			if _, err := os.Stat(pathPrefix + path); err == nil || os.IsPermission(err) {
				if (k8sRuntime == "docker" && strings.Contains(path, "containerd")) || k8sRuntime == "containerd" {
					if nriPath, err := DetectNRI(pathPrefix); err == nil {
						return k8sRuntime, path, nriPath
					} else {
						log.Warnf("%s", err)
					}
				}
				return k8sRuntime, path, ""
			} else {
				log.Warnf("%s", err)
			}
		}
	}
	log.Warn("Couldn't detect k8s runtime location, searching for other runtime sockets")
	for runtime, paths := range common.ContainerRuntimeSocketMap {
		for _, path := range paths {
			if _, err := os.Stat(pathPrefix + path); err == nil || os.IsPermission(err) {
				return runtime, path, ""
			} else {
				log.Warnf("%s", err)
			}
		}
	}
	log.Warn("Couldn't detect runtime")
	return "NA", "NA", "NA"
}
