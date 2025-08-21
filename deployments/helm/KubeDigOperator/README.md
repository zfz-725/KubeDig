# Install KubeDigOperator

Install KubeDigOperator using the official `kubedig` Helm chart repo. Also see [values](#values) for your respective environment.

```bash
helm repo add kubedig https://kubedig.github.io/charts
helm repo update kubedig
helm upgrade --install kubedig-operator kubedig/kubedig-operator -n kubedig --create-namespace
```

Install KubeDigOperator using Helm charts locally (for testing)

```bash
cd deployments/helm/KubeDigOperator
helm upgrade --install kubedig-operator . -n kubedig --create-namespace
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| kubedigOperator.name | string | kubedig-operator | name of the operator's deployment |
| kubedigOperator.image.repository | string | kubedig/kubedig-operator | image repository to pull KubeDigOperator from |
| kubedigOperator.image.tag | string | latest | KubeDigOperator image tag |
| kubedigOperator.imagePullPolicy | string | IfNotPresent | pull policy for operator image |
| kubedigOperator.podLabels | object | {} | additional pod labels |
| kubedigOperator.podAnnotations | object | {} | additional pod annotations |
| kubedigOperator.resources | object | {} | operator container resources |
| kubedigOperator.podSecurityContext | object | {} | pod security context |
| kubedigOperator.securityContext | object | {} | operator container security context |
| kubedigConfig | object | [values.yaml](values.yaml) | KubeDig default configurations |
| kubedigOperator.annotateResource | bool | false | flag to control RBAC permissions conditionally, use `--annotateResource=<value>` arg as well to pass the same value to operator configuration |
| autoDeploy | bool | false | Auto deploy KubeDig with default configurations |

The operator needs a `KubeDigConfig` object in order to create resources related to KubeDig. A default config is present in Helm `values.yaml` which can be overridden during Helm install. To install KubeDig with default configuration use `--set autoDeploy=true` flag with helm install/upgrade command. It is possible to specify configuration even after KubeDig resources have been installed by directly editing the created `KubeDigConfig` CR.

By Default the helm does not deploys the default KubeDig Configurations (KubeDigConfig CR) and once installed, the operator waits for the user to create a `KubeDigConfig` object.
## KubeDigConfig specification

```yaml
apiVersion: operator.kubedig.com/v1
kind: KubeDigConfig
metadata:
    labels:
        app.kubernetes.io/name: kubedigconfig
        app.kubernetes.io/instance: kubedigconfig-sample
        app.kubernetes.io/part-of: kubedigoperator
        app.kubernetes.io/managed-by: kustomize
        app.kubernetes.io/created-by: kubedigoperator
    name: [config name]
    namespace: [namespace name]
spec:
    # default global posture
    defaultCapabilitiesPosture: audit|block                    # DEFAULT - audit
    defaultFilePosture: audit|block                            # DEFAULT - audit
    defaultNetworkPosture: audit|block                         # DEFAULT - audit

    enableStdOutLogs: [show stdout logs for relay server]      # DEFAULT - false
    enableStdOutAlerts: [show stdout alerts for relay server]  # DEFAULT - false
    enableStdOutMsgs: [show stdout messages for relay server]  # DEFAULT - false 

    # default visibility configuration
    defaultVisibility: [comma separated: process|file|network] # DEFAULT - process,network

    # enabling NRI
    # Naming convention for kubedig daemonset in case of NRI will be effective only when initally NRI is available & enabled. 
    # In case snitch service account token is already present before its deployment, the naming convention won't show NRI, 
    # it will be based on the runtime present. This happens because operator won't get KubedigConfig event(initially).
    enableNRI: [true|false] # DEFAULT - false

    # KubeDig image and pull policy
    kubedigImage:
        image: [image-repo:tag]                                # DEFAULT - kubedig/kubedig:stable
        imagePullPolicy: [image pull policy]                   # DEFAULT - Always

    # KubeDig init image and pull policy
    kubedigInitImage:
        image: [image-repo:tag]                                # DEFAULT - kubedig/kubedig-init:stable
        imagePullPolicy: [image pull policy]                   # DEFAULT - Always

    # KubeDig relay image and pull policy
    kubedigRelayImage:
        image: [image-repo:tag]                                # DEFAULT - kubedig/kubedig-relay-server:latest
        imagePullPolicy: [image pull policy]                   # DEFAULT - Always

    # KubeDig controller image and pull policy
    kubedigControllerImage:
        image: [image-repo:tag]                                # DEFAULT - kubedig/kubedig-controller:latest
        imagePullPolicy: [image pull policy]                   # DEFAULT - Always

    # kube-rbac-proxy image and pull policy
    kubeRbacProxyImage:
        image: [image-repo:tag]                                # DEFAULT - gcr.io/kubebuilder/kube-rbac-proxy:v0.15.0
        imagePullPolicy: [image pull policy]                   # DEFAULT - Always
```

## Verify if all the resources are up and running
If a valid configuration is received, the operator will deploy jobs to your nodes to get the environment information and then start installing KubeDig components.

Once done, the following resources related to KubeDig will exist in your cluster:
```
$ kubectl get all -n kubedig -l kubedig-app
NAME                                        READY   STATUS      RESTARTS   AGE
pod/kubedig-operator-66fbff5559-qb7dh     1/1     Running     0          11m
pod/kubedig-relay-557dfcc57b-c8t55        1/1     Running     0          2m53s
pod/kubedig-controller-7879755b58-t4v8m   2/2     Running     0          2m53s
pod/kubedig-snitch-lglbd-z92gb            0/1     Completed   0          31s
pod/kubedig-bpf-docker-d4651-r5n7q        1/1     Running     0          30s

NAME                                           TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)     AGE
service/kubedig-controller-metrics-service   ClusterIP   10.43.241.84    <none>        8443/TCP    2m53s
service/kubedig                              ClusterIP   10.43.216.156   <none>        32767/TCP   2m53s

NAME                                        DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR                                                                                                                                                                       AGE
daemonset.apps/kubedig-bpf-docker-d4651   1         1         1       1            1           kubedig.io/btf=yes,kubedig.io/enforcer=bpf,kubedig.io/runtime=docker,kubedig.io/socket=run_docker.sock,kubernetes.io/os=linux   30s

NAME                                   READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/kubedig-operator     1/1     1            1           11m
deployment.apps/kubedig-relay        1/1     1            1           2m53s
deployment.apps/kubedig-controller   1/1     1            1           2m53s

NAME                                              DESIRED   CURRENT   READY   AGE
replicaset.apps/kubedig-operator-66fbff5559     1         1         1       11m
replicaset.apps/kubedig-relay-557dfcc57b        1         1         1       2m53s
replicaset.apps/kubedig-controller-7879755b58   1         1         1       2m53s

NAME                               COMPLETIONS   DURATION   AGE
job.batch/kubedig-snitch-lglbd   1/1           3s         11m
```

## Uninstall the Operator

Uninstalling the Operator will also uninstall KubeDig from all your nodes. To uninstall, just run:

```bash
helm uninstall kubedig-operator -n kubedig
```
