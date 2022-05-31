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
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const GroupWithAssumableRolesPolicyIdentifier = "aws-iam:index:GroupWithAssumableRolesPolicy"

type GroupWithAssumableRolesPolicyArgs struct {
	// Name of IAM policy and IAM group.
	Name string `pulumi:"name"`

	// List of IAM roles ARNs which can be assumed by the group.
	AssumableRoles []string `pulumi:"assumableRoles"`

	// List of IAM users to have in an IAM group which can assume the role.
	GroupUsers []string `pulumi:"groupUsers"`

	// A map of tags to add to all resources.
	Tags map[string]string `pulumi:"tags"`
}

type GroupWithAssumableRolesPolicy struct {
	pulumi.ResourceState

	// List of IAM users in IAM group.
	GroupUsers pulumi.StringArrayOutput `pulumi:"groupUsers"`

	// List of ARNs of IAM roles which members of IAM group can assume.
	AssumableRoles pulumi.StringArrayOutput `pulumi:"assumableRoles"`

	// Assume role policy ARN of IAM group.
	PolicyARN pulumi.StringOutput `pulumi:"policyArn"`

	// IAM group name.
	GroupName pulumi.StringOutput `pulumi:"groupName"`

	// IAM group arn.
	GroupARN pulumi.StringOutput `pulumi:"groupArn"`
}

func NewGroupWithAssumableRolesPolicy(ctx *pulumi.Context, name string, args *GroupWithAssumableRolesPolicyArgs, opts ...pulumi.ResourceOption) (*GroupWithAssumableRolesPolicy, error) {
	if args == nil {
		args = &GroupWithAssumableRolesPolicyArgs{}
	}

	component := &GroupWithAssumableRolesPolicy{}
	err := ctx.RegisterComponentResource(GroupWithAssumableRolesPolicyIdentifier, name, component, opts...)
	if err != nil {
		return nil, err
	}

	opts = append(opts, pulumi.Parent(component))

	policyDocArgs := newIAMPolicyDocumentStatementConstructor("Allow", []string{"sts:AssumeRole"}).
		AddResources(args.AssumableRoles).
		Build()

	assumeRole, err := iam.GetPolicyDocument(ctx, policyDocArgs)
	if err != nil {
		return nil, err
	}

	policy, err := iam.NewPolicy(ctx, name, &iam.PolicyArgs{
		Name:        pulumi.String(args.Name),
		Description: pulumi.String("Allows to assume role in another AWS account"),
		Policy:      pulumi.String(assumeRole.Json),
		Tags:        pulumi.ToStringMap(args.Tags),
	}, opts...)
	if err != nil {
		return nil, err
	}

	group, err := iam.NewGroup(ctx, name, &iam.GroupArgs{
		Name: pulumi.String(name),
	}, opts...)
	if err != nil {
		return nil, err
	}

	_, err = iam.NewGroupPolicyAttachment(ctx, name, &iam.GroupPolicyAttachmentArgs{
		Group:     group.ID(),
		PolicyArn: policy.ID(),
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

	component.GroupUsers = transformStringArrayToStringArrayOutput(args.GroupUsers)
	component.AssumableRoles = transformStringArrayToStringArrayOutput(args.AssumableRoles)
	component.GroupARN = group.Arn
	component.GroupName = group.Name
	component.PolicyARN = policy.Arn

	return component, nil
}
