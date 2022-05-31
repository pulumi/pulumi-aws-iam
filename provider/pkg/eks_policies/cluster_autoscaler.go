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
	clusterAutoscalerNamePrefix  = "Cluster_Autoscaler_Policy-"
	clusterAutoscalerDescription = "Cluster autoscaler policy to allow examination and modification of EC2 Auto Scaling Groups"
)

type ClusterAutoScalingPolicyArgs struct {
	// Determines whether to attach the Cluster Autoscaler IAM policy to the role.
	Attach bool `pulumi:"attach"`

	// List of cluster IDs to appropriately scope permissions within the Cluster Autoscaler IAM policy.
	ClusterIDs []string `pulumi:"clusterIds"`
}

func AttachClusterAutoscalerPolicy(policyBuilder *EKSRoleBuilder, args ClusterAutoScalingPolicyArgs) error {
	var policyStatements []iam.GetPolicyDocumentStatement
	for _, id := range args.ClusterIDs {
		policyStatements = append(policyStatements, iam.GetPolicyDocumentStatement{
			Actions: []string{
				"autoscaling:SetDesiredCapacity",
				"autoscaling:TerminateInstanceInAutoScalingGroup",
				"autoscaling:UpdateAutoScalingGroup",
			},
			Resources: []string{"*"},
			Conditions: []iam.GetPolicyDocumentStatementCondition{
				{
					Test:     "StringEquals",
					Variable: fmt.Sprintf("autoscaling:ResourceTag/kubernetes.io/cluster/%s", id),
					Values:   []string{"owned"},
				},
			},
		})
	}

	policyStatements = append(policyStatements, iam.GetPolicyDocumentStatement{
		Actions: []string{
			"autoscaling:DescribeAutoScalingGroups",
			"autoscaling:DescribeAutoScalingInstances",
			"autoscaling:DescribeLaunchConfigurations",
			"autoscaling:DescribeTags",
			"ec2:DescribeLaunchTemplateVersions",
			"ec2:DescribeInstanceTypes",
		},
		Resources: []string{"*"},
	})

	return policyBuilder.CreatePolicyWithAttachment(clusterAutoscalerNamePrefix, clusterAutoscalerDescription, policyStatements)
}
