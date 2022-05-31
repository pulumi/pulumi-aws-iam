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
	nodeTerminationHandlerNamePrefix  = "Node_Termination_Handler_Policy-"
	nodeTerminationHandlerDescription = "Provides permissions to handle node termination events via the Node Termination Handler"

	nodeTerminationHandlerDefaultSQSQueueARN = "*"
)

type NodeTerminationHandlerPolicyArgs struct {
	// Determines whether to attach the Node Termination Handler policy to the role.
	Attach bool `pulumi:"attach"`

	// List of SQS ARNs that contain node termination events.
	SQSQueueARNs []string `pulumi:"sqsQueueArns"`
}

func AttachNodeTerminationPolicy(policyBuilder *EKSRoleBuilder, args NodeTerminationHandlerPolicyArgs) error {
	if len(args.SQSQueueARNs) == 0 {
		args.SQSQueueARNs = append(args.SQSQueueARNs, nodeTerminationHandlerDefaultSQSQueueARN)
	}

	policyStatements := []iam.GetPolicyDocumentStatement{
		{
			Resources: []string{"*"},
			Actions: []string{
				"autoscaling:CompleteLifecycleAction",
				"autoscaling:DescribeAutoScalingInstances",
				"autoscaling:DescribeTags",
				"ec2:DescribeInstances",
			},
		},
		{
			Actions:   []string{"sqs:DeleteMessage", "sqs:ReceiveMessage"},
			Resources: args.SQSQueueARNs,
		},
	}

	return policyBuilder.CreatePolicyWithAttachment(nodeTerminationHandlerNamePrefix, nodeTerminationHandlerDescription, policyStatements)
}
