# Testing Guide

There are two ways to check the functionalities of KubeDig: 1) testing KubeDig manually and 2) using the testing framework.

# 0. Make sure that the annotation controller is installed on the cluster (Applicable for Steps 1 and 2)

- To install the controller from KubeDig docker repository to your cluster run

```text
$ cd KubeDig/pkg/KubeDigAnnotation
~/KubeDig/pkg/KubeDigAnnotation$ make deploy
```
- To install the controller (local version) to your cluster run

```text
$ cd KubeDig/pkg/KubeDigAnnotation
~/KubeDig/pkg/KubeDigAnnotation$ make docker-build deploy
```

# 1.  Test KubeDig manually

## 1.1. Run 'kubectl proxy' in background

```text
$ kubectl proxy &
```

## 1.2. Compile KubeDig

```text
$ cd KubeDig/KubeDig
~/KubeDig/KubeDig$ make clean && make
```

## 1.3. Run KubeDig

```text
~/KubeDig/KubeDig$ sudo -E ./kubedig -gRPC=[gRPC port number]
                                           -logPath=[log file path]
                                           -enableKubeDigPolicy=[true|false]
                                           -enableKubeDigHostPolicy=[true|false]
```

## 1.4. Apply security policies into Kubernetes

Beforehand, check if the KubeDigPolicy and KubeDigHostPolicy CRDs are already applied.

```text
$ kubectl explain KubeDigPolicy
```

If they are still not applied, do so.

```text
$ kubectl apply -f ~/KubeDig/deployments/CRD/
```

Now you can apply specific policies.

```text
$ kubectl apply -f [policy file]
```

You can refer to security policies defined for example microservices in [examples](../examples).

## 1.5. Trigger policy violations to generate alerts

```text
$ kubectl -n [namespace name] exec -it [pod name] -- bash -c [command]
```

## 1.6. Check generated alerts

- Watch alerts using [karmor](https://github.com/zfz-725/kubedig-client) cli tool

    ```text
    $ karmor log [flags]
    ```
    
    flags:

    ```text
    --gRPC string        gRPC server information
    --help               help for log
    --json               Flag to print alerts and logs in the JSON format
    --logFilter string   What kinds of alerts and logs to receive, {policy|system|all} (default "policy")
    --logPath string     Output location for alerts and logs, {path|stdout|none} (default "stdout")
    --msgPath string     Output location for messages, {path|stdout|none} (default "none")
    ```
    
    Note that you will see alerts and logs generated right after `karmor` runs logs; thus, we recommend to run the above command in other terminal to see logs live.
    

# 2.  Test KubeDig using the auto-testing framework

## 2.1. Prepare microservices and test scenarios

The auto-testing framework operates based on two things: microservices and test scenarios for each microservice.

- Microservices

    Create a directory for a microservice in [microservices](../tests/microservices)

    ```text
    $ cd KubeDig/tests/microservices
    ~/KubeDig/tests/microservices$ mkdir [microservice name]
    ```

    Then, create YAML files for the microservice

    ```text
    $ cd KubeDig/tests/microservices/[microservice name]
    ~/KubeDig/tests/microservices/[microservice name]$ ...
    ```

    As an example, we created 'multiubuntu' in [microservices](../tests/microservices) and defined 'multiubuntu-deployment.yaml' in [multiubuntu](../examples/multiubuntu).

- Test scenarios

    Create a directory whose name is like '[microservice name]_[scenario name]' in [scenarios](../tests/scenarios)
    
    ```text
    $ cd KubeDig/tests/scenarios
    ~/KubeDig/tests/scenarios$ mkdir [microservice name]_[scenario name]
    ```
    
    Then, define a YAML file for a test policy in the directory
    
    ```text
    ~/KubeDig/tests/scenarios$ cd [microservice name]_[scenario name]
    .../[microservice name]_[scenario name]$ vi [policy name].yaml
    ```

    Create cmd files whose names are like 'cmd#'
    
    ```text
    .../[microservice name]_[scenario name]$ vi cmd1 / cmd2 / ...
    ```
    
    Here is a template for a cmd file.

    ```text
    source: [pod name]
    cmd: [command to trigger a policy violation]
    result: [expected result], { passed | failed }
    ---
    operation: [operation], { Process | File | Network }
    condition: [matching string]
    action: [action in a policy] { Allow | Audit | Block }
    ```

    This is a cmd example of a test scenario.

    ```text
    source: ubuntu-1-deployment
    cmd: sleep 1
    result: failed
    ---
    operation: Process
    condition: sleep
    action: Block
    ```

    You can refer to predefined testcases in [scenarios](../tests/scenarios).

## 2.2. Test KubeDig
    
- The case that KubeDig is directly running in a host

    Compile KubeDig

    ```text
    $ cd KubeDig/KubeDig
    ~/KubeDig/KubeDig$ make clean && make
    ```

    Run the auto-testing framework

    ```text
    $ cd KubeDig/tests
    ~/KubeDig/tests$ ./test-scenarios-local.sh
    ```

    Check the test report

    ```text
    ~/KubeDig/tests$ cat /tmp/kubedig.test
    ```

- The case that KubeDig is running as a daemonset in Kubernetes

    Run the testing framework

    ```text
    $ cd KubeDig/tests
    ~/KubeDig/tests$ ./test-scenarios-in-runtime.sh
    ```

    Check the test report

    ```text
    ~/KubeDig/tests$ cat /tmp/kubedig.test
    ```
