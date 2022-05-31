using Pulumi;
using Pulumi.AwsIam;
using Pulumi.AwsIam.Inputs;
using System.Collections.Immutable;

class MyStack : Stack
{
    public MyStack()
    {
        // Account
        var account = new Account("account", new AccountArgs
        {
            AccountAlias = "cool-alias",
            PasswordPolicy=new AccountPasswordPolicyArgs
            {
                MinimumLength = 37,
                RequireNumbers = false,
                AllowUsersToChange = true,
                HardExpiry = true,
                RequireSymbols = true,
                RequireLowercaseCharacters = true,
                RequireUppercaseCharacters = true,
            }

        });

        this.Account = Output.Create<Account>(account);

        // Assumable Role
        var assumableRole = new AssumableRole("assumable-role", new AssumableRoleArgs
        {
            TrustedRoleArns = {"arn:aws:iam::307990089504:root", "arn:aws:iam::835367859851:user/pulumipus"},
            Role = new RoleWithMFAArgs
            {
                Name = "custom",
                RequiresMfa = true,
                PolicyArns = {"arn:aws:iam::aws:policy/AmazonCognitoReadOnly","arn:aws:iam::aws:policy/AlexaForBusinessFullAccess"},
            },
        });

        this.AssumableRole = Output.Create<AssumableRole>(assumableRole);

        // Assumable Role With OIDC
        var assumableRoleWithOidc = new AssumableRoleWithOIDC("assumable-role-with-oidc", new AssumableRoleWithOIDCArgs
        {
            Role = new RoleArgs
            {
                Name = "oidc-role",
                PolicyArns = {"arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"},
            },
            Tags = new InputMap<string>
            {
                {"Role", "odic-role"},
            },
            ProviderUrls = {"oidc.eks.eu-west-1.amazonaws.com/id/BA9E170D464AF7B92084EF72A69B9DC8"},
        });

        this.AssumableRoleWithOidc = Output.Create<AssumableRoleWithOIDC>(assumableRoleWithOidc);

        // Assumable Role With SAML
        var assumableRoleWithSaml = new AssumableRoleWithSAML("assumable-role-with-saml", new AssumableRoleWithSAMLArgs
        {
            Role = new RoleArgs
            {
                Name = "saml-role",
                PolicyArns = {"arn:aws:iam::aws:policy/ReadOnlyAccess"},
            },
            Tags = new InputMap<string>
            {
                {"Role", "saml-role"},
            },
            ProviderIds = {"arn:aws:iam::235367859851:saml-provider/idp_saml"},
        });

        this.AssumableRoleWithSaml = Output.Create<AssumableRoleWithSAML>(assumableRoleWithSaml);

        // Assumable Roles
        var assumableRoles = new AssumableRoles("assumable-roles", new AssumableRolesArgs
        {
            TrustedRoleArns = {"arn:aws:iam::307990089504:root", "arn:aws:iam::835367859851:user/anton"},
            Admin = new AdminRoleWithMFAArgs(),
            Poweruser = new PoweruserRoleWithMFAArgs
            {
                Name = "developer",
            },
            Readonly = new ReadonlyRoleWithMFAArgs
            {
                RequiresMfa = true,
            },
        });

        this.AssumableRoles = Output.Create<AssumableRoles>(assumableRoles);

        // Assumable Roles With SAML
        var assumableRolesWithSaml = new AssumableRolesWithSAML("assumable-roles-with-saml", new AssumableRolesWithSAMLArgs
        {
            ProviderIds = {"arn:aws:iam::235367859851:saml-provider/idp_saml"},
            Admin = new AdminRoleArgs(),
            Readonly = new ReadonlyRoleArgs(),
            Poweruser = new PoweruserRoleArgs
            {
                Name = "developer",
            },
        });

        this.AssumableRolesWithSaml = Output.Create<AssumableRolesWithSAML>(assumableRolesWithSaml);

        // EKS Role
        var eksRole = new EKSRole("eks-role", new EKSRoleArgs
        {
            Role = new RoleArgs
            {
                Name = "eks-role",
                PolicyArns = {"arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"},
            },
            Tags = new InputMap<string>
            {
                {"Name", "eks-role"},
            },
            // Uncomment the below and replace actual cluster values.
            // ClusterServiceAccounts = {
            //     {"cluster1", ImmutableArray.Create<string>(new string[] {"default:my-app"})},
            //     {"cluster2", ImmutableArray.Create<string>(new string[] {"default:my-app", "canary:my-app"})}
            // },
        });

        this.EksRole = Output.Create<EKSRole>(eksRole);

        // Group With Assumable Roles Policy
        var groupWithAssumableRolePolicy = new GroupWithAssumableRolesPolicy("group-with-assumable-roles-policy", new GroupWithAssumableRolesPolicyArgs
        {
            Name = "production-readonly",
            AssumableRoles = {"arn:aws:iam::835367859855:role/readonly"},
            GroupUsers = {"user1", "user2"},
        });

        this.GroupWithAssumableRolesPolicy = Output.Create<GroupWithAssumableRolesPolicy>(groupWithAssumableRolePolicy);

        // Group With Policies
        var groupWithPolicies = new GroupWithPolicies("group-with-policies", new GroupWithPoliciesArgs
        {
            Name = "superadmins",
            GroupUsers = {"user1", "user2"},
            AttachIamSelfManagementPolicy = true,
            CustomGroupPolicyArns = {"arn:aws:iam::aws:policy/AdministratorAccess"},
            CustomGroupPolicies = new InputList<ImmutableDictionary<string, string>>
            {
                ImmutableDictionary.Create<string, string>()
                    .Add("name", "AllowS3Listing")
                    .Add("policy", "{}"),
            },
        });

        this.GroupWithPolicies = Output.Create<GroupWithPolicies>(groupWithPolicies);

        // Policy
        var policy = new Policy("policy", new PolicyArgs
        {
            Name = "example",
            Path = "/",
            Description = "My example policy",
            PolicyDocument =
                @"{
                ""Version"": ""2012-10-17"",
                ""Statement"": [
                {
                    ""Action"": [
                    ""ec2:Describe*""
                    ],
                    ""Effect"": ""Allow"",
                    ""Resource"": ""*""
                }
                ]
            }"
        });

        this.Policy = Output.Create<Policy>(policy);

        // Read Only Policy
        var readOnlyPolicy = new ReadOnlyPolicy("read-only-policy", new ReadOnlyPolicyArgs
        {
            Name = "example",
            Path = "/",
            Description = "My example read only policy",
            AllowedServices = {"rds", "dynamo"},
        });

        this.ReadOnlyPolicy = Output.Create<ReadOnlyPolicy>(readOnlyPolicy);

        // Role For Service Accounts EKS
        var roleForServiceAccountEks = new RoleForServiceAccountsEks("role-for-service-account-eks", new RoleForServiceAccountsEksArgs
        {
            Role = new EKSServiceAccountRoleArgs
            {
                Name = "vpn-cni",
            },
            Tags = {
                {"Name", "vpc-cni-irsa"},
            },
            OidcProviders = {
                {"main", new OIDCProviderArgs
                {
                    ProviderArn = "arn:aws:iam::012345678901:oidc-provider/oidc.eks.us-east-1.amazonaws.com/id/5C54DDF35ER19312844C7333374CC09D",
                    NamespaceServiceAccounts = {"default:my-app", "canary:my-app"},
                }},
            },
            Policies = new EKSRolePoliciesArgs
            {
                VpnCni = new EKSVPNCNIPolicyArgs
                {
                    Attach = true,
                    EnableIpv4 = true,
                },
            },
        });

        this.RoleForServiceAccountEks = Output.Create<RoleForServiceAccountsEks>(roleForServiceAccountEks);

        // User
        var user = new User("user", new UserArgs
        {
            Name = "pulumipus",
            ForceDestroy = true,
            PgpKey = "keybase:test",
            PasswordResetRequired = false,
        });

        this.User = Output.Create<User>(user);
    }

    [Output]
    public Output<Account> Account { get; set; }

    [Output]
    public Output<AssumableRole> AssumableRole { get; set; }

    [Output]
    public Output<AssumableRoleWithOIDC> AssumableRoleWithOidc { get; set; }

    [Output]
    public Output<AssumableRoleWithSAML> AssumableRoleWithSaml { get; set; }

    [Output]
    public Output<AssumableRoles> AssumableRoles { get; set; }

    [Output]
    public Output<AssumableRolesWithSAML> AssumableRolesWithSaml { get; set; }

    [Output]
    public Output<EKSRole> EksRole { get; set; }

    [Output]
    public Output<GroupWithAssumableRolesPolicy> GroupWithAssumableRolesPolicy { get; set; }

    [Output]
    public Output<GroupWithPolicies> GroupWithPolicies { get; set; }

    [Output]
    public Output<Policy> Policy { get; set; }

    [Output]
    public Output<ReadOnlyPolicy> ReadOnlyPolicy { get; set; }

    [Output]
    public Output<RoleForServiceAccountsEks> RoleForServiceAccountEks { get; set; }

    [Output]
    public Output<User> User { get; set; }
}
