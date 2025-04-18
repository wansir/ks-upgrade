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

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubesphere.io/ks-upgrade/v3/api/notification/v2beta2"
)

const (
	ResourcePluralFederatedNotificationSilence   = "federatednotificationsilences"
	ResourceSingularFederatedNotificationSilence = "federatednotificationsilence"
	FederatedNotificationSilenceKind             = "FederatedNotificationSilence"
)

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:object:root=true
// +k8s:openapi-gen=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:subresource:status

type FederatedNotificationSilence struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FederatedNotificationSilenceSpec `json:"spec"`

	Status *GenericFederatedStatus `json:"status,omitempty"`
}

type FederatedNotificationSilenceSpec struct {
	Template  NotificationSilenceTemplate `json:"template"`
	Placement GenericPlacementFields      `json:"placement"`
	Overrides []GenericOverrideItem       `json:"overrides,omitempty"`
}

type NotificationSilenceTemplate struct {
	// +kubebuilder:pruning:PreserveUnknownFields
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              v2beta2.SilenceSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:openapi-gen=true

// FederatedNotificationSilenceList contains a list of federatednotificationsilencelists
type FederatedNotificationSilenceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedNotificationSilence `json:"items"`
}
