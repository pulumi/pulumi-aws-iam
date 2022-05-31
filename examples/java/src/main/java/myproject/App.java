package myproject;

import java.util.Arrays;
import java.util.HashMap;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.awsiam.*;
import com.pulumi.awsiam.inputs.*;

public class App {
    public static void main(String[] args) {
        Pulumi.run(ctx -> {
            // Account
            var account = new Account("account", AccountArgs.builder()
                .accountAlias("cool-alias")
                .passwordPolicy(AccountPasswordPolicyArgs.builder()
                    .minimumLength(37)
                    .requireNumbers(false)
                    .allowUsersToChange(true)
                    .hardExpiry(true)
                    .requireSymbols(true)
                    .requireLowercaseCharacters(true)
                    .requireUppercaseCharacters(true)
                    .build())
                .build());

            ctx.export("account", Output.of(account));

            // Assumable Role
            var assumableRole = new AssumableRole("assumable-role", AssumableRoleArgs.builder()
                .trustedRoleActions("arn:aws:iam::307990089504:root", "arn:aws:iam::835367859851:user/pulumipus")
                .role(RoleWithMFAArgs.builder()
                    .name("custom")
                    .requiresMfa(true)
                    .policyArns("arn:aws:iam::aws:policy/AmazonCognitoReadOnly","arn:aws:iam::aws:policy/AlexaForBusinessFullAccess")
                    .build())
                .build());

            ctx.export("assumableRole", Output.of(assumableRole));

            // Assumable Role With OIDC
            var assumableRoleWithOidc = new AssumableRoleWithOIDC("assumable-role-with-oidc", AssumableRoleWithOIDCArgs.builder()
                .role(RoleArgs.builder()
                    .name("oidc-role")
                    .policyArns("arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy")
                    .build())
                .tags(new HashMap<String, String>() {{
                    put("Role", "oidc-role");
                }})
                .providerUrls("oidc.eks.eu-west-1.amazonaws.com/id/BA9E170D464AF7B92084EF72A69B9DC8")
                .build());

            ctx.export("assumableRoleWithOidc", Output.of(assumableRoleWithOidc));

            // Assumable Role With SAML
            var assumableRoleWithSaml = new AssumableRoleWithSAML("assumable-role-with-saml", AssumableRoleWithSAMLArgs.builder()
                .role(RoleArgs.builder()
                    .name("saml-role")
                    .policyArns("arn:aws:iam::aws:policy/ReadOnlyAccess")
                    .build())
                .tags(new HashMap<String, String>() {{
                    put("Role", "saml-role");
                }})
                .providerIds("arn:aws:iam::235367859851:saml-provider/idp_saml")
                .build());

            ctx.export("assumableRoleWithSaml", Output.of(assumableRoleWithSaml));

            // Assumable Roles
            var assumableRoles = new AssumableRoles("assumable-roles", AssumableRolesArgs.builder()
                .trustedRoleArns("arn:aws:iam::307990089504:root", "arn:aws:iam::835367859851:user/anton")
                .admin(AdminRoleWithMFAArgs.builder().build())
                .poweruser(PoweruserRoleWithMFAArgs.builder()
                    .name("developer")
                    .build())
                .readonly(ReadonlyRoleWithMFAArgs.builder()
                    .requiresMfa(true)
                    .build())
                .build());

            ctx.export("assumableRoles", Output.of(assumableRoles));

            // Assumable Roles with SAML
            var assumableRolesWithSaml = new AssumableRolesWithSAML("assumable-roles-with-saml", AssumableRolesWithSAMLArgs.builder()
                .providerIds("arn:aws:iam::235367859851:saml-provider/idp_saml")
                .admin(AdminRoleArgs.builder().build())
                .readonly(ReadonlyRoleArgs.builder().build())
                .poweruser(PoweruserRoleArgs.builder()
                    .name("developer")
                    .build())
                .build());

            ctx.export("assumableRolesWithSaml", Output.of(assumableRolesWithSaml));

            // EKS Role
            var eksRole = new EKSRole("eks-role", EKSRoleArgs.builder()
                .role(RoleArgs.builder()
                    .name("eks-role")
                    .policyArns("arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy")
                    .build())
                .tags(new HashMap<String, String>() {{
                    put("Name", "eks-role");
                }})
                // Uncomment the below and replace actual cluster values.
                // .clusterServiceAccounts(new HashMap<String, List<String>>() {{
                //     put("cluster1", Arrays.asList("default:my-app"));
                //     put("cluster2", Arrays.asList("default:my-app", "canary:my-app"));
                // }})
                .build());

            ctx.export("eksRole", Output.of(eksRole));

            // Group With Assumable Roles Policy
            var groupWithAssumableRolesPolicy = new GroupWithAssumableRolesPolicy("group-with-assumable-roles-policy", GroupWithAssumableRolesPolicyArgs.builder()
                .name("production-readonly")
                .assumableRoles("arn:aws:iam::835367859855:role/readonly")
                .groupUsers("user1", "user2")
                .build());

            ctx.export("groupWithAssumableRolesPolicy", Output.of(groupWithAssumableRolesPolicy));

            // Group With Policies
            var groupWithPolicies = new GroupWithPolicies("group-with-policies", GroupWithPoliciesArgs.builder()
                .name("superadmins")
                .groupUsers("user1", "user2")
                .attachIamSelfManagementPolicy(true)
                .customGroupPolicyArns("arn:aws:iam::aws:policy/AdministratorAccess")
                .customGroupPolicies(Arrays.asList(
                    new HashMap<String, String>() {{
                        put("name", "AllowS3Listing");
                        put("policy", "{}");
                    }}
                ))
                .build());

            ctx.export("groupWithPolicies", Output.of(groupWithPolicies));

            // Policy
            var policy = new Policy("policy", PolicyArgs.builder()
                .name("example")
                .path("/")
                .description("My example policy")
                .policyDocument("{" +
                    "\"Version\": \"2012-10-17\"," +
                    "\"Statement\": [" +
                      "{" +
                        "\"Action\": [" +
                          "\"ec2:Describe*\"" +
                        "]," +
                        "\"Effect\": \"Allow\"," +
                        "\"Resource\": \"*\"" +
                      "}" +
                    "]" +
                "}")
                .build());

            ctx.export("policy", Output.of(policy));

            // Read Only Policy
            var readOnlyPolicy = new ReadOnlyPolicy("read-only-policy", ReadOnlyPolicyArgs.builder()
                .name("example")
                .path("/")
                .description("My example read only policy")
                .allowedServices("rds", "dynamo")
                .build());

            ctx.export("readOnlyPolicy", Output.of(readOnlyPolicy));

            // Role For Service Accounts EKS
            var roleForServiceAccountsEks = new RoleForServiceAccountsEks("role-for-eks-service-account", RoleForServiceAccountsEksArgs.builder()
                .role(EKSServiceAccountRoleArgs.builder()
                    .name("vpc-cni")
                    .build())
                .tags(new HashMap<String, String>() {{
                    put("Name", "vpc-cni-irsa");
                }})
                .oidcProviders(new HashMap<String, OIDCProviderArgs>() {{
                    put("main", OIDCProviderArgs.builder()
                        .providerArn("arn:aws:iam::012345678901:oidc-provider/oidc.eks.us-east-1.amazonaws.com/id/5C54DDF35ER19312844C7333374CC09D")
                        .namespaceServiceAccounts("default:my-app", "canary:my-app")
                        .build());
                }})
                .policies(EKSRolePoliciesArgs.builder()
                    .vpnCni(EKSVPNCNIPolicyArgs.builder()
                        .attach(true)
                        .enableIpv4(true)
                        .build())
                    .build())
                .build());

            ctx.export("roleForServiceAccountsEks", Output.of(roleForServiceAccountsEks));

            // User
            var user = new User("user", UserArgs.builder()
                .name("pulumipus")
                .forceDestroy(true)
                .pgpKey("keybase:test")
                .passwordResetRequired(true)
                .build());

            ctx.export("user", Output.of(user));
        });
    }
}
