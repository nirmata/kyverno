/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	types "k8s.io/apimachinery/pkg/types"
)

// GenerationApplyConfiguration represents an declarative configuration of the Generation type for use
// with apply.
type GenerationApplyConfiguration struct {
	GenerateExisting                *bool `json:"generateExisting,omitempty"`
	*ResourceSpecApplyConfiguration `json:"ResourceSpec,omitempty"`
	Synchronize                     *bool                        `json:"synchronize,omitempty"`
	OrphanDownstreamOnPolicyDelete  *bool                        `json:"orphanDownstreamOnPolicyDelete,omitempty"`
	RawData                         *apiextensionsv1.JSON        `json:"data,omitempty"`
	Clone                           *CloneFromApplyConfiguration `json:"clone,omitempty"`
	CloneList                       *CloneListApplyConfiguration `json:"cloneList,omitempty"`
}

// GenerationApplyConfiguration constructs an declarative configuration of the Generation type for use with
// apply.
func Generation() *GenerationApplyConfiguration {
	return &GenerationApplyConfiguration{}
}

// WithGenerateExisting sets the GenerateExisting field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateExisting field is set to the value of the last call.
func (b *GenerationApplyConfiguration) WithGenerateExisting(value bool) *GenerationApplyConfiguration {
	b.GenerateExisting = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *GenerationApplyConfiguration) WithAPIVersion(value string) *GenerationApplyConfiguration {
	b.ensureResourceSpecApplyConfigurationExists()
	b.APIVersion = &value
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *GenerationApplyConfiguration) WithKind(value string) *GenerationApplyConfiguration {
	b.ensureResourceSpecApplyConfigurationExists()
	b.Kind = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *GenerationApplyConfiguration) WithNamespace(value string) *GenerationApplyConfiguration {
	b.ensureResourceSpecApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *GenerationApplyConfiguration) WithName(value string) *GenerationApplyConfiguration {
	b.ensureResourceSpecApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *GenerationApplyConfiguration) WithUID(value types.UID) *GenerationApplyConfiguration {
	b.ensureResourceSpecApplyConfigurationExists()
	b.UID = &value
	return b
}

func (b *GenerationApplyConfiguration) ensureResourceSpecApplyConfigurationExists() {
	if b.ResourceSpecApplyConfiguration == nil {
		b.ResourceSpecApplyConfiguration = &ResourceSpecApplyConfiguration{}
	}
}

// WithSynchronize sets the Synchronize field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Synchronize field is set to the value of the last call.
func (b *GenerationApplyConfiguration) WithSynchronize(value bool) *GenerationApplyConfiguration {
	b.Synchronize = &value
	return b
}

// WithOrphanDownstreamOnPolicyDelete sets the OrphanDownstreamOnPolicyDelete field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OrphanDownstreamOnPolicyDelete field is set to the value of the last call.
func (b *GenerationApplyConfiguration) WithOrphanDownstreamOnPolicyDelete(value bool) *GenerationApplyConfiguration {
	b.OrphanDownstreamOnPolicyDelete = &value
	return b
}

// WithRawData sets the RawData field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RawData field is set to the value of the last call.
func (b *GenerationApplyConfiguration) WithRawData(value apiextensionsv1.JSON) *GenerationApplyConfiguration {
	b.RawData = &value
	return b
}

// WithClone sets the Clone field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Clone field is set to the value of the last call.
func (b *GenerationApplyConfiguration) WithClone(value *CloneFromApplyConfiguration) *GenerationApplyConfiguration {
	b.Clone = value
	return b
}

// WithCloneList sets the CloneList field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CloneList field is set to the value of the last call.
func (b *GenerationApplyConfiguration) WithCloneList(value *CloneListApplyConfiguration) *GenerationApplyConfiguration {
	b.CloneList = value
	return b
}
