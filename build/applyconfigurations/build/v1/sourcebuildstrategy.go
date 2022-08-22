// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// SourceBuildStrategyApplyConfiguration represents an declarative configuration of the SourceBuildStrategy type for use
// with apply.
type SourceBuildStrategyApplyConfiguration struct {
	From        *v1.ObjectReference             `json:"from,omitempty"`
	PullSecret  *v1.LocalObjectReference        `json:"pullSecret,omitempty"`
	Env         []v1.EnvVar                     `json:"env,omitempty"`
	Scripts     *string                         `json:"scripts,omitempty"`
	Incremental *bool                           `json:"incremental,omitempty"`
	ForcePull   *bool                           `json:"forcePull,omitempty"`
	Volumes     []BuildVolumeApplyConfiguration `json:"volumes,omitempty"`
}

// SourceBuildStrategyApplyConfiguration constructs an declarative configuration of the SourceBuildStrategy type for use with
// apply.
func SourceBuildStrategy() *SourceBuildStrategyApplyConfiguration {
	return &SourceBuildStrategyApplyConfiguration{}
}

// WithFrom sets the From field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the From field is set to the value of the last call.
func (b *SourceBuildStrategyApplyConfiguration) WithFrom(value v1.ObjectReference) *SourceBuildStrategyApplyConfiguration {
	b.From = &value
	return b
}

// WithPullSecret sets the PullSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PullSecret field is set to the value of the last call.
func (b *SourceBuildStrategyApplyConfiguration) WithPullSecret(value v1.LocalObjectReference) *SourceBuildStrategyApplyConfiguration {
	b.PullSecret = &value
	return b
}

// WithEnv adds the given value to the Env field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Env field.
func (b *SourceBuildStrategyApplyConfiguration) WithEnv(values ...v1.EnvVar) *SourceBuildStrategyApplyConfiguration {
	for i := range values {
		b.Env = append(b.Env, values[i])
	}
	return b
}

// WithScripts sets the Scripts field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Scripts field is set to the value of the last call.
func (b *SourceBuildStrategyApplyConfiguration) WithScripts(value string) *SourceBuildStrategyApplyConfiguration {
	b.Scripts = &value
	return b
}

// WithIncremental sets the Incremental field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Incremental field is set to the value of the last call.
func (b *SourceBuildStrategyApplyConfiguration) WithIncremental(value bool) *SourceBuildStrategyApplyConfiguration {
	b.Incremental = &value
	return b
}

// WithForcePull sets the ForcePull field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ForcePull field is set to the value of the last call.
func (b *SourceBuildStrategyApplyConfiguration) WithForcePull(value bool) *SourceBuildStrategyApplyConfiguration {
	b.ForcePull = &value
	return b
}

// WithVolumes adds the given value to the Volumes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Volumes field.
func (b *SourceBuildStrategyApplyConfiguration) WithVolumes(values ...*BuildVolumeApplyConfiguration) *SourceBuildStrategyApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithVolumes")
		}
		b.Volumes = append(b.Volumes, *values[i])
	}
	return b
}
