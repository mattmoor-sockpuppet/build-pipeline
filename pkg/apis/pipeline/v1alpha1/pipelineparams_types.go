/*
Copyright 2018 The Knative Authors.

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
	"github.com/knative/pkg/webhook"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PipelineParamsSpec is the spec for a Pipeline resource
type PipelineParamsSpec struct {
	ServiceAccount string  `json:"serviceAccount"`
	Results        Results `json:"results"`
	// +optional
	Generation int64 `json:"generation,omitempty"`
}

// PipelineParamStatus does not contain anything because Params on their own
// do not have a status, they just hold data which is later used by PipelineRun.
type PipelineParamsStatus struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PipelineParams is the Schema for the pipelineparams API
// +k8s:openapi-gen=true
type PipelineParams struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec holds the desired state of the PipelineParams from the client
	// +optional
	Spec PipelineParamsSpec `json:"spec,omitempty"`
	// Status communicates the observed state of the PipelineParams from the controller
	// +optional
	Status PipelineParamsStatus `json:"status,omitempty"`
}

// Assert that PipelineParams implements the GenericCRD interface.
var _ webhook.GenericCRD = (*PipelineParams)(nil)

// Results tells a pipeline where to persist the results of runnign the pipeline.
type Results struct {
	// Runs is used to store the yaml/json of TaskRuns and PipelineRuns.
	Runs ResultTarget `json:"runs"`

	// Logs will store all logs output from running a task.
	Logs ResultTarget `json:"logs"`

	// Tests will store test results, if a task provides them.
	// +optional
	Tests *ResultTarget `json:"tests,omitempty"`
}

// AllResultTargetTypes is a list of all ResultTargetTypes, used for validation
var AllResultTargetTypes = []ResultTargetType{ResultTargetTypeGCS}

// ResultTargetType represents the type of endpoint that this result target is,
// so that the controller will know how to write results to it.
type ResultTargetType string

const (
	// ResultTargetTypeGCS indicates that the URL endpoint is a GCS bucket.
	ResultTargetTypeGCS = "gcs"
)

// ResultTarget is used to identify an endpoint where results can be uploaded. The
// serviceaccount used for the pipeline must have access to this endpoint.
type ResultTarget struct {
	Name string           `json:"name"`
	Type ResultTargetType `json:"type"`
	URL  string           `json:"url"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PipelineParamsList contains a list of PipelineParams
type PipelineParamsList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PipelineParams `json:"items"`
}
