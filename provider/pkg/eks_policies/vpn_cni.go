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
	vpnCNINamePrefix  = "CNI_Policy-"
	vpnCNIDescription = "Provides the Amazon VPC CNI Plugin (amazon-vpc-cni-k8s) the permissions it requires to modify the IPv4/IPv6 address configuration on your EKS worker nodes"
)

type VPNCNIPolicyArgs struct {
	// Determines whether to attach the VPC CNI IAM policy to the role.
	Attach bool `pulumi:"attach"`

	// Determines whether to enable IPv4 permissions for VPC CNI policy.
	EnableIPV4 bool `pulumi:"enableIpv4"`

	// Determines whether to enable IPv6 permissions for VPC CNI policy.
	EnableIpv6 bool `pulumi:"enableIpv6"`
}

func AttachVPNCNIPolicy(policyBuilder *EKSRoleBuilder, partition string, args VPNCNIPolicyArgs) error {
	policyStatements := []iam.GetPolicyDocumentStatement{
		{
			Sid:       pulumi.StringRef("CreateTags"),
			Actions:   []string{"ec2:CreateTags"},
			Resources: []string{fmt.Sprintf("arn:%s:ec2:*:*:network-interface/*", partition)},
		},
	}

	if args.EnableIPV4 {
		policyStatements = append(policyStatements, iam.GetPolicyDocumentStatement{
			Sid: pulumi.StringRef("IPV4"),
			Actions: []string{
				"ec2:AssignPrivateIpAddresses",
				"ec2:AttachNetworkInterface",
				"ec2:CreateNetworkInterface",
				"ec2:DeleteNetworkInterface",
				"ec2:DescribeInstances",
				"ec2:DescribeTags",
				"ec2:DescribeNetworkInterfaces",
				"ec2:DescribeInstanceTypes",
				"ec2:DetachNetworkInterface",
				"ec2:ModifyNetworkInterfaceAttribute",
				"ec2:UnassignPrivateIpAddresses",
			},
			Resources: []string{"*"},
		})
	}

	if args.EnableIpv6 {
		policyStatements = append(policyStatements, iam.GetPolicyDocumentStatement{
			Sid: pulumi.StringRef("IPV6"),
			Actions: []string{
				"ec2:AssignIpv6Addresses",
				"ec2:DescribeInstances",
				"ec2:DescribeTags",
				"ec2:DescribeNetworkInterfaces",
				"ec2:DescribeInstanceTypes",
			},
			Resources: []string{"*"},
		})
	}

	return policyBuilder.CreatePolicyWithAttachment(vpnCNINamePrefix, vpnCNIDescription, policyStatements)
}
