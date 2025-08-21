# Minikube Installation

If virtualbox and vagrant are not installed on your machine, please run the following commands in advance.

```text
$ cd KubeDig/contribution/minikube
~/KubeDig/contribution/minikube$ ./install_virtualbox.sh
~/KubeDig/contribution/minikube$ sudo reboot
```

After rebooting the machine, please keep running the following commands.

```text
$ cd KubeDig/contribution/minikube
~/KubeDig/contribution/minikube$ ./install_minikube.sh
```

Ensure to use virtualbox driver when running minikube. This step is necessary in order to mount roofs as read/write.

```text
$ minikube config set driver virtualbox
```

In order to use KubeDig, Minikube needs to support eBPF capabilities. Unfortunately, Minikube doesn't suport them by default. We have compiled Minikube's Kernel with eBPF capablities and AppArmor which is required to enforce security policies. Thus, please run the following command rather than simply running "minikube start".

```text
~/KubeDig/contribution/minikube$ ./start_minikube.sh
```

It will use the minikube image with Linux kernel 5.4.40 with AppArmor service enabled by default.

If you see no error, you're ready to test KubeDig.
