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

// EndpointSpec defines the desired state of Endpoint.
//
// A hosted endpoint for real-time inference.
type EndpointSpec struct {
	DeploymentConfig *DeploymentConfig `json:"deploymentConfig,omitempty"`
	// The name of an endpoint configuration. For more information, see CreateEndpointConfig
	// (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpointConfig.html).
	// +kubebuilder:validation:Required
	EndpointConfigName *string `json:"endpointConfigName"`
	// The name of the endpoint.The name must be unique within an Amazon Web Services
	// Region in your Amazon Web Services account. The name is case-insensitive
	// in CreateEndpoint, but the case is preserved and must be matched in InvokeEndpoint
	// (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_runtime_InvokeEndpoint.html).
	// +kubebuilder:validation:Required
	EndpointName *string `json:"endpointName"`
	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, for example, by purpose, owner, or
	// environment. For more information, see Tagging Amazon Web Services Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).
	Tags []*Tag `json:"tags,omitempty"`
}

// EndpointStatus defines the observed state of Endpoint
type EndpointStatus struct {
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
	// A timestamp that shows when the endpoint was created.
	// +kubebuilder:validation:Optional
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
	// The status of the endpoint.
	//
	//    * OutOfService: Endpoint is not available to take incoming requests.
	//
	//    * Creating: CreateEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpoint.html)
	//    is executing.
	//
	//    * Updating: UpdateEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_UpdateEndpoint.html)
	//    or UpdateEndpointWeightsAndCapacities (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_UpdateEndpointWeightsAndCapacities.html)
	//    is executing.
	//
	//    * SystemUpdating: Endpoint is undergoing maintenance and cannot be updated
	//    or deleted or re-scaled until it has completed. This maintenance operation
	//    does not change any customer-specified values such as VPC config, KMS
	//    encryption, model, instance type, or instance count.
	//
	//    * RollingBack: Endpoint fails to scale up or down or change its variant
	//    weight and is in the process of rolling back to its previous configuration.
	//    Once the rollback completes, endpoint returns to an InService status.
	//    This transitional status only applies to an endpoint that has autoscaling
	//    enabled and is undergoing variant weight or capacity changes as part of
	//    an UpdateEndpointWeightsAndCapacities (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_UpdateEndpointWeightsAndCapacities.html)
	//    call or when the UpdateEndpointWeightsAndCapacities (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_UpdateEndpointWeightsAndCapacities.html)
	//    operation is called explicitly.
	//
	//    * InService: Endpoint is available to process incoming requests.
	//
	//    * Deleting: DeleteEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DeleteEndpoint.html)
	//    is executing.
	//
	//    * Failed: Endpoint could not be created, updated, or re-scaled. Use the
	//    FailureReason value returned by DescribeEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeEndpoint.html)
	//    for information about the failure. DeleteEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DeleteEndpoint.html)
	//    is the only operation that can be performed on a failed endpoint.
	//
	//    * UpdateRollbackFailed: Both the rolling deployment and auto-rollback
	//    failed. Your endpoint is in service with a mix of the old and new endpoint
	//    configurations. For information about how to remedy this issue and restore
	//    the endpoint's status to InService, see Rolling Deployments (https://docs.aws.amazon.com/sagemaker/latest/dg/deployment-guardrails-rolling.html).
	// +kubebuilder:validation:Optional
	EndpointStatus *string `json:"endpointStatus,omitempty"`
	// If the status of the endpoint is Failed, the reason why it failed.
	// +kubebuilder:validation:Optional
	FailureReason *string `json:"failureReason,omitempty"`
	// A timestamp that shows when the endpoint was last modified.
	// +kubebuilder:validation:Optional
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
	// Returns the summary of an in-progress deployment. This field is only returned
	// when the endpoint is creating or updating with a new endpoint configuration.
	// +kubebuilder:validation:Optional
	PendingDeploymentSummary *PendingDeploymentSummary `json:"pendingDeploymentSummary,omitempty"`
	// An array of ProductionVariantSummary (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ProductionVariantSummary.html)
	// objects, one for each model hosted behind this endpoint.
	// +kubebuilder:validation:Optional
	ProductionVariants []*ProductionVariantSummary `json:"productionVariants,omitempty"`
}

// Endpoint is the Schema for the Endpoints API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="FAILURE-REASON",type=string,priority=1,JSONPath=`.status.failureReason`
// +kubebuilder:printcolumn:name="STATUS",type=string,priority=0,JSONPath=`.status.endpointStatus`
type Endpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointSpec   `json:"spec,omitempty"`
	Status            EndpointStatus `json:"status,omitempty"`
}

// EndpointList contains a list of Endpoint
// +kubebuilder:object:root=true
type EndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Endpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Endpoint{}, &EndpointList{})
}
