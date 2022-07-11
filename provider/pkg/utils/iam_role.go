package utils

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RoleArgs struct {
	// IAM role name.
	Name pulumi.StringPtrInput `pulumi:"name"`

	// IAM role name prefix.
	NamePrefix pulumi.StringPtrInput `pulumi:"namePrefix"`

	// IAM Role description.
	Description pulumi.StringInput `pulumi:"description"`

	// Path of IAM role.
	Path pulumi.StringInput `pulumi:"path"`

	// Permissions boundary ARN to use for IAM role.
	PermissionsBoundaryArn pulumi.StringInput `pulumi:"permissionsBoundaryArn"`

	// List of ARNs of IAM policies to attach to IAM role.
	PolicyArns []pulumi.StringInput `pulumi:"policyArns"`

	// Whether role requires MFA.
	RequiresMFA pulumi.BoolInput `pulumi:"requiresMfa"`

	// A map of tags to add.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

type IAMRoleArgs struct {
	AssumeRolePolicy    pulumi.StringInput
	ForceDetachPolicies pulumi.BoolInput
	MaxSessionDuration  pulumi.IntInput
	Tags                pulumi.StringMapInput
	Role                RoleArgs
}

func NewIAMRole(ctx *pulumi.Context, name string, args *IAMRoleArgs, opts ...pulumi.ResourceOption) (*iam.Role, error) {
	roleResourceName := fmt.Sprintf("%s-role", name)

	// Set MaxSession Duration to default value if not set.
	if args.MaxSessionDuration == nil {
		args.MaxSessionDuration = pulumi.Int(3600)
	}

	if args.Role.Path == nil {
		args.Role.Path = pulumi.String("/")
	}

	role, err := iam.NewRole(ctx, roleResourceName, &iam.RoleArgs{
		AssumeRolePolicy:    args.AssumeRolePolicy,
		Description:         args.Role.Description,
		ForceDetachPolicies: args.ForceDetachPolicies,
		MaxSessionDuration:  args.MaxSessionDuration,
		Name:                args.Role.Name,
		NamePrefix:          args.Role.NamePrefix,
		Path:                args.Role.Path,
		PermissionsBoundary: args.Role.PermissionsBoundaryArn,
		Tags:                args.Tags,
	}, opts...)
	if err != nil {
		return nil, err
	}

	for i, policyARN := range args.Role.PolicyArns {
		policyAttachmentName := fmt.Sprintf("%s-policy-attachment-%v", roleResourceName, i)
		_, err = iam.NewRolePolicyAttachment(ctx, policyAttachmentName, &iam.RolePolicyAttachmentArgs{
			Role:      role.Name,
			PolicyArn: policyARN,
		}, opts...)
	}

	return role, nil
}
