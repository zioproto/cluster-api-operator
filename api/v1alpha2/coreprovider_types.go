/*
Copyright 2023 The Kubernetes Authors.

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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CoreProviderSpec defines the desired state of CoreProvider.
type CoreProviderSpec struct {
	ProviderSpec `json:",inline"`
}

// CoreProviderStatus defines the observed state of CoreProvider.
type CoreProviderStatus struct {
	ProviderStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=coreproviders,shortName=cacp,scope=Namespaced
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="InstalledVersion",type="string",JSONPath=".status.installedVersion"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:storageversion

// CoreProvider is the Schema for the coreproviders API.
type CoreProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CoreProviderSpec   `json:"spec,omitempty"`
	Status CoreProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CoreProviderList contains a list of CoreProvider.
type CoreProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CoreProvider `json:"items"`
}

func init() {
	objectTypes = append(objectTypes, &CoreProvider{}, &CoreProviderList{})
}
