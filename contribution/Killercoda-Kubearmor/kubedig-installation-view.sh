#!/bin/bash
# SPDX-License-Identifier: Apache-2.0
# Copyright 2023 Authors of KubeDig

namespace="kubedig"

echo "Waiting for all pods in namespace '$namespace' to be in the 'Running' state"

kubectl wait --for=condition=ready --timeout=5m -n kubedig pod -l kubedig-app=kubedig-operator
kubectl get po -n $namespace
kubectl wait -n kubedig --timeout=5m --for=jsonpath='{.status.phase}'=Running kubedigconfigs/kubedig-default
kubectl wait --timeout=5m --for=condition=ready pod -l kubedig-app,kubedig-app!=kubedig-snitch -n kubedig
kubectl wait --timeout=5m --for=condition=ready pod -l kubedig-app=kubedig,kubedig-app!=kubedig-snitch -n kubedig

echo "All pods in namespace '$namespace' are now in the 'Running' state"

kubectl get po -n $namespace
