package utils

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RoleTypeIdentifier string

const (
	AdminRoleType       RoleTypeIdentifier = "admin"
	AdminRoleDefaultARN                    = "arn:aws:iam::aws:policy/AdministratorAccess"

	PoweruserRoleType   RoleTypeIdentifier = "poweruser"
	PoweruserDefaultARN                    = "arn:aws:iam::aws:policy/PowerUserAccess"

	ReadonlyRoleType   = "readonly"
	ReadonlyDefaultARN = "arn:aws:iam::aws:policy/ReadOnlyAccess"
)

type IAMAssumableRolesArgs struct {
	MaxSessionDuration  int
	ForceDetachPolicies bool
	Admin               RoleArgs
	Poweruser           RoleArgs
	Readonly            RoleArgs
	AssumeRolePolicy    string
	AssumeRoleWithMFA   string
}

func NewAssumableRoles(ctx *pulumi.Context, name string, args *IAMAssumableRolesArgs, opts ...pulumi.ResourceOption) (map[RoleTypeIdentifier]*iam.Role, error) {
	rolesToCreate := map[RoleTypeIdentifier]RoleArgs{
		AdminRoleType:     args.Admin,
		PoweruserRoleType: args.Poweruser,
		ReadonlyRoleType:  args.Readonly,
	}

	roleOutput := make(map[RoleTypeIdentifier]*iam.Role)
	for typ, roleArgs := range rolesToCreate {
		rolePolicy := args.AssumeRolePolicy
		if roleArgs.RequiresMFA {
			rolePolicy = args.AssumeRoleWithMFA
		}

		if len(roleArgs.PolicyArns) == 0 {
			switch typ {
			case AdminRoleType:
				roleArgs.PolicyArns = append(roleArgs.PolicyArns, AdminRoleDefaultARN)
			case PoweruserRoleType:
				roleArgs.PolicyArns = append(roleArgs.PolicyArns, PoweruserDefaultARN)
			case ReadonlyRoleType:
				roleArgs.PolicyArns = append(roleArgs.PolicyArns, ReadonlyDefaultARN)
			}
		}

		roleResourceName := fmt.Sprintf("%s-%s", name, string(typ))

		// Set the role name to the resource name if a name was
		// not provided.
		if roleArgs.Name == "" {
			roleArgs.Name = roleResourceName
		}

		role, err := NewIAMRole(ctx, roleResourceName, &IAMRoleArgs{
			MaxSessionDuration:  args.MaxSessionDuration,
			ForceDetachPolicies: args.ForceDetachPolicies,
			AssumeRolePolicy:    rolePolicy,
			Tags:                roleArgs.Tags,
			Role:                roleArgs,
		}, opts...)
		if err != nil {
			return nil, err
		}

		roleOutput[typ] = role
	}

	return roleOutput, nil
}
