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
	"github.com/pulumi/pulumi-aws-iam/pkg/utils"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const AssumableRoleWithSAMLIdentifier = "aws-iam:index:AssumableRoleWithSAML"

type AssumableRoleWithSAMLArgs struct {
	// List of SAML Provider IDs.
	ProviderIDs pulumi.StringArrayInput `pulumi:"providerIds"`

	// AWS SAML Endpoint.
	AWSSAMLEndpoint string `pulumi:"awsSamlEndpoint"`

	// A map of tags to add.
	Tags map[string]string `pulumi:"tags"`

	// IAM role.
	Role utils.RoleArgs `pulumi:"role"`

	// Maximum CLI/API session duration in seconds between 3600 and 43200.
	MaxSessionDuration int `pulumi:"maxSessionDuration"`

	// Whether policies should be detached from this role when destroying.
	ForceDetachPolicies bool `pulumi:"forceDetachPolicies"`
}

type AssumableRoleWithSAML struct {
	pulumi.ResourceState

	// ARN of IAM role.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// Name of IAM role.
	Name pulumi.StringOutput `pulumi:"name"`

	// Path of IAM role.
	Path pulumi.StringPtrOutput `pulumi:"path"`

	// Unique ID of IAM role.
	UniqueID pulumi.StringOutput `pulumi:"uniqueId"`
}

func NewAssumableRoleWithSAML(ctx *pulumi.Context, name string, args *AssumableRoleWithSAMLArgs, opts ...pulumi.ResourceOption) (*AssumableRoleWithSAML, error) {
	if args == nil {
		args = &AssumableRoleWithSAMLArgs{}
	}

	component := &AssumableRoleWithSAML{}
	err := ctx.RegisterComponentResource(AssumableRoleWithSAMLIdentifier, name, component, opts...)
	if err != nil {
		return nil, err
	}

	opts = append(opts, pulumi.Parent(component))

	policyJSON := args.ProviderIDs.ToStringArrayOutput().ApplyT(func(ids []string) (string, error) {
		policyDocArgs := newIAMPolicyDocumentStatementConstructor("Allow", []string{"sts:AssumeRoleWithSAML"}).
			AddFederatedPrincipal(ids).
			AddCondition("StringEquals", "SAML:aud", []string{args.AWSSAMLEndpoint}).
			Build()

		policyDoc, err := iam.GetPolicyDocument(ctx, policyDocArgs)
		if err != nil {
			return "", err
		}

		return policyDoc.Json, nil
	}).(pulumi.StringOutput)

	role, err := utils.NewIAMRole(ctx, name, &utils.IAMRoleArgs{
		Role:                args.Role,
		MaxSessionDuration:  args.MaxSessionDuration,
		ForceDetachPolicies: args.ForceDetachPolicies,
		AssumeRolePolicy:    policyJSON,
		Tags:                args.Tags,
	}, opts...)
	if err != nil {
		return nil, err
	}

	component.Arn = role.Arn
	component.Name = role.Name
	component.Path = role.Path
	component.UniqueID = role.UniqueId

	return component, nil
}
