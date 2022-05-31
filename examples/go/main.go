package main

import (
	"encoding/json"

	iam "github.com/pulumi/pulumi-aws-iam/sdk/go/aws-iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Account
		account, err := iam.NewAccount(ctx, "account", &iam.AccountArgs{
			AccountAlias: pulumi.String("cool-alias"),
			PasswordPolicy: iam.AccountPasswordPolicyArgs{
				MinimumLength:              pulumi.IntPtr(37),
				RequireNumbers:             pulumi.Bool(false),
				AllowUsersToChange:         pulumi.Bool(true),
				HardExpiry:                 pulumi.Bool(true),
				RequireSymbols:             pulumi.Bool(true),
				RequireLowercaseCharacters: pulumi.Bool(true),
				RequireUppercaseCharacters: pulumi.Bool(true),
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("account", account)

		// Assumable Role
		assumableRole, err := iam.NewAssumableRole(ctx, "assumable-role", &iam.AssumableRoleArgs{
			TrustedRoleArns: pulumi.ToStringArray([]string{"arn:aws:iam::307990089504:root", "arn:aws:iam::835367859851:user/pulumipus"}),
			Role: &iam.RoleWithMFAArgs{
				Name:        pulumi.String("custom"),
				RequiresMfa: pulumi.BoolPtr(true),
				PolicyArns:  pulumi.ToStringArray([]string{"arn:aws:iam::aws:policy/AmazonCognitoReadOnly", "arn:aws:iam::aws:policy/AlexaForBusinessFullAccess"}),
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("assumableRole", assumableRole)

		// Assumable Role With OIDC
		assumableRoleWithOIDC, err := iam.NewAssumableRoleWithOIDC(ctx, "assumable-role-with-oidc", &iam.AssumableRoleWithOIDCArgs{
			Role: iam.RoleArgs{
				Name:       pulumi.String("oidc-role"),
				PolicyArns: pulumi.ToStringArray([]string{"arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"}),
			},
			Tags: pulumi.ToStringMap(map[string]string{
				"Role": "oidc-role",
			}),
			ProviderUrls: pulumi.ToStringArray([]string{"oidc.eks.eu-west-1.amazonaws.com/id/BA9E170D464AF7B92084EF72A69B9DC8"}),
		})
		if err != nil {
			return err
		}

		ctx.Export("assumableRoleWithOIDC", assumableRoleWithOIDC)

		// Assumable Role With SAML
		assumableRoleWithSAML, err := iam.NewAssumableRoleWithSAML(ctx, "assumable-role-with-saml", &iam.AssumableRoleWithSAMLArgs{
			Role: iam.RoleArgs{
				Name:       pulumi.String("saml-role"),
				PolicyArns: pulumi.ToStringArray([]string{"arn:aws:iam::aws:policy/ReadOnlyAccess"}),
			},
			Tags: pulumi.ToStringMap(map[string]string{
				"Role": "saml-role",
			}),
			ProviderIds: pulumi.ToStringArray([]string{"arn:aws:iam::235367859851:saml-provider/idp_saml"}),
		})
		if err != nil {
			return err
		}

		ctx.Export("assumableRoleWithSAML", assumableRoleWithSAML)

		// Assumable Roles
		assumableRoles, err := iam.NewAssumableRoles(ctx, "assumable-roles", &iam.AssumableRolesArgs{
			TrustedRoleArns: pulumi.ToStringArray([]string{"arn:aws:iam::307990089504:root", "arn:aws:iam::835367859851:user/anton"}),
			Admin:           iam.AdminRoleWithMFAArgs{},
			Poweruser: iam.PoweruserRoleWithMFAArgs{
				Name: pulumi.String("developer"),
			},
			Readonly: iam.ReadonlyRoleWithMFAArgs{
				RequiresMfa: pulumi.BoolPtr(true),
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("assumableRoles", assumableRoles)

		// Assumable Roles With SAML
		assumableRolesWithSAML, err := iam.NewAssumableRolesWithSAML(ctx, "assumable-roles-with-saml", &iam.AssumableRolesWithSAMLArgs{
			ProviderIds: pulumi.ToStringArray([]string{"arn:aws:iam::235367859851:saml-provider/idp_saml"}),
			Admin:       iam.AdminRoleArgs{},
			Readonly:    iam.ReadonlyRoleArgs{},
			Poweruser: iam.PoweruserRoleArgs{
				Name: pulumi.String("developer"),
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("assumableRolesWithSAML", assumableRolesWithSAML)

		// EKS Role
		eksRole, err := iam.NewEKSRole(ctx, "eks-role", &iam.EKSRoleArgs{
			Role: iam.RoleArgs{
				Name:       pulumi.String("eks-role"),
				PolicyArns: pulumi.ToStringArray([]string{"arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"}),
			},
			Tags: pulumi.ToStringMap(map[string]string{
				"Role": "eks-role",
			}),
			// Uncomment the below and replace actual cluster values.
			// ClusterServiceAccounts: pulumi.ToStringArrayMap(map[string][]string{
			// 	"cluster1": {"default:my-app"},
			// 	"cluster2": {"default:my-app", "canary:my-app"},
			// }),
		})
		if err != nil {
			return err
		}

		ctx.Export("eksRole", eksRole)

		// Group With Assumable Roles Policy
		groupWithAssumableRolesPolicy, err := iam.NewGroupWithAssumableRolesPolicy(ctx, "group-with-assumable-roles-policy", &iam.GroupWithAssumableRolesPolicyArgs{
			Name:           pulumi.String("production-readonly"),
			AssumableRoles: pulumi.ToStringArray([]string{"arn:aws:iam::835367859855:role/readonly"}),
			GroupUsers:     pulumi.ToStringArray([]string{"user1", "user2"}),
		})
		if err != nil {
			return err
		}

		ctx.Export("groupWithAssumableRolesPolicy", groupWithAssumableRolesPolicy)

		// Group With Policies
		groupWithPolicies, err := iam.NewGroupWithPolicies(ctx, "group-with-policies", &iam.GroupWithPoliciesArgs{
			Name:                          pulumi.String("superadmins"),
			GroupUsers:                    pulumi.ToStringArray([]string{"user1", "user2"}),
			AttachIamSelfManagementPolicy: pulumi.BoolPtr(true),
			CustomGroupPolicyArns:         pulumi.ToStringArray([]string{"arn:aws:iam::aws:policy/AdministratorAccess"}),
			CustomGroupPolicies: pulumi.ToStringMapArray([]map[string]string{
				{
					"name":   "AllowS3Listing",
					"policy": "{}",
				},
			}),
		})
		if err != nil {
			return err
		}

		ctx.Export("groupWithPolicies", groupWithPolicies)

		// Policy
		policyJSON, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []interface{}{
				map[string]interface{}{
					"Effect":   "Allow",
					"Action":   []string{"ec2:Describe"},
					"Resource": []string{"*"},
				},
			},
		})
		if err != nil {
			return err
		}

		policy, err := iam.NewPolicy(ctx, "policy", &iam.PolicyArgs{
			Name:           pulumi.String("example"),
			Path:           pulumi.String("/"),
			Description:    pulumi.String("My example policy"),
			PolicyDocument: pulumi.String(string(policyJSON)),
		})
		if err != nil {
			return err
		}

		ctx.Export("policy", policy)

		// Read Only Policy
		readOnlyPolicy, err := iam.NewReadOnlyPolicy(ctx, "read-only-policy", &iam.ReadOnlyPolicyArgs{
			Name:            pulumi.String("example"),
			Path:            pulumi.String("/"),
			Description:     pulumi.String("My example policy"),
			AllowedServices: pulumi.ToStringArray([]string{"rds", "dynamo"}),
		})
		if err != nil {
			return err
		}

		ctx.Export("readOnlyPolicy", readOnlyPolicy)

		// Role For Service Accounts EKS
		roleForServiceAccountsEKS, err := iam.NewRoleForServiceAccountsEks(ctx, "role-for-service-accounts-eks", &iam.RoleForServiceAccountsEksArgs{
			Role: iam.EKSServiceAccountRolePtr(&iam.EKSServiceAccountRoleArgs{
				Name: pulumi.String("vpc-cni"),
			}),
			Tags: pulumi.ToStringMap(map[string]string{
				"Name": "vpc-cni-irsa",
			}),
			OidcProviders: iam.OIDCProviderMap{
				"main": iam.OIDCProviderArgs{
					ProviderArn:              pulumi.String("arn:aws:iam::012345678901:oidc-provider/oidc.eks.us-east-1.amazonaws.com/id/5C54DDF35ER19312844C7333374CC09D"),
					NamespaceServiceAccounts: pulumi.ToStringArray([]string{"default:my-app", "canary:my-app"}),
				},
			},
			Policies: iam.EKSRolePoliciesPtr(&iam.EKSRolePoliciesArgs{
				VpnCni: iam.EKSVPNCNIPolicyPtr(&iam.EKSVPNCNIPolicyArgs{
					Attach:     pulumi.Bool(true),
					EnableIpv4: pulumi.BoolPtr(true),
				}),
			}),
		})
		if err != nil {
			return err
		}

		ctx.Export("roleForServiceAccountsEKS", roleForServiceAccountsEKS)

		// User
		user, err := iam.NewUser(ctx, "user", &iam.UserArgs{
			Name:                  pulumi.String("pulumipus"),
			ForceDestroy:          pulumi.BoolPtr(true),
			PgpKey:                pulumi.String("keybase:test"),
			PasswordResetRequired: pulumi.BoolPtr(false),
		})
		if err != nil {
			return err
		}

		ctx.Export("user", user)

		return nil
	})
}
