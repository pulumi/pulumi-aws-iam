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
	appmeshControllerNamePrefix  = "Appmesh_Controller-"
	appmeshControllerDescription = "Provides permissions to for appmesh controller"

	appmeshEnvoyProxyNamePrefix  = "Appmesh_Envoy_Proxy-"
	appmeshEnvoyProxyDescription = "Provides permissions to for appmesh envoy proxy"
)

type AppmeshPolicyArgs struct {
	// Determines whether to attach the Appmesh Controller policy to the role.
	Controller bool `pulumi:"controller"`

	// Determines whether to attach the Appmesh envoy proxy policy to the role.
	EnvoyProxy bool `pulumi:"envoyProxy"`
}

func AttachAppmeshControllerPolicy(policyBuilder *EKSRoleBuilder, partition, dnsSuffix string) error {
	policyStatements := []iam.GetPolicyDocumentStatement{
		{
			Resources: []string{"*"},
			Actions: []string{
				"appmesh:ListVirtualRouters",
				"appmesh:ListVirtualServices",
				"appmesh:ListRoutes",
				"appmesh:ListGatewayRoutes",
				"appmesh:ListMeshes",
				"appmesh:ListVirtualNodes",
				"appmesh:ListVirtualGateways",
				"appmesh:DescribeMesh",
				"appmesh:DescribeVirtualRouter",
				"appmesh:DescribeRoute",
				"appmesh:DescribeVirtualNode",
				"appmesh:DescribeVirtualGateway",
				"appmesh:DescribeGatewayRoute",
				"appmesh:DescribeVirtualService",
				"appmesh:CreateMesh",
				"appmesh:CreateVirtualRouter",
				"appmesh:CreateVirtualGateway",
				"appmesh:CreateVirtualService",
				"appmesh:CreateGatewayRoute",
				"appmesh:CreateRoute",
				"appmesh:CreateVirtualNode",
				"appmesh:UpdateMesh",
				"appmesh:UpdateRoute",
				"appmesh:UpdateVirtualGateway",
				"appmesh:UpdateVirtualRouter",
				"appmesh:UpdateGatewayRoute",
				"appmesh:UpdateVirtualService",
				"appmesh:UpdateVirtualNode",
				"appmesh:DeleteMesh",
				"appmesh:DeleteRoute",
				"appmesh:DeleteVirtualRouter",
				"appmesh:DeleteGatewayRoute",
				"appmesh:DeleteVirtualService",
				"appmesh:DeleteVirtualNode",
				"appmesh:DeleteVirtualGateway",
			},
		},
		{
			Actions: []string{"iam:CreateServiceLinkedRole"},
			Resources: []string{
				fmt.Sprintf("arn:%s:iam::*:role/aws-service-role/appmesh.%s/AWSServiceRoleForAppMesh", partition, dnsSuffix),
			},
			Conditions: []iam.GetPolicyDocumentStatementCondition{
				NewPolicyDocCondition("StringLike", "iam:AWSServiceName", fmt.Sprintf("appmesh.%s", dnsSuffix)),
			},
		},
		{
			Resources: []string{"*"},
			Actions: []string{
				"acm:ListCertificates",
				"acm:DescribeCertificate",
				"acm-pca:DescribeCertificateAuthority",
				"acm-pca:ListCertificateAuthorities",
			},
		},
		{
			Resources: []string{"*"},
			Actions: []string{
				"servicediscovery:CreateService",
				"servicediscovery:DeleteService",
				"servicediscovery:GetService",
				"servicediscovery:GetInstance",
				"servicediscovery:RegisterInstance",
				"servicediscovery:DeregisterInstance",
				"servicediscovery:ListInstances",
				"servicediscovery:ListNamespaces",
				"servicediscovery:ListServices",
				"servicediscovery:GetInstancesHealthStatus",
				"servicediscovery:UpdateInstanceCustomHealthStatus",
				"servicediscovery:GetOperation",
				"route53:GetHealthCheck",
				"route53:CreateHealthCheck",
				"route53:UpdateHealthCheck",
				"route53:ChangeResourceRecordSets",
				"route53:DeleteHealthCheck",
			},
		},
	}

	return policyBuilder.CreatePolicyWithAttachment(appmeshControllerNamePrefix, appmeshControllerDescription, policyStatements)
}

func AttachAppmeshEnvoyProxyPolicy(policyBuilder *EKSRoleBuilder) error {
	policyStatements := []iam.GetPolicyDocumentStatement{
		{
			Resources: []string{"*"},
			Actions:   []string{"appmesh:StreamAggregatedResources"},
		},
		{
			Resources: []string{"*"},
			Actions:   []string{"acm:ExportCertificate", "acm-pca:GetCertificateAuthorityCertificate"},
		},
	}

	return policyBuilder.CreatePolicyWithAttachment(appmeshEnvoyProxyNamePrefix, appmeshEnvoyProxyDescription, policyStatements)
}
