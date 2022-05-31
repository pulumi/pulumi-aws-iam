package utils

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RoleArgs struct {
	// IAM role name.
	Name string `pulumi:"name"`

	// IAM role name prefix.
	NamePrefix string `pulumi:"namePrefix"`

	// IAM Role description.
	Description string `pulumi:"description"`

	// Path of IAM role.
	Path string `pulumi:"path"`

	// Permissions boundary ARN to use for IAM role.
	PermissionsBoundaryArn string `pulumi:"permissionsBoundaryArn"`

	// List of ARNs of IAM policies to attach to IAM role.
	PolicyArns []string `pulumi:"policyArns"`

	// Whether role requires MFA.
	RequiresMFA bool `pulumi:"requiresMfa"`

	// A map of tags to add.
	Tags map[string]string `pulumi:"tags"`
}

type IAMRoleArgs struct {
	AssumeRolePolicy    string
	ForceDetachPolicies bool
	MaxSessionDuration  int
	Tags                map[string]string
	Role                RoleArgs
}

func NewIAMRole(ctx *pulumi.Context, name string, args *IAMRoleArgs, opts ...pulumi.ResourceOption) (*iam.Role, error) {
	roleResourceName := fmt.Sprintf("%s-role", name)

	// If the Role NamePrefix is set prefer that over the Name value.
	var roleNamePrefix pulumi.StringPtrInput
	roleName := pulumi.StringPtr(args.Role.Name)
	if args.Role.NamePrefix != "" {
		roleNamePrefix = pulumi.StringPtr(args.Role.NamePrefix)
		roleName = nil
	}

	// Set MaxSession Duration to default value if not set.
	if args.MaxSessionDuration == 0 {
		args.MaxSessionDuration = 36000
	}

	// Set the Path to "/" if it is not provided.
	if args.Role.Path == "" {
		args.Role.Path = "/"
	}

	role, err := iam.NewRole(ctx, roleResourceName, &iam.RoleArgs{
		AssumeRolePolicy:    pulumi.String(args.AssumeRolePolicy),
		Description:         pulumi.String(args.Role.Description),
		ForceDetachPolicies: pulumi.BoolPtr(args.ForceDetachPolicies),
		MaxSessionDuration:  pulumi.IntPtr(args.MaxSessionDuration),
		Name:                roleName,
		NamePrefix:          roleNamePrefix,
		Path:                pulumi.String(args.Role.Path),
		PermissionsBoundary: pulumi.String(args.Role.PermissionsBoundaryArn),
		Tags:                pulumi.ToStringMap(args.Tags),
	}, opts...)
	if err != nil {
		return nil, err
	}

	for i, policyARN := range args.Role.PolicyArns {
		policyAttachmentName := fmt.Sprintf("%s-policy-attachment-%v", roleResourceName, i)
		_, err = iam.NewRolePolicyAttachment(ctx, policyAttachmentName, &iam.RolePolicyAttachmentArgs{
			Role:      role.Name,
			PolicyArn: pulumi.String(policyARN),
		}, opts...)
	}

	return role, nil
}
