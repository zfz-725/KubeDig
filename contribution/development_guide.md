# Development Guide

## Development

### 1. Vagrant Environment (Recommended)

   > **Note** Skip the steps for the vagrant setup if you're directly compiling KubeDig on the Linux host.
    Proceed [here](#2-self-managed-kubernetes) to setup K8s on the same host by resolving any dependencies.

   * Requirements

     Here is the list of requirements for a Vagrant environment

     ```text
     Vagrant - v2.2.9
     VirtualBox - v6.1
     ```

     Clone the KubeDig github repository in your system

     ```text
     $ git clone https://github.com/zfz-725/KubeDig.git
     ```

     Install Vagrant and VirtualBox in your environment, go to the vagrant path and run the setup.sh file

     ```text
     $ cd KubeDig/contribution/vagrant
     ~/KubeDig/contribution/vagrant$ ./setup.sh
     ~/KubeDig/contribution/vagrant$ sudo reboot
     ```

  * VM Setup using Vagrant

      Now, it is time to prepare a VM for development.

      To create a vagrant VM

      ```text
      ~/KubeDig/KubeDig$ make vagrant-up
      ```

      Output will show up as ...

      <details>
      <summary>Click to expand!</summary>

      ```text
      cd /home/gourav/KubeDig/contribution/vagrant; NETNEXT=0 DLV_RPORT=2345 vagrant up; true
      Bringing machine 'kubedig-dev' up with 'virtualbox' provider...
      ==> kubedig-dev: Importing base box 'ubuntu/bionic64'...
      ==> kubedig-dev: Matching MAC address for NAT networking...
      ==> kubedig-dev: Checking if box 'ubuntu/bionic64' version '20220131.0.0' is up to date...
      ==> kubedig-dev: Setting the name of the VM: kubedig-dev
      ==> kubedig-dev: Clearing any previously set network interfaces...
      ==> kubedig-dev: Preparing network interfaces based on configuration...
          kubedig-dev: Adapter 1: nat
      ==> kubedig-dev: Forwarding ports...
          kubedig-dev: 2345 (guest) => 2345 (host) (adapter 1)
          kubedig-dev: 22 (guest) => 2222 (host) (adapter 1)
      ==> kubedig-dev: Running 'pre-boot' VM customizations...
      ==> kubedig-dev: Booting VM...
      ==> kubedig-dev: Waiting for machine to boot. This may take a few minutes...
          kubedig-dev: SSH address: 127.0.0.1:2222
          kubedig-dev: SSH username: vagrant
          kubedig-dev: SSH auth method: private key
          kubedig-dev: Warning: Connection reset. Retrying...
          kubedig-dev: Warning: Remote connection disconnect. Retrying...
          kubedig-dev:
          kubedig-dev: Vagrant insecure key detected. Vagrant will automatically replace
          kubedig-dev: this with a newly generated keypair for better security.
          kubedig-dev:
          kubedig-dev: Inserting generated public key within guest...
          kubedig-dev: Removing insecure key from the guest if it's present...
          kubedig-dev: Key inserted! Disconnecting and reconnecting using new SSH key...
      ==> kubedig-dev: Machine booted and ready!
      ==> kubedig-dev: Checking for guest additions in VM...
          kubedig-dev: The guest additions on this VM do not match the installed version of
          kubedig-dev: VirtualBox! In most cases this is fine, but in rare cases it can
          kubedig-dev: prevent things such as shared folders from working properly. If you see
          kubedig-dev: shared folder errors, please make sure the guest additions within the
          kubedig-dev: virtual machine match the version of VirtualBox you have installed on
          kubedig-dev: your host and reload your VM.
          kubedig-dev:
          kubedig-dev: Guest Additions Version: 5.2.42
          kubedig-dev: VirtualBox Version: 6.1
      ==> kubedig-dev: Setting hostname...
      ==> kubedig-dev: Mounting shared folders...
          kubedig-dev: /vagrant => /home/gourav/KubeDig/contribution/vagrant
          kubedig-dev: /home/vagrant/KubeDig => /home/gourav/KubeDig
      ==> kubedig-dev: Running provisioner: file...
          kubedig-dev: ~/.ssh/id_rsa.pub => /home/vagrant/.ssh/id_rsa.pub
      ==> kubedig-dev: Running provisioner: shell...
          kubedig-dev: Running: inline script
      ==> kubedig-dev: Running provisioner: file...
          kubedig-dev: ~/.gitconfig => $HOME/.gitconfig
      ==> kubedig-dev: Running provisioner: shell...
          kubedig-dev: Running: /tmp/vagrant-shell20220202-55671-bn8u0f.sh
          ...
      ```

      </details>

      To get into the vagrant VM

      ```text
      ~/KubeDig/KubeDig$ make vagrant-ssh
      ```

      Output will show up as ...

      <details>
      <summary>Click to expand!</summary>

      ```text
      d /home/gourav/KubeDig/contribution/vagrant; NETNEXT=0 DLV_RPORT=2345 vagrant ssh; true
      Welcome to Ubuntu 18.04.6 LTS (GNU/Linux 4.15.0-167-generic x86_64)

       * Documentation:  https://help.ubuntu.com
       * Management:     https://landscape.canonical.com
       * Support:        https://ubuntu.com/advantage

        System information as of Wed Feb  2 10:35:55 UTC 2022

        System load:  0.06               Processes:              128
        Usage of /:   11.1% of 38.71GB   Users logged in:        0
        Memory usage: 10%                IP address for enp0s3:  10.0.2.15
        Swap usage:   0%                 IP address for docker0: 172.17.0.1


      5 updates can be applied immediately.
      1 of these updates is a standard security update.
      To see these additional updates run: apt list --upgradable

      New release '20.04.3 LTS' available.
      Run 'do-release-upgrade' to upgrade to it.


      vagrant@kubedig-dev:~$
      ```

      </details>

      To destroy the vagrant VM

      ```text
      ~/KubeDig/KubeDig$ make vagrant-destroy
      ```

    * VM Setup using Vagrant with Ubuntu 21.10 (v5.13)

      To use the recent Linux kernel v5.13 for dev env, you can run `make` with the `NETNEXT` flag set to `1` for the respective make option.

      ```text
      ~/KubeDig/KubeDig$ make vagrant-up NETNEXT=1
      ```

       You can also make the setting static by changing `NETNEXT=0` to `NETNEXT=1` in the Makefile.

      ```text
      ~/KubeDig/KubeDig$ vi Makefile
      ```

### 2. Self-managed Kubernetes
   * Requirements

     Here is the list of minimum requirements for self-managed Kubernetes.

     ```text
     OS - Ubuntu 18.04
     Kubernetes - v1.19
     Docker - 18.09 or Containerd - 1.3.7
     Linux Kernel - v4.15
     LSM - AppArmor
     ```

     KubeDig is designed for Kubernetes environment. If Kubernetes is not setup yet, please refer to [Kubernetes installation guide](self-managed-k8s/README.md).
     KubeDig leverages CRI (Container Runtime Interfaces) APIs and works with Docker or Containerd or CRIO based container runtimes. KubeDig uses LSMs for policy enforcement; thus, please make sure that your environment supports LSMs \(either AppArmor or bpf-lsm\). Otherwise, KubeDig will operate in Audit-Mode with no policy "enforcement" support.

        #### Alternative Setup
        You can try the following alternative if you face any difficulty in the above Kubernetes (kubeadm) setup.

        > **Note** Please make sure to set up the alternative k8s environment on the same host where the KubeDig development environment is running.
      * K3s

        You can also develop and test KubeDig on K3s instead of the self-managed Kubernetes.
        Please follow the instructions in [K3s installation guide](k3s/README.md).

      * MicroK8s

        You can also develop and test KubeDig on MicroK8s instead of the self-managed Kubernetes.
        Please follow the instructions in [MicroK8s installation guide](microk8s/README.md).

      * No Support - Docker Desktops

        KubeDig does not work with Docker Desktops on Windows and macOS because KubeDig integrates with Linux-kernel native primitives (including LSMs).


   * Development Setup

     In order to install all dependencies, please run the following command.

     ```text
     $ cd KubeDig/contribution/self-managed-k8s
     ~/KubeDig/contribution/self-managed-k8s$ ./setup.sh
     ```

     [setup.sh](self-managed-k8s/setup.sh) will automatically install [BCC](https://github.com/iovisor/bcc/blob/master/INSTALL.md), [Go](https://go.dev/doc/install), [Protobuf](https://grpc.io/docs/protoc-installation/), and some other dependencies.

     Now, you are ready to develop any code for KubeDig. Enjoy your journey with KubeDig.

### 3.  Environment Check
   * Compilation

        Check if KubeDig can be compiled on your environment without any problems.

        ```text
        $ cd KubeDig/KubeDig
        ~/KubeDig/KubeDig$ make
        ```

        If you see any error messages, please let us know the issue with the full error messages through #kubedig-development channel on CNCF slack.

   * Execution

        In order to directly run KubeDig in a host (not as a container), you need to run a local proxy in advance.

        ```text
        $ kubectl proxy &
        ```

        Then, run KubeDig on your environment.

        ```text
        $ cd KubeDig/KubeDig
        ~/KubeDig/KubeDig$ make run
        ```
        > **Note** If you have followed all the above steps and still getting the warning `The node information is not available`, then this could be due to the case-sensitivity discrepancy in the actual hostname (obtained by running `hostname`) and the hostname used by Kubernetes (under `kubectl get nodes -o wide`).<br>
        K8s converts the hostname to lowercase, which results in a mismatch with the actual hostname.<br>
        To resolve this, change the hostname to lowercase using the command `hostnamectl set-hostname <lowercase-hostname>`.

   * KubeDig Controller

      Starting from KubeDig v0.11 - annotations, container policies, and host policies are handled via kubedig controller, the controller code can be found under `pkg/KubeDigController`.

      To install the controller from KubeDig docker repository run
      ```text
      $ cd KubeDig/pkg/KubeDigController
      ~/KubeDig/pkg/KubeDigController$ make deploy
      ```

      To install the controller (local version) to your cluster run
      ```text
      $ cd KubeDig/pkg/KubeDigController
      ~/KubeDig/pkg/KubeDigController$ make docker-build deploy
      ```

      if you need to setup a local registry to push you image, use `docker-registry.sh` script under `~/KubeDig/contribution/local-registry` directory

## Code Directories

Here, we briefly give you an overview of KubeDig's directories.

* Source code for KubeDig \(/KubeDig\)

  ```text
  KubeDig/
    BPF                  - eBPF code for system monitor
    common               - Libraries internally used
    config               - Configuration loader
    core                 - The main body (start point) of KubeDig
    enforcer             - Runtime policy enforcer (enforcing security policies into LSMs)
    feeder               - gRPC-based feeder (sending audit/system logs to a log server)
    kvmAgent             - KubeDig VM agent
    log                  - Message logger (stdout)
    monitor              - eBPF-based system monitor (mapping process IDs to container IDs)
    policy               - gRPC service to manage Host Policies for VM environments
    types                - Type definitions
  protobuf/              - Protocol buffer
  ```

* Source code for KubeDig Controller \(CRD\)

  ```text
  pkg/KubeDigController/  - KubeDigController generated by Kube-Builder for KubeDig Annotations, KubeDigPolicy and KubeDigHostPolicy
  ```

* Deployment tools and files for KubeDig
  ```text
  deployments/
    <cloud-platform-name>   - Deployments specific to respective cloud platform (deprecated - use karmor install or helm)
    controller              - Deployments for installing KubeDigController alongwith cert-manager
    CRD                     - KubeDigPollicy and KubeDigHostPolicy CRDs
    get                     - Stores source code for deploygen, a tool used for specifying kubedig deployments
    helm/
        KubeDig           - KubeDig's Helm chart
        KubeDigOperator   - KubeDigOperator's Helm chart
  ```

* Files for testing

  ```text
  examples/     - Example microservices for testing
  tests/        - Automated test framework for KubeDig
  ```
