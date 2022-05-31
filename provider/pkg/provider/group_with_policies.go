// Copyright 2016-2022, Pulumi Corporation.
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

package provider

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const GroupWithPoliciesIdentifier = "aws-iam:index:GroupWithPolicies"

type GroupWithPoliciesArgs struct {
	// Name of IAM group.
	Name string `pulumi:"name"`

	// List of IAM users to have in an IAM group which can assume the role.
	GroupUsers []string `pulumi:"groupUsers"`

	// List of IAM policies ARNs to attach to IAM group.
	CustomGroupPolicyARNs []string `pulumi:"customGroupPolicyArns"`

	// List of maps of inline IAM policies to attach to IAM group. Should have `name` and `policy` keys in each element.
	CustomGroupPolicies []map[string]string `pulumi:"customGroupPolicies"`

	// Whether to attach IAM policy which allows IAM users to manage their credentials and MFA.
	AttachIAMSelfManagementPolicy bool `pulumi:"attachIamSelfManagementPolicy"`

	// Name prefix for IAM policy to create with IAM self-management permissions.
	IAMSelfManagementPolicyNamePrefix string `pulumi:"iamSelfManagementPolicyNamePrefix"`

	// AWS account id to use inside IAM policies. If empty, current AWS account ID will be used.
	AWSAccountID string `pulumi:"awsAccountId"`

	// A map of tags to add to all resources.
	Tags map[string]string `pulumi:"tags"`
}

type GroupWithPolicies struct {
	pulumi.ResourceState

	// IAM AWS account id.
	AWSAccountID pulumi.StringOutput `pulumi:"awsAccountId"`

	// IAM group arn.
	GroupARN pulumi.StringOutput `pulumi:"groupArn"`

	// List of IAM users in IAM group.
	GroupUsers pulumi.StringArrayOutput `pulumi:"groupUsers"`

	// IAM group name.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
}

func NewGroupWithPolicies(ctx *pulumi.Context, name string, args *GroupWithPoliciesArgs, opts ...pulumi.ResourceOption) (*GroupWithPolicies, error) {
	if args == nil {
		args = &GroupWithPoliciesArgs{}
	}

	component := &GroupWithPolicies{}
	err := ctx.RegisterComponentResource(GroupWithPoliciesIdentifier, name, component, opts...)
	if err != nil {
		return nil, err
	}

	opts = append(opts, pulumi.Parent(component))

	awsAccountID := args.AWSAccountID
	if awsAccountID == "" {
		account, err := aws.GetCallerIdentity(ctx)
		if err != nil {
			return nil, err
		}

		awsAccountID = account.Id
	}

	currentPartition, err := aws.GetPartition(ctx, nil, nil)
	if err != nil {
		return nil, err
	}

	effect := "Allow"
	allowSelfManagement := "AllowSelfManagement"
	allowIAMReadOnly := "AllowIAMReadOnly"
	allowDeactivateMFADevice := "AllowDeactivateMFADevice"
	policyArgs := &iam.GetPolicyDocumentArgs{
		Statements: []iam.GetPolicyDocumentStatement{
			{
				Sid:    &allowSelfManagement,
				Effect: &effect,
				Actions: []string{
					"iam:ChangePassword",
					"iam:CreateAccessKey",
					"iam:CreateLoginProfile",
					"iam:CreateVirtualMFADevice",
					"iam:DeleteAccessKey",
					"iam:DeleteLoginProfile",
					"iam:DeleteVirtualMFADevice",
					"iam:EnableMFADevice",
					"iam:GenerateCredentialReport",
					"iam:GenerateServiceLastAccessedDetails",
					"iam:Get*",
					"iam:List*",
					"iam:ResyncMFADevice",
					"iam:UpdateAccessKey",
					"iam:UpdateLoginProfile",
					"iam:UpdateUser",
					"iam:UploadSigningCertificate",
					"iam:UploadSSHPublicKey",
				},
				Resources: []string{
					fmt.Sprintf("arn:%s:iam::%s:user/*/${aws:username}", currentPartition.Partition, awsAccountID),
					fmt.Sprintf("arn:%s:iam::%s:user/${aws:username}", currentPartition.Partition, awsAccountID),
					fmt.Sprintf("arn:%s:iam::%s:mfa/${aws:username}", currentPartition.Partition, awsAccountID),
				},
			},
			{
				Sid:       &allowIAMReadOnly,
				Effect:    &effect,
				Actions:   []string{"iam:Get*", "iam:List*"},
				Resources: []string{"*"},
			},
			{
				Sid:     &allowDeactivateMFADevice,
				Effect:  &effect,
				Actions: []string{"iam:DeactivateMFADevice"},
				Resources: []string{
					fmt.Sprintf("arn:%s:iam::%s:user/*/${aws:username}", currentPartition.Partition, awsAccountID),
					fmt.Sprintf("arn:%s:iam::%s:user/${aws:username}", currentPartition.Partition, awsAccountID),
					fmt.Sprintf("arn:%s:iam::%s:mfa/${aws:username}", currentPartition.Partition, awsAccountID),
				},
				Conditions: []iam.GetPolicyDocumentStatementCondition{
					{
						Test:     "Bool",
						Variable: "aws:MultiFactorAuthPresent",
						Values:   []string{"true"},
					},
					{
						Test:     "NumericLessThan",
						Variable: "aws:MultiFactorAuthAge",
						Values:   []string{"3600"},
					},
				},
			},
		},
	}

	policyDoc, err := iam.GetPolicyDocument(ctx, policyArgs)
	if err != nil {
		return nil, err
	}

	group, err := iam.NewGroup(ctx, name, &iam.GroupArgs{
		Name: pulumi.String(args.Name),
	}, opts...)
	if err != nil {
		return nil, err
	}

	_, err = iam.NewGroupMembership(ctx, name, &iam.GroupMembershipArgs{
		Group: group.ID(),
		Name:  pulumi.String(args.Name),
		Users: pulumi.ToStringArray(args.GroupUsers),
	}, opts...)
	if err != nil {
		return nil, err
	}

	if args.AttachIAMSelfManagementPolicy {
		if args.IAMSelfManagementPolicyNamePrefix == "" {
			args.IAMSelfManagementPolicyNamePrefix = "IAMSelfManagement-"
		}

		iamSelfManagementPolicy, err := iam.NewPolicy(ctx, name, &iam.PolicyArgs{
			NamePrefix: pulumi.String(args.IAMSelfManagementPolicyNamePrefix),
			Policy:     pulumi.String(policyDoc.Json),
			Tags:       pulumi.ToStringMap(args.Tags),
		}, opts...)
		if err != nil {
			return nil, err
		}

		_, err = iam.NewGroupPolicyAttachment(ctx, name, &iam.GroupPolicyAttachmentArgs{
			Group:     group.ID(),
			PolicyArn: iamSelfManagementPolicy.Arn,
		}, opts...)
		if err != nil {
			return nil, err
		}
	}

	for _, arn := range args.CustomGroupPolicyARNs {
		groupPolicyAttachmentName := fmt.Sprintf("%s-group-policy-attachment-%s", name, arn)
		_, err = iam.NewGroupPolicyAttachment(ctx, groupPolicyAttachmentName, &iam.GroupPolicyAttachmentArgs{
			Group:     group.ID(),
			PolicyArn: pulumi.String(arn),
		}, opts...)
		if err != nil {
			return nil, err
		}
	}

	for _, policyValues := range args.CustomGroupPolicies {
		policyDescription, ok := policyValues["description"]
		if !ok {
			policyDescription = ""
		}

		policyName := fmt.Sprintf("%s-policy-%s", name, policyValues["name"])
		customPolicy, err := iam.NewPolicy(ctx, policyName, &iam.PolicyArgs{
			Name:        pulumi.String(policyValues["name"]),
			Policy:      pulumi.String(policyValues["policy"]),
			Description: pulumi.String(policyDescription),
			Tags:        pulumi.ToStringMap(args.Tags),
		}, opts...)
		if err != nil {
			return nil, err
		}

		groupPolicyAttachmentName := fmt.Sprintf("%s-group-policy-attachment-%s", name, policyValues["name"])
		_, err = iam.NewGroupPolicyAttachment(ctx, groupPolicyAttachmentName, &iam.GroupPolicyAttachmentArgs{
			Group:     group.ID(),
			PolicyArn: customPolicy.Arn,
		}, opts...)
		if err != nil {
			return nil, err
		}
	}

	component.AWSAccountID = pulumi.Sprintf("%s", awsAccountID)
	component.GroupARN = group.Arn
	component.GroupName = group.Name
	component.GroupUsers = transformStringArrayToStringArrayOutput(args.GroupUsers)

	return component, nil
}
