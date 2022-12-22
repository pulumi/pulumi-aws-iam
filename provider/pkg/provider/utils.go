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

func setDefaultStringPtr(output pulumi.StringPtrInput, value string) pulumi.StringPtrOutput {
	if output == nil {
		output = pulumi.StringPtr("").ToStringPtrOutput()
	}

	return output.ToStringPtrOutput().ApplyT(func(v *string) *string {
		if v == nil || *v == "" {
			v = &value
		}
		return v
	}).(pulumi.StringPtrOutput)
}

func createAssumableRoleOutput(role *iam.Role, requiresMFA pulumi.BoolInput) AssumableRoleOutput {
	return AssumableRoleOutput{
		RoleARN:      role.Arn,
		RoleName:     role.Name,
		RolePath:     role.Path,
		RoleUniqueID: role.UniqueId,
		RequiresMFA:  requiresMFA,
	}
}

// IAM Policy Document
type IAMPolicyDocumentEffect string
type IAMPolicyDocumentPrincipalType string

const (
	iamPolicyDocumentAWSPrincipal       IAMPolicyDocumentPrincipalType = "AWS"
	iamPolicyDocumentFederatedPrincipal IAMPolicyDocumentPrincipalType = "Federated"
	iamPolicyDocumentServicePrincipal   IAMPolicyDocumentPrincipalType = "Service"
)

type IAMPolicyDocumentStatementConstructor struct {
	Effect     *string
	Actions    []string
	Principals []iam.GetPolicyDocumentStatementPrincipal
	Conditions []iam.GetPolicyDocumentStatementCondition
	Resources  []string
}

func newIAMPolicyDocumentStatementConstructor(effectName IAMPolicyDocumentEffect, actions []string) *IAMPolicyDocumentStatementConstructor {
	effect := string(effectName)

	return &IAMPolicyDocumentStatementConstructor{
		Effect:  &effect,
		Actions: actions,
	}
}

func (i *IAMPolicyDocumentStatementConstructor) AddAWSPrincipal(identifiers []string) *IAMPolicyDocumentStatementConstructor {
	if len(identifiers) == 0 {
		return i
	}

	typ := string(iamPolicyDocumentAWSPrincipal)

	i.Principals = append(i.Principals, iam.GetPolicyDocumentStatementPrincipal{
		Type:        typ,
		Identifiers: identifiers,
	})

	return i
}

func (i *IAMPolicyDocumentStatementConstructor) AddFederatedPrincipal(identifiers []string) *IAMPolicyDocumentStatementConstructor {
	if len(identifiers) == 0 {
		return i
	}

	typ := string(iamPolicyDocumentFederatedPrincipal)

	i.Principals = append(i.Principals, iam.GetPolicyDocumentStatementPrincipal{
		Type:        typ,
		Identifiers: identifiers,
	})

	return i
}

func (i *IAMPolicyDocumentStatementConstructor) AddServicePrincipal(identifiers []string) *IAMPolicyDocumentStatementConstructor {
	if len(identifiers) == 0 {
		return i
	}

	typ := string(iamPolicyDocumentServicePrincipal)

	i.Principals = append(i.Principals, iam.GetPolicyDocumentStatementPrincipal{
		Type:        typ,
		Identifiers: identifiers,
	})

	return i
}

func (i *IAMPolicyDocumentStatementConstructor) AddCondition(test, variable string, values []string) *IAMPolicyDocumentStatementConstructor {
	i.Conditions = append(i.Conditions, iam.GetPolicyDocumentStatementCondition{
		Test:     test,
		Variable: variable,
		Values:   values,
	})

	return i
}

func (i *IAMPolicyDocumentStatementConstructor) AddResources(resources []string) *IAMPolicyDocumentStatementConstructor {
	i.Resources = append(i.Resources, resources...)
	return i
}

func (i *IAMPolicyDocumentStatementConstructor) Build() *iam.GetPolicyDocumentArgs {
	return &iam.GetPolicyDocumentArgs{
		Statements: []iam.GetPolicyDocumentStatement{
			{
				Effect:     i.Effect,
				Actions:    i.Actions,
				Principals: i.Principals,
				Conditions: i.Conditions,
				Resources:  i.Resources,
			},
		},
	}
}

func NewPolicyDocCondition(test, variable string, values ...string) iam.GetPolicyDocumentStatementCondition {
	return iam.GetPolicyDocumentStatementCondition{
		Test:     test,
		Variable: variable,
		Values:   values,
	}
}
