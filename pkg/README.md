# Kubernetes Controllers

## KubeDigController

KubeDigController provides CRDs for container, host policy specifications and also verifies the policies given by a user through kubectl. It also provides the admission controller that automatically adds the annotations for kubedig-policy, kubedig-visibilities, and apparmor.

```
cd KubeDigController
make              # compile the kubedig-controller
make manifests    # create the KubeDigPolicy and KubeDigHostPolicy CRD, WebhookConfiguration and ClusterRole
make docker-build # create a local image for the kubedig-controller
make deploy       # deploy the created local image for testing
make delete       # delete the controller deployed for testing
```
