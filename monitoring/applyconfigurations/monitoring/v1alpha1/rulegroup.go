// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// RuleGroupApplyConfiguration represents an declarative configuration of the RuleGroup type for use
// with apply.
type RuleGroupApplyConfiguration struct {
	Name     *string                  `json:"name,omitempty"`
	Interval *string                  `json:"interval,omitempty"`
	Rules    []RuleApplyConfiguration `json:"rules,omitempty"`
}

// RuleGroupApplyConfiguration constructs an declarative configuration of the RuleGroup type for use with
// apply.
func RuleGroup() *RuleGroupApplyConfiguration {
	return &RuleGroupApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *RuleGroupApplyConfiguration) WithName(value string) *RuleGroupApplyConfiguration {
	b.Name = &value
	return b
}

// WithInterval sets the Interval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Interval field is set to the value of the last call.
func (b *RuleGroupApplyConfiguration) WithInterval(value string) *RuleGroupApplyConfiguration {
	b.Interval = &value
	return b
}

// WithRules adds the given value to the Rules field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Rules field.
func (b *RuleGroupApplyConfiguration) WithRules(values ...*RuleApplyConfiguration) *RuleGroupApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRules")
		}
		b.Rules = append(b.Rules, *values[i])
	}
	return b
}
