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
	ebsCSINamePrefix  = "EBS_CSI_Policy-"
	ebsCSIDescription = "Provides permissions to manage EBS volumes via the container storage interface driver"
)

type EBSCSIPolicyArgs struct {
	// Determines whether to attach the EBS CSI IAM policy to the role.
	Attach bool `pulumi:"attach"`

	// KMS CMK IDs to allow EBS CSI to manage encrypted volumes.
	KMSCMKIDs pulumi.StringArrayInput `pulumi:"kmsCmkIds"`
}

func AttachEBSCSIPolicy(ctx *pulumi.Context, policyBuilder *EKSRoleBuilder, partition string, args EBSCSIPolicyArgs) error {
	policyJSON := args.KMSCMKIDs.ToStringArrayOutput().ApplyT(func(ids []string) (string, error) {
		policyStatements := []iam.GetPolicyDocumentStatement{
			{
				Resources: []string{"*"},
				Actions: []string{
					"ec2:CreateSnapshot",
					"ec2:AttachVolume",
					"ec2:DetachVolume",
					"ec2:ModifyVolume",
					"ec2:DescribeAvailabilityZones",
					"ec2:DescribeInstances",
					"ec2:DescribeSnapshots",
					"ec2:DescribeTags",
					"ec2:DescribeVolumes",
					"ec2:DescribeVolumesModifications",
				},
			},
			{
				Actions: []string{"ec2:CreateTags"},
				Resources: []string{
					fmt.Sprintf("arn:%s:ec2:*:*:volume/*", partition),
					fmt.Sprintf("arn:%s:ec2:*:*:snapshot/*", partition),
				},
				Conditions: []iam.GetPolicyDocumentStatementCondition{
					NewPolicyDocCondition("StringEquals", "ec2:CreateAction", "CreateVolume", "CreateSnapShot"),
				},
			},
			{
				Actions: []string{"ec2:DeleteTags"},
				Resources: []string{
					fmt.Sprintf("arn:%s:ec2:*:*:volume/*", partition),
					fmt.Sprintf("arn:%s:ec2:*:*:snapshot/*", partition),
				},
			},
			{
				Actions:   []string{"ec2:CreateVolume"},
				Resources: []string{"*"},
				Conditions: []iam.GetPolicyDocumentStatementCondition{
					NewPolicyDocCondition("StringLike", "aws:RequestTag/ebs.csi.aws.com/cluster", "true"),
				},
			},
			{
				Actions:   []string{"ec2:CreateVolume"},
				Resources: []string{"*"},
				Conditions: []iam.GetPolicyDocumentStatementCondition{
					NewPolicyDocCondition("StringLike", "aws:RequestTag/CSIVolumeName", "*"),
				},
			},
			{
				Actions:   []string{"ec2:CreateVolume"},
				Resources: []string{"*"},
				Conditions: []iam.GetPolicyDocumentStatementCondition{
					NewPolicyDocCondition("StringLike", "aws:RequestTag/kubernetes.io/cluster/*", "owned"),
				},
			},
			{
				Actions:   []string{"ec2:DeleteVolume"},
				Resources: []string{"*"},
				Conditions: []iam.GetPolicyDocumentStatementCondition{
					NewPolicyDocCondition("StringLike", "aws:RequestTag/ebs.csi.aws.com/cluster", "true"),
				},
			},
			{
				Actions:   []string{"ec2:DeleteVolume"},
				Resources: []string{"*"},
				Conditions: []iam.GetPolicyDocumentStatementCondition{
					NewPolicyDocCondition("StringLike", "aws:RequestTag/CSIVolumeName", "*"),
				},
			},
			{
				Actions:   []string{"ec2:DeleteVolume"},
				Resources: []string{"*"},
				Conditions: []iam.GetPolicyDocumentStatementCondition{
					NewPolicyDocCondition("StringLike", "aws:RequestTag/kubernetes.io/cluster/*", "owned"),
				},
			},
			{
				Actions:   []string{"ec2:DeleteSnapshot"},
				Resources: []string{"*"},
				Conditions: []iam.GetPolicyDocumentStatementCondition{
					NewPolicyDocCondition("StringLike", "ec2:ResourceTag/CSIVolumeSnapshotName", "*"),
				},
			},
			{
				Actions:   []string{"ec2:DeleteSnapshot"},
				Resources: []string{"*"},
				Conditions: []iam.GetPolicyDocumentStatementCondition{
					NewPolicyDocCondition("StringLike", "ec2:ResourceTag/ebs.csi.aws.com/cluster", "true"),
				},
			},
		}

		if len(ids) > 0 {
			policyStatements = append(policyStatements, iam.GetPolicyDocumentStatement{
				Actions: []string{
					"kms:CreateGrant",
					"kms:ListGrants",
					"kms:RevokeGrant",
				},
				Resources: ids,
			})
			policyStatements = append(policyStatements, iam.GetPolicyDocumentStatement{
				Actions: []string{
					"kms:Encrypt",
					"kms:Decrypt",
					"kms:ReEncrypt*",
					"kms:GenerateDataKey*",
					"kms:DescribeKey",
				},
				Resources: ids,
			})
		}

		policyDoc, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: policyStatements,
		})
		if err != nil {
			return "", err
		}

		return policyDoc.Json, err
	}).(pulumi.StringOutput)

	return policyBuilder.CreatePolicyWithAttachment(ebsCSINamePrefix, ebsCSIDescription, policyJSON)
}
