// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PipelineExecutionSpec defines the desired state of PipelineExecution.
//
// An execution of a pipeline.
type PipelineExecutionSpec struct {

	// This configuration, if specified, overrides the parallelism configuration
	// of the parent pipeline for this specific run.

	ParallelismConfiguration *ParallelismConfiguration `json:"parallelismConfiguration,omitempty"`
	// The description of the pipeline execution.

	PipelineExecutionDescription *string `json:"pipelineExecutionDescription,omitempty"`
	// The display name of the pipeline execution.

	PipelineExecutionDisplayName *string `json:"pipelineExecutionDisplayName,omitempty"`
	// The name or Amazon Resource Name (ARN) of the pipeline.

	// +kubebuilder:validation:Required

	PipelineName *string `json:"pipelineName"`
	// Contains a list of pipeline parameters. This list can be empty.

	PipelineParameters []*Parameter `json:"pipelineParameters,omitempty"`
	// The selective execution configuration applied to the pipeline run.

	SelectiveExecutionConfig *SelectiveExecutionConfig `json:"selectiveExecutionConfig,omitempty"`
}

// PipelineExecutionStatus defines the observed state of PipelineExecution
type PipelineExecutionStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRs managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The time when the pipeline execution was created.
	// +kubebuilder:validation:Optional
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
	// If the execution failed, a message describing why.
	// +kubebuilder:validation:Optional
	FailureReason *string `json:"failureReason,omitempty"`
	// The time when the pipeline execution was modified last.
	// +kubebuilder:validation:Optional
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
	// The status of the pipeline execution.
	// +kubebuilder:validation:Optional
	PipelineExecutionStatus *string `json:"pipelineExecutionStatus,omitempty"`
}

// PipelineExecution is the Schema for the PipelineExecutions API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="FAILURE-REASON",type=string,priority=1,JSONPath=`.status.failureReason`
// +kubebuilder:printcolumn:name="STATUS",type=string,priority=0,JSONPath=`.status.pipelineExecutionStatus`
type PipelineExecution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PipelineExecutionSpec   `json:"spec,omitempty"`
	Status            PipelineExecutionStatus `json:"status,omitempty"`
}

// PipelineExecutionList contains a list of PipelineExecution
// +kubebuilder:object:root=true
type PipelineExecutionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PipelineExecution `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PipelineExecution{}, &PipelineExecutionList{})
}
