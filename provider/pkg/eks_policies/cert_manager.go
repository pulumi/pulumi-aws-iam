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
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
)

const (
	eksPolicyNamePrefix  = "Cert_Manager_Policy-"
	eksPolicyDescription = "Cert Manager policy to allow management of Route53 hosted zone records"

	defaultHostedZoneARN = "arn:aws:route53:::hostedzone/*"
)

type CertManagerPolicyArgs struct {
	// Determines whether to attach the Cert Manager IAM policy to the role.
	Attach bool `pulumi:"attach"`

	// Route53 hosted zone ARNs to allow Cert manager to manage records.
	HostedZoneARNs []string `pulumi:"hostedZoneArns"`
}

func AttachCertManagerPolicy(policyBuilder *EKSRoleBuilder, partition string, args CertManagerPolicyArgs) error {
	if len(args.HostedZoneARNs) == 0 {
		args.HostedZoneARNs = append(args.HostedZoneARNs, defaultHostedZoneARN)
	}

	policyStatements := []iam.GetPolicyDocumentStatement{
		{
			Actions:   []string{"route53:GetChange"},
			Resources: []string{fmt.Sprintf("arn:%s:route53:::change/*", partition)},
		},
		{
			Actions:   []string{"route53:ChangeResourceRecordSets", "route53:ListResourceRecordSets"},
			Resources: args.HostedZoneARNs,
		},
		{
			Actions:   []string{"route53:ListHostedZonesByName"},
			Resources: []string{"*"},
		},
	}

	return policyBuilder.CreatePolicyWithAttachment(eksPolicyNamePrefix, eksPolicyDescription, policyStatements)
}
