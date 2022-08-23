// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	consolev1 "github.com/openshift/api/console/v1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// ConsoleYAMLSampleSpecApplyConfiguration represents an declarative configuration of the ConsoleYAMLSampleSpec type for use
// with apply.
type ConsoleYAMLSampleSpecApplyConfiguration struct {
	TargetResource *v1.TypeMetaApplyConfiguration          `json:"targetResource,omitempty"`
	Title          *consolev1.ConsoleYAMLSampleTitle       `json:"title,omitempty"`
	Description    *consolev1.ConsoleYAMLSampleDescription `json:"description,omitempty"`
	YAML           *consolev1.ConsoleYAMLSampleYAML        `json:"yaml,omitempty"`
	Snippet        *bool                                   `json:"snippet,omitempty"`
}

// ConsoleYAMLSampleSpecApplyConfiguration constructs an declarative configuration of the ConsoleYAMLSampleSpec type for use with
// apply.
func ConsoleYAMLSampleSpec() *ConsoleYAMLSampleSpecApplyConfiguration {
	return &ConsoleYAMLSampleSpecApplyConfiguration{}
}

// WithTargetResource sets the TargetResource field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetResource field is set to the value of the last call.
func (b *ConsoleYAMLSampleSpecApplyConfiguration) WithTargetResource(value *v1.TypeMetaApplyConfiguration) *ConsoleYAMLSampleSpecApplyConfiguration {
	b.TargetResource = value
	return b
}

// WithTitle sets the Title field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Title field is set to the value of the last call.
func (b *ConsoleYAMLSampleSpecApplyConfiguration) WithTitle(value consolev1.ConsoleYAMLSampleTitle) *ConsoleYAMLSampleSpecApplyConfiguration {
	b.Title = &value
	return b
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *ConsoleYAMLSampleSpecApplyConfiguration) WithDescription(value consolev1.ConsoleYAMLSampleDescription) *ConsoleYAMLSampleSpecApplyConfiguration {
	b.Description = &value
	return b
}

// WithYAML sets the YAML field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the YAML field is set to the value of the last call.
func (b *ConsoleYAMLSampleSpecApplyConfiguration) WithYAML(value consolev1.ConsoleYAMLSampleYAML) *ConsoleYAMLSampleSpecApplyConfiguration {
	b.YAML = &value
	return b
}

// WithSnippet sets the Snippet field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Snippet field is set to the value of the last call.
func (b *ConsoleYAMLSampleSpecApplyConfiguration) WithSnippet(value bool) *ConsoleYAMLSampleSpecApplyConfiguration {
	b.Snippet = &value
	return b
}