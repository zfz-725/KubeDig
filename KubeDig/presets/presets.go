// SPDX-License-Identifier: Apache-2.0
// Copyright 2021 Authors of KubeDig

// Package presets contains preset rules components
package presets

import (
	"errors"

	fd "github.com/zfz-725/KubeDig/KubeDig/feeder"
	mon "github.com/zfz-725/KubeDig/KubeDig/monitor"
	anonmap "github.com/zfz-725/KubeDig/KubeDig/presets/anonmapexec"
	"github.com/zfz-725/KubeDig/KubeDig/presets/base"
	filelessexec "github.com/zfz-725/KubeDig/KubeDig/presets/filelessexec"
	tp "github.com/zfz-725/KubeDig/KubeDig/types"
)

// Preset struct
type Preset struct {
	base.Preset

	List map[string]base.PresetInterface
}

// NewPreset returns an instance of Preset
func NewPreset(logger *fd.Feeder, monitor *mon.SystemMonitor) *Preset {
	p := &Preset{}

	p.List = make(map[string]base.PresetInterface)
	p.Logger = logger
	p.Monitor = monitor

	// add all presets
	p.List[anonmap.NAME] = anonmap.NewAnonMapExecPreset()
	p.List[filelessexec.NAME] = filelessexec.NewFilelessExecPreset()

	// register all presets
	p.RegisterPresets()

	if len(p.List) > 0 {
		return p
	}
	return nil
}

// RegisterPresets initiates and adds presets to map
func (p *Preset) RegisterPresets() {
	for k, v := range p.List {
		_, err := v.RegisterPreset(p.Logger, p.Monitor)
		if err != nil {
			delete(p.List, k)
		}
	}
}

// RegisterContainer registers container identifiers
func (p *Preset) RegisterContainer(containerID string, pidns, mntns uint32) {
	for _, v := range p.List {
		v.RegisterContainer(containerID, pidns, mntns)
	}
}

// UnregisterContainer removes container identifiers
func (p *Preset) UnregisterContainer(containerID string) {
	for _, v := range p.List {
		v.UnregisterContainer(containerID)
	}
}

// UpdateSecurityPolicies Function
func (p *Preset) UpdateSecurityPolicies(endPoint tp.EndPoint) {
	for _, v := range p.List {
		v.UpdateSecurityPolicies(endPoint)
	}
}

// Destroy Function
func (p *Preset) Destroy() error {
	var destroyErr error
	for _, v := range p.List {
		err := v.Destroy()
		if err != nil {
			destroyErr = errors.Join(destroyErr, err)
		}
	}
	return destroyErr
}
