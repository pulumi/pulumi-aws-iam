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
	nodeTerminationHandlerNamePrefix  = "Node_Termination_Handler_Policy-"
	nodeTerminationHandlerDescription = "Provides permissions to handle node termination events via the Node Termination Handler"

	nodeTerminationHandlerDefaultSQSQueueARN = "*"
)

type NodeTerminationHandlerPolicyArgs struct {
	// Determines whether to attach the Node Termination Handler policy to the role.
	Attach bool `pulumi:"attach"`

	// List of SQS ARNs that contain node termination events.
	SQSQueueARNs pulumi.StringArrayInput `pulumi:"sqsQueueArns"`
}

func AttachNodeTerminationPolicy(ctx *pulumi.Context, policyBuilder *EKSRoleBuilder, args NodeTerminationHandlerPolicyArgs) error {
	policyJSON := args.SQSQueueARNs.ToStringArrayOutput().ApplyT(func(arns []string) (string, error) {
		if len(arns) == 0 {
			arns = append(arns, nodeTerminationHandlerDefaultSQSQueueARN)
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
				Resources: arns,
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

	return policyBuilder.CreatePolicyWithAttachment(nodeTerminationHandlerNamePrefix, nodeTerminationHandlerDescription, policyJSON)
}
