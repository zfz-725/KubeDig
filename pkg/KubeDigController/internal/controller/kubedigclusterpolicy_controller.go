// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 Authors of KubeDig

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	securityv1 "github.com/zfz-725/KubeDig/pkg/KubeDigController/api/security.kubedig.com/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// KubeDigClusterPolicyReconciler reconciles a KubeDigClusterPolicy object
type KubeDigClusterPolicyReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=security.kubedig.com,resources=kubedigclusterpolicies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=security.kubedig.com,resources=kubedigclusterpolicies/status,verbs=get;update;patch

func (r *KubeDigClusterPolicyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	return ctrl.Result{}, nil
}

func (r *KubeDigClusterPolicyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&securityv1.KubeDigClusterPolicy{}).
		Complete(r)
}
