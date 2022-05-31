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
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
)

const (
	efsCSINamePrefix  = "EFS_CSI_Policy-"
	efsCSIDescription = "Provides permissions to manage EFS volumes via the container storage interface driver"
)

type EFSCSIPolicyArgs struct {
	// Determines whether to attach the EFS CSI IAM policy to the role.
	Attach bool `pulumi:"attach"`
}

func AttachEFSCSIPolicy(policyBuilder *EKSRoleBuilder) error {
	policyStatements := []iam.GetPolicyDocumentStatement{
		{
			Resources: []string{"*"},
			Actions: []string{
				"ec2:DescribeAvailabilityZones",
				"elasticfilesystem:DescribeAccessPoints",
				"elasticfilesystem:DescribeFileSystems",
				"elasticfilesystem:DescribeMountTargets",
			},
		},
		{
			Resources: []string{"*"},
			Actions:   []string{"elasticfilesystem:CreateAccessPoint"},
			Conditions: []iam.GetPolicyDocumentStatementCondition{
				NewPolicyDocCondition("StringLike", "aws:RequestTag/efs.csi.aws.com/cluster", "true"),
			},
		},
		{
			Resources: []string{"*"},
			Actions:   []string{"elasticfilesystem:DeleteAccessPoint"},
			Conditions: []iam.GetPolicyDocumentStatementCondition{
				NewPolicyDocCondition("StringEquals", "aws:ResourceTag/efs.csi.aws.com/cluster", "true"),
			},
		},
	}

	return policyBuilder.CreatePolicyWithAttachment(efsCSINamePrefix, efsCSIDescription, policyStatements)
}
