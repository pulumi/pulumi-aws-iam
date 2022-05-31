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

const (
	AssumableRoleIdentifier = "aws-iam:index:AssumableRole"

	AdminRolePolicyARN     = "arn:aws:iam::aws:policy/AdministratorAccess"
	PoweruserRolePolicyARN = "arn:aws:iam::aws:policy/PowerUserAccess"
	ReadonlyRolePolicyARN  = "arn:aws:iam::aws:policy/ReadOnlyAccess"
)

type AssumableRoleArgs struct {
	// Actions of STS.
	TrustedRoleActions []string `pulumi:"trustedRoleActions"`

	// ARNs of AWS entities who can assume these roles.
	TrustedRoleArns []string `pulumi:"trustedRoleArns"`

	// AWS Services that can assume these roles.
	TrustedRoleServices []string `pulumi:"trustedRoleServices"`

	// Max age of valid MFA (in seconds) for roles which require MFA.
	MFAAge int `pulumi:"mfaAge"`

	// Maximum CLI/API session duration in seconds between 3600 and 43200.
	MaxSessionDuration int `pulumi:"maxSessionDuration"`

	// IAM role.
	Role utils.RoleArgs `pulumi:"role"`

	// A map of tags to add.
	Tags map[string]string `pulumi:"tags"`

	// A custom role trust policy.
	CustomRoleTrustPolicy string `pulumi:"customRoleTrustPolicy"`

	// Whether to attach an admin policy to a role.
	AttachAdminPolicy bool `pulumi:"attachAdminPolicy"`

	// Whether to attach a poweruser policy to a role.
	AttachPoweruserPolicy bool `pulumi:"attachPoweruserPolicy"`

	// Whether to attach a readonly policy to a role.
	AttachReadonlyPolicy bool `pulumi:"attachReadonlyPolicy"`

	// Whether policies should be detached from this role when destroying.
	ForceDetachPolicies bool `pulumi:"forceDetachPolicies"`

	// STS ExternalId condition values to use with a role (when MFA is not required).
	RoleSTSExternalIDs []string `pulumi:"roleStsExternalIds"`
}

type AssumableRoleRoleOutput struct {
	// ARN of IAM role.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// Name of IAM role.
	Name pulumi.StringOutput `pulumi:"name"`

	// Path of IAM role.
	Path pulumi.StringPtrOutput `pulumi:"path"`

	// Unique ID of IAM role.
	UniqueID pulumi.StringOutput `pulumi:"uniqueId"`

	// Whether IAM role requires MFA.
	RequiresMFA bool `pulumi:"requiresMfa"`

	// STS ExternalId condition value to use with a role.
	STSExternalIDs []string `pulumi:"stsExternalIds"`
}

type AssumableRoleInstanceProfileOutput struct {
	// ARN of IAM instance profile.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// Name of IAM instance profile.
	Name pulumi.StringOutput `pulumi:"name"`

	// IAM Instance profile's ID.
	ID pulumi.StringOutput `pulumi:"id"`

	// Path of IAM instance profile.
	Path pulumi.StringPtrOutput `pulumi:"path"`
}

type AssumableRole struct {
	pulumi.ResourceState

	// IAM Role
	Role AssumableRoleRoleOutput `pulumi:"role"`

	// IAM instance profile.
	InstanceProfile AssumableRoleInstanceProfileOutput `pulumi:"instanceProfile"`
}

func NewAssumableRole(ctx *pulumi.Context, name string, args *AssumableRoleArgs, opts ...pulumi.ResourceOption) (*AssumableRole, error) {
	if args == nil {
		args = &AssumableRoleArgs{}
	}

	component := &AssumableRole{}
	err := ctx.RegisterComponentResource(AssumableRoleIdentifier, name, component, opts...)
	if err != nil {
		return nil, err
	}

	opts = append(opts, pulumi.Parent(component))

	if len(args.TrustedRoleActions) == 0 {
		args.TrustedRoleActions = append(args.TrustedRoleActions, "sts:AssumeRole")
	}

	policyDocumentStatementConditions := []iam.GetPolicyDocumentStatementCondition{
		NewPolicyDocCondition("StringEquals", "sts:ExternalId", args.RoleSTSExternalIDs...),
	}

	if args.Role.RequiresMFA {
		mfaAge := args.MFAAge
		if mfaAge == 0 {
			mfaAge = 86400
		}

		policyDocumentStatementConditions = []iam.GetPolicyDocumentStatementCondition{
			NewPolicyDocCondition("Bool", "aws:MultiFactorAuthPresent", "true"),
			NewPolicyDocCondition("NumericLessThan", "aws:MultiFactorAuthAge", fmt.Sprintf("%v", mfaAge)),
		}
	} else {

	}

	policyArgs := &iam.GetPolicyDocumentArgs{
		Statements: []iam.GetPolicyDocumentStatement{
			{
				Effect:     pulumi.StringRef("Allow"),
				Actions:    args.TrustedRoleActions,
				Conditions: policyDocumentStatementConditions,
				Principals: []iam.GetPolicyDocumentStatementPrincipal{
					{
						Type:        "AWS",
						Identifiers: args.TrustedRoleArns,
					},
					{
						Type:        "Service",
						Identifiers: args.TrustedRoleServices,
					},
				},
			},
		},
	}

	rolePolicy := args.CustomRoleTrustPolicy
	if rolePolicy == "" {
		assumeRolePolicy, err := utils.GetIAMPolicyDocument(ctx, policyArgs)
		if err != nil {
			return nil, err
		}

		rolePolicy = assumeRolePolicy.Json
	}

	if args.AttachAdminPolicy {
		args.Role.PolicyArns = append(args.Role.PolicyArns, AdminRolePolicyARN)
	}

	if args.AttachPoweruserPolicy {
		args.Role.PolicyArns = append(args.Role.PolicyArns, PoweruserRolePolicyARN)
	}

	if args.AttachReadonlyPolicy {
		args.Role.PolicyArns = append(args.Role.PolicyArns, ReadonlyRolePolicyARN)
	}

	role, err := utils.NewIAMRole(ctx, name, &utils.IAMRoleArgs{
		Role:                args.Role,
		MaxSessionDuration:  args.MaxSessionDuration,
		ForceDetachPolicies: args.ForceDetachPolicies,
		AssumeRolePolicy:    rolePolicy,
		Tags:                args.Tags,
	}, opts...)
	if err != nil {
		return nil, err
	}

	instanceProfileName := fmt.Sprintf("%s-instance-profile", name)
	instanceProfile, err := iam.NewInstanceProfile(ctx, instanceProfileName, &iam.InstanceProfileArgs{
		Name: pulumi.String(args.Role.Name),
		Path: pulumi.String(args.Role.Path),
		Role: role.Name,
		Tags: pulumi.ToStringMap(args.Tags),
	}, opts...)
	if err != nil {
		return nil, err
	}

	component.Role.Arn = role.Arn
	component.Role.Name = role.Name
	component.Role.Path = role.Path
	component.Role.UniqueID = role.UniqueId
	component.Role.RequiresMFA = args.Role.RequiresMFA
	component.Role.STSExternalIDs = args.RoleSTSExternalIDs
	component.InstanceProfile.Arn = instanceProfile.Arn
	component.InstanceProfile.ID = instanceProfile.UniqueId
	component.InstanceProfile.Name = instanceProfile.Name
	component.InstanceProfile.Path = instanceProfile.Path

	return component, nil
}
