// Copyright 2016-2021, Pulumi Corporation.
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

const (
	externalSecretsNamePrefix  = "External_Secrets_Policy-"
	externalSecretsDescription = "Provides permissions to for External Secrets to retrieve secrets from AWS SSM and AWS Secrets Manager"

	externalSecretsDefaultSSMParameterARN   = "arn:aws:ssm:*:*:parameter/*"
	externalSecretsDefaultSecretsManagerARN = "arn:aws:secretsmanager:*:*:secret:*"
)

type ExternalSecretsPolicyArgs struct {
	// Determines whether to attach the External Secrets policy to the role.
	Attach bool `pulumi:"attach"`

	// List of Systems Manager Parameter ARNs that contain secrets to mount using External Secrets.
	SSMParameterARNs pulumi.StringArrayInput `pulumi:"ssmParameterArns"`

	// List of Secrets Manager ARNs that contain secrets to mount using External Secrets.
	SecretsMangerARNs pulumi.StringArrayInput `pulumi:"secretsManagerArns"`
}

func AttachExternalSecretsPolicy(ctx *pulumi.Context, policyBuilder *EKSRoleBuilder, args ExternalSecretsPolicyArgs) error {
	policyJSON := pulumi.All(args.SSMParameterARNs, args.SecretsMangerARNs).ApplyT(func(x []interface{}) (string, error) {
		ssmParameterARNs := x[0].([]string)
		secretsManagerARNs := x[1].([]string)

		if len(ssmParameterARNs) == 0 {
			ssmParameterARNs = append(ssmParameterARNs, externalSecretsDefaultSSMParameterARN)
		}

		if len(secretsManagerARNs) == 0 {
			secretsManagerARNs = append(secretsManagerARNs, externalSecretsDefaultSecretsManagerARN)
		}

		policyStatements := []iam.GetPolicyDocumentStatement{
			{
				Actions:   []string{"ssm:GetParameter"},
				Resources: ssmParameterARNs,
			},
			{
				Actions: []string{
					"secretsmanager:GetResourcePolicy",
					"secretsmanager:GetSecretValue",
					"secretsmanager:DescribeSecret",
					"secretsmanager:ListSecretVersionIds",
				},
				Resources: secretsManagerARNs,
			},
		}

		policyDoc, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: policyStatements,
		})
		if err != nil {
			return "", err
		}

		return policyDoc.Json, err
	}).(pulumi.StringOutput)

	return policyBuilder.CreatePolicyWithAttachment(externalSecretsNamePrefix, externalSecretsDescription, policyJSON)
}
