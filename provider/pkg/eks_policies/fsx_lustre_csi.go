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
)

const (
	fsxLustreCSINamePrefix  = "FSx_Lustre_CSI_Policy-"
	fsxLustreCSIDescription = "Provides permissions to manage FSx Lustre volumes via the container storage interface driver"

	fsxLustreCSIDefaultServiceRoleARN = "arn:aws:iam::*:role/aws-service-role/s3.data-source.lustre.fsx.amazonaws.com/*"
)

type FSXLustreCSIPolicyArgs struct {
	// Determines whether to attach the FSx for Lustre CSI Driver IAM policy to the role.
	Attach bool `pulumi:"attach"`

	// Service role ARNs to allow FSx for Lustre CSI create and manage FSX for Lustre service linked roles.
	ServiceRoleARNs []string `pulumi:"serviceRoleArns"`
}

func AttachFSXLustreCSIPolicy(policyBuilder *EKSRoleBuilder, dnsSuffix string, args FSXLustreCSIPolicyArgs) error {
	if len(args.ServiceRoleARNs) == 0 {
		args.ServiceRoleARNs = append(args.ServiceRoleARNs, fsxLustreCSIDefaultServiceRoleARN)
	}

	policyStatements := []iam.GetPolicyDocumentStatement{
		{
			Actions: []string{
				"iam:CreateServiceLinkedRole",
				"iam:AttachRolePolicy",
				"iam:PutRolePolicy",
			},
			Resources: args.ServiceRoleARNs,
		},
		{
			Actions:   []string{"iam:CreateServiceLinkedRole"},
			Resources: []string{"*"},
			Conditions: []iam.GetPolicyDocumentStatementCondition{
				NewPolicyDocCondition("StringLike", "iam:AWSServiceName", fmt.Sprintf("fsx.%s", dnsSuffix)),
			},
		},
		{
			Actions: []string{
				"s3:ListBucket",
				"fsx:CreateFileSystem",
				"fsx:DeleteFileSystem",
				"fsx:DescribeFileSystems",
				"fsx:TagResource",
			},
			Resources: []string{"*"},
		},
	}

	return policyBuilder.CreatePolicyWithAttachment(fsxLustreCSINamePrefix, fsxLustreCSIDescription, policyStatements)
}
