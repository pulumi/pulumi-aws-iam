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
	"strings"

	"github.com/pulumi/pulumi-aws-iam/pkg/utils"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const AssumableRoleWithOIDCIdentifier = "aws-iam:index:AssumableRoleWithOIDC"

type AssumableRoleWithOIDCArgs struct {
	// List of URLs of the OIDC Providers.
	ProviderURLs []string `pulumi:"providerUrls"`

	// The AWS account ID where the OIDC provider lives, leave empty to use the account for the AWS provider.
	AWSAccountID string `pulumi:"awsAccountId"`

	// A map of tags to add.
	Tags map[string]string `pulumi:"tags"`

	// IAM role.
	Role utils.RoleArgs `pulumi:"role"`

	// Maximum CLI/API session duration in seconds between 3600 and 43200.
	MaxSessionDuration int `pulumi:"maxSessionDuration"`

	// The fully qualified OIDC subjects to be added to the role policy.
	OIDCFullyQualifiedSubjects []string `pulumi:"oidcFullyQualifiedSubjects"`

	// The OIDC subject using wildcards to be added to the role policy.
	OIDCSubjectsWithWildcards []string `pulumi:"oidcSubjectsWithWildcards"`

	// The audience to be added to the role policy. Set to sts.amazonaws.com for cross-account assumable role. Leave empty otherwise.
	OIDCFullyQualifiedAudiences []string `pulumi:"oidcFullyQualifiedAudiences"`

	// Whether policies should be detached from this role when destroying.
	ForceDetachPolicies bool `pulumi:"forceDetachPolicies"`
}

type AssumableRoleWithOIDC struct {
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

func NewIAMAssumableRoleWithOIDC(ctx *pulumi.Context, name string, args *AssumableRoleWithOIDCArgs, opts ...pulumi.ResourceOption) (*AssumableRoleWithOIDC, error) {
	if args == nil {
		args = &AssumableRoleWithOIDCArgs{}
	}

	component := &AssumableRoleWithOIDC{}
	err := ctx.RegisterComponentResource(AssumableRoleWithOIDCIdentifier, name, component, opts...)
	if err != nil {
		return nil, err
	}

	opts = append(opts, pulumi.Parent(component))

	if args.AWSAccountID == "" {
		account, err := aws.GetCallerIdentity(ctx)
		if err != nil {
			return nil, err
		}
		args.AWSAccountID = account.AccountId
	}

	currentPartition, err := aws.GetPartition(ctx, nil, nil)
	if err != nil {
		return nil, err
	}

	for index, url := range args.ProviderURLs {
		args.ProviderURLs[index] = strings.ReplaceAll(url, "https://", "")
	}

	var policies []string
	for _, url := range args.ProviderURLs {
		effect := "Allow"
		principalIdentifier := fmt.Sprintf("arn:%s:iam::%s:oidc-provider/%s", currentPartition.Partition, args.AWSAccountID, url)

		var policyConditions []iam.GetPolicyDocumentStatementCondition
		if len(args.OIDCFullyQualifiedSubjects) > 0 {
			policyConditions = append(policyConditions, NewPolicyDocCondition("StringEquals", fmt.Sprintf("%s:sub", url), args.OIDCFullyQualifiedSubjects...))
		}

		if len(args.OIDCSubjectsWithWildcards) > 0 {
			policyConditions = append(policyConditions, NewPolicyDocCondition("StringLike", fmt.Sprintf("%s:sub", url), args.OIDCSubjectsWithWildcards...))
		}

		if len(args.OIDCFullyQualifiedAudiences) > 0 {
			policyConditions = append(policyConditions, NewPolicyDocCondition("StringLike", fmt.Sprintf("%s:aud", url), args.OIDCFullyQualifiedAudiences...))
		}

		policyDoc, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect:  &effect,
					Actions: []string{"sts:AssumeRoleWithWebIdentity"},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type:        "Federated",
							Identifiers: []string{principalIdentifier},
						},
					},
					Conditions: policyConditions,
				},
			},
		})
		if err != nil {
			return nil, err
		}

		policies = append(policies, policyDoc.Json)
	}

	role, err := utils.NewIAMRole(ctx, name, &utils.IAMRoleArgs{
		Role:                args.Role,
		MaxSessionDuration:  args.MaxSessionDuration,
		ForceDetachPolicies: args.ForceDetachPolicies,
		AssumeRolePolicy:    strings.Join(policies, ""),
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
