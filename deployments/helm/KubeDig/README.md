## Install KubeDig
Install KubeDig using Helm chart repo. Also see [values](#Values) for your respective environment.
```
helm repo add kubedig https://kubedig.github.io/charts
helm repo update kubedig
helm upgrade --install kubedig kubedig/kubedig -n kubedig --create-namespace
```

Install KubeDig using Helm charts locally (for testing)
```
cd deployments/helm/KubeDig
helm upgrade --install kubedig . -n kubedig --create-namespace
```

## Values
| Key | Type | Default | Description |
|-----|------|---------|-------------|
| environment.name | string | generic | The target environment to install KubeDig in. Possible values: generic, GKE, EKS, BottleRocket, k0s, k3s, minikube, microk8s |
| kubedig.image.repository | string | kubedig/kubedig | kubedig image repo |
| kubedig.image.tag | string | stable | kubedig image tag |
| kubedig.imagePullPolicy | string | Always | kubedig imagePullPolicy |
| kubedig.args | list | [] | Specify additional args to the kubedig daemon. See [kubedig-args](#kubedig-args) |
| kubedig.configMap.defaultFilePosture | string | audit | Default file posture for KubeDig |
| kubedig.configMap.defaultNetworkPosture | string | audit | Default network posture for KubeDig |
| kubedig.configMap.defaultCapabilitiesPosture | string | audit | Default capabilities posture for KubeDig |
| kubedig.configMap.visibility | string | audit | Default visibility for KubeDig |
| kubedigRelay.enable | bool | true | to enable/disable kubedig-relay |
| kubedigRelay.image.repository | string | kubedig/kubedig-relay | kubedig-relay image repo |
| kubedigRelay.image.tag | string | latest | kubedig-relay image tag |
| kubedigRelay.imagePullPolicy | string | Always | kubedig-relay imagePullPolicy |
| kubedigInit.image.repository | string | kubedig/kubedig-init | kubedig-init image repo |
| kubedigInit.image.tag | string | stable | kubedig-init image tag |
| kubedigInit.imagePullPolicy | string | Always | kubedig-init imagePullPolicy |
| kubeRbacProxy.image.repository | string | gcr.io/kubebuilder/kube-rbac-proxy | kube-rbac-proxy image repo |
| kubeRbacProxy.image.tag | string | v0.15.0 | kube-rbac-proxy image tag |
| kubeRbacProxy.imagePullPolicy | string | Always | kube-rbac-proxy imagePullPolicy |
| kubedigController.replicas | int | 1 | kubedig-controller replicas |
| kubedigController.image.repository | string | kubedig/kubedig-controller | kubedig-controller image repo |
| kubedigController.image.tag | string | latest | kubedig-controller image tag |
| kubedigController.mutation.failurePolicy | string | Ignore | kubedig-controller failure policy |
| kubedigController.imagePullPolicy | string | Always | kubedig-controller imagePullPolicy |

## kubedig-args
```
$ sudo ./kubedig -h
Usage of ./kubedig:
  -alertThrottling
        enabling Alert Throttling
  -bpfFsPath string
        Path to the BPF filesystem to use for storing maps (default "/sys/fs/bpf")
  -cluster string
        cluster name (default "default")
  -coverageTest
        enabling CoverageTest
  -criSocket string
        path to CRI socket (format: unix:///path/to/file.sock)
  -defaultCapabilitiesPosture string
        configuring default enforcement action in global capability context {allow|audit|block} (default "audit")
  -defaultFilePosture string
        configuring default enforcement action in global file context {allow|audit|block} (default "audit")
  -defaultNetworkPosture string
        configuring default enforcement action in global network context {allow|audit|block} (default "audit")
  -enableKubeDigHostPolicy
        enabling KubeDigHostPolicy
  -enableKubeDigPolicy
        enabling KubeDigPolicy (default true)
  -enableKubeDigVm
        enabling KubeDigVM
  -gRPC string
        gRPC port number (default "32767")
  -host string
        host name (default "kubedig-dev-next")
  -hostDefaultCapabilitiesPosture string
        configuring default enforcement action in global capability context {allow|audit|block} (default "audit")
  -hostDefaultFilePosture string
        configuring default enforcement action in global file context {allow|audit|block} (default "audit")
  -hostDefaultNetworkPosture string
        configuring default enforcement action in global network context {allow|audit|block} (default "audit")
  -hostVisibility string
        Host Visibility to use [process,file,network,capabilities,none] (default "none" for k8s, "process,file,network,capabilities" for VM) (default "default")
  -k8s
        is k8s env? (default true)
  -kubeconfig string
        Paths to a kubeconfig. Only required if out-of-cluster.
  -logPath string
        log file path, {path|stdout|none} (default "none")
  -lsm string
        lsm preference order to use, available lsms [bpf, apparmor, selinux] (default "bpf,apparmor,selinux")
  -maxAlertPerSec int
        Maximum alerts allowed per second (default 10)
  -seLinuxProfileDir string
        SELinux profile directory (default "/tmp/kubedig.selinux")
  -throttleSec int
        Time period for which subsequent alerts will be dropped (in sec) (default 30)
  -visibility string
        Container Visibility to use, available visibility [process,file,network,capabilities,none] (default "process,network")
```

## Verify if all the resources are up and running
```
kubectl get all -n kubedig -l kubedig-app
NAME                                        READY   STATUS    RESTARTS   AGE
pod/kubedig-controller-7b48cf777f-bn7d8   2/2     Running   0          24s
pod/kubedig-relay-5656cc5bf7-jl56q        1/1     Running   0          24s
pod/kubedig-cnc7b                         1/1     Running   0          24s

NAME                                           TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
service/kubedig-controller-metrics-service   ClusterIP   10.43.208.188   <none>        8443/TCP   24s

NAME                       DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR            AGE
daemonset.apps/kubedig   1         1         1       1            1           kubernetes.io/os=linux   24s

NAME                                   READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/kubedig-controller   1/1     1            1           24s
deployment.apps/kubedig-relay        1/1     1            1           24s

NAME                                              DESIRED   CURRENT   READY   AGE
replicaset.apps/kubedig-controller-7b48cf777f   1         1         1       24s
replicaset.apps/kubedig-relay-5656cc5bf7        1         1         1       24s
```

## Remove KubeDig
Uninstall KubeDig using helm
```
helm uninstall kubedig -n kubedig
```
