// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eks_policies

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	veleroNamePrefix  = "Velero_Policy-"
	veleroDescription = "Provides Velero permissions to backup and restore cluster resources"

	veleroDefaultS3BucketARN = "*"
)

type VeleroPolicyArgs struct {
	// Determines whether to attach the Velero IAM policy to the role.
	Attach bool `pulumi:"attach"`

	// List of S3 Bucket ARNs that Velero needs access to in order to backup and restore cluster resources.
	S3BucketARNs pulumi.StringArrayInput `pulumi:"s3BucketArns"`
}

func AttachVeleroPolicy(ctx *pulumi.Context, policyBuilder *EKSRoleBuilder, args VeleroPolicyArgs) error {
	policyJSON := args.S3BucketARNs.ToStringArrayOutput().ApplyT(func(arns []string) (string, error) {
		if len(arns) == 0 {
			arns = append(arns, veleroDefaultS3BucketARN)
		}

		var s3ReadWriteResources []string
		for _, bucket := range arns {
			s3ReadWriteResources = append(s3ReadWriteResources, fmt.Sprintf("%s/*", bucket))
		}

		policyStatements := []iam.GetPolicyDocumentStatement{
			{
				Sid:       pulumi.StringRef("Ec2ReadWrite"),
				Resources: []string{"*"},
				Actions: []string{
					"ec2:DescribeVolumes",
					"ec2:DescribeSnapshots",
					"ec2:CreateTags",
					"ec2:CreateVolume",
					"ec2:CreateSnapshot",
					"ec2:DeleteSnapshot",
				},
			},
			{
				Sid:       pulumi.StringRef("S3ReadWrite"),
				Resources: s3ReadWriteResources,
				Actions: []string{
					"s3:GetObject",
					"s3:DeleteObject",
					"s3:PutObject",
					"s3:AbortMultipartUpload",
					"s3:ListMultipartUploadParts",
				},
			},
			{
				Sid:       pulumi.StringRef("S3List"),
				Resources: arns,
				Actions:   []string{"s3:ListBucket"},
			},
		}

		policyDoc, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: policyStatements,
		})
		if err != nil {
			return "", err
		}

		return policyDoc.Json, err
	}).(pulumi.StringOutput)

	return policyBuilder.CreatePolicyWithAttachment(veleroNamePrefix, veleroDescription, policyJSON)
}
