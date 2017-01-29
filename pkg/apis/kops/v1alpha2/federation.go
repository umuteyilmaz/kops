/*
Copyright 2016 The Kubernetes Authors.

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
	"k8s.io/kubernetes/pkg/api/v1"
	meta_v1 "k8s.io/kubernetes/pkg/apis/meta/v1"
)

// Federation represents a federated set of kubernetes clusters
type Federation struct {
	meta_v1.TypeMeta `json:",inline"`
	ObjectMeta       v1.ObjectMeta `json:"metadata,omitempty"`

	Spec FederationSpec `json:"spec,omitempty"`
}

type FederationSpec struct {
	Controllers []string `json:"controllers,omitempty"`
	Members     []string `json:"members,omitempty"`

	DNSName string `json:"dnsName,omitempty"`
}

type FederationList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`

	Items []Federation `json:"items"`
}

func (f *Federation) Validate() error {
	return nil
}