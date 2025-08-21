// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 Authors of KubeDig

package state

import (
	"encoding/json"
	"time"

	kg "github.com/zfz-725/KubeDig/KubeDig/log"
	"github.com/zfz-725/KubeDig/KubeDig/types"
	tp "github.com/zfz-725/KubeDig/KubeDig/types"
	pb "github.com/zfz-725/KubeDig/protobuf"
)

// PushContainerEvent function pushes (container + pod) event
func (sa *StateAgent) PushContainerEvent(container tp.Container, event string) {
	if container.ContainerID == "" {
		kg.Debug("Error while pushing container event. Missing data.")
		return
	}

	// create ns first
	namespace := container.NamespaceName
	sa.KubeDigNamespacesLock.Lock()
	if event == EventAdded {
		// create this kubedig ns if it doesn't exist
		// currently only "container_namespace" until we have config agent
		if ns, ok := sa.KubeDigNamespaces[namespace]; !ok {
			nsObj := types.Namespace{
				Name: namespace,
				// no way to change ns annotations in non-kubernetes mode right now
				// thus use defaults
				KubedigFilePosture:    "audit",
				KubedigNetworkPosture: "audit",
				LastUpdatedAt:           container.LastUpdatedAt,

				ContainerCount: 1,
			}
			sa.KubeDigNamespaces[namespace] = nsObj

			sa.PushNamespaceEvent(nsObj, EventAdded)
		} else {
			// update the container count
			ns.ContainerCount++
			sa.KubeDigNamespaces[namespace] = ns
		}

	} else if event == EventDeleted {
		if ns, ok := sa.KubeDigNamespaces[namespace]; ok {
			ns.ContainerCount--
			sa.KubeDigNamespaces[namespace] = ns

			// delete ns if no containers left in it
			if ns.ContainerCount == 0 {
				ns.LastUpdatedAt = time.Now().UTC().String()
				sa.PushNamespaceEvent(ns, EventDeleted)
				delete(sa.KubeDigNamespaces, namespace)
			}
		}

	}
	sa.KubeDigNamespacesLock.Unlock()

	containerBytes, err := json.Marshal(container)
	if err != nil {
		kg.Warnf("Error while trying to marshal container data: %s", err.Error())
		return
	}

	containerEvent := &pb.StateEvent{
		Kind:   KindContainer,
		Type:   event,
		Name:   container.ContainerName,
		Object: containerBytes,
	}

	// skip sending message as no state receiver is connected
	if sa.StateEventChans == nil {
		return
	}

	for uid, conn := range sa.StateEventChans {
		select {
		case conn <- containerEvent:
		default:
			kg.Debugf("Failed to send container %s event to connection: %s", event, uid)
			return
		}
	}

	return
}

// PushNodeEvent function pushes node event
func (sa *StateAgent) PushNodeEvent(node tp.Node, event string) {
	if node.NodeName == "" {
		kg.Warn("Received empty node event")
		return
	}

	nodeData, err := json.Marshal(node)
	if err != nil {
		kg.Warnf("Error while trying to marshal node data: %s", err.Error())
		return
	}

	nodeEvent := &pb.StateEvent{
		Kind:   KindNode,
		Type:   event,
		Name:   node.NodeName,
		Object: nodeData,
	}

	// skip sending message as no state receiver is connected
	if sa.StateEventChans == nil {
		return
	}

	for uid, conn := range sa.StateEventChans {
		select {
		case conn <- nodeEvent:
		default:
			kg.Debugf("Failed to send node %s event to connection: %s", event, uid)
			return
		}
	}

	return
}

// PushNamespaceEvent function pushes namespace event
func (sa *StateAgent) PushNamespaceEvent(namespace tp.Namespace, event string) {
	nsBytes, err := json.Marshal(namespace)
	if err != nil {
		kg.Warnf("Failed to marshal ns event: %s", err.Error())
		return
	}

	nsEvent := &pb.StateEvent{
		Kind:   KindNamespace,
		Type:   event,
		Name:   namespace.Name,
		Object: nsBytes,
	}

	// skip sending message as no state receiver is connected
	if sa.StateEventChans == nil {
		return
	}

	for uid, conn := range sa.StateEventChans {
		select {
		case conn <- nsEvent:
		default:
			kg.Debugf("Failed to send namespace %s event to connection: %s", event, uid)
			return
		}
	}
}
