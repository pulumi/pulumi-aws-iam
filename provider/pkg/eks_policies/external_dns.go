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
	externalDNSNamePrefix  = "External_DNS_Policy-"
	externalDNSDescription = "External DNS policy to allow management of Route53 hosted zone records"

	externalDNSDefaultHostedZoneARN = "arn:aws:route53:::hostedzone/*"
)

type ExternalDNSPolicyArgs struct {
	// Determines whether to attach the External DNS IAM policy to the role.
	Attach bool `pulumi:"attach"`

	// Route53 hosted zone ARNs to allow External DNS to manage records.
	HostedZoneARNs []string `pulumi:"hostedZoneArns"`
}

func AttachExternalDNSPolicy(policyBuilder *EKSRoleBuilder, args ExternalDNSPolicyArgs) error {
	if len(args.HostedZoneARNs) == 0 {
		args.HostedZoneARNs = append(args.HostedZoneARNs, externalDNSDefaultHostedZoneARN)
	}

	policyStatements := []iam.GetPolicyDocumentStatement{
		{
			Actions:   []string{"route53:ChangeResourceRecordSets"},
			Resources: args.HostedZoneARNs,
		},
		{
			Actions:   []string{"route53:ListHostedZones", "route53:ListResourceRecordSets"},
			Resources: []string{"*"},
		},
	}

	return policyBuilder.CreatePolicyWithAttachment(externalDNSNamePrefix, externalDNSDescription, policyStatements)
}
