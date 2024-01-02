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

const PolicyIdentifier = "aws-iam:index:Policy"

type PolicyArgs struct {
	// Whether to create the IAM policy.
	CreatePolicy bool `pulumi:"createPolicy"`

	// The name of the policy.
	Name string `pulumi:"name"`

	// The path of the policy in IAM.
	Path string `pulumi:"path"`

	// The description of the policy.
	Description string `pulumi:"description"`

	// The policy document.
	PolicyDocument string `pulumi:"policyDocument"`

	// A map of tags to add to all resources.
	Tags map[string]string `pulumi:"tags"`
}

type Policy struct {
	pulumi.ResourceState

	// The policy's ID.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ARN assigned by AWS to this policy.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The description of the policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`

	// The name of the policy.
	Name pulumi.StringOutput `pulumi:"name"`

	// The path of the policy in IAM.
	Path pulumi.StringPtrOutput `pulumi:"path"`

	// The policy document.
	PolicyDocument pulumi.StringOutput `pulumi:"policyDocument"`
}

func NewPolicy(ctx *pulumi.Context, name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		args = &PolicyArgs{}
	}

	component := &Policy{}
	err := ctx.RegisterComponentResource(PolicyIdentifier, name, component, opts...)
	if err != nil {
		return nil, err
	}

	opts = append(opts, pulumi.Parent(component))

	policy, err := iam.NewPolicy(ctx, name, &iam.PolicyArgs{
		Name:        pulumi.String(args.Name),
		Path:        pulumi.String(args.Path),
		Policy:      pulumi.String(args.PolicyDocument),
		Description: pulumi.String(args.Description),
		Tags:        pulumi.ToStringMap(args.Tags),
	}, opts...)

	component.ID = policy.ID()
	component.Arn = policy.Arn
	component.Description = policy.Description
	component.Name = policy.Name
	component.Path = policy.Path
	component.PolicyDocument = policy.Policy

	return component, nil
}
