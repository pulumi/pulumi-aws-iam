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
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const AssumableRolesWithSAMLIdentifier = "aws-iam:index:AssumableRolesWithSAML"

type AssumableRolesWithSAMLArgs struct {
	// List of SAML Provider IDs.
	ProviderIDs pulumi.StringArrayInput `pulumi:"providerIds"`

	// AWS SAML Endpoint.
	AWSSAMLEndpoint string `pulumi:"awsSamlEndpoint"`

	// Maximum CLI/API session duration in seconds between 3600 and 43200.
	MaxSessionDuration pulumi.IntInput `pulumi:"maxSessionDuration"`

	// Whether policies should be detached from this role when destroying.
	ForceDetachPolicies pulumi.BoolInput `pulumi:"forceDetachPolicies"`

	// IAM role with admin access.
	Admin utils.RoleArgs `pulumi:"admin"`

	// IAM role with poweruser access.
	Poweruser utils.RoleArgs `pulumi:"poweruser"`

	// IAM role with readonly access.
	Readonly utils.RoleArgs `pulumi:"readonly"`
}

type AssumableRolesWithSAML struct {
	pulumi.ResourceState

	// Admin role.
	Admin AssumableRoleOutput `pulumi:"admin"`

	// Poweruser role.
	Poweruser AssumableRoleOutput `pulumi:"poweruser"`

	// Readonly role.
	Readonly AssumableRoleOutput `pulumi:"readonly"`
}

func NewAssumableRolesWithSAML(ctx *pulumi.Context, name string, args *AssumableRolesWithSAMLArgs, opts ...pulumi.ResourceOption) (*AssumableRolesWithSAML, error) {
	if args == nil {
		args = &AssumableRolesWithSAMLArgs{}
	}

	component := &AssumableRolesWithSAML{}
	err := ctx.RegisterComponentResource(AssumableRolesWithSAMLIdentifier, name, component, opts...)
	if err != nil {
		return nil, err
	}

	opts = append(opts, pulumi.Parent(component))

	assumeRoleJSON := args.ProviderIDs.ToStringArrayOutput().ApplyT(func(ids []string) (string, error) {
		assumableRoleWithSAMLArgs := newIAMPolicyDocumentStatementConstructor("Allow", []string{"sts:AssumeRoleWithSAML"}).
			AddFederatedPrincipal(ids).
			AddCondition("StringEquals", "SAML:aud", []string{args.AWSSAMLEndpoint}).
			Build()

		assumeRole, err := utils.GetIAMPolicyDocument(ctx, assumableRoleWithSAMLArgs)
		if err != nil {
			return "", err
		}

		return assumeRole.Json, nil
	}).(pulumi.StringOutput)

	roleOutput, err := utils.NewAssumableRoles(ctx, name, &utils.IAMAssumableRolesArgs{
		MaxSessionDuration:  args.MaxSessionDuration,
		ForceDetachPolicies: args.ForceDetachPolicies,
		AssumeRolePolicy:    assumeRoleJSON,
		Admin:               args.Admin,
		Poweruser:           args.Poweruser,
		Readonly:            args.Readonly,
	}, opts...)
	if err != nil {
		return nil, err
	}

	component.Admin = createAssumableRoleOutput(roleOutput[utils.AdminRoleType], pulumi.Bool(false))
	component.Poweruser = createAssumableRoleOutput(roleOutput[utils.PoweruserRoleType], pulumi.Bool(false))
	component.Readonly = createAssumableRoleOutput(roleOutput[utils.PoweruserRoleType], pulumi.Bool(false))

	return component, nil
}
