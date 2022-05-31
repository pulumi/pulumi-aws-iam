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

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const ReadOnlyPolicyIdentifier = "aws-iam:index:ReadOnlyPolicy"

type ReadOnlyPolicyArgs struct {
	// The name of the policy.
	Name string `pulumi:"name"`

	// The path of the policy in IAM.
	Path string `pulumi:"path"`

	// The description of the policy.
	Description string `pulumi:"description"`

	// A map of tags to add.
	Tags map[string]string `pulumi:"tags"`

	// List of services to allow Get/List/Describe/View options. Service name should be the
	// same as corresponding service IAM prefix. See what it is for each service here
	// https://docs.aws.amazon.com/service-authorization/latest/reference/reference_policies_actions-resources-contextkeys.html.
	AllowedServices []string `pulumi:"allowedServices"`

	// JSON policy document if you want to add custom actions.
	AdditionalPolicyJSON string `pulumi:"additionalPolicyJson"`

	// Allows StartQuery/StopQuery/FilterLogEvents CloudWatch actions.
	AllowCloudwatchLogsQuery bool `pulumi:"allowCloudwatchLogsQuery"`

	// Allows GetCallerIdentity/GetSessionToken/GetAccessKeyInfo sts actions.
	AllowPredefinedStsActions bool `pulumi:"allowPredefinedStsActions"`

	// Allows List/Get/Describe/View actions for services used when browsing AWS console (e.g. resource-groups, tag, health services).
	AllowWebConsoleServices bool `pulumi:"allowWebConsoleServices"`

	// List of web console services to allow.
	WebConsoleServices []string `pulumi:"webConsoleServices"`
}

type ReadOnlyPolicy struct {
	pulumi.ResourceState

	// Policy document as json. Useful if you need document but do not want to create IAM
	// policy itself. For example for SSO Permission Set inline policies.
	PolicyJSON pulumi.StringOutput `pulumi:"policyJson"`

	// The policy's ID.
	ID pulumi.StringOutput `pulumi:"id"`

	// The name of the policy.
	Name pulumi.StringOutput `pulumi:"name"`

	// The ARN assigned by AWS to this policy.
	ARN pulumi.StringOutput `pulumi:"arn"`

	// The description of the policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`

	// The path of the policy in IAM.
	Path pulumi.StringPtrOutput `pulumi:"path"`

	// The policy document.
	Policy pulumi.StringOutput `pulumi:"policy"`
}

func NewReadOnlyPolicy(ctx *pulumi.Context, name string, args *ReadOnlyPolicyArgs, opts ...pulumi.ResourceOption) (*ReadOnlyPolicy, error) {
	if args == nil {
		args = &ReadOnlyPolicyArgs{}
	}

	component := &ReadOnlyPolicy{}
	err := ctx.RegisterComponentResource(ReadOnlyPolicyIdentifier, name, component, opts...)
	if err != nil {
		return nil, err
	}

	opts = append(opts, pulumi.Parent(component))

	var policyDocStatements []iam.GetPolicyDocumentStatement
	for _, service := range args.AllowedServices {
		sid := strings.ReplaceAll(service, "-", "")
		policyDocStatements = append(policyDocStatements, iam.GetPolicyDocumentStatement{
			Sid:       &sid,
			Resources: []string{"*"},
			Actions: []string{
				fmt.Sprintf("%s:List*", service),
				fmt.Sprintf("%s:Get*", service),
				fmt.Sprintf("%s:Describe*", service),
				fmt.Sprintf("%s:View*", service),
			},
		})
	}

	if args.AllowWebConsoleServices {
		if len(args.WebConsoleServices) == 0 {
			args.WebConsoleServices = append(args.WebConsoleServices, "resource-groups", "tag", "health", "ce")
		}

		for _, service := range args.WebConsoleServices {
			sid := strings.ReplaceAll(service, "-", "")
			policyDocStatements = append(policyDocStatements, iam.GetPolicyDocumentStatement{
				Sid:       &sid,
				Resources: []string{"*"},
				Actions: []string{
					fmt.Sprintf("%s:List*", service),
					fmt.Sprintf("%s:Get*", service),
					fmt.Sprintf("%s:Describe*", service),
					fmt.Sprintf("%s:View*", service),
				},
			})
		}
	}

	if args.AllowPredefinedStsActions {
		policyDocStatements = append(policyDocStatements, iam.GetPolicyDocumentStatement{
			Sid:       pulumi.StringRef("STS"),
			Resources: []string{"*"},
			Actions:   []string{"sts:GetAccessKeyInfo", "sts:GetCallerIdentity", "sts:GetSessionToken"},
		})
	}

	if args.AllowCloudwatchLogsQuery {
		policyDocStatements = append(policyDocStatements, iam.GetPolicyDocumentStatement{
			Sid:       pulumi.StringRef("AllowLogsQuery"),
			Resources: []string{"*"},
			Actions:   []string{"logs:StartQuery", "logs:StopQuery", "logs:FilterLogEvents"},
		})
	}

	policyDoc, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
		Statements: policyDocStatements,
	})
	if err != nil {
		return nil, err
	}

	policy, err := iam.NewPolicy(ctx, name, &iam.PolicyArgs{
		Name:        pulumi.String(args.Name),
		Path:        pulumi.String(args.Path),
		Description: pulumi.String(args.Description),
		Policy:      pulumi.String(policyDoc.Json),
	}, opts...)
	if err != nil {
		return nil, err
	}

	component.PolicyJSON = pulumi.Sprintf("%s", policyDoc.Json)
	component.ID = policy.ID().ToStringOutput()
	component.Name = policy.Name
	component.ARN = policy.Arn
	component.Description = policy.Description
	component.Path = policy.Path
	component.Policy = policy.Policy

	return component, nil
}
