// SPDX-License-Identifier: Apache-2.0
// Copyright 2022 Authors of KubeDig

package common

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"os"
	"strings"

	deployments "github.com/zfz-725/KubeDig/deployments/get"
	securityv1 "github.com/zfz-725/KubeDig/pkg/KubeDigController/api/security.kubedig.com/v1"
	opv1 "github.com/zfz-725/KubeDig/pkg/KubeDigOperator/api/operator.kubedig.com/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/client-go/kubernetes"
)

const (
	// constants for CRD status
	CREATED  string = "Created"
	PENDING  string = "Pending"
	RUNNING  string = "Running"
	UPDATING string = "Updating"
	ERROR    string = "Error"

	// Status Messages
	CREATED_MSG  string = "Installation has been created"
	PENDING_MSG  string = "Kubedig Installation is in-progress"
	RUNNING_MSG  string = "Kubedig Application is Up and Running"
	UPDATING_MSG string = "Updating the Application Configuration"

	// Error Messages
	INSTALLATION_ERR_MSG    string = "Failed to install KubeDig component(s)"
	MULTIPLE_CRD_ERR_MSG    string = "There's already a CRD exists to manage KubeDig"
	UPDATION_FAILED_ERR_MSG string = "Failed to update KubeDig configuration"
)

var OperatorConfigCrd *opv1.KubeDigConfig

var (
	// node labels
	EnforcerLabel   string = "kubedig.io/enforcer"
	RuntimeLabel    string = "kubedig.io/runtime"
	SocketLabel     string = "kubedig.io/socket"
	NRISocketLabel  string = "kubedig.io/nri-socket"
	RandLabel       string = "kubedig.io/rand"
	OsLabel         string = "kubernetes.io/os"
	ArchLabel       string = "kubernetes.io/arch"
	BTFLabel        string = "kubedig.io/btf"
	ApparmorFsLabel string = "kubedig.io/apparmorfs"
	SecurityFsLabel string = "kubedig.io/securityfs"
	SeccompLabel    string = "kubedig.io/seccomp"

	// node taints label
	NotreadyTaint      string = "node.kubernetes.io/not-ready"
	UnreachableTaint   string = "node.kubernetes.io/unreachable"
	UnschedulableTaint string = "node.kubernetes.io/unschedulable"

	// if any node with securityfs/lsm present
	IfNodeWithSecurtiyFs bool = false

	DeleteAction            string = "DELETE"
	AddAction               string = "ADD"
	Namespace               string = "kubedig"
	Privileged              bool   = false
	HostPID                 bool   = false
	SnitchName              string = "kubedig-snitch"
	SnitchImage             string = "kubedig/kubedig-snitch"
	SnitchImageTag          string = "latest"
	KubeDigSnitchRoleName string = "kubedig-snitch"

	// KubeDigConfigMapName string = "kubedig-config"

	// ConfigMap Data
	ConfigGRPC                       string = "gRPC"
	ConfigVisibility                 string = "visibility"
	ConfigCluster                    string = "cluster"
	ConfigDefaultFilePosture         string = "defaultFilePosture"
	ConfigDefaultCapabilitiesPosture string = "defaultCapabilitiesPosture"
	ConfigDefaultNetworkPosture      string = "defaultNetworkPosture"
	ConfigDefaultPostureLogs         string = "defaultPostureLogs"
	ConfigAlertThrottling            string = "alertThrottling"
	ConfigMaxAlertPerSec             string = "maxAlertPerSec"
	ConfigThrottleSec                string = "throttleSec"
	ConfigEnableNRI                  string = "enableNRI"

	GlobalImagePullSecrets []corev1.LocalObjectReference = []corev1.LocalObjectReference{}
	GlobalTolerations      []corev1.Toleration           = []corev1.Toleration{}
	//KubedigRelayEnvVariables

	EnableStdOutAlerts string = "enableStdOutAlerts"
	EnableStdOutLogs   string = "enableStdOutLogs"
	EnableStdOutMsgs   string = "enableStdOutMsgs"

	// Images
	KubeDigName string   = "kubedig"
	KubeDigArgs []string = []string{
		"-gRPC=32767",
		"-procfsMount=/host/procfs",
		"-tlsEnabled=false",
	}
	KubeDigImage            string                        = "kubedig/kubedig:stable"
	KubeDigImagePullPolicy  string                        = "Always"
	KubeDigImagePullSecrets []corev1.LocalObjectReference = []corev1.LocalObjectReference{}
	KubeDigTolerations      []corev1.Toleration           = []corev1.Toleration{}

	KubeDigInitName             string                        = "kubedig-init"
	KubeDigInitArgs             []string                      = []string{}
	KubeDigInitImage            string                        = "kubedig/kubedig-init:stable"
	KubeDigInitImagePullPolicy  string                        = "Always"
	KubeDigInitImagePullSecrets []corev1.LocalObjectReference = []corev1.LocalObjectReference{}
	KubeDigInitTolerations      []corev1.Toleration           = []corev1.Toleration{}

	KubeDigRelayName string   = "kubedig-relay"
	KubeDigRelayArgs []string = []string{
		"-tlsEnabled=false",
	}
	KubeDigRelayImage            string                        = "kubedig/kubedig-relay-server:latest"
	KubeDigRelayImagePullPolicy  string                        = "Always"
	KubeDigRelayImagePullSecrets []corev1.LocalObjectReference = []corev1.LocalObjectReference{}
	KubeDigRelayTolerations      []corev1.Toleration           = []corev1.Toleration{}

	KubeDigControllerName string   = "kubedig-controller"
	KubeDigControllerArgs []string = []string{
		"--leader-elect",
		"--health-probe-bind-address=:8081",
		"--annotateExisting=false",
	}
	KubeDigControllerImage            string                        = "kubedig/kubedig-controller:latest"
	KubeDigControllerImagePullPolicy  string                        = "Always"
	KubeDigControllerImagePullSecrets []corev1.LocalObjectReference = []corev1.LocalObjectReference{}
	KubeDigControllerTolerations      []corev1.Toleration           = []corev1.Toleration{}

	SeccompProfile     = "kubedig-seccomp.json"
	SeccompInitProfile = "kubedig-init-seccomp.json"

	// tls
	EnableTls                      bool     = false
	ExtraDnsNames                  []string = []string{"localhost"}
	ExtraIpAddresses               []string = []string{"127.0.0.1"}
	KubeDigCaSecretName          string   = "kubedig-ca"
	KubeDigClientSecretName      string   = "kubedig-client-certs"
	KubeDigRelayServerSecretName string   = "kubedig-relay-server-certs"
	DefaultTlsCertPath             string   = "/var/lib/kubedig/tls"
	DefaultMode                    int32    = 420 // deciaml representation of octal value 644

	// throttling
	AlertThrottling       bool   = true
	DefaultMaxAlertPerSec string = "10"
	DefaultThrottleSec    string = "30"

	// recommend policies
	RecommendedPolicies opv1.RecommendedPolicies = opv1.RecommendedPolicies{
		MatchExpressions: []securityv1.MatchExpressionsType{
			{
				Key:      "namespace",
				Operator: "NotIn",
				Values: []string{
					"kube-system",
					"kubedig",
				},
			},
		},
	}

	Adapter opv1.Adapters = opv1.Adapters{
		ElasticSearch: opv1.ElasticSearchAdapter{
			Enabled:         false,
			Url:             "",
			AlertsIndexName: "kubedig-alerts",
			Auth: opv1.ElasticSearchAuth{
				SecretName:       "elastic-secret",
				UserNameKey:      "username",
				PasswordKey:      "password",
				AllowTlsInsecure: false,
				CAcertSecretName: "",
				CaCertKey:        "ca.crt",
			},
		},
	}

	ElasticSearchAdapterCaCertPath = "/cert"
)
var Pointer2True bool = true

var ConfigMapData = map[string]string{
	ConfigGRPC:                       "32767",
	ConfigCluster:                    "default",
	ConfigDefaultFilePosture:         "audit",
	ConfigDefaultCapabilitiesPosture: "audit",
	ConfigDefaultNetworkPosture:      "audit",
	ConfigVisibility:                 "process,network,capabilities",
	ConfigDefaultPostureLogs:         "true",
	ConfigAlertThrottling:            "true",
	ConfigMaxAlertPerSec:             "10",
	ConfigThrottleSec:                "30",
}

var ConfigDefaultSeccompEnabled = "false"

var KubedigRelayEnvMap = map[string]string{
	EnableStdOutAlerts: "false",
	EnableStdOutLogs:   "false",
	EnableStdOutMsgs:   "false",
}

var ContainerRuntimeSocketMap = map[string][]string{
	"docker": {
		"/run/containerd/containerd.sock",
		"/var/run/containerd/containerd.sock",
		"/var/run/docker.sock",
		"/run/docker.sock",
	},
	"containerd": {
		"/var/snap/microk8s/common/run/containerd.sock",
		"/run/k0s/containerd.sock",
		"/run/k3s/containerd/containerd.sock",
		"/run/containerd/containerd.sock",
		"/var/run/containerd/containerd.sock",
		"/run/dockershim.sock",
	},
	"cri-o": {
		"/var/run/crio/crio.sock",
		"/run/crio/crio.sock",
	},
	"nri": {
		"/var/run/nri/nri.sock",
		"/run/nri/nri.sock",
	},
}

var NRIEnabled = false

var HostPathDirectory = corev1.HostPathDirectory
var HostPathDirectoryOrCreate = corev1.HostPathDirectoryOrCreate
var HostPathSocket = corev1.HostPathSocket
var HostPathFile = corev1.HostPathFile

var EnforcerVolumesMounts = map[string][]corev1.VolumeMount{
	"apparmor": {
		{
			Name:      "etc-apparmor-d-path",
			MountPath: "/etc/apparmor.d",
		},
	},
}

var EnforcerVolumes = map[string][]corev1.Volume{
	"apparmor": {
		{
			Name: "etc-apparmor-d-path",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/etc/apparmor.d",
					Type: &HostPathDirectory,
				},
			},
		},
	},
}

var RuntimeSocketLocation = map[string]string{
	"docker":     "/var/run/docker.sock",
	"containerd": "/var/run/containerd/containerd.sock",
	"cri-o":      "/var/run/crio/crio.sock",
	"nri":        "/var/run/nri/nri.sock",
}

func ShortSHA(s string) string {
	sBytes := []byte(s)

	shaFunc := sha512.New()
	shaFunc.Write(sBytes)
	res := shaFunc.Sum(nil)
	return hex.EncodeToString(res)[:5]
}

var BPFVolumes = []corev1.Volume{
	{
		Name: "bpf",
		VolumeSource: corev1.VolumeSource{
			EmptyDir: &corev1.EmptyDirVolumeSource{},
		},
	},
}

var BPFVolumesMount = []corev1.VolumeMount{
	{
		Name:      "bpf",
		MountPath: "/opt/kubedig/BPF",
	},
}

var CommonVolumes = []corev1.Volume{
	{
		Name: "sys-kernel-debug-path",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/sys/kernel/debug",
				Type: &HostPathDirectory,
			},
		},
	},
	{
		Name: "proc-fs-mount",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/proc",
				Type: &HostPathDirectory,
			},
		},
	},
}

var CommonVolumesMount = []corev1.VolumeMount{
	{
		Name:      "sys-kernel-debug-path",
		MountPath: "/sys/kernel/debug",
	},
	{
		Name:      "proc-fs-mount",
		MountPath: "/host/procfs",
		ReadOnly:  true,
	},
}

var KubeDigCaVolume = []corev1.Volume{
	{
		Name: "kubedig-ca-secret",
		VolumeSource: corev1.VolumeSource{
			Secret: &corev1.SecretVolumeSource{
				SecretName: KubeDigCaSecretName,
				Items: []corev1.KeyToPath{
					{
						Key:  "tls.crt",
						Path: "ca.crt",
					},
					{
						Key:  "tls.key",
						Path: "ca.key",
					},
				},
				DefaultMode: &DefaultMode,
			},
		},
	},
}

var KubeDigCaVolumeMount = []corev1.VolumeMount{
	{
		Name:      "kubedig-ca-secret",
		MountPath: DefaultTlsCertPath,
		ReadOnly:  true,
	},
}

var KubeDigRelayTlsVolume = []corev1.Volume{
	{
		Name: "kubedig-relay-certs-secrets",
		VolumeSource: corev1.VolumeSource{
			Projected: &corev1.ProjectedVolumeSource{
				Sources: []corev1.VolumeProjection{
					{
						Secret: &corev1.SecretProjection{
							LocalObjectReference: corev1.LocalObjectReference{
								Name: KubeDigClientSecretName,
							},
							Items: []corev1.KeyToPath{
								{
									Key:  "tls.crt",
									Path: "client.crt",
								},
								{
									Key:  "tls.key",
									Path: "client.key",
								},
							},
						},
					},
					{
						Secret: &corev1.SecretProjection{
							LocalObjectReference: corev1.LocalObjectReference{
								Name: KubeDigRelayServerSecretName,
							},
							Items: []corev1.KeyToPath{
								{
									Key:  "tls.crt",
									Path: "server.crt",
								},
								{
									Key:  "tls.key",
									Path: "server.key",
								},
								{
									Key:  "ca.crt",
									Path: "ca.crt",
								},
							},
						},
					},
				},
				DefaultMode: &DefaultMode,
			},
		},
	},
}

var KubeDigRelayTlsVolumeMount = []corev1.VolumeMount{
	{
		Name:      "kubedig-relay-certs-secrets",
		MountPath: DefaultTlsCertPath,
		ReadOnly:  true,
	},
}

var KernelHeaderVolumes = []corev1.Volume{
	{
		Name: "lib-modules-path",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/lib/modules",
				Type: &HostPathDirectory,
			},
		},
	},
	{
		Name: "usr-src-path",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/usr/src",
				Type: &HostPathDirectory,
			},
		},
	},
	{
		Name: "os-release-path",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/etc/os-release",
				Type: &HostPathFile,
			},
		},
	},
}

var KernelHeaderVolumesMount = []corev1.VolumeMount{
	{
		Name:      "usr-src-path",
		MountPath: "/usr/src",
		ReadOnly:  true,
	},
	{
		Name:      "lib-modules-path",
		MountPath: "/lib/modules",
		ReadOnly:  true,
	},
	{
		Name:      "os-release-path",
		MountPath: "/media/root/etc/os-release",
		ReadOnly:  true,
	},
}

func GetFreeRandSuffix(c *kubernetes.Clientset, namespace string) (suffix string, err error) {
	var found bool
	for {
		suffix = rand.String(5)
		found = false
		if _, err = c.CoreV1().Secrets(namespace).Get(context.Background(), deployments.KubeDigControllerSecretName+"-"+suffix, metav1.GetOptions{}); err != nil {
			if !strings.Contains(err.Error(), "not found") {
				return "", err
			}
		} else {
			found = true
		}

		if !found {
			break
		}
	}
	return suffix, nil
}

func GetOperatorNamespace() string {
	ns := os.Getenv("KUBEDIG_OPERATOR_NS")

	if ns == "" {
		return Namespace
	}

	return ns
}

func GetApplicationImage(app string) string {
	// RELATED_IMAGE_* env variables will be present in case of redhat certified operator
	switch app {
	case KubeDigName:
		if image := os.Getenv("RELATED_IMAGE_KUBEDIG"); image != "" {
			return image
		}
		return KubeDigImage
	case KubeDigInitName:
		if image := os.Getenv("RELATED_IMAGE_KUBEDIG_INIT"); image != "" {
			return image
		}
		return KubeDigInitImage
	case KubeDigRelayName:
		if image := os.Getenv("RELATED_IMAGE_KUBEDIG_RELAY_SERVER"); image != "" {
			return image
		}
		return KubeDigRelayImage
	case KubeDigControllerName:
		if image := os.Getenv("RELATED_IMAGE_KUBEDIG_CONTROLLER"); image != "" {
			return image
		}
		return KubeDigControllerImage
	case SnitchName:
		if image := os.Getenv("RELATED_IMAGE_KUBEDIG_SNITCH"); image != "" {
			return image
		}
		return SnitchImage + ":" + SnitchImageTag
	}
	return ""
}

func IsCertifiedOperator() bool {
	certified := os.Getenv("REDHAT_CERTIFIED_OP")
	if certified == "" {
		return false
	}
	return true
}

func CopyStrMap(src map[string]string) map[string]string {
	newMap := make(map[string]string)
	for key, value := range src {
		newMap[key] = value
	}
	return newMap
}

func init() {
	Namespace = GetOperatorNamespace()
	if IsCertifiedOperator() {
		HostPID = true
	}
}

func AddOrReplaceArg(add, replace string, args *[]string) {
	added := false
	for i, arg := range *args {
		if arg == replace || arg == add {
			(*args)[i] = add
			added = true
			break
		}
	}
	if !added {
		*args = append(*args, add)
	}
}

func GetTlsState() bool {
	return EnableTls
}

func AddOrRemoveVolumeMount(src *[]corev1.VolumeMount, dest *[]corev1.VolumeMount, action string) {
	for i, mnt := range *dest {
		for _, m := range *src {
			if mnt.Name == m.Name {
				(*dest)[i] = (*dest)[len(*dest)-1]
				*dest = (*dest)[:len(*dest)-1]
			}
		}
	}
	if action == AddAction {
		*dest = append(*dest, *src...)
	}
}

func AddOrRemoveVolume(src *[]corev1.Volume, dest *[]corev1.Volume, action string) {
	for i, mnt := range *dest {
		for _, m := range *src {
			if mnt.Name == m.Name {
				(*dest)[i] = (*dest)[len(*dest)-1]
				*dest = (*dest)[:len(*dest)-1]
			}
		}
	}
	if action == AddAction {
		*dest = append(*dest, *src...)
	}
}

func ParseArgument(arg string) (key string, value string, found bool) {
	arg = strings.TrimLeft(arg, "-")

	parts := strings.SplitN(arg, "=", 2)
	if len(parts) != 2 {
		return "", "", false
	}

	return parts[0], parts[1], true
}

func GenerateNRIvol(nriSocket string) (vol []corev1.Volume, volMnt []corev1.VolumeMount) {
	if nriSocket != "" {
		for _, socket := range ContainerRuntimeSocketMap["nri"] {
			if strings.ReplaceAll(socket[1:], "/", "_") == nriSocket {
				vol = append(vol, corev1.Volume{
					Name: "nri-socket",
					VolumeSource: corev1.VolumeSource{
						HostPath: &corev1.HostPathVolumeSource{
							Path: socket,
							Type: &HostPathSocket,
						},
					},
				})

				socket = RuntimeSocketLocation["nri"]
				volMnt = append(volMnt, corev1.VolumeMount{
					Name:      "nri-socket",
					MountPath: socket,
					ReadOnly:  true,
				})
				break
			}
		}
	}
	return
}
