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
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/eks"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const EKSRoleIdentifier = "aws-iam:index:EKSRole"

type EKSClusterServiceAccount struct {
	Name            pulumi.StringInput      `pulumi:"name"`
	ServiceAccounts pulumi.StringArrayInput `pulumi:"serviceAccounts"`
}

type EKSRoleArgs struct {
	// A map of tags to add.
	Tags pulumi.StringMapInput `pulumi:"tags"`

	// IAM role.
	Role utils.RoleArgs `pulumi:"role"`

	// Maximum CLI/API session duration in seconds between 3600 and 43200.
	MaxSessionDuration pulumi.IntInput `pulumi:"maxSessionDuration"`

	// Whether policies should be detached from this role when destroying.
	ForceDetachPolicies pulumi.BoolInput `pulumi:"forceDetachPolicies"`

	// EKS cluster and k8s ServiceAccount pairs. Each EKS cluster can have multiple k8s ServiceAccount. See README for details.
	ClusterServiceAccounts []EKSClusterServiceAccount `pulumi:"clusterServiceAccounts"`

	// ARNs of any policies to attach to the IAM role
	RolePolicyARNs []pulumi.StringInput `pulumi:"rolePolicyArns"`
}

type EKSRole struct {
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

func NewEKSRole(ctx *pulumi.Context, name string, args *EKSRoleArgs, opts ...pulumi.ResourceOption) (*EKSRole, error) {
	if args == nil {
		args = &EKSRoleArgs{}
	}

	component := &EKSRole{}
	err := ctx.RegisterComponentResource(EKSRoleIdentifier, name, component, opts...)
	if err != nil {
		return nil, err
	}

	opts = append(opts, pulumi.Parent(component))

	account, err := aws.GetCallerIdentity(ctx)
	if err != nil {
		return nil, err
	}

	currentPartition, err := aws.GetPartition(ctx, nil, nil)
	if err != nil {
		return nil, err
	}

	var policyDocStatements iam.GetPolicyDocumentStatementArray
	for _, sAccount := range args.ClusterServiceAccounts {
		cluster := eks.LookupClusterOutput(ctx, eks.LookupClusterOutputArgs{
			Name: sAccount.Name,
		})

		issuer := cluster.Identities().ApplyT(func(identities []eks.GetClusterIdentity) string {
			return strings.ReplaceAll(identities[0].Oidcs[0].Issuer, "https://", "")
		}).(pulumi.StringOutput)

		serviceAccounts := sAccount.ServiceAccounts.ToStringArrayOutput().ApplyT(func(accts []string) []string {
			var result []string
			for _, acct := range accts {
				result = append(result, fmt.Sprintf("system:serviceaccount:%s", acct))
			}
			return result
		}).(pulumi.StringArrayInput)

		principalIdentifier := pulumi.Sprintf("arn:%s:iam::%s:oidc-provider/%s", currentPartition.Partition, account.Id, issuer)

		policyDocStatements = append(policyDocStatements, iam.GetPolicyDocumentStatementArgs{
			Effect:  pulumi.StringPtr("Allow"),
			Actions: pulumi.ToStringArray([]string{"sts:AssumeRoleWithWebIdentity"}),
			Principals: iam.GetPolicyDocumentStatementPrincipalArray{
				iam.GetPolicyDocumentStatementPrincipalArgs{
					Type:        pulumi.String("Federated"),
					Identifiers: pulumi.ToStringArrayOutput([]pulumi.StringOutput{principalIdentifier}),
				},
			},
			Conditions: iam.GetPolicyDocumentStatementConditionArray{
				iam.GetPolicyDocumentStatementConditionArgs{
					Test:     pulumi.String("StringEquals"),
					Variable: pulumi.Sprintf("%s:sub", issuer),
					Values:   serviceAccounts,
				},
			},
		})
	}

	assumeRoleWithOIDC := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
		Statements: policyDocStatements,
	})

	assumeRoleWithOIDC.Json().ApplyT(func(data string) error {
		fmt.Println(data)
		return nil
	})

	role, err := utils.NewIAMRole(ctx, name, &utils.IAMRoleArgs{
		Role:                args.Role,
		AssumeRolePolicy:    assumeRoleWithOIDC.Json(),
		ForceDetachPolicies: args.ForceDetachPolicies,
		MaxSessionDuration:  args.MaxSessionDuration,
		Tags:                args.Tags,
	}, opts...)
	if err != nil {
		return nil, err
	}

	for _, arn := range args.RolePolicyARNs {
		_, err := iam.NewRolePolicyAttachment(ctx, fmt.Sprintf("%s-custom", name), &iam.RolePolicyAttachmentArgs{
			Role:      role.Name,
			PolicyArn: arn,
		}, opts...)
		if err != nil {
			return nil, err
		}
	}

	component.Arn = role.Arn
	component.Name = role.Name
	component.UniqueID = role.UniqueId
	component.Path = role.Path

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"arn":      role.Arn,
		"name":     role.Name,
		"uniqueId": role.UniqueId,
		"path":     role.Path,
	}); err != nil {
		return nil, err
	}

	return component, nil
}
