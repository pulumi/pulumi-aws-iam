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

package provider

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-aws-iam/pkg/eks_policies"
	"github.com/pulumi/pulumi-aws-iam/pkg/utils"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const RoleForServiceAccountsEksIdentifier = "aws-iam:index:RoleForServiceAccountsEks"

type OIDCServiceProviderEKS struct {
	ProviderARN              string   `pulumi:"providerArn"`
	NamespaceServiceAccounts []string `pulumi:"namespaceServiceAccounts"`
}

type EKSServiceAccountPolicies struct {
	// The Cert Manager IAM policy to attach to the role.
	CertManager eks_policies.CertManagerPolicyArgs `pulumi:"certManager"`

	// The Cluster Autoscaler IAM policy to the role.
	ClusterAutoScaling eks_policies.ClusterAutoScalingPolicyArgs `pulumi:"clusterAutoScaling"`

	// The EBS CSI IAM policy to the role.
	EBSCSI eks_policies.EBSCSIPolicyArgs `pulumi:"ebsCsi"`

	// The EFS CSI IAM policy to the role.
	EFSCSI eks_policies.EFSCSIPolicyArgs `pulumi:"efsCsi"`

	// The External DNS IAM policy to the role.
	ExternalDNS eks_policies.ExternalDNSPolicyArgs `pulumi:"externalDns"`

	// The External Secrets policy to the role.
	ExternalSecrets eks_policies.ExternalSecretsPolicyArgs `pulumi:"externalSecrets"`

	// The FSx for Lustre CSI Driver IAM policy to the role.
	FSxLustreCSI eks_policies.FSXLustreCSIPolicyArgs `pulumi:"fsxLustreCsi"`

	// The Karpenter Controller policy to the role.
	KarpenterController eks_policies.KarpenterControllerPolicyArgs `pulumi:"karpenterController"`

	// The Load Balancer Controller policy to the role.
	LoadBalancer eks_policies.LoadBalancerPolicyArgs `pulumi:"loadBalancer"`

	// The Appmesh policies.
	Appmesh eks_policies.AppmeshPolicyArgs `pulumi:"appmesh"`

	// The Amazon Managed Service for Prometheus IAM policy to the role.
	AmazonManagedServicePrometheus eks_policies.AmazonManagedServicePrometheusPolicyArgs `pulumi:"amazonManagedServicePrometheus"`

	// The Velero IAM policy to the role.
	Velero eks_policies.VeleroPolicyArgs `pulumi:"velero"`

	// The VPC CNI IAM policy to the role.
	VPNCNI eks_policies.VPNCNIPolicyArgs `pulumi:"vpnCni"`

	// The Node Termination Handler policy to the role.
	NodeTerminationHandler eks_policies.NodeTerminationHandlerPolicyArgs `pulumi:"nodeTerminationHandler"`
}

type RoleForServiceAccountsEksArgs struct {
	// A map of tags to add.
	Tags map[string]string `pulumi:"tags"`

	// IAM role.
	Role utils.RoleArgs `pulumi:"role"`

	// IAM policy name prefix.
	PolicyNamePrefix string `pulumi:"policyNamePrefix"`

	// Maximum CLI/API session duration in seconds between 3600 and 43200.
	MaxSessionDuration int `pulumi:"maxSessionDuration"`

	// Whether policies should be detached from this role when destroying.
	ForceDetachPolicies bool `pulumi:"forceDetachPolicies"`

	// Map of OIDC providers.
	OIDCProviders map[string]OIDCServiceProviderEKS `pulumi:"oidcProviders"`

	// Name of the IAM condition operator to evaluate when assuming the role.
	AssumeRoleConditionTest string `pulumi:"assumeRoleConditionTest"`

	// The different policies to attach to the role.
	Policies EKSServiceAccountPolicies `pulumi:"policies"`
}

type RoleForServiceAccountsEks struct {
	pulumi.ResourceState

	Role struct {
		// ARN of IAM role.
		Arn pulumi.StringOutput `pulumi:"arn"`

		// Name of IAM role.
		Name pulumi.StringOutput `pulumi:"name"`

		// Path of IAM role.
		Path pulumi.StringPtrOutput `pulumi:"path"`

		// Unique ID of IAM role.
		UniqueID pulumi.StringOutput `pulumi:"uniqueId"`
	} `pulumi:"role"`
}

func NewRoleForServiceAccountsEks(ctx *pulumi.Context, name string, args *RoleForServiceAccountsEksArgs, opts ...pulumi.ResourceOption) (*RoleForServiceAccountsEks, error) {
	if args == nil {
		args = &RoleForServiceAccountsEksArgs{}
	}

	component := &RoleForServiceAccountsEks{}
	err := ctx.RegisterComponentResource(RoleForServiceAccountsEksIdentifier, name, component, opts...)
	if err != nil {
		return nil, err
	}

	opts = append(opts, pulumi.Parent(component))

	account, err := aws.GetCallerIdentity(ctx)
	if err != nil {
		return nil, err
	}

	currentPartition, err := aws.GetPartition(ctx, nil, nil)
	if err != nil {
		return nil, err
	}

	var oidcPolicyDocStatements []iam.GetPolicyDocumentStatement
	for _, provider := range args.OIDCProviders {
		effect := "Allow"

		var serviceAccounts []string
		for _, sa := range provider.NamespaceServiceAccounts {
			serviceAccounts = append(serviceAccounts, fmt.Sprintf("system:serviceaccount:%s", sa))
		}

		oidcPolicyDocStatements = append(oidcPolicyDocStatements, iam.GetPolicyDocumentStatement{
			Effect:  &effect,
			Actions: []string{"sts:AssumeRoleWithWebIdentity"},
			Principals: []iam.GetPolicyDocumentStatementPrincipal{
				{
					Type:        "Federated",
					Identifiers: []string{provider.ProviderARN},
				},
			},
			Conditions: []iam.GetPolicyDocumentStatementCondition{
				{
					Test:     args.AssumeRoleConditionTest,
					Variable: fmt.Sprintf("%s:sub", strings.ReplaceAll(provider.ProviderARN, "/^(.*provider/)/", "")),
					Values:   serviceAccounts,
				},
				{
					Test:     args.AssumeRoleConditionTest,
					Variable: fmt.Sprintf("%s:aud", strings.ReplaceAll(provider.ProviderARN, "/^(.*provider/)/", "")),
					Values:   []string{"sts.amazonaws.com"},
				},
			},
		})
	}

	policyDoc, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
		Statements: oidcPolicyDocStatements,
	})
	if err != nil {
		return nil, err
	}

	eksRole, err := utils.NewIAMRole(ctx, name, &utils.IAMRoleArgs{
		Role:                args.Role,
		MaxSessionDuration:  args.MaxSessionDuration,
		ForceDetachPolicies: args.ForceDetachPolicies,
		AssumeRolePolicy:    policyDoc.Json,
		Tags:                args.Tags,
	}, opts...)
	if err != nil {
		return nil, err
	}

	for _, policyARN := range args.Role.PolicyArns {
		_, err := iam.NewRolePolicyAttachment(ctx, fmt.Sprintf("%s-custom", name), &iam.RolePolicyAttachmentArgs{
			Role:      eksRole.Name,
			PolicyArn: pulumi.String(policyARN),
		}, opts...)
		if err != nil {
			return nil, err
		}
	}

	policyBuilder := eks_policies.CreateNewRoleBuilder(ctx, eksRole, name, args.PolicyNamePrefix, args.Role.Path, args.Tags, opts...)

	// Cert Manager
	if args.Policies.CertManager.Attach {
		err = eks_policies.AttachCertManagerPolicy(policyBuilder, currentPartition.Partition, args.Policies.CertManager)
		if err != nil {
			return nil, err
		}
	}

	// Cluster Autoscaler
	if args.Policies.ClusterAutoScaling.Attach {
		err = eks_policies.AttachClusterAutoscalerPolicy(policyBuilder, args.Policies.ClusterAutoScaling)
		if err != nil {
			return nil, err
		}
	}

	// EBS CSI
	if args.Policies.EBSCSI.Attach {
		err = eks_policies.AttachEBSCSIPolicy(policyBuilder, currentPartition.Partition, args.Policies.EBSCSI)
		if err != nil {
			return nil, err
		}
	}

	// EFS CSI
	if args.Policies.EFSCSI.Attach {
		err = eks_policies.AttachEFSCSIPolicy(policyBuilder)
		if err != nil {
			return nil, err
		}
	}

	// External DNS
	if args.Policies.ExternalDNS.Attach {
		err = eks_policies.AttachExternalDNSPolicy(policyBuilder, args.Policies.ExternalDNS)
		if err != nil {
			return nil, err
		}
	}

	// External Secrets
	if args.Policies.ExternalSecrets.Attach {
		err = eks_policies.AttachExternalSecretsPolicy(policyBuilder, args.Policies.ExternalSecrets)
		if err != nil {
			return nil, err
		}
	}

	// FSx Lustre CSI
	if args.Policies.FSxLustreCSI.Attach {
		err = eks_policies.AttachFSXLustreCSIPolicy(policyBuilder, currentPartition.DnsSuffix, args.Policies.FSxLustreCSI)
		if err != nil {
			return nil, err
		}
	}

	if args.Policies.KarpenterController.Attach {
		err = eks_policies.AttachKarpenterControllerPolicy(policyBuilder, currentPartition.Partition, account.AccountId, args.Policies.KarpenterController)
		if err != nil {
			return nil, err
		}
	}

	if args.Policies.LoadBalancer.Controller {
		err = eks_policies.AttachLoadBalancerControllerPolicy(policyBuilder, currentPartition.Partition, currentPartition.DnsSuffix)
		if err != nil {
			return nil, err
		}
	}

	if args.Policies.LoadBalancer.TargetGroupBindingOnly {
		err = eks_policies.AttachLoadBalancerTargetGroupBindingOnlyPolicy(policyBuilder)
		if err != nil {
			return nil, err
		}
	}

	if args.Policies.Appmesh.Controller {
		err = eks_policies.AttachAppmeshControllerPolicy(policyBuilder, currentPartition.Partition, currentPartition.DnsSuffix)
		if err != nil {
			return nil, err
		}
	}

	if args.Policies.Appmesh.EnvoyProxy {
		err = eks_policies.AttachAppmeshEnvoyProxyPolicy(policyBuilder)
		if err != nil {
			return nil, err
		}
	}

	if args.Policies.AmazonManagedServicePrometheus.Attach {
		err = eks_policies.AttachAmazonManagedServicePrometheusPolicy(policyBuilder, args.Policies.AmazonManagedServicePrometheus)
		if err != nil {
			return nil, err
		}
	}

	if args.Policies.Velero.Attach {
		err = eks_policies.AttachVeleroPolicy(policyBuilder, args.Policies.Velero)
		if err != nil {
			return nil, err
		}
	}

	if args.Policies.VPNCNI.Attach {
		err = eks_policies.AttachVPNCNIPolicy(policyBuilder, currentPartition.Partition, args.Policies.VPNCNI)
		if err != nil {
			return nil, err
		}
	}

	if args.Policies.NodeTerminationHandler.Attach {
		err = eks_policies.AttachNodeTerminationPolicy(policyBuilder, args.Policies.NodeTerminationHandler)
		if err != nil {
			return nil, err
		}
	}

	component.Role.Arn = eksRole.Arn
	component.Role.Name = eksRole.Name
	component.Role.Path = eksRole.Path
	component.Role.UniqueID = eksRole.UniqueId

	return component, nil
}
