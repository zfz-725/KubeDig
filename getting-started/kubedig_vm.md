# KubeDig on VM/Bare-Metal

This recipe explains how to use KubeDig directly on a VM/Bare-Metal machine, and we tested the following steps on Ubuntu hosts.

The recipe installs `kubedig` as systemd process and `karmor` cli tool to manage policies and show alerts/telemetry.

## Download and Install KubeDig

1. Download the [latest release](https://github.com/zfz-725/KubeDig/releases) or KubeDig.
2. Install KubeDig (VER is the kubedig release version)
  ```
  sudo apt --no-install-recommends install ./kubedig_${VER}_linux-amd64.deb
  ```
  > Note that the above command doesn't installs the recommended packages, as we ship object files along with the package file. In case you don't have BTF, consider removing `--no-install-recommends` flag.
  
<details><summary>For distributions other than Ubuntu/Debian</summary>
<p>

1. Refer [Installing BCC](https://github.com/iovisor/bcc/blob/master/INSTALL.md#installing-bcc) to install pre-requisites.

2. Download release tarball from KubeDig releases for the version you want
  ```
  wget https://github.com/KubeDig/KubeDig/releases/download/v${VER}/kubedig_${VER}_linux-amd64.tar.gz
  ```

3. Unpack the tarball to the root directory:
  ```
  sudo tar --no-overwrite-dir -C / -xzf kubedig_${VER}_linux-amd64.tar.gz
  sudo systemctl daemon-reload
  ```
</p>
</details>

## Start KubeDig

```
sudo systemctl start kubedig
```

Check the status of KubeDig using `sudo systemctl status kubedig` or use `sudo journalctl -u kubedig -f` to continuously monitor kubedig logs.

## Apply sample policy

Following policy is to deny execution of `sleep` binary on the host:

```yaml=
apiVersion: security.kubedig.com/v1
kind: KubeDigHostPolicy
metadata:
  name: hsp-kubedig-dev-proc-path-block
spec:
  nodeSelector:
    matchLabels:
      kubedig.io/hostname: "*" # Apply to all hosts
  process:
    matchPaths:
    - path: /usr/bin/sleep # try sleep 1
  action:
    Block
```

Save the above policy to _`hostpolicy.yaml`_ and apply:
```
karmor vm policy add hostpolicy.yaml
```

**Now if you run `sleep` command, the process would be denied execution.**

> Note that `sleep` may not be blocked if you run it in the same terminal where you apply the above policy. In that case, please open a new terminal and run `sleep` again to see if the command is blocked.

## Get Alerts for policies and telemetry

```
karmor logs --gRPC=:32767 --json
```

```json
{
"Timestamp":1717259989,
"UpdatedTime":"2024-06-01T16:39:49.360067Z",
"HostName":"kubedig-dev",
"HostPPID":1582,
"HostPID":2420,
"PPID":1582,
"PID":2420,
"UID":1000,
"ParentProcessName":"/usr/bin/bash",
"ProcessName":"/usr/bin/sleep",
"PolicyName":"hsp-kubedig-dev-proc-path-block",
"Severity":"1",
"Type":"MatchedHostPolicy",
"Source":"/usr/bin/bash",
"Operation":"Process",
"Resource":"/usr/bin/sleep",
"Data":"lsm=SECURITY_BPRM_CHECK",
"Enforcer":"BPFLSM",
"Action":"Block",
"Result":"Permission denied",
"Cwd":"/"
}
```
