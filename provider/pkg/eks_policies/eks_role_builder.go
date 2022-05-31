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

package eks_policies

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func NewPolicyDocCondition(test, variable string, values ...string) iam.GetPolicyDocumentStatementCondition {
	return iam.GetPolicyDocumentStatementCondition{
		Test:     test,
		Variable: variable,
		Values:   values,
	}
}

type EKSRoleBuilder struct {
	Ctx                 *pulumi.Context
	AWSAccountID        string
	AWSCurrentPartition string
	DNSSuffix           string
	Name                string
	ResourceOpts        []pulumi.ResourceOption
	Role                *iam.Role

	BaseNamePrefix string
	Path           string
	Tags           map[string]string
}

func CreateNewRoleBuilder(ctx *pulumi.Context, role *iam.Role, name, baseNamePrefix,
	path string, tags map[string]string, opts ...pulumi.ResourceOption) *EKSRoleBuilder {
	return &EKSRoleBuilder{
		Ctx:            ctx,
		Name:           name,
		ResourceOpts:   opts,
		Role:           role,
		BaseNamePrefix: baseNamePrefix,
		Path:           path,
		Tags:           tags,
	}
}

func (r *EKSRoleBuilder) CreatePolicyWithAttachment(namePrefix, description string, statements []iam.GetPolicyDocumentStatement) error {
	policyDoc, err := iam.GetPolicyDocument(r.Ctx, &iam.GetPolicyDocumentArgs{
		Statements: statements,
	})
	if err != nil {
		return err
	}

	policy, err := iam.NewPolicy(r.Ctx, r.Name, &iam.PolicyArgs{
		NamePrefix:  pulumi.Sprintf("%s%s", r.BaseNamePrefix, namePrefix),
		Path:        pulumi.String(r.Path),
		Description: pulumi.String(description),
		Policy:      pulumi.String(policyDoc.Json),
		Tags:        pulumi.ToStringMap(r.Tags),
	}, r.ResourceOpts...)
	if err != nil {
		return err
	}

	_, err = iam.NewRolePolicyAttachment(r.Ctx, r.Name, &iam.RolePolicyAttachmentArgs{
		Role:      r.Role.Name,
		PolicyArn: policy.Arn,
	}, r.ResourceOpts...)
	return err
}
