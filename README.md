![](.gitbook/assets/logo.png)

[![Build Status](https://github.com/zfz-725/KubeDig/actions/workflows/ci-go.yml/badge.svg)](https://github.com/zfz-725/KubeDig/actions/workflows/ci-go.yml/)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/5401/badge)](https://bestpractices.coreinfrastructure.org/projects/5401)
[![CLOMonitor](https://img.shields.io/endpoint?url=https://clomonitor.io/api/projects/cncf/kubedig/badge)](https://clomonitor.io/projects/cncf/kubedig)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/zfz-725/kubedig/badge)](https://securityscorecards.dev/viewer/?uri=github.com/zfz-725/KubeDig)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fkubedig%2FKubeDig.svg?type=shield&issueType=license)](https://app.fossa.com/projects/git%2Bgithub.com%2Fkubedig%2FKubeDig?ref=badge_shield)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fkubedig%2FKubeDig.svg?type=shield&issueType=security)](https://app.fossa.com/projects/git%2Bgithub.com%2Fkubedig%2FKubeDig?ref=badge_shield)
[![Slack](https://img.shields.io/badge/Join%20Our%20Community-Slack-blue)](https://cloud-native.slack.com/archives/C02R319HVL3)
[![Discussions](https://img.shields.io/badge/Got%20Questions%3F-Chat-Violet)](https://github.com/zfz-725/KubeDig/discussions)
[![Docker Downloads](https://img.shields.io/docker/pulls/kubedig/kubedig)](https://hub.docker.com/r/kubedig/kubedig)
[![ArtifactHub](https://img.shields.io/badge/ArtifactHub-KubeDig-blue?logo=artifacthub&labelColor=grey&color=green)](https://artifacthub.io/packages/search?kind=19)

KubeDig is a cloud-native runtime security enforcement system that restricts the behavior \(such as process execution, file access, and networking operations\) of pods, containers, and nodes (VMs) at the system level.

KubeDig leverages [Linux security modules \(LSMs\)](https://en.wikipedia.org/wiki/Linux_Security_Modules) such as [AppArmor](https://en.wikipedia.org/wiki/AppArmor), [SELinux](https://en.wikipedia.org/wiki/Security-Enhanced_Linux), or [BPF-LSM](https://docs.kernel.org/bpf/prog_lsm.html) to enforce the user-specified policies. KubeDig generates rich alerts/telemetry events with container/pod/namespace identities by leveraging eBPF.

|  |   |
|:---|:---|
| :muscle: **[Harden Infrastructure](getting-started/hardening_guide.md)** <hr>:chains: Protect critical paths such as cert bundles <br>:clipboard: MITRE, STIGs, CIS based rules <br>:left_luggage: Restrict access to raw DB table | :ring: **[Least Permissive Access](getting-started/least_permissive_access.md)** <hr>:traffic_light: Process Whitelisting <br>:traffic_light: Network Whitelisting <br>:control_knobs: Control access to sensitive assets |
| :telescope: **[Application Behavior](getting-started/workload_visibility.md)** <hr>:dna: Process execs, File System accesses <br>:compass: Service binds, Ingress, Egress connections <br>:microscope: Sensitive system call profiling | :snowflake: **[Deployment Models](getting-started/deployment_models.md)** <hr>:wheel_of_dharma: Kubernetes Deployment<br>:whale2: Containerized Deployment<br>:computer: VM/Bare-Metal Deployment |

## Architecture Overview

![KubeDig High Level Design](.gitbook/assets/kubedig_overview.png)

## Documentation :notebook:

* :point_right: [Getting Started](getting-started/deployment_guide.md)
* :dart: [Use Cases](getting-started/use-cases/hardening.md)
* :heavy_check_mark: [KubeDig Support Matrix](getting-started/support_matrix.md)
* :chess_pawn: [How is KubeDig different?](getting-started/differentiation.md)
* :scroll: Security Policy for Pods/Containers [[Spec](getting-started/security_policy_specification.md)] [[Examples](getting-started/security_policy_examples.md)]
* :scroll: Cluster level security Policy for Pods/Containers [[Spec](getting-started/cluster_security_policy_specification.md)] [[Examples](getting-started/cluster_security_policy_examples.md)]
* :scroll: Security Policy for Hosts/Nodes [[Spec](getting-started/host_security_policy_specification.md)] [[Examples](getting-started/host_security_policy_examples.md)]<br>
... [detailed documentation](https://docs.kubedig.io/kubedig/)

### Contributors :busts_in_silhouette:

* :blue_book: [Contribution Guide](contribution/contribution_guide.md)
* :technologist: [Development Guide](contribution/development_guide.md), [Testing Guide](contribution/testing_guide.md)
* :raised_hand: [Join KubeDig Slack](https://cloud-native.slack.com/archives/C02R319HVL3)
* :question: [FAQs](getting-started/FAQ.md)

### Biweekly Meeting

- :speaking_head: [Zoom Link](http://zoom.kubedig.io)
- :page_facing_up: Minutes: [Document](https://docs.google.com/document/d/1IqIIG9Vz-PYpbUwrH0u99KYEM1mtnYe6BHrson4NqEs/edit)
- :calendar: Calendar invite: [Google Calendar](http://www.google.com/calendar/event?action=TEMPLATE&dates=20220210T150000Z%2F20220210T153000Z&text=KubeDig%20Community%20Call&location=&details=%3Ca%20href%3D%22https%3A%2F%2Fdocs.google.com%2Fdocument%2Fd%2F1IqIIG9Vz-PYpbUwrH0u99KYEM1mtnYe6BHrson4NqEs%2Fedit%22%3EMinutes%20of%20Meeting%3C%2Fa%3E%0A%0A%3Ca%20href%3D%22%20http%3A%2F%2Fzoom.kubedig.io%22%3EZoom%20Link%3C%2Fa%3E&recur=RRULE:FREQ=WEEKLY;INTERVAL=2;BYDAY=TH&ctz=Asia/Calcutta), [ICS file](getting-started/resources/KubeDigMeetup.ics)

## Notice/Credits :handshake:

- KubeDig uses [Tracee](https://github.com/aquasecurity/tracee/)'s system call utility functions.

## CNCF

KubeDig is [Sandbox Project](https://www.cncf.io/projects/kubedig/) of the Cloud Native Computing Foundation.
![CNCF SandBox Project](.gitbook/assets/cncf-sandbox.png)

## ROADMAP

KubeDig roadmap is tracked via [KubeDig Projects](https://github.com/orgs/kubedig/projects?query=is%3Aopen)
