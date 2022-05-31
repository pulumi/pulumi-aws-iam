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
	amazonManagedServicePrometheusNamePrefix  = "Managed_Service_Prometheus_Policy-"
	amazonManagedServicePrometheusDescription = "Provides permissions to for Amazon Managed Service for Prometheus"

	amazonManagedServicePrometheusDefaultWorkspaceARN = "*"
)

type AmazonManagedServicePrometheusPolicyArgs struct {
	// Determines whether to attach the Amazon Managed Service for Prometheus IAM policy to the role.
	Attach bool `pulumi:"attach"`

	// List of AMP Workspace ARNs to read and write metrics.
	WorkspaceARNs []string `pulumi:"workspaceArns"`
}

func AttachAmazonManagedServicePrometheusPolicy(policyBuilder *EKSRoleBuilder, args AmazonManagedServicePrometheusPolicyArgs) error {
	if len(args.WorkspaceARNs) == 0 {
		args.WorkspaceARNs = append(args.WorkspaceARNs, amazonManagedServicePrometheusDefaultWorkspaceARN)
	}

	policyStatements := []iam.GetPolicyDocumentStatement{
		{
			Actions: []string{
				"aps:RemoteWrite",
				"aps:QueryMetrics",
				"aps:GetSeries",
				"aps:GetLabels",
				"aps:GetMetricMetadata",
			},
			Resources: args.WorkspaceARNs,
		},
	}

	return policyBuilder.CreatePolicyWithAttachment(amazonManagedServicePrometheusNamePrefix, amazonManagedServicePrometheusDescription, policyStatements)
}
