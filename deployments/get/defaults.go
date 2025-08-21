// SPDX-License-Identifier: Apache-2.0
// Copyright 2021 Authors of KubeDig

package deployments

import (
	corev1 "k8s.io/api/core/v1"
)

var kubedig = "kubedig"
var port int32 = 32767

// K8s Object Name Defaults
var (
	KubeDigServiceAccountName                      = kubedig
	KubeDigClusterRoleBindingName                  = "kubedig-clusterrolebinding"
	KubeDigClusterRoleName                         = "kubedig-clusterrole"
	RelayServiceName                                 = kubedig
	RelayServiceAccountName                          = "kubedig-relay"
	RelayClusterRoleName                             = "kubedig-relay-clusterrole"
	RelayClusterRoleBindingName                      = "kubedig-relay-clusterrolebinding"
	RelayDeploymentName                              = "kubedig-relay"
	KubeDigConfigMapName                           = "kubedig-config"
	KubeDigControllerDeploymentName                = "kubedig-controller"
	KubeDigControllerServiceAccountName            = KubeDigControllerDeploymentName
	KubeDigControllerClusterRoleName               = "kubedig-controller-clusterrole"
	KubeDigControllerClusterRoleBindingName        = "kubedig-controller-clusterrolebinding"
	KubeDigControllerLeaderElectionRoleName        = "kubedig-controller-leader-election-role"
	KubeDigControllerLeaderElectionRoleBindingName = "kubedig-controller-leader-election-rolebinding"
	KubeDigControllerWebhookServiceName            = "kubedig-controller-webhook-service"
	KubeDigControllerSecretName                    = "kubedig-controller-webhook-server-cert"
	KubeDigControllerMutatingWebhookConfiguration  = "kubedig-controller-mutating-webhook-configuration"
)

// DaemonSetConfig Structure
type DaemonSetConfig struct {
	Args         []string
	Envs         []corev1.EnvVar
	VolumeMounts []corev1.VolumeMount
	Volumes      []corev1.Volume
}

var hostPathDirectory = corev1.HostPathDirectory
var hostPathDirectoryOrCreate = corev1.HostPathDirectoryOrCreate
var hostPathFile = corev1.HostPathFile
var hostPathSocket = corev1.HostPathSocket

var gkeHostUsrVolMnt = corev1.VolumeMount{
	Name:      "usr-src-path", // /usr -> /media/root/usr (read-only) check issue #579 for details
	MountPath: "/media/root/usr",
	ReadOnly:  true,
}

var gkeHostUsrVol = corev1.Volume{ // check #579 why GKE is handled separately
	Name: "usr-src-path",
	VolumeSource: corev1.VolumeSource{
		HostPath: &corev1.HostPathVolumeSource{
			Path: "/usr",
			Type: &hostPathDirectory,
		},
	},
}

var hostUsrVolMnt = corev1.VolumeMount{
	Name:      "usr-src-path", // /usr/src (read-only)
	MountPath: "/usr/src",
	ReadOnly:  true,
}

var hostUsrVol = corev1.Volume{
	Name: "usr-src-path",
	VolumeSource: corev1.VolumeSource{
		HostPath: &corev1.HostPathVolumeSource{
			Path: "/usr/src",
			Type: &hostPathDirectory,
		},
	},
}

var apparmorVolMnt = corev1.VolumeMount{
	Name:      "etc-apparmor-d-path",
	MountPath: "/etc/apparmor.d",
}

var apparmorVol = corev1.Volume{
	Name: "etc-apparmor-d-path",
	VolumeSource: corev1.VolumeSource{
		HostPath: &corev1.HostPathVolumeSource{
			Path: "/etc/apparmor.d",
			Type: &hostPathDirectoryOrCreate,
		},
	},
}

var envVar = []corev1.EnvVar{
	{
		Name: "KUBEDIG_NODENAME",
		ValueFrom: &corev1.EnvVarSource{
			FieldRef: &corev1.ObjectFieldSelector{
				FieldPath: "spec.nodeName",
			},
		},
	},
	{
		Name: "KUBEDIG_NAMESPACE",
		ValueFrom: &corev1.EnvVarSource{
			FieldRef: &corev1.ObjectFieldSelector{
				FieldPath: "metadata.namespace",
			},
		},
	},
}

// Environment Specific Daemonset Configuration
var defaultConfigs = map[string]DaemonSetConfig{
	"generic": {
		Args: []string{},
		Envs: envVar,
		VolumeMounts: []corev1.VolumeMount{
			apparmorVolMnt,
			{
				Name:      "containerd-sock-path", // containerd
				MountPath: "/var/run/containerd/containerd.sock",
				ReadOnly:  true,
			},
		},
		Volumes: []corev1.Volume{
			apparmorVol,
			{
				Name: "containerd-sock-path",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: "/var/run/containerd/containerd.sock",
						Type: &hostPathSocket,
					},
				},
			},
		},
	},
	"oke": {
		Args: []string{},
		Envs: envVar,
		VolumeMounts: []corev1.VolumeMount{
			apparmorVolMnt,
			{
				Name:      "crio-sock-path", // crio socket
				MountPath: "/var/run/crio/crio.sock",
				ReadOnly:  true,
			},
		},
		Volumes: []corev1.Volume{
			apparmorVol,
			{
				Name: "crio-sock-path",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: "/var/run/crio/crio.sock",
						Type: &hostPathSocket,
					},
				},
			},
		},
	},
	"docker": {
		Args: []string{},
		Envs: envVar,
		VolumeMounts: []corev1.VolumeMount{
			apparmorVolMnt,
			{
				Name:      "docker-sock-path", // docker
				MountPath: "/var/run/docker.sock",
				ReadOnly:  true,
			},
		},
		Volumes: []corev1.Volume{
			apparmorVol,
			{
				Name: "docker-sock-path",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: "/var/run/docker.sock",
						Type: &hostPathSocket,
					},
				},
			},
		},
	},
	"minikube": {
		Args: []string{},
		Envs: envVar,
		VolumeMounts: []corev1.VolumeMount{
			apparmorVolMnt,
			{
				Name:      "docker-sock-path", // docker
				MountPath: "/var/run/docker.sock",
				ReadOnly:  true,
			},
		},
		Volumes: []corev1.Volume{
			apparmorVol,
			{
				Name: "docker-sock-path",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: "/var/run/docker.sock",
						Type: &hostPathSocket,
					},
				},
			},
		},
	},
	"microk8s": {
		Args: []string{},
		Envs: envVar,
		VolumeMounts: []corev1.VolumeMount{
			apparmorVolMnt,
			{
				Name:      "containerd-sock-path", // containerd
				MountPath: "/var/snap/microk8s/common/run/containerd.sock",
				ReadOnly:  true,
			},
		},
		Volumes: []corev1.Volume{
			apparmorVol,
			{
				Name: "containerd-sock-path",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: "/var/snap/microk8s/common/run/containerd.sock",
						Type: &hostPathSocket,
					},
				},
			},
		},
	},
	"k0s": {
		Args: []string{},
		Envs: envVar,
		VolumeMounts: []corev1.VolumeMount{
			apparmorVolMnt,
			{
				Name:      "containerd-sock-path",
				MountPath: "/var/run/containerd/containerd.sock",
				ReadOnly:  true,
			},
		},
		Volumes: []corev1.Volume{
			apparmorVol,
			{
				Name: "containerd-sock-path",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: "/run/k0s/containerd.sock",
						Type: &hostPathSocket,
					},
				},
			},
		},
	},
	"k3s": {
		Args: []string{},
		Envs: envVar,
		VolumeMounts: []corev1.VolumeMount{
			apparmorVolMnt,
			{
				Name:      "containerd-sock-path", // containerd
				MountPath: "/var/run/containerd/containerd.sock",
				ReadOnly:  true,
			},
		},
		Volumes: []corev1.Volume{
			apparmorVol,
			{
				Name: "containerd-sock-path",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: "/run/k3s/containerd/containerd.sock",
						Type: &hostPathSocket,
					},
				},
			},
		},
	},
	"gke": {
		Args: []string{},
		Envs: envVar,
		VolumeMounts: []corev1.VolumeMount{
			apparmorVolMnt,
			{
				Name:      "containerd-sock-path", // containerd
				MountPath: "/var/run/containerd/containerd.sock",
				ReadOnly:  true,
			},
		},
		Volumes: []corev1.Volume{
			apparmorVol,
			{
				Name: "containerd-sock-path",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: "/var/run/containerd/containerd.sock",
						Type: &hostPathSocket,
					},
				},
			},
		},
	},
	"eks": {
		Args: []string{},
		Envs: envVar,
		VolumeMounts: []corev1.VolumeMount{
			apparmorVolMnt,
			{
				Name:      "containerd-sock-path", // containerd
				MountPath: "/var/run/containerd/containerd.sock",
				ReadOnly:  true,
			},
		},
		Volumes: []corev1.Volume{
			apparmorVol,
			{
				Name: "containerd-sock-path",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: "/var/run/containerd/containerd.sock",
						Type: &hostPathSocket,
					},
				},
			},
		},
	},
	"bottlerocket": {
		Args: []string{},
		Envs: envVar,
		VolumeMounts: []corev1.VolumeMount{
			apparmorVolMnt,
			{
				Name:      "containerd-sock-path", // containerd
				MountPath: "/var/run/containerd/containerd.sock",
				ReadOnly:  true,
			},
		},
		Volumes: []corev1.Volume{
			apparmorVol,
			{
				Name: "containerd-sock-path",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: "/run/containerd/containerd.sock",
						Type: &hostPathSocket,
					},
				},
			},
		},
	},
	"aks": {
		Args: []string{},
		Envs: envVar,
		VolumeMounts: []corev1.VolumeMount{
			apparmorVolMnt,
			{
				Name:      "containerd-sock-path", // containerd
				MountPath: "/var/run/containerd/containerd.sock",
				ReadOnly:  true,
			},
		},
		Volumes: []corev1.Volume{
			apparmorVol,
			{
				Name: "containerd-sock-path",
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: "/var/run/containerd/containerd.sock",
						Type: &hostPathSocket,
					},
				},
			},
		},
	},
}
