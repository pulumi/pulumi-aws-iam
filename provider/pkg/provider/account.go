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
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const AccountIdentifier = "aws-iam:index:Account"

type AccountPasswordPolicy struct {
	// The number of days that an user password is valid.
	MaxAge int `pulumi:"maxAge,optional"`

	// Minimum length to require for user passwords.
	MinimumLength int `pulumi:"minimumLength,optional"`

	// The number of previous passwords that users are prevented from reusing.
	ReusePrevention int `pulumi:"reusePrevention,optional"`

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

func (c *AccountPasswordPolicy) Annotate(a infer.Annotator) {
	a.Describe(&c, "Options to specify complexity requirements and mandatory rotation periods for your IAM users' passwords.")
	a.Describe(&c.MaxAge, "The number of days that an user password is valid. If not set or a value of `0` is provided, then\npasswords will not expire.\n")
	a.Describe(&c.MinimumLength, "Minimum length to require for user passwords. Defaults to `8` if not set or\nthe provided value is invalid. Valid values are between 6 and 128.\n")
	a.Describe(&c.AllowUsersToChange, "Whether to allow users to change their own password.\n")
	a.Describe(&c.HardExpiry, "Whether users are prevented from setting a new password after their password has\nexpired (i.e. require administrator reset).\n")
	a.Describe(&c.ReusePrevention, "The number of previous passwords that users are prevented from reusing. If not set or a\nvalue of `0` is provided, no reuse prevention policy will be used.\n")
	a.Describe(&c.RequireLowercaseCharacters, "Whether to require lowercase characters for user passwords.\n")
	a.Describe(&c.ReusePrevention, "Whether to require uppercase characters for user passwords.\n")
	a.Describe(&c.RequireNumbers, "Whether to require numbers for user passwords.\n")
	a.Describe(&c.RequireSymbols, "Whether to require symbols for user passwords.\n")
}

type AccountArgs struct {
	// AWS IAM account alias for this account.
	AccountAlias pulumi.StringInput `pulumi:"accountAlias"`

	// Options to specify complexity requirements and mandatory rotation periods for
	// your IAM users' passwords. If left empty the default AWS password policy will be applied.
	PasswordPolicy AccountPasswordPolicy `pulumi:"passwordPolicy"`
}

func (c *AccountArgs) Annotate(a infer.Annotator) {
	a.Describe(&c.AccountAlias, "AWS IAM account alias for this account.")
	a.Describe(&c.PasswordPolicy, "Options to specify complexity requirements and mandatory rotation periods for your IAM users' passwords. If\n"+
		"left empty the default AWS password policy will be applied.\n")
}

func (this *AccountArgs) Defaults() error {
	return nil
}

func (this *AccountArgs) Validate() error {
	return nil
}

type Account struct{}

func (c *Account) Annotate(a infer.Annotator) {
	a.Describe(&c, `This resource helps you manage an Iam Account's Alias and Password Policy. If your IAM Account Alias was previously
gset (either via the AWS console or when AWS created your Account) you will see an error like the below:

`+"```"+`
    * Aws_iam_account_alias.this: Error creating account alias with name my-account-alias
`+"```"+`

If you want to manage you Alias using Pulumi you will need to import this resource.

{{% examples %}}
## Example Usage

{{% example %}}
## Account

`+"```"+`typescript
import * as iam from \"@pulumi/aws-iam\";

export const account = new iam.Account(\"account\", {
    accountAlias: \"cool-alias\",
    passwordPolicy: {
        minimumLength: 37,
        requireNumbers: false,
        allowUsersToChange: true,
        hardExpiry: true,
        requireSymbols: true,
        requireLowercaseCharacters: true,
        requireUppercaseCharacters: true,
    },
});
`+"```"+`

`+"```"+`python
import pulumi
import pulumi_aws_iam as iam

account = iam.Account(
    'account',
    account_alias='cool-alias',
    password_policy=iam.AccountPasswordPolicyArgs(
        minimum_length=37,
        require_numbers=False,
        allow_users_to_change=True,
        hard_expiry=True,
        require_symbols=True,
        require_lowercase_characters=True,
        require_uppercase_characters=True,
    )
)

pulumi.export('account', account)
`+"```"+`

`+"```"+`go
package main

import (
    iam \"github.com/pulumi/pulumi-aws-iam/sdk/go/aws-iam\"
    \"github.com/pulumi/pulumi/sdk/v3/go/pulumi\"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        account, err := iam.NewAccount(ctx, \"account\", &iam.AccountArgs{
            AccountAlias: pulumi.String(\"cool-alias\"),
            PasswordPolicy: iam.AccountPasswordPolicyArgs{
                MinimumLength:              pulumi.IntPtr(37),
                RequireNumbers:             pulumi.Bool(false),
                AllowUsersToChange:         pulumi.Bool(true),
                HardExpiry:                 pulumi.Bool(true),
                RequireSymbols:             pulumi.Bool(true),
                RequireLowercaseCharacters: pulumi.Bool(true),
                RequireUppercaseCharacters: pulumi.Bool(true),
            },
        })
        if err != nil {
            return err
        }

        ctx.Export(\"account\", account)

        return nil
    })
}
`+"```"+`

`+"```"+`csharp
using Pulumi;
using Pulumi.AwsIam;
using Pulumi.AwsIam.Inputs;

class MyStack : Stack
{
    public MyStack()
    {
        var account = new Account(\"account\", new AccountArgs
        {
            AccountAlias = \"cool-alias\",
            PasswordPolicy=new AccountPasswordPolicyArgs
            {
                MinimumLength = 37,
                RequireNumbers = false,
                AllowUsersToChange = true,
                HardExpiry = true,
                RequireSymbols = true,
                RequireLowercaseCharacters = true,
                RequireUppercaseCharacters = true,
            }

        });

        this.Account = Output.Create<Account>(account);
    }

    [Output]
    public Output<Account> Account { get; set; }
}
`+"```"+`

`+"```"+`yaml
name: awsiam-yaml
runtime: yaml
resources:
    account:
        type: \"aws-iam:index:Account\"
        properties:
            accountAlias: \"cool-alias\"
            passwordPolicy:
                minimumLength: 37
                requireNumbers: false
                allowUsersToChange: true
                hardExpiry: true
                requireSymbols: true
                requireLowercaseCharacters: true
                requireUppercaseCharacters: true
outputs:
    account: ${account}
`+"```"+`
{{ /example }}

{{% examples %}}
`)
}

func (c *Account) Construct(ctx *pulumi.Context, name string, typ string, args AccountArgs, opts pulumi.ResourceOption) (*AccountState, error) {
	// Return an error if MinimumLength is not valid.
	if args.PasswordPolicy.MinimumLength < 6 || args.PasswordPolicy.MinimumLength > 128 {
		return nil, fmt.Errorf("Invalid MinimumLength for PasswordPolicy provided for resource with name [%s]. Valid values are between 6 and 128.", name)
	}

	component := &AccountState{}
	err := ctx.RegisterComponentResource(AccountIdentifier, name, component, opts)
	if err != nil {
		return nil, err
	}

	opts = pulumi.Composite(opts, pulumi.Parent(component))

	account, err := aws.GetCallerIdentity(ctx)
	if err != nil {
		return nil, err
	}

	aliasName := fmt.Sprintf("%s-account-alias", name)
	_, err = iam.NewAccountAlias(ctx, aliasName, &iam.AccountAliasArgs{
		AccountAlias: args.AccountAlias,
	}, opts)
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
	}, opts)
	if err != nil {
		return nil, err
	}

	component.Id = pulumi.String(account.AccountId)
	component.Arn = pulumi.String(account.Arn)
	component.UserId = pulumi.String(account.UserId)
	component.PasswordPolicyExpirePasswords = passwordPolicy.ExpirePasswords

	return component, nil
}

type AccountState struct {
	pulumi.ResourceState

	// The AWS Account ID number of the account that owns or contains the calling entity.
	Id pulumi.String `pulumi:"id"`

	// The AWS ARN associated with the calling entity.
	Arn pulumi.String `pulumi:"arn"`

	// The unique identifier of the calling entity.
	UserId pulumi.String `pulumi:"userId"`

	// Indicates whether passwords in the account expire. Returns true if max password age contains a value greater than 0. Returns false if it is 0 or not present.
	PasswordPolicyExpirePasswords pulumi.BoolOutput `pulumi:"passwordPolicyExpirePasswords"`
}

func (s *AccountState) Annotate(a infer.Annotator) {
	a.Describe(&s.Id, "The AWS Account ID number of the account that owns or contains the calling entity.\n")
	a.Describe(&s.Arn, "The AWS ARN associated with the calling entity.\n")
	a.Describe(&s.UserId, "The unique identifier of the calling entity.\n")
	a.Describe(&s.PasswordPolicyExpirePasswords, "Indicates whether passwords in the account expire. "+
		"Returns true if max password age contains a value greater than 0. Returns false if it is 0 or not present.\n")
}
