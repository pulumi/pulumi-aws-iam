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
	loadBalancerControllerNamePrefix  = "AWS_Load_Balancer_Controller-"
	loadBalancerControllerDescription = "Provides permissions for AWS Load Balancer Controller addon"

	loadBalancerTargetGroupBindingOnlyNamePrefix  = "AWS_Load_Balancer_Controller_TargetGroup_Only-"
	loadBalancerTargetGroupBindingOnlyDescription = "Provides permissions for AWS Load Balancer Controller addon in TargetGroup binding only scenario"
)

type LoadBalancerPolicyArgs struct {
	// Determines whether to attach the Load Balancer Controller policy to the role.
	Controller bool `pulumi:"controller"`

	// Determines whether to attach the Load Balancer Controller policy for the TargetGroupBinding only.
	TargetGroupBindingOnly bool `pulumi:"targetGroupBindingOnly"`
}

func AttachLoadBalancerControllerPolicy(policyBuilder *EKSRoleBuilder, partition, dnsSuffix string) error {
	policyStatements := []iam.GetPolicyDocumentStatement{
		{
			Actions:   []string{"iam:CreateServiceLinkedRole"},
			Resources: []string{"*"},
			Conditions: []iam.GetPolicyDocumentStatementCondition{
				NewPolicyDocCondition("StringEquals", "iam:AWSServiceName", fmt.Sprintf("elasticloadbalancing.%s", dnsSuffix)),
			},
		},
		{
			Resources: []string{"*"},
			Actions: []string{
				"ec2:DescribeAccountAttributes",
				"ec2:DescribeAddresses",
				"ec2:DescribeAvailabilityZones",
				"ec2:DescribeInternetGateways",
				"ec2:DescribeVpcs",
				"ec2:DescribeVpcPeeringConnections",
				"ec2:DescribeSubnets",
				"ec2:DescribeSecurityGroups",
				"ec2:DescribeInstances",
				"ec2:DescribeNetworkInterfaces",
				"ec2:DescribeTags",
				"ec2:GetCoipPoolUsage",
				"ec2:DescribeCoipPools",
				"elasticloadbalancing:DescribeLoadBalancers",
				"elasticloadbalancing:DescribeLoadBalancerAttributes",
				"elasticloadbalancing:DescribeListeners",
				"elasticloadbalancing:DescribeListenerCertificates",
				"elasticloadbalancing:DescribeSSLPolicies",
				"elasticloadbalancing:DescribeRules",
				"elasticloadbalancing:DescribeTargetGroups",
				"elasticloadbalancing:DescribeTargetGroupAttributes",
				"elasticloadbalancing:DescribeTargetHealth",
				"elasticloadbalancing:DescribeTags",
			},
		},
		{
			Resources: []string{"*"},
			Actions: []string{
				"cognito-idp:DescribeUserPoolClient",
				"acm:ListCertificates",
				"acm:DescribeCertificate",
				"iam:ListServerCertificates",
				"iam:GetServerCertificate",
				"waf-regional:GetWebACL",
				"waf-regional:GetWebACLForResource",
				"waf-regional:AssociateWebACL",
				"waf-regional:DisassociateWebACL",
				"wafv2:GetWebACL",
				"wafv2:GetWebACLForResource",
				"wafv2:AssociateWebACL",
				"wafv2:DisassociateWebACL",
				"shield:GetSubscriptionState",
				"shield:DescribeProtection",
				"shield:CreateProtection",
				"shield:DeleteProtection",
			},
		},
		{
			Resources: []string{"*"},
			Actions: []string{
				"ec2:AuthorizeSecurityGroupIngress",
				"ec2:RevokeSecurityGroupIngress",
				"ec2:CreateSecurityGroup",
			},
		},
		{
			Resources: []string{fmt.Sprintf("arn:%s:ec2:*:*:security-group/*", partition)},
			Actions:   []string{"ec2:CreateTags"},
			Conditions: []iam.GetPolicyDocumentStatementCondition{
				NewPolicyDocCondition("StringEquals", "ec2:CreateAction", "CreateSecurityGroup"),
				NewPolicyDocCondition("Null", "aws:RequestTag/elbv2.k8s.aws/cluster", "false"),
			},
		},
		{
			Actions:   []string{"ec2:CreateTags", "ec2:DeleteTags"},
			Resources: []string{fmt.Sprintf("arn:%s:ec2:*:*:security-group/*", partition)},
			Conditions: []iam.GetPolicyDocumentStatementCondition{
				NewPolicyDocCondition("Null", "aws:RequestTag/elbv2.k8s.aws/cluster", "true"),
				NewPolicyDocCondition("Null", "aws:ResourceTag/elbv2.k8s.aws/cluster", "false"),
			},
		},
		{
			Resources: []string{"*"},
			Actions: []string{
				"ec2:AuthorizeSecurityGroupIngress",
				"ec2:RevokeSecurityGroupIngress",
				"ec2:DeleteSecurityGroup",
			},
			Conditions: []iam.GetPolicyDocumentStatementCondition{
				NewPolicyDocCondition("Null", "aws:ResourceTag/elbv2.k8s.aws/cluster", "false"),
			},
		},
		{
			Resources: []string{"*"},
			Actions: []string{
				"elasticloadbalancing:CreateLoadBalancer",
				"elasticloadbalancing:CreateTargetGroup",
			},
			Conditions: []iam.GetPolicyDocumentStatementCondition{
				NewPolicyDocCondition("Null", "aws:RequestTag/elbv2.k8s.aws/cluster", "false"),
			},
		},
		{
			Resources: []string{"*"},
			Actions: []string{
				"elasticloadbalancing:CreateListener",
				"elasticloadbalancing:DeleteListener",
				"elasticloadbalancing:CreateRule",
				"elasticloadbalancing:DeleteRule",
			},
		},
		{
			Actions: []string{
				"elasticloadbalancing:AddTags",
				"elasticloadbalancing:RemoveTags",
			},
			Resources: []string{
				fmt.Sprintf("arn:%s:elasticloadbalancing:*:*:targetgroup/*/*", partition),
				fmt.Sprintf("arn:%s:elasticloadbalancing:*:*:loadbalancer/net/*/*", partition),
				fmt.Sprintf("arn:%s:elasticloadbalancing:*:*:loadbalancer/app/*/*", partition),
			},
			Conditions: []iam.GetPolicyDocumentStatementCondition{
				NewPolicyDocCondition("Null", "aws:RequestTag/elbv2.k8s.aws/cluster", "true"),
				NewPolicyDocCondition("Null", "aws:ResourceTag/elbv2.k8s.aws/cluster", "false"),
			},
		},
		{
			Actions: []string{
				"elasticloadbalancing:AddTags",
				"elasticloadbalancing:RemoveTags",
			},
			Resources: []string{
				fmt.Sprintf("arn:%s:elasticloadbalancing:*:*:listener/net/*/*/*", partition),
				fmt.Sprintf("arn:%s:elasticloadbalancing:*:*:listener/app/*/*/*", partition),
				fmt.Sprintf("arn:%s:elasticloadbalancing:*:*:listener-rule/net/*/*/*", partition),
				fmt.Sprintf("arn:%s:elasticloadbalancing:*:*:listener-rule/app/*/*/*", partition),
			},
		},
		{
			Resources: []string{"*"},
			Actions: []string{
				"elasticloadbalancing:ModifyLoadBalancerAttributes",
				"elasticloadbalancing:SetIpAddressType",
				"elasticloadbalancing:SetSecurityGroups",
				"elasticloadbalancing:SetSubnets",
				"elasticloadbalancing:DeleteLoadBalancer",
				"elasticloadbalancing:ModifyTargetGroup",
				"elasticloadbalancing:ModifyTargetGroupAttributes",
				"elasticloadbalancing:DeleteTargetGroup",
			},
			Conditions: []iam.GetPolicyDocumentStatementCondition{
				NewPolicyDocCondition("Null", "aws:ResourceTag/elbv2.k8s.aws/cluster", "false"),
			},
		},
		{
			Actions: []string{
				"elasticloadbalancing:RegisterTargets",
				"elasticloadbalancing:DeregisterTargets",
			},
			Resources: []string{fmt.Sprintf("arn:%s:elasticloadbalancing:*:*:targetgroup/*/*", partition)},
		},
		{
			Resources: []string{"*"},
			Actions: []string{
				"elasticloadbalancing:SetWebAcl",
				"elasticloadbalancing:ModifyListener",
				"elasticloadbalancing:AddListenerCertificates",
				"elasticloadbalancing:RemoveListenerCertificates",
				"elasticloadbalancing:ModifyRule",
			},
		},
	}

	return policyBuilder.CreatePolicyWithAttachment(loadBalancerControllerNamePrefix, loadBalancerControllerDescription, policyStatements)
}

func AttachLoadBalancerTargetGroupBindingOnlyPolicy(policyBuilder *EKSRoleBuilder) error {
	policyStatements := []iam.GetPolicyDocumentStatement{
		{
			Resources: []string{"*"},
			Actions: []string{
				"ec2:DescribeSecurityGroups",
				"ec2:DescribeInstances",
				"ec2:DescribeVpcs",
				"elasticloadbalancing:DescribeTargetGroups",
				"elasticloadbalancing:DescribeTargetHealth",
				"elasticloadbalancing:ModifyTargetGroup",
				"elasticloadbalancing:ModifyTargetGroupAttributes",
				"elasticloadbalancing:RegisterTargets",
				"elasticloadbalancing:DeregisterTargets",
			},
		},
	}

	return policyBuilder.CreatePolicyWithAttachment(loadBalancerTargetGroupBindingOnlyNamePrefix, loadBalancerTargetGroupBindingOnlyDescription, policyStatements)
}
