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

package processing_job

import (
	"context"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/sagemaker"
	svcsdktypes "github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/aws/smithy-go"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/sagemaker-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &svcsdk.Client{}
	_ = &svcapitypes.ProcessingJob{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
	_ = fmt.Sprintf("")
	_ = &ackrequeue.NoRequeue{}
	_ = &aws.Config{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer func() {
		exit(err)
	}()
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.DescribeProcessingJobOutput
	resp, err = rm.sdkapi.DescribeProcessingJob(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeProcessingJob", err)
	if err != nil {
		var awsErr smithy.APIError
		if errors.As(err, &awsErr) && awsErr.ErrorCode() == "ValidationException" && strings.HasPrefix(awsErr.ErrorMessage(), "Could not find requested job") {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.AppSpecification != nil {
		f0 := &svcapitypes.AppSpecification{}
		if resp.AppSpecification.ContainerArguments != nil {
			f0.ContainerArguments = aws.StringSlice(resp.AppSpecification.ContainerArguments)
		}
		if resp.AppSpecification.ContainerEntrypoint != nil {
			f0.ContainerEntrypoint = aws.StringSlice(resp.AppSpecification.ContainerEntrypoint)
		}
		if resp.AppSpecification.ImageUri != nil {
			f0.ImageURI = resp.AppSpecification.ImageUri
		}
		ko.Spec.AppSpecification = f0
	} else {
		ko.Spec.AppSpecification = nil
	}
	if resp.Environment != nil {
		ko.Spec.Environment = aws.StringMap(resp.Environment)
	} else {
		ko.Spec.Environment = nil
	}
	if resp.ExperimentConfig != nil {
		f5 := &svcapitypes.ExperimentConfig{}
		if resp.ExperimentConfig.ExperimentName != nil {
			f5.ExperimentName = resp.ExperimentConfig.ExperimentName
		}
		if resp.ExperimentConfig.TrialComponentDisplayName != nil {
			f5.TrialComponentDisplayName = resp.ExperimentConfig.TrialComponentDisplayName
		}
		if resp.ExperimentConfig.TrialName != nil {
			f5.TrialName = resp.ExperimentConfig.TrialName
		}
		ko.Spec.ExperimentConfig = f5
	} else {
		ko.Spec.ExperimentConfig = nil
	}
	if resp.FailureReason != nil {
		ko.Status.FailureReason = resp.FailureReason
	} else {
		ko.Status.FailureReason = nil
	}
	if resp.NetworkConfig != nil {
		f9 := &svcapitypes.NetworkConfig{}
		if resp.NetworkConfig.EnableInterContainerTrafficEncryption != nil {
			f9.EnableInterContainerTrafficEncryption = resp.NetworkConfig.EnableInterContainerTrafficEncryption
		}
		if resp.NetworkConfig.EnableNetworkIsolation != nil {
			f9.EnableNetworkIsolation = resp.NetworkConfig.EnableNetworkIsolation
		}
		if resp.NetworkConfig.VpcConfig != nil {
			f9f2 := &svcapitypes.VPCConfig{}
			if resp.NetworkConfig.VpcConfig.SecurityGroupIds != nil {
				f9f2.SecurityGroupIDs = aws.StringSlice(resp.NetworkConfig.VpcConfig.SecurityGroupIds)
			}
			if resp.NetworkConfig.VpcConfig.Subnets != nil {
				f9f2.Subnets = aws.StringSlice(resp.NetworkConfig.VpcConfig.Subnets)
			}
			f9.VPCConfig = f9f2
		}
		ko.Spec.NetworkConfig = f9
	} else {
		ko.Spec.NetworkConfig = nil
	}
	if resp.ProcessingInputs != nil {
		f11 := []*svcapitypes.ProcessingInput{}
		for _, f11iter := range resp.ProcessingInputs {
			f11elem := &svcapitypes.ProcessingInput{}
			if f11iter.AppManaged != nil {
				f11elem.AppManaged = f11iter.AppManaged
			}
			if f11iter.DatasetDefinition != nil {
				f11elemf1 := &svcapitypes.DatasetDefinition{}
				if f11iter.DatasetDefinition.AthenaDatasetDefinition != nil {
					f11elemf1f0 := &svcapitypes.AthenaDatasetDefinition{}
					if f11iter.DatasetDefinition.AthenaDatasetDefinition.Catalog != nil {
						f11elemf1f0.Catalog = f11iter.DatasetDefinition.AthenaDatasetDefinition.Catalog
					}
					if f11iter.DatasetDefinition.AthenaDatasetDefinition.Database != nil {
						f11elemf1f0.Database = f11iter.DatasetDefinition.AthenaDatasetDefinition.Database
					}
					if f11iter.DatasetDefinition.AthenaDatasetDefinition.KmsKeyId != nil {
						f11elemf1f0.KMSKeyID = f11iter.DatasetDefinition.AthenaDatasetDefinition.KmsKeyId
					}
					if f11iter.DatasetDefinition.AthenaDatasetDefinition.OutputCompression != "" {
						f11elemf1f0.OutputCompression = aws.String(string(f11iter.DatasetDefinition.AthenaDatasetDefinition.OutputCompression))
					}
					if f11iter.DatasetDefinition.AthenaDatasetDefinition.OutputFormat != "" {
						f11elemf1f0.OutputFormat = aws.String(string(f11iter.DatasetDefinition.AthenaDatasetDefinition.OutputFormat))
					}
					if f11iter.DatasetDefinition.AthenaDatasetDefinition.OutputS3Uri != nil {
						f11elemf1f0.OutputS3URI = f11iter.DatasetDefinition.AthenaDatasetDefinition.OutputS3Uri
					}
					if f11iter.DatasetDefinition.AthenaDatasetDefinition.QueryString != nil {
						f11elemf1f0.QueryString = f11iter.DatasetDefinition.AthenaDatasetDefinition.QueryString
					}
					if f11iter.DatasetDefinition.AthenaDatasetDefinition.WorkGroup != nil {
						f11elemf1f0.WorkGroup = f11iter.DatasetDefinition.AthenaDatasetDefinition.WorkGroup
					}
					f11elemf1.AthenaDatasetDefinition = f11elemf1f0
				}
				if f11iter.DatasetDefinition.DataDistributionType != "" {
					f11elemf1.DataDistributionType = aws.String(string(f11iter.DatasetDefinition.DataDistributionType))
				}
				if f11iter.DatasetDefinition.InputMode != "" {
					f11elemf1.InputMode = aws.String(string(f11iter.DatasetDefinition.InputMode))
				}
				if f11iter.DatasetDefinition.LocalPath != nil {
					f11elemf1.LocalPath = f11iter.DatasetDefinition.LocalPath
				}
				if f11iter.DatasetDefinition.RedshiftDatasetDefinition != nil {
					f11elemf1f4 := &svcapitypes.RedshiftDatasetDefinition{}
					if f11iter.DatasetDefinition.RedshiftDatasetDefinition.ClusterId != nil {
						f11elemf1f4.ClusterID = f11iter.DatasetDefinition.RedshiftDatasetDefinition.ClusterId
					}
					if f11iter.DatasetDefinition.RedshiftDatasetDefinition.ClusterRoleArn != nil {
						f11elemf1f4.ClusterRoleARN = f11iter.DatasetDefinition.RedshiftDatasetDefinition.ClusterRoleArn
					}
					if f11iter.DatasetDefinition.RedshiftDatasetDefinition.Database != nil {
						f11elemf1f4.Database = f11iter.DatasetDefinition.RedshiftDatasetDefinition.Database
					}
					if f11iter.DatasetDefinition.RedshiftDatasetDefinition.DbUser != nil {
						f11elemf1f4.DBUser = f11iter.DatasetDefinition.RedshiftDatasetDefinition.DbUser
					}
					if f11iter.DatasetDefinition.RedshiftDatasetDefinition.KmsKeyId != nil {
						f11elemf1f4.KMSKeyID = f11iter.DatasetDefinition.RedshiftDatasetDefinition.KmsKeyId
					}
					if f11iter.DatasetDefinition.RedshiftDatasetDefinition.OutputCompression != "" {
						f11elemf1f4.OutputCompression = aws.String(string(f11iter.DatasetDefinition.RedshiftDatasetDefinition.OutputCompression))
					}
					if f11iter.DatasetDefinition.RedshiftDatasetDefinition.OutputFormat != "" {
						f11elemf1f4.OutputFormat = aws.String(string(f11iter.DatasetDefinition.RedshiftDatasetDefinition.OutputFormat))
					}
					if f11iter.DatasetDefinition.RedshiftDatasetDefinition.OutputS3Uri != nil {
						f11elemf1f4.OutputS3URI = f11iter.DatasetDefinition.RedshiftDatasetDefinition.OutputS3Uri
					}
					if f11iter.DatasetDefinition.RedshiftDatasetDefinition.QueryString != nil {
						f11elemf1f4.QueryString = f11iter.DatasetDefinition.RedshiftDatasetDefinition.QueryString
					}
					f11elemf1.RedshiftDatasetDefinition = f11elemf1f4
				}
				f11elem.DatasetDefinition = f11elemf1
			}
			if f11iter.InputName != nil {
				f11elem.InputName = f11iter.InputName
			}
			if f11iter.S3Input != nil {
				f11elemf3 := &svcapitypes.ProcessingS3Input{}
				if f11iter.S3Input.LocalPath != nil {
					f11elemf3.LocalPath = f11iter.S3Input.LocalPath
				}
				if f11iter.S3Input.S3CompressionType != "" {
					f11elemf3.S3CompressionType = aws.String(string(f11iter.S3Input.S3CompressionType))
				}
				if f11iter.S3Input.S3DataDistributionType != "" {
					f11elemf3.S3DataDistributionType = aws.String(string(f11iter.S3Input.S3DataDistributionType))
				}
				if f11iter.S3Input.S3DataType != "" {
					f11elemf3.S3DataType = aws.String(string(f11iter.S3Input.S3DataType))
				}
				if f11iter.S3Input.S3InputMode != "" {
					f11elemf3.S3InputMode = aws.String(string(f11iter.S3Input.S3InputMode))
				}
				if f11iter.S3Input.S3Uri != nil {
					f11elemf3.S3URI = f11iter.S3Input.S3Uri
				}
				f11elem.S3Input = f11elemf3
			}
			f11 = append(f11, f11elem)
		}
		ko.Spec.ProcessingInputs = f11
	} else {
		ko.Spec.ProcessingInputs = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.ProcessingJobArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.ProcessingJobArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.ProcessingJobName != nil {
		ko.Spec.ProcessingJobName = resp.ProcessingJobName
	} else {
		ko.Spec.ProcessingJobName = nil
	}
	if resp.ProcessingJobStatus != "" {
		ko.Status.ProcessingJobStatus = aws.String(string(resp.ProcessingJobStatus))
	} else {
		ko.Status.ProcessingJobStatus = nil
	}
	if resp.ProcessingOutputConfig != nil {
		f15 := &svcapitypes.ProcessingOutputConfig{}
		if resp.ProcessingOutputConfig.KmsKeyId != nil {
			f15.KMSKeyID = resp.ProcessingOutputConfig.KmsKeyId
		}
		if resp.ProcessingOutputConfig.Outputs != nil {
			f15f1 := []*svcapitypes.ProcessingOutput{}
			for _, f15f1iter := range resp.ProcessingOutputConfig.Outputs {
				f15f1elem := &svcapitypes.ProcessingOutput{}
				if f15f1iter.AppManaged != nil {
					f15f1elem.AppManaged = f15f1iter.AppManaged
				}
				if f15f1iter.FeatureStoreOutput != nil {
					f15f1elemf1 := &svcapitypes.ProcessingFeatureStoreOutput{}
					if f15f1iter.FeatureStoreOutput.FeatureGroupName != nil {
						f15f1elemf1.FeatureGroupName = f15f1iter.FeatureStoreOutput.FeatureGroupName
					}
					f15f1elem.FeatureStoreOutput = f15f1elemf1
				}
				if f15f1iter.OutputName != nil {
					f15f1elem.OutputName = f15f1iter.OutputName
				}
				if f15f1iter.S3Output != nil {
					f15f1elemf3 := &svcapitypes.ProcessingS3Output{}
					if f15f1iter.S3Output.LocalPath != nil {
						f15f1elemf3.LocalPath = f15f1iter.S3Output.LocalPath
					}
					if f15f1iter.S3Output.S3UploadMode != "" {
						f15f1elemf3.S3UploadMode = aws.String(string(f15f1iter.S3Output.S3UploadMode))
					}
					if f15f1iter.S3Output.S3Uri != nil {
						f15f1elemf3.S3URI = f15f1iter.S3Output.S3Uri
					}
					f15f1elem.S3Output = f15f1elemf3
				}
				f15f1 = append(f15f1, f15f1elem)
			}
			f15.Outputs = f15f1
		}
		ko.Spec.ProcessingOutputConfig = f15
	} else {
		ko.Spec.ProcessingOutputConfig = nil
	}
	if resp.ProcessingResources != nil {
		f16 := &svcapitypes.ProcessingResources{}
		if resp.ProcessingResources.ClusterConfig != nil {
			f16f0 := &svcapitypes.ProcessingClusterConfig{}
			if resp.ProcessingResources.ClusterConfig.InstanceCount != nil {
				instanceCountCopy := int64(*resp.ProcessingResources.ClusterConfig.InstanceCount)
				f16f0.InstanceCount = &instanceCountCopy
			}
			if resp.ProcessingResources.ClusterConfig.InstanceType != "" {
				f16f0.InstanceType = aws.String(string(resp.ProcessingResources.ClusterConfig.InstanceType))
			}
			if resp.ProcessingResources.ClusterConfig.VolumeKmsKeyId != nil {
				f16f0.VolumeKMSKeyID = resp.ProcessingResources.ClusterConfig.VolumeKmsKeyId
			}
			if resp.ProcessingResources.ClusterConfig.VolumeSizeInGB != nil {
				volumeSizeInGBCopy := int64(*resp.ProcessingResources.ClusterConfig.VolumeSizeInGB)
				f16f0.VolumeSizeInGB = &volumeSizeInGBCopy
			}
			f16.ClusterConfig = f16f0
		}
		ko.Spec.ProcessingResources = f16
	} else {
		ko.Spec.ProcessingResources = nil
	}
	if resp.RoleArn != nil {
		ko.Spec.RoleARN = resp.RoleArn
	} else {
		ko.Spec.RoleARN = nil
	}
	if resp.StoppingCondition != nil {
		f19 := &svcapitypes.ProcessingStoppingCondition{}
		if resp.StoppingCondition.MaxRuntimeInSeconds != nil {
			maxRuntimeInSecondsCopy := int64(*resp.StoppingCondition.MaxRuntimeInSeconds)
			f19.MaxRuntimeInSeconds = &maxRuntimeInSecondsCopy
		}
		ko.Spec.StoppingCondition = f19
	} else {
		ko.Spec.StoppingCondition = nil
	}

	rm.setStatusDefaults(ko)
	rm.customSetOutput(&resource{ko})
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.ProcessingJobName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeProcessingJobInput, error) {
	res := &svcsdk.DescribeProcessingJobInput{}

	if r.ko.Spec.ProcessingJobName != nil {
		res.ProcessingJobName = r.ko.Spec.ProcessingJobName
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateProcessingJobOutput
	_ = resp
	resp, err = rm.sdkapi.CreateProcessingJob(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateProcessingJob", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.ProcessingJobArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.ProcessingJobArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateProcessingJobInput, error) {
	res := &svcsdk.CreateProcessingJobInput{}

	if r.ko.Spec.AppSpecification != nil {
		f0 := &svcsdktypes.AppSpecification{}
		if r.ko.Spec.AppSpecification.ContainerArguments != nil {
			f0.ContainerArguments = aws.ToStringSlice(r.ko.Spec.AppSpecification.ContainerArguments)
		}
		if r.ko.Spec.AppSpecification.ContainerEntrypoint != nil {
			f0.ContainerEntrypoint = aws.ToStringSlice(r.ko.Spec.AppSpecification.ContainerEntrypoint)
		}
		if r.ko.Spec.AppSpecification.ImageURI != nil {
			f0.ImageUri = r.ko.Spec.AppSpecification.ImageURI
		}
		res.AppSpecification = f0
	}
	if r.ko.Spec.Environment != nil {
		res.Environment = aws.ToStringMap(r.ko.Spec.Environment)
	}
	if r.ko.Spec.ExperimentConfig != nil {
		f2 := &svcsdktypes.ExperimentConfig{}
		if r.ko.Spec.ExperimentConfig.ExperimentName != nil {
			f2.ExperimentName = r.ko.Spec.ExperimentConfig.ExperimentName
		}
		if r.ko.Spec.ExperimentConfig.TrialComponentDisplayName != nil {
			f2.TrialComponentDisplayName = r.ko.Spec.ExperimentConfig.TrialComponentDisplayName
		}
		if r.ko.Spec.ExperimentConfig.TrialName != nil {
			f2.TrialName = r.ko.Spec.ExperimentConfig.TrialName
		}
		res.ExperimentConfig = f2
	}
	if r.ko.Spec.NetworkConfig != nil {
		f3 := &svcsdktypes.NetworkConfig{}
		if r.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption != nil {
			f3.EnableInterContainerTrafficEncryption = r.ko.Spec.NetworkConfig.EnableInterContainerTrafficEncryption
		}
		if r.ko.Spec.NetworkConfig.EnableNetworkIsolation != nil {
			f3.EnableNetworkIsolation = r.ko.Spec.NetworkConfig.EnableNetworkIsolation
		}
		if r.ko.Spec.NetworkConfig.VPCConfig != nil {
			f3f2 := &svcsdktypes.VpcConfig{}
			if r.ko.Spec.NetworkConfig.VPCConfig.SecurityGroupIDs != nil {
				f3f2.SecurityGroupIds = aws.ToStringSlice(r.ko.Spec.NetworkConfig.VPCConfig.SecurityGroupIDs)
			}
			if r.ko.Spec.NetworkConfig.VPCConfig.Subnets != nil {
				f3f2.Subnets = aws.ToStringSlice(r.ko.Spec.NetworkConfig.VPCConfig.Subnets)
			}
			f3.VpcConfig = f3f2
		}
		res.NetworkConfig = f3
	}
	if r.ko.Spec.ProcessingInputs != nil {
		f4 := []svcsdktypes.ProcessingInput{}
		for _, f4iter := range r.ko.Spec.ProcessingInputs {
			f4elem := &svcsdktypes.ProcessingInput{}
			if f4iter.AppManaged != nil {
				f4elem.AppManaged = f4iter.AppManaged
			}
			if f4iter.DatasetDefinition != nil {
				f4elemf1 := &svcsdktypes.DatasetDefinition{}
				if f4iter.DatasetDefinition.AthenaDatasetDefinition != nil {
					f4elemf1f0 := &svcsdktypes.AthenaDatasetDefinition{}
					if f4iter.DatasetDefinition.AthenaDatasetDefinition.Catalog != nil {
						f4elemf1f0.Catalog = f4iter.DatasetDefinition.AthenaDatasetDefinition.Catalog
					}
					if f4iter.DatasetDefinition.AthenaDatasetDefinition.Database != nil {
						f4elemf1f0.Database = f4iter.DatasetDefinition.AthenaDatasetDefinition.Database
					}
					if f4iter.DatasetDefinition.AthenaDatasetDefinition.KMSKeyID != nil {
						f4elemf1f0.KmsKeyId = f4iter.DatasetDefinition.AthenaDatasetDefinition.KMSKeyID
					}
					if f4iter.DatasetDefinition.AthenaDatasetDefinition.OutputCompression != nil {
						f4elemf1f0.OutputCompression = svcsdktypes.AthenaResultCompressionType(*f4iter.DatasetDefinition.AthenaDatasetDefinition.OutputCompression)
					}
					if f4iter.DatasetDefinition.AthenaDatasetDefinition.OutputFormat != nil {
						f4elemf1f0.OutputFormat = svcsdktypes.AthenaResultFormat(*f4iter.DatasetDefinition.AthenaDatasetDefinition.OutputFormat)
					}
					if f4iter.DatasetDefinition.AthenaDatasetDefinition.OutputS3URI != nil {
						f4elemf1f0.OutputS3Uri = f4iter.DatasetDefinition.AthenaDatasetDefinition.OutputS3URI
					}
					if f4iter.DatasetDefinition.AthenaDatasetDefinition.QueryString != nil {
						f4elemf1f0.QueryString = f4iter.DatasetDefinition.AthenaDatasetDefinition.QueryString
					}
					if f4iter.DatasetDefinition.AthenaDatasetDefinition.WorkGroup != nil {
						f4elemf1f0.WorkGroup = f4iter.DatasetDefinition.AthenaDatasetDefinition.WorkGroup
					}
					f4elemf1.AthenaDatasetDefinition = f4elemf1f0
				}
				if f4iter.DatasetDefinition.DataDistributionType != nil {
					f4elemf1.DataDistributionType = svcsdktypes.DataDistributionType(*f4iter.DatasetDefinition.DataDistributionType)
				}
				if f4iter.DatasetDefinition.InputMode != nil {
					f4elemf1.InputMode = svcsdktypes.InputMode(*f4iter.DatasetDefinition.InputMode)
				}
				if f4iter.DatasetDefinition.LocalPath != nil {
					f4elemf1.LocalPath = f4iter.DatasetDefinition.LocalPath
				}
				if f4iter.DatasetDefinition.RedshiftDatasetDefinition != nil {
					f4elemf1f4 := &svcsdktypes.RedshiftDatasetDefinition{}
					if f4iter.DatasetDefinition.RedshiftDatasetDefinition.ClusterID != nil {
						f4elemf1f4.ClusterId = f4iter.DatasetDefinition.RedshiftDatasetDefinition.ClusterID
					}
					if f4iter.DatasetDefinition.RedshiftDatasetDefinition.ClusterRoleARN != nil {
						f4elemf1f4.ClusterRoleArn = f4iter.DatasetDefinition.RedshiftDatasetDefinition.ClusterRoleARN
					}
					if f4iter.DatasetDefinition.RedshiftDatasetDefinition.Database != nil {
						f4elemf1f4.Database = f4iter.DatasetDefinition.RedshiftDatasetDefinition.Database
					}
					if f4iter.DatasetDefinition.RedshiftDatasetDefinition.DBUser != nil {
						f4elemf1f4.DbUser = f4iter.DatasetDefinition.RedshiftDatasetDefinition.DBUser
					}
					if f4iter.DatasetDefinition.RedshiftDatasetDefinition.KMSKeyID != nil {
						f4elemf1f4.KmsKeyId = f4iter.DatasetDefinition.RedshiftDatasetDefinition.KMSKeyID
					}
					if f4iter.DatasetDefinition.RedshiftDatasetDefinition.OutputCompression != nil {
						f4elemf1f4.OutputCompression = svcsdktypes.RedshiftResultCompressionType(*f4iter.DatasetDefinition.RedshiftDatasetDefinition.OutputCompression)
					}
					if f4iter.DatasetDefinition.RedshiftDatasetDefinition.OutputFormat != nil {
						f4elemf1f4.OutputFormat = svcsdktypes.RedshiftResultFormat(*f4iter.DatasetDefinition.RedshiftDatasetDefinition.OutputFormat)
					}
					if f4iter.DatasetDefinition.RedshiftDatasetDefinition.OutputS3URI != nil {
						f4elemf1f4.OutputS3Uri = f4iter.DatasetDefinition.RedshiftDatasetDefinition.OutputS3URI
					}
					if f4iter.DatasetDefinition.RedshiftDatasetDefinition.QueryString != nil {
						f4elemf1f4.QueryString = f4iter.DatasetDefinition.RedshiftDatasetDefinition.QueryString
					}
					f4elemf1.RedshiftDatasetDefinition = f4elemf1f4
				}
				f4elem.DatasetDefinition = f4elemf1
			}
			if f4iter.InputName != nil {
				f4elem.InputName = f4iter.InputName
			}
			if f4iter.S3Input != nil {
				f4elemf3 := &svcsdktypes.ProcessingS3Input{}
				if f4iter.S3Input.LocalPath != nil {
					f4elemf3.LocalPath = f4iter.S3Input.LocalPath
				}
				if f4iter.S3Input.S3CompressionType != nil {
					f4elemf3.S3CompressionType = svcsdktypes.ProcessingS3CompressionType(*f4iter.S3Input.S3CompressionType)
				}
				if f4iter.S3Input.S3DataDistributionType != nil {
					f4elemf3.S3DataDistributionType = svcsdktypes.ProcessingS3DataDistributionType(*f4iter.S3Input.S3DataDistributionType)
				}
				if f4iter.S3Input.S3DataType != nil {
					f4elemf3.S3DataType = svcsdktypes.ProcessingS3DataType(*f4iter.S3Input.S3DataType)
				}
				if f4iter.S3Input.S3InputMode != nil {
					f4elemf3.S3InputMode = svcsdktypes.ProcessingS3InputMode(*f4iter.S3Input.S3InputMode)
				}
				if f4iter.S3Input.S3URI != nil {
					f4elemf3.S3Uri = f4iter.S3Input.S3URI
				}
				f4elem.S3Input = f4elemf3
			}
			f4 = append(f4, *f4elem)
		}
		res.ProcessingInputs = f4
	}
	if r.ko.Spec.ProcessingJobName != nil {
		res.ProcessingJobName = r.ko.Spec.ProcessingJobName
	}
	if r.ko.Spec.ProcessingOutputConfig != nil {
		f6 := &svcsdktypes.ProcessingOutputConfig{}
		if r.ko.Spec.ProcessingOutputConfig.KMSKeyID != nil {
			f6.KmsKeyId = r.ko.Spec.ProcessingOutputConfig.KMSKeyID
		}
		if r.ko.Spec.ProcessingOutputConfig.Outputs != nil {
			f6f1 := []svcsdktypes.ProcessingOutput{}
			for _, f6f1iter := range r.ko.Spec.ProcessingOutputConfig.Outputs {
				f6f1elem := &svcsdktypes.ProcessingOutput{}
				if f6f1iter.AppManaged != nil {
					f6f1elem.AppManaged = f6f1iter.AppManaged
				}
				if f6f1iter.FeatureStoreOutput != nil {
					f6f1elemf1 := &svcsdktypes.ProcessingFeatureStoreOutput{}
					if f6f1iter.FeatureStoreOutput.FeatureGroupName != nil {
						f6f1elemf1.FeatureGroupName = f6f1iter.FeatureStoreOutput.FeatureGroupName
					}
					f6f1elem.FeatureStoreOutput = f6f1elemf1
				}
				if f6f1iter.OutputName != nil {
					f6f1elem.OutputName = f6f1iter.OutputName
				}
				if f6f1iter.S3Output != nil {
					f6f1elemf3 := &svcsdktypes.ProcessingS3Output{}
					if f6f1iter.S3Output.LocalPath != nil {
						f6f1elemf3.LocalPath = f6f1iter.S3Output.LocalPath
					}
					if f6f1iter.S3Output.S3UploadMode != nil {
						f6f1elemf3.S3UploadMode = svcsdktypes.ProcessingS3UploadMode(*f6f1iter.S3Output.S3UploadMode)
					}
					if f6f1iter.S3Output.S3URI != nil {
						f6f1elemf3.S3Uri = f6f1iter.S3Output.S3URI
					}
					f6f1elem.S3Output = f6f1elemf3
				}
				f6f1 = append(f6f1, *f6f1elem)
			}
			f6.Outputs = f6f1
		}
		res.ProcessingOutputConfig = f6
	}
	if r.ko.Spec.ProcessingResources != nil {
		f7 := &svcsdktypes.ProcessingResources{}
		if r.ko.Spec.ProcessingResources.ClusterConfig != nil {
			f7f0 := &svcsdktypes.ProcessingClusterConfig{}
			if r.ko.Spec.ProcessingResources.ClusterConfig.InstanceCount != nil {
				instanceCountCopy0 := *r.ko.Spec.ProcessingResources.ClusterConfig.InstanceCount
				if instanceCountCopy0 > math.MaxInt32 || instanceCountCopy0 < math.MinInt32 {
					return nil, fmt.Errorf("error: field InstanceCount is of type int32")
				}
				instanceCountCopy := int32(instanceCountCopy0)
				f7f0.InstanceCount = &instanceCountCopy
			}
			if r.ko.Spec.ProcessingResources.ClusterConfig.InstanceType != nil {
				f7f0.InstanceType = svcsdktypes.ProcessingInstanceType(*r.ko.Spec.ProcessingResources.ClusterConfig.InstanceType)
			}
			if r.ko.Spec.ProcessingResources.ClusterConfig.VolumeKMSKeyID != nil {
				f7f0.VolumeKmsKeyId = r.ko.Spec.ProcessingResources.ClusterConfig.VolumeKMSKeyID
			}
			if r.ko.Spec.ProcessingResources.ClusterConfig.VolumeSizeInGB != nil {
				volumeSizeInGBCopy0 := *r.ko.Spec.ProcessingResources.ClusterConfig.VolumeSizeInGB
				if volumeSizeInGBCopy0 > math.MaxInt32 || volumeSizeInGBCopy0 < math.MinInt32 {
					return nil, fmt.Errorf("error: field VolumeSizeInGB is of type int32")
				}
				volumeSizeInGBCopy := int32(volumeSizeInGBCopy0)
				f7f0.VolumeSizeInGB = &volumeSizeInGBCopy
			}
			f7.ClusterConfig = f7f0
		}
		res.ProcessingResources = f7
	}
	if r.ko.Spec.RoleARN != nil {
		res.RoleArn = r.ko.Spec.RoleARN
	}
	if r.ko.Spec.StoppingCondition != nil {
		f9 := &svcsdktypes.ProcessingStoppingCondition{}
		if r.ko.Spec.StoppingCondition.MaxRuntimeInSeconds != nil {
			maxRuntimeInSecondsCopy0 := *r.ko.Spec.StoppingCondition.MaxRuntimeInSeconds
			if maxRuntimeInSecondsCopy0 > math.MaxInt32 || maxRuntimeInSecondsCopy0 < math.MinInt32 {
				return nil, fmt.Errorf("error: field MaxRuntimeInSeconds is of type int32")
			}
			maxRuntimeInSecondsCopy := int32(maxRuntimeInSecondsCopy0)
			f9.MaxRuntimeInSeconds = &maxRuntimeInSecondsCopy
		}
		res.StoppingCondition = f9
	}
	if r.ko.Spec.Tags != nil {
		f10 := []svcsdktypes.Tag{}
		for _, f10iter := range r.ko.Spec.Tags {
			f10elem := &svcsdktypes.Tag{}
			if f10iter.Key != nil {
				f10elem.Key = f10iter.Key
			}
			if f10iter.Value != nil {
				f10elem.Value = f10iter.Value
			}
			f10 = append(f10, *f10elem)
		}
		res.Tags = f10
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	return nil, ackerr.NewTerminalError(ackerr.NotImplemented)
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer func() {
		exit(err)
	}()
	latestStatus := r.ko.Status.ProcessingJobStatus
	if latestStatus != nil {
		if *latestStatus == string(svcsdktypes.ProcessingJobStatusStopping) {
			return r, requeueWaitWhileDeleting
		}

		// Call StopProcessingJob only if the job is InProgress, otherwise just
		// return nil to mark the resource Unmanaged
		if *latestStatus != string(svcsdktypes.ProcessingJobStatusInProgress) {
			return r, err
		}
	}
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.StopProcessingJobOutput
	_ = resp
	resp, err = rm.sdkapi.StopProcessingJob(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "StopProcessingJob", err)

	if err == nil {
		if observed, err := rm.sdkFind(ctx, r); err != ackerr.NotFound {
			if err != nil {
				return nil, err
			}
			r.SetStatus(observed)
			return r, requeueWaitWhileDeleting
		}
	}

	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.StopProcessingJobInput, error) {
	res := &svcsdk.StopProcessingJobInput{}

	if r.ko.Spec.ProcessingJobName != nil {
		res.ProcessingJobName = r.ko.Spec.ProcessingJobName
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.ProcessingJob,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}

	var terminalErr smithy.APIError
	if !errors.As(err, &terminalErr) {
		return false
	}
	switch terminalErr.ErrorCode() {
	case "ResourceNotFound",
		"ResourceInUse",
		"InvalidParameterCombination",
		"InvalidParameterValue",
		"MissingParameter":
		return true
	default:
		return false
	}
}
