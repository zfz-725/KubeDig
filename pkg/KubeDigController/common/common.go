// SPDX-License-Identifier: Apache-2.0
// Copyright 2022 Authors of KubeDig

package common

import (
	"context"
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const k8sVisibility = "process,file,network,capabilities"
const appArmorAnnotation = "container.apparmor.security.beta.kubernetes.io/"
const KubeDigRestartedAnnotation = "kubedig.kubernetes.io/restartedAt"

// == Add AppArmor annotations == //
func AppArmorAnnotator(pod *corev1.Pod, binding *corev1.Binding, isBinding bool) {
	podAnnotations := map[string]string{}
	var podOwnerName string

	// podOwnerName is the pod name for static pods and parent object's name
	// in other cases
	for _, ownerRef := range pod.ObjectMeta.OwnerReferences {
		// pod is owned by a replicaset, daemonset etc thus we use the managing
		// controller's name
		if ownerRef.Controller != nil && *ownerRef.Controller {
			podOwnerName = ownerRef.Name

			if ownerRef.Kind == "ReplicaSet" {
				// if it belongs to a replicaset, we also remove the pod template hash
				podOwnerName = strings.TrimSuffix(podOwnerName, fmt.Sprintf("-%s", pod.ObjectMeta.Labels["pod-template-hash"]))
			}
		}
	}

	if podOwnerName == "" {
		// pod is standalone, name remains constant
		podOwnerName = pod.ObjectMeta.Name
	}

	// Get existant kubedig annotations
	for k, v := range pod.Annotations {
		if strings.HasPrefix(k, appArmorAnnotation) {
			if v == "unconfined" {
				containerName := strings.Split(k, "/")[1]
				podAnnotations[containerName] = v
			} else {
				containerName := strings.Split(k, "/")[1]
				podAnnotations[containerName] = strings.Split(v, "/")[1]
			}
		}
	}

	// Get the remaining containers / not addressed explecitly in the annotation
	for _, container := range pod.Spec.Containers {
		if _, ok := podAnnotations[container.Name]; !ok {
			podAnnotations[container.Name] = "kubedig-" + pod.Namespace + "-" + podOwnerName + "-" + container.Name
		}
	}
	// Add kubedig annotations to the pod
	for k, v := range podAnnotations {
		if v == "unconfined" {
			continue
		}
		if isBinding {
			binding.Annotations[appArmorAnnotation+k] = "localhost/" + v
		} else {
			pod.Annotations[appArmorAnnotation+k] = "localhost/" + v
		}
	}
}
func AddCommonAnnotations(obj *metav1.ObjectMeta) {

	if obj.Annotations == nil {
		obj.Annotations = map[string]string{}
	}

	// == Policy == //

	if _, ok := obj.Annotations["kubedig-policy"]; !ok {
		// if no annotation is set enable kubedig by default
		obj.Annotations["kubedig-policy"] = "enabled"
	} else if obj.Annotations["kubedig-policy"] != "enabled" && obj.Annotations["kubedig-policy"] != "disabled" && obj.Annotations["kubedig-policy"] != "audited" {
		// if kubedig policy is not set correctly, default it to enabled
		obj.Annotations["kubedig-policy"] = "enabled"
	}
	// == Exception == //

	// exception: kubernetes app
	if obj.Namespace == "kube-system" {
		if _, ok := obj.Labels["k8s-app"]; ok {
			obj.Annotations["kubedig-policy"] = "audited"
		}

		if value, ok := obj.Labels["component"]; ok {
			if value == "etcd" || value == "kube-apiserver" || value == "kube-controller-manager" || value == "kube-scheduler" || value == "kube-proxy" {
				obj.Annotations["kubedig-policy"] = "audited"
			}
		}
	}

	// exception: cilium-operator
	if _, ok := obj.Labels["io.cilium/app"]; ok {
		obj.Annotations["kubedig-policy"] = "audited"
	}

	// exception: kubedig
	if _, ok := obj.Labels["kubedig-app"]; ok {
		obj.Annotations["kubedig-policy"] = "audited"
	}

	// == Visibility == //

	if _, ok := obj.Annotations["kubedig-visibility"]; !ok {
		obj.Annotations["kubedig-visibility"] = k8sVisibility
	}
}

func RemoveApparmorAnnotation(pod *corev1.Pod) {
	annotations := []string{}

	for key := range pod.Annotations {
		if strings.HasPrefix(key, "container.apparmor.security.beta.kubernetes.io/") {
			annotations = append(annotations, key)
		}
	}
	for _, key := range annotations {
		delete(pod.Annotations, key)
	}
}

func CheckKubedigStatus(nodeName string, c *kubernetes.Clientset) (bool, error) {
	pods, err := c.CoreV1().Pods("kubedig").List(context.TODO(), metav1.ListOptions{
		LabelSelector: "kubedig-app=kubedig",
	})
	if err != nil {
		return false, fmt.Errorf("failed to list pods: %v", err)
	}
	// Filter Pods by nodeName and return their status.phase
	for _, pod := range pods.Items {
		if pod.Spec.NodeName == nodeName {
			return true, nil
		}
	}

	return false, nil

}
func hasApparmorAnnotation(annotations map[string]string) bool {
	for key := range annotations {
		if strings.HasPrefix(key, "container.apparmor.security.beta.kubernetes.io/") {
			return true
		}
	}
	return false
}

func HandleAppArmor(annotations map[string]string) bool {
	return !hasApparmorAnnotation(annotations)
}

func HandleBPF(annotations map[string]string) bool {
	return hasApparmorAnnotation(annotations)
}

func IsAppArmorExempt(labels map[string]string, namespace string) bool {

	// exception: kubernetes app
	if namespace == "kube-system" {
		if _, ok := labels["k8s-app"]; ok {
			return true
		}

		if value, ok := labels["component"]; ok {
			if value == "etcd" || value == "kube-apiserver" || value == "kube-controller-manager" || value == "kube-scheduler" || value == "kube-proxy" {
				return true
			}
		}
	}

	// exception: cilium-operator
	if _, ok := labels["io.cilium/app"]; ok {
		return true
	}

	// exception: kubedig
	if _, ok := labels["kubedig-app"]; ok {
		return true
	}
	return false
}
