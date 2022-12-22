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

	"github.com/pulumi/pulumi-aws-iam/pkg/utils"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RoleTypeIdentifier string

const AssumableRolesIdentifier = "aws-iam:index:AssumableRoles"

type AssumableRolesArgs struct {
	// ARNs of AWS entities who can assume these roles.
	TrustedRoleArns pulumi.StringArrayInput `pulumi:"trustedRoleArns"`

	// AWS Services that can assume these roles.
	TrustedRoleServices []string `pulumi:"trustedRoleServices"`

	// Max age of valid MFA (in seconds) for roles which require MFA.
	MFAAge int `pulumi:"mfaAge"`

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

type AssumableRoleOutput struct {
	// ARN of the IAM role.
	RoleARN pulumi.StringOutput `pulumi:"roleArn"`

	// Name of the IAM role.
	RoleName pulumi.StringOutput `pulumi:"roleName"`

	// Path of the IAM role.
	RolePath pulumi.StringPtrOutput `pulumi:"rolePath"`

	// Unique ID of IAM role.
	RoleUniqueID pulumi.StringOutput `pulumi:"roleUniqueId"`

	// Whether readonly IAM role requires MFA.
	RequiresMFA pulumi.BoolInput `pulumi:"requiresMfa"`
}

type AssumableRoles struct {
	pulumi.ResourceState

	// Admin role.
	Admin AssumableRoleOutput `pulumi:"admin"`

	// Poweruser role.
	Poweruser AssumableRoleOutput `pulumi:"poweruser"`

	// Readonly role.
	Readonly AssumableRoleOutput `pulumi:"readonly"`
}

func newAssumableRolePolicyDocumentArgs(trustedRoleARNs []string, trustedRoleServices []string, requiresMFA bool, mfaAge int) *iam.GetPolicyDocumentArgs {
	var conditions []iam.GetPolicyDocumentStatementCondition
	if requiresMFA {
		if mfaAge == 0 {
			mfaAge = 86400
		}

		conditions = append(conditions, []iam.GetPolicyDocumentStatementCondition{
			NewPolicyDocCondition("Bool", "aws:MultiFactorAuthPresent", "true"),
			NewPolicyDocCondition("NumericLessThan", "aws:MultiFactorAuthAge", fmt.Sprintf("%v", mfaAge)),
		}...)
	}

	return &iam.GetPolicyDocumentArgs{
		Statements: []iam.GetPolicyDocumentStatement{
			{
				Effect:  pulumi.StringRef("Allow"),
				Actions: []string{"sts:AssumeRole"},
				Principals: []iam.GetPolicyDocumentStatementPrincipal{
					{
						Type:        "AWS",
						Identifiers: trustedRoleARNs,
					},
					{
						Type:        "Federated",
						Identifiers: trustedRoleServices,
					},
				},
				Conditions: conditions,
			},
		},
	}
}

func NewAssumableRoles(ctx *pulumi.Context, name string, args *AssumableRolesArgs, opts ...pulumi.ResourceOption) (*AssumableRoles, error) {
	if args == nil {
		args = &AssumableRolesArgs{}
	}

	component := &AssumableRoles{}
	err := ctx.RegisterComponentResource(AssumableRolesIdentifier, name, component, opts...)
	if err != nil {
		return nil, err
	}

	opts = append(opts, pulumi.Parent(component))

	assumeRoleJSON := args.TrustedRoleArns.ToStringArrayOutput().ApplyT(func(arns []string) (string, error) {
		assumeRoleArgs := newAssumableRolePolicyDocumentArgs(arns, args.TrustedRoleServices, false, 0)

		assumeRole, err := utils.GetIAMPolicyDocument(ctx, assumeRoleArgs)
		if err != nil {
			return "", err
		}

		return assumeRole.Json, nil
	}).(pulumi.StringOutput)

	assumeRoleWithMFAJSON := args.TrustedRoleArns.ToStringArrayOutput().ApplyT(func(arns []string) (string, error) {
		assumeRoleWithMFAArgs := newAssumableRolePolicyDocumentArgs(arns, args.TrustedRoleServices, true, args.MFAAge)

		assumeRoleMFA, err := utils.GetIAMPolicyDocument(ctx, assumeRoleWithMFAArgs)
		if err != nil {
			return "", err
		}

		return assumeRoleMFA.Json, nil
	}).(pulumi.StringOutput)

	args.Admin.Name = setDefaultStringPtr(args.Admin.Name, "admin")
	if args.Admin.RequiresMFA == nil {
		args.Admin.RequiresMFA = pulumi.Bool(true)
	}

	args.Poweruser.Name = setDefaultStringPtr(args.Poweruser.Name, "poweruser")
	if args.Poweruser.RequiresMFA == nil {
		args.Poweruser.RequiresMFA = pulumi.Bool(true)
	}

	args.Readonly.Name = setDefaultStringPtr(args.Readonly.Name, "readonly")
	if args.Readonly.RequiresMFA == nil {
		args.Readonly.RequiresMFA = pulumi.Bool(true)
	}

	roleOutput, err := utils.NewAssumableRoles(ctx, name, &utils.IAMAssumableRolesArgs{
		MaxSessionDuration:  args.MaxSessionDuration,
		ForceDetachPolicies: args.ForceDetachPolicies,
		AssumeRolePolicy:    assumeRoleJSON,
		AssumeRoleWithMFA:   assumeRoleWithMFAJSON,
		Admin:               args.Admin,
		Poweruser:           args.Poweruser,
		Readonly:            args.Readonly,
	}, opts...)
	if err != nil {
		return nil, err
	}

	component.Admin = createAssumableRoleOutput(roleOutput[utils.AdminRoleType], args.Admin.RequiresMFA)
	component.Poweruser = createAssumableRoleOutput(roleOutput[utils.PoweruserRoleType], args.Poweruser.RequiresMFA)
	component.Readonly = createAssumableRoleOutput(roleOutput[utils.ReadonlyRoleType], args.Readonly.RequiresMFA)

	return component, nil
}
