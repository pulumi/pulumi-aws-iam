import * as iam from "@pulumi/aws-iam";

// Account
// export const account = new iam.Account("account", {
//     accountAlias: "cool-alias",
//     passwordPolicy: {
//         minimumLength: 37,
//         requireNumbers: false,
//         allowUsersToChange: true,
//         hardExpiry: true,
//         requireSymbols: true,
//         requireLowercaseCharacters: true,
//         requireUppercaseCharacters: true,
//     },
// });

// User
export const user = new iam.User("aws-iam-example-user", {
    name: "pulumipus",
    forceDestroy: true,
    pgpKey: "keybase:test",
    passwordResetRequired: false,
});

// Assumable Role
export const assumableRole = new iam.AssumableRole("aws-iam-example-assumable-role", {
    trustedRoleArns: [ "arn:aws:iam::307990089504:root", "arn:aws:iam::835367859851:user/pulumipus" ],
    role: {
        name: "custom",
        requiresMfa: true,
        policyArns: [ "arn:aws:iam::aws:policy/AmazonCognitoReadOnly","arn:aws:iam::aws:policy/AlexaForBusinessFullAccess" ],
    },
});

// // Assumable Role With OIDC
export const assumableRoleWithOidc = new iam.AssumableRoleWithOIDC("aws-iam-example-assumable-role-with-oidc", {
    providerUrls: ["oidc.eks.eu-west-1.amazonaws.com/id/BA9E170D464AF7B92084EF72A69B9DC8"],
    role: {
        name: "oidc-role",
        policyArns: [ "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy" ],
    },
    tags: {
        Role: "oidc-role",
    },
});

// // Assumable Role With SAML
export const assumableRoleWithSaml = new iam.AssumableRoleWithSAML("aws-iam-example-assumable-role-with-saml", {
    providerIds: [ "arn:aws:iam::235367859851:saml-provider/idp_saml" ],
    role: {
        name: "saml-role",
        policyArns: [ "arn:aws:iam::aws:policy/ReadOnlyAccess" ],
    },
    tags: {
        Role: "saml-role",
    },
});

// // Assumable Roles
export const assumableRoles = new iam.AssumableRoles("aws-iam-example-assumable-roles", {
    trustedRoleArns: [ "arn:aws:iam::307990089504:root", "arn:aws:iam::835367859851:user/pulumipus" ],
    admin: {},
    poweruser: {
        name: "developer",
    },
    readonly: {
        requiresMfa: true,
    },
});

// // Assumable Roles With SAML
export const assumableRolesWithSaml = new iam.AssumableRolesWithSAML("aws-iam-example-assumable-role-with-saml", {
    providerIds: [ "arn:aws:iam::235367859851:saml-provider/idp_saml" ],
    admin: {},
    poweruser: {
        name: "developer",
    },
    readonly: {},
});

// // EKS Role
export const eksRole = new iam.EKSRole("aws-iam-example-eks-role", {
    role: {
        name: "eks-role",
        policyArns: [ "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy" ],
    },
    tags: {
        Name: "eks-role",
    },
    // Uncomment the below and replace actual cluster values.
    // clusterServiceAccounts: {
    //     "cluster1": [ "default:my-app" ],
    //     "cluster2": [ "default:my-app", "canary:my-app" ],
    // },
});

// // Group With Assumable Roles Policy
export const groupWithAssumableRolesPolicy = new iam.GroupWithAssumableRolesPolicy("aws-iam-example-group-with-assumable-roles-policy", {
    name: "production-readonly",
    assumableRoles: [ "arn:aws:iam::835367859855:role/readonly" ],
    groupUsers: [ "pulumipus" ],
}, { dependsOn: [user] });

// // Group With Policies
export const groupWithPolicies = new iam.GroupWithPolicies("aws-iam-example-group-with-policies", {
    name: "superadmins",
    groupUsers: [ "pulumipus" ],
    attachIamSelfManagementPolicy: true,
    customGroupPolicyArns: [ "arn:aws:iam::aws:policy/AdministratorAccess" ],
    customGroupPolicies: [{
        "name": "AllowS3Listing",
        "policy": "{}",
    }],
}, { dependsOn: [user] });

// // Policy
export const policy = new iam.Policy("aws-iam-example-policy", {
    name: "aws-iam-example-policy",
    path: "/",
    description: "My example policy",
    policyDocument: `{
        "Version": "2012-10-17",
        "Statement": [
          {
            "Action": [
              "ec2:Describe*"
            ],
            "Effect": "Allow",
            "Resource": "*"
          }
        ]
    }`,
});

// // Read Only Policy
export const readOnlyPolicy = new iam.ReadOnlyPolicy("aws-iam-example-read-only-policy", {
    name: "aws-iam-example-read-only",
    path: "/",
    description: "My example read only policy",
    allowedServices: [ "rds", "dynamodb" ],
});

// // Role For Service Accounts EKS
export const roleForServiceAccountsEks = new iam.RoleForServiceAccountsEks("aws-iam-example-role-for-service-accounts-eks", {
    role: {
        name: "vpc-cni"
    },
    tags: {
        Name: "vpc-cni-irsa",
    },
    oidcProviders: {
        main: {
            providerArn: "arn:aws:iam::012345678901:oidc-provider/oidc.eks.us-east-1.amazonaws.com/id/5C54DDF35ER19312844C7333374CC09D",
            namespaceServiceAccounts: ["default:my-app", "canary:my-app"],
        }
    },
    policies: {
        vpnCni: {
            attach: true,
            enableIpv4: true,
        },
    },
});
