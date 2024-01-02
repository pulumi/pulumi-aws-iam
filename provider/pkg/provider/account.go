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

const AccountIdentifier = "aws-iam:index:Account"

type AccountPasswordPolicyArgs struct {
	// The number of days that an user password is valid.
	MaxAge int `pulumi:"maxAge"`

	// Minimum length to require for user passwords.
	MinimumLength int `pulumi:"minimumLength"`

	// The number of previous passwords that users are prevented from reusing.
	ReusePrevention int `pulumi:"reusePrevention"`

	// Whether to allow users to change their own password.
	AllowUsersToChange bool `pulumi:"allowUsersToChange"`

	// Whether users are prevented from setting a new password after their password
	// has expired (i.e. require administrator reset).
	HardExpiry bool `pulumi:"hardExpiry"`

	// Whether to require lowercase characters for user passwords.
	RequireLowercaseCharacters bool `pulumi:"requireLowercaseCharacters"`

	// Whether to require uppercase characters for user passwords.
	RequireUppercaseCharacters bool `pulumi:"requireUppercaseCharacters"`

	// Whether to require numbers for user passwords.
	RequireNumbers bool `pulumi:"requireNumbers"`

	// Whether to require symbols for user passwords.
	RequireSymbols bool `pulumi:"requireSymbols"`
}

type AccountArgs struct {
	// AWS IAM account alias for this account.
	AccountAlias string `pulumi:"accountAlias"`

	// Options to specify complexity requirements and mandatory rotation periods for
	// your IAM users' passwords. If left empty the default AWS password policy will be applied.
	PasswordPolicy AccountPasswordPolicyArgs `pulumi:"passwordPolicy"`
}

func (this *AccountArgs) Defaults() error {
	return nil
}

func (this *AccountArgs) Validate() error {
	return nil
}

type Account struct {
	pulumi.ResourceState

	// The AWS Account ID number of the account that owns or contains the calling entity.
	Id string `pulumi:"id"`

	// The AWS ARN associated with the calling entity.
	Arn string `pulumi:"arn"`

	// The unique identifier of the calling entity.
	UserId string `pulumi:"userId"`

	// Indicates whether passwords in the account expire. Returns true if max password age contains a value greater than 0. Returns false if it is 0 or not present.
	PasswordPolicyExpirePasswords pulumi.BoolOutput `pulumi:"passwordPolicyExpirePasswords"`
}

func NewIAMAccount(ctx *pulumi.Context, name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		args = &AccountArgs{}
	}

	// Return an error if MinimumLength is not valid.
	if args.PasswordPolicy.MinimumLength < 6 || args.PasswordPolicy.MinimumLength > 128 {
		return nil, fmt.Errorf("Invalid MinimumLength for PasswordPolicy provided for resource with name [%s]. Valid values are between 6 and 128.", name)
	}

	component := &Account{}
	err := ctx.RegisterComponentResource(AccountIdentifier, name, component, opts...)
	if err != nil {
		return nil, err
	}

	opts = append(opts, pulumi.Parent(component))

	account, err := aws.GetCallerIdentity(ctx)
	if err != nil {
		return nil, err
	}

	aliasName := fmt.Sprintf("%s-account-alias", name)
	_, err = iam.NewAccountAlias(ctx, aliasName, &iam.AccountAliasArgs{
		AccountAlias: pulumi.String(args.AccountAlias),
	}, opts...)
	if err != nil {
		return nil, err
	}

	passwordPolicyName := fmt.Sprintf("%s-password-policy", name)
	passwordPolicy, err := iam.NewAccountPasswordPolicy(ctx, passwordPolicyName, &iam.AccountPasswordPolicyArgs{
		MaxPasswordAge:             pulumi.Int(args.PasswordPolicy.MaxAge),
		MinimumPasswordLength:      pulumi.Int(args.PasswordPolicy.MinimumLength),
		AllowUsersToChangePassword: pulumi.Bool(args.PasswordPolicy.AllowUsersToChange),
		HardExpiry:                 pulumi.Bool(args.PasswordPolicy.HardExpiry),
		PasswordReusePrevention:    pulumi.Int(args.PasswordPolicy.ReusePrevention),
		RequireLowercaseCharacters: pulumi.Bool(args.PasswordPolicy.RequireLowercaseCharacters),
		RequireUppercaseCharacters: pulumi.Bool(args.PasswordPolicy.RequireUppercaseCharacters),
		RequireNumbers:             pulumi.Bool(args.PasswordPolicy.RequireNumbers),
		RequireSymbols:             pulumi.Bool(args.PasswordPolicy.RequireSymbols),
	}, opts...)
	if err != nil {
		return nil, err
	}

	component.Id = account.AccountId
	component.Arn = account.Arn
	component.UserId = account.UserId
	component.PasswordPolicyExpirePasswords = passwordPolicy.ExpirePasswords

	return component, nil
}
