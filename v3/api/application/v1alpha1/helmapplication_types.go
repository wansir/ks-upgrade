/*
Copyright 2020 The KubeSphere Authors.

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

package v1alpha1

import (
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"kubesphere.io/ks-upgrade/v3/api/constants"
)

const (
	ResourceKindHelmApplication     = "HelmApplication"
	ResourceSingularHelmApplication = "helmapplication"
	ResourcePluralHelmApplication   = "helmapplications"
)

// HelmApplicationSpec defines the desired state of HelmApplication
type HelmApplicationSpec struct {
	// the name of the helm application
	Name string `json:"name"`
	// description from chart's description or frontend
	Description string `json:"description,omitempty"`
	// attachments id
	Attachments []string `json:"attachments,omitempty"`
	// info from frontend
	Abstraction string `json:"abstraction,omitempty"`
	AppHome     string `json:"appHome,omitempty"`
	// The attachment id of the icon
	Icon string `json:"icon,omitempty"`
}

// HelmApplicationStatus defines the observed state of HelmApplication
type HelmApplicationStatus struct {
	// If this application belong to appStore, latestVersion is the the latest version of the active application version.
	// otherwise latestVersion is the latest version of all application version
	LatestVersion string `json:"latestVersion,omitempty"`
	// the state of the helm application: draft, submitted, passed, rejected, suspended, active
	State      string       `json:"state,omitempty"`
	UpdateTime *metav1.Time `json:"updateTime"`
	StatusTime *metav1.Time `json:"statusTime"`
}

// +kubebuilder:resource:scope=Cluster,shortName=happ
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="application name",type=string,JSONPath=`.spec.name`
// +kubebuilder:printcolumn:name="workspace",type="string",JSONPath=".metadata.labels.kubesphere\\.io/workspace"
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +genclient
// +genclient:nonNamespaced
// +kubebuilder:object:root=true

// HelmApplication is the Schema for the helmapplications API
type HelmApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HelmApplicationSpec   `json:"spec,omitempty"`
	Status HelmApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HelmApplicationList contains a list of HelmApplication
type HelmApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HelmApplication `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HelmApplication{}, &HelmApplicationList{})
}

func (in *HelmApplication) GetTrueName() string {
	return in.Spec.Name
}

func (in *HelmApplication) GetHelmRepoId() string {
	return getValue(in.Labels, constants.ChartRepoIdLabelKey)
}

func (in *HelmApplication) GetApplicationId() string {
	return strings.TrimSuffix(in.Name, HelmApplicationAppStoreSuffix)
}

func (in *HelmApplication) GetCategoryId() string {
	return getValue(in.Labels, constants.CategoryIdLabelKey)
}

func (in *HelmApplication) GetWorkspace() string {
	ws := getValue(in.Labels, constants.WorkspaceLabelKey)
	if ws == "" {
		return getValue(in.Labels, OriginWorkspaceLabelKey)
	}
	return ws
}

func getValue(m map[string]string, key string) string {
	if m == nil {
		return ""
	}
	return m[key]
}

func (in *HelmApplication) GetState() string {
	if in.Status.State == "" {
		return StateDraft
	}
	return in.Status.State
}

func (in *HelmApplication) GetCreator() string {
	return getValue(in.Annotations, constants.CreatorAnnotationKey)
}

func (in *HelmApplication) GetAppName() string {
	return in.Name
}

func (in *HelmApplication) GetCreationTime() time.Time {
	return in.CreationTimestamp.Time
}

func (in *HelmApplication) GetLatestVersion() string {
	return in.Status.LatestVersion
}

func (in *HelmApplication) GetUpdateTime() *metav1.Time {
	return in.Status.UpdateTime
}

func (in *HelmApplication) GetStatusTime() *metav1.Time {
	return in.Status.StatusTime
}

func (in *HelmApplication) GetAbstraction() string {
	return in.Spec.Abstraction
}

func (in *HelmApplication) GetDescription() string {
	return in.Spec.Description
}

func (in *HelmApplication) GetAttachments() []string {
	return in.Spec.Attachments
}

func (in *HelmApplication) GetAppHome() string {
	return in.Spec.AppHome
}

func (in *HelmApplication) GetIcon() string {
	return in.Spec.Icon
}

func (in *HelmApplication) GetAnnotations() map[string]string {
	return in.Annotations
}
