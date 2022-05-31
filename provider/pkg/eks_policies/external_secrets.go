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
	SSMParameterARNs []string `pulumi:"ssmParameterArns"`

	// List of Secrets Manager ARNs that contain secrets to mount using External Secrets.
	SecretsMangerARNs []string `pulumi:"secretsManagerArns"`
}

func AttachExternalSecretsPolicy(policyBuilder *EKSRoleBuilder, args ExternalSecretsPolicyArgs) error {
	if len(args.SSMParameterARNs) == 0 {
		args.SSMParameterARNs = append(args.SSMParameterARNs, externalSecretsDefaultSSMParameterARN)
	}

	if len(args.SecretsMangerARNs) == 0 {
		args.SecretsMangerARNs = append(args.SecretsMangerARNs, externalSecretsDefaultSecretsManagerARN)
	}

	policyStatements := []iam.GetPolicyDocumentStatement{
		{
			Actions:   []string{"ssm:GetParameter"},
			Resources: args.SSMParameterARNs,
		},
		{
			Actions: []string{
				"secretsmanager:GetResourcePolicy",
				"secretsmanager:GetSecretValue",
				"secretsmanager:DescribeSecret",
				"secretsmanager:ListSecretVersionIds",
			},
			Resources: args.SecretsMangerARNs,
		},
	}

	return policyBuilder.CreatePolicyWithAttachment(externalSecretsNamePrefix, externalSecretsDescription, policyStatements)
}
