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
	karpenterControllerNamePrefix  = "Karpenter_Controller_Policy-"
	karpenterControllerDescription = "Provides permissions to the Krpenter Controller."

	karpenterControllerDefaultSSMParameterARN = "arn:aws:ssm:*:*:parameter/aws/service/*"
	karpenterControllerDefaultNodeIAMRoleARN  = "*"
)

type KarpenterControllerPolicyArgs struct {
	// Determines whether to attach the Karpenter Controller policy to the role.
	Attach bool `pulumi:"attach"`

	// Cluster ID where the Karpenter controller is provisioned/managing.
	ClusterID string `pulumi:"clusterId"`

	// Tag key (`{key = value}`) applied to resources launched by Karpenter through the Karpenter provisioner.
	TagKey string `pulumi:"tagKey"`

	// List of SSM Parameter ARNs that contain AMI IDs launched by Karpenter.
	SSMParameterARNs []string `pulumi:"ssmParameterArns"`

	// List of node IAM role ARNs Karpenter can use to launch nodes.
	NodeIAMRoleARNS []string `pulumi:"nodeIamRoleArns"`

	// Account ID of where the subnets Karpenter will utilize resides. Used when subnets are shared from another account.
	SubnetAccountID string `pulumi:"subnetAccountId"`
}

func AttachKarpenterControllerPolicy(policyBuilder *EKSRoleBuilder, partition, awsAccountID string, args KarpenterControllerPolicyArgs) error {
	karpenterSubnetId := args.SubnetAccountID
	if karpenterSubnetId == "" {
		karpenterSubnetId = awsAccountID
	}

	if len(args.SSMParameterARNs) == 0 {
		args.SSMParameterARNs = append(args.SSMParameterARNs, karpenterControllerDefaultSSMParameterARN)
	}

	if len(args.NodeIAMRoleARNS) == 0 {
		args.NodeIAMRoleARNS = append(args.NodeIAMRoleARNS, karpenterControllerDefaultNodeIAMRoleARN)
	}

	policyStatements := []iam.GetPolicyDocumentStatement{
		{
			Actions: []string{
				"ec2:CreateLaunchTemplate",
				"ec2:CreateFleet",
				"ec2:CreateTags",
				"ec2:DescribeLaunchTemplates",
				"ec2:DescribeInstances",
				"ec2:DescribeSecurityGroups",
				"ec2:DescribeSubnets",
				"ec2:DescribeInstanceTypes",
				"ec2:DescribeInstanceTypeOfferings",
				"ec2:DescribeAvailabilityZones",
			},
			Resources: []string{"*"},
		},
		{
			Actions: []string{
				"ec2:TerminateInstances",
				"ec2:DeleteLaunchTemplate",
			},
			Resources: []string{"*"},
			Conditions: []iam.GetPolicyDocumentStatementCondition{
				NewPolicyDocCondition("StringEquals", fmt.Sprintf("ec2:ResourceTag/%s", args.TagKey), args.ClusterID),
			},
		},
		{
			Actions: []string{"ec2:RunInstances"},
			Resources: []string{
				fmt.Sprintf("arn:%s:ec2:*:%s:launch-template/*", partition, awsAccountID),
				fmt.Sprintf("arn:%s:ec2:*:%s:security-group/*", partition, awsAccountID),
				fmt.Sprintf("arn:%s:ec2:*:%s:subnet/*", partition, karpenterSubnetId),
			},
			Conditions: []iam.GetPolicyDocumentStatementCondition{
				NewPolicyDocCondition("StringEquals", fmt.Sprintf("ec2:ResourceTag/%s", args.TagKey), args.ClusterID),
			},
		},
		{
			Actions: []string{"ec2:RunInstances"},
			Resources: []string{
				fmt.Sprintf("arn:%s:ec2:*::image/*", partition),
				fmt.Sprintf("arn:%s:ec2:*:%s:instance/*", partition, awsAccountID),
				fmt.Sprintf("arn:%s:ec2:*:%s:volume/*", partition, awsAccountID),
				fmt.Sprintf("arn:%s:ec2:*:%s:network-interface/*", partition, awsAccountID),
			},
		},
		{
			Actions:   []string{"ssm:GetParameter"},
			Resources: args.SSMParameterARNs,
		},
		{
			Actions:   []string{"iam:PassRole"},
			Resources: args.NodeIAMRoleARNS,
		},
	}

	return policyBuilder.CreatePolicyWithAttachment(karpenterControllerNamePrefix, karpenterControllerDescription, policyStatements)
}
