/*
Copyright 2020 KubeSphere Authors

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	workspacev1alpha1 "kubesphere.io/ks-upgrade/v3/api/tenant/v1alpha1"
)

const (
	ResourcePluralFederatedWorkspace   = "federatedworkspaces"
	ResourceSingularFederatedWorkspace = "federatedworkspace"
	FederatedWorkspaceKind             = "FederatedWorkspace"
)

// +kubebuilder:object:root=true
// +k8s:openapi-gen=true
type FederatedWorkspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FederatedWorkspaceSpec `json:"spec"`

	Status *GenericFederatedStatus `json:"status,omitempty"`
}

type FederatedWorkspaceSpec struct {
	Template  WorkspaceTemplate      `json:"template"`
	Placement GenericPlacementFields `json:"placement"`
	Overrides []GenericOverrideItem  `json:"overrides,omitempty"`
}

type WorkspaceTemplate struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              workspacev1alpha1.WorkspaceSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// FederatedWorkspaceList contains a list of FederatedWorkspace
type FederatedWorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedWorkspace `json:"items"`
}
