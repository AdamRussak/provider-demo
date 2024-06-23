/*
Copyright 2022 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// JokeParameters are the configurable fields of a Joke.
type JokeParameters struct {
	Url    *string `json:"url,omitempty"`
	Path   *string `json:"path,omitempty"`
	Format *string `json:"format,omitempty"`
}

// JokeObservation are the observable fields of a Joke.
type JokeObservation struct {
	Joke string `json:"joke,omitempty"`
}

// A JokeSpec defines the desired state of a Joke.
type JokeSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       JokeParameters `json:"forProvider"`
}

// A JokeStatus represents the observed state of a Joke.
type JokeStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          JokeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Joke is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,demo}
type Joke struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JokeSpec   `json:"spec"`
	Status JokeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JokeList contains a list of Joke
type JokeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Joke `json:"items"`
}

// Joke type metadata.
var (
	JokeKind             = reflect.TypeOf(Joke{}).Name()
	JokeGroupKind        = schema.GroupKind{Group: Group, Kind: JokeKind}.String()
	JokeKindAPIVersion   = JokeKind + "." + SchemeGroupVersion.String()
	JokeGroupVersionKind = SchemeGroupVersion.WithKind(JokeKind)
)

func init() {
	SchemeBuilder.Register(&Joke{}, &JokeList{})
}
