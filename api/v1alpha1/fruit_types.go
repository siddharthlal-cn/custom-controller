/*


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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FruitSpec defines the desired state of Fruit
type FruitSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Name is the name of fruit
	Name string `json:"name,omitempty"`

	// Type is the species of fruit
	Type string `json:"type,omitempty"`

	// Amount is the number of fruit
	Amount int32 `json:"amount,omitempty"`
}

// FruitStatus defines the observed state of Fruit
type FruitStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Created bool `json:"created,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Fruit is the Schema for the fruits API
type Fruit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FruitSpec   `json:"spec,omitempty"`
	Status FruitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FruitList contains a list of Fruit
type FruitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Fruit `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Fruit{}, &FruitList{})
}
