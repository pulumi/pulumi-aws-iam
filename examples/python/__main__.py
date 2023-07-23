"""An AWS Python Pulumi program"""

import json
import pulumi
import pulumi_aws as aws

# Account
account = aws.iam.Account(
    'account',
    account_alias='cool-alias',
    password_policy=aws.iam.AccountPasswordPolicyArgs(
        minimum_length=37,
        require_numbers=False,
        allow_users_to_change=True,
        hard_expiry=True,
        require_symbols=True,
        require_lowercase_characters=True,
        require_uppercase_characters=True,
    )
)

pulumi.export('account', account)

# Assumable Role
assumable_role = aws.iam.AssumableRole(
    'assumable_role',
    trusted_role_arns=['arn:aws:iam::307990089504:root','arn:aws:iam::835367859851:user/pulumipus'],
    role=iam.RoleWithMFAArgs(
        name='custom',
        requires_mfa=True,
        policy_arns=['arn:aws:iam::aws:policy/AmazonCognitoReadOnly','arn:aws:iam::aws:policy/AlexaForBusinessFullAccess'],
    ),
)

pulumi.export('assumable_role', assumable_role)

# Assumable Role With OIDC
assumable_role_with_oidc = aws.iam.AssumableRoleWithOIDC(
    'assumable_role_with_oidc',
    role=aws.iam.RoleArgs(
        name='oidc-role',
        policy_arns=['arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy']
    ),
    tags={
        'Role': 'oidc-role',
    },
    provider_urls=['oidc.eks.eu-west-1.amazonaws.com/id/BA9E170D464AF7B92084EF72A69B9DC8']
)

pulumi.export('assumable_role_with_oidc', assumable_role_with_oidc)

# Assumable Role With SAML
assumable_role_with_saml = aws.iam.AssumableRoleWithSAML(
    'assumable_role_with_saml',
    role=iam.RoleArgs(
        name='saml-role',
        policy_arns=['arn:aws:iam::aws:policy/ReadOnlyAccess'],
    ),
    tags={
        'Role': 'saml-role',
    },
    provider_ids=['arn:aws:iam::235367859851:saml-provider/idp_saml']
)

pulumi.export('assumable_role_with_saml', assumable_role_with_saml)

# Assumable Roles
assumable_roles = aws.iam.AssumableRoles(
    'assumable_roles',
    trusted_role_arns=['arn:aws:iam::307990089504:root','arn:aws:iam::835367859851:user/anton'],
    admin=aws.iam.AdminRoleArgs(),
    poweruser=iam.PoweruserRoleArgs(
        name='developer',
    ),
    readonly=aws.iam.ReadonlyRoleWithMFAArgs(
        requires_mfa=True,
    ),
)

pulumi.export('assumable_roles', assumable_roles)

# Assumable Roles With SAML
assumable_roles_with_saml = aws.iam.AssumableRolesWithSAML(
    'assumable_roles_with_saml',
    provider_ids=['arn:aws:iam::235367859851:saml-provider/idp_saml'],
    admin=aws.iam.AdminRoleArgs(),
    readonly=aws.iam.ReadonlyRoleArgs(),
    poweruser=aws.iam.PoweruserRoleArgs(
        name='developer',
    ),
)

pulumi.export('assumable_roles_with_saml', assumable_roles_with_saml)

# EKS Role
eks_role = aws.iam.EKSRole(
    'eks_role',
    role=aws.iam.RoleArgs(
        name='eks-role',
        policy_arns=['arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy'],
    ),
    tags={
        'Name': 'eks-role',
    },
    # cluster_service_acccounts={
    #     'cluster1': [ 'default:my-app' ],
    #     'cluster2': [ 'default:my-app', 'canary:my-app' ],
    # },
)

pulumi.export('eks_role', eks_role)

# Group With Assumable Roles Policy
group_with_assume_roles_policy = aws.iam.GroupWithAssumableRolesPolicy(
    'group_with_assume_roles_policy',
    name='production-readonly',
    assumable_roles=['arn:aws:iam::835367859855:role/readonly'],
    group_users=['user1','user2'],
)

pulumi.export('group_with_assume_roles_policy', group_with_assume_roles_policy)

# Group With Policies
group_with_policies = aws.iam.GroupWithPolicies(
    'group_with_policies',
    name='superadmins',
    group_users=['user1','user2'],
    attach_iam_self_management_policy=True,
    custom_group_policy_arns=['arn:aws:iam::aws:policy/AdministratorAccess'],
    custom_group_policies=[{
        'name': 'AllowS3Listing',
        'policy': '{}',
    }],
)

pulumi.export('group_with_policies', group_with_policies)

# Policy
policy = iam.Policy(
    'policy',
    name='example',
    path='/',
    description='My example policy',
    policy_document=json.dumps({
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
    })
)

pulumi.export('policy', policy)

# Read Only Policy
read_only_policy = aws.iam.ReadOnlyPolicy(
    'read_only_policy',
    name='example',
    path='/',
    description='My example read only policy',
    allowed_services=['rds','dynamo'],
)

pulumi.export('read_only_policy', read_only_policy)

# Role For Service Accounts EKS
role_for_service_account_eks = aws.iam.RoleForServiceAccountsEks(
    'role_for_service_account_eks',
    role=aws.iam.RoleArgs(
        name='vpc-cni'
    ),
    tags={
        'Name': 'vpc-cni-irsa',
    },
    oidc_providers={
        'main': aws.iam.OIDCProviderArgs(
            provider_arn='arn:aws:iam::012345678901:oidc-provider/oidc.eks.us-east-1.amazonaws.com/id/5C54DDF35ER19312844C7333374CC09D',
            namespace_service_accounts=['default:my-app', 'canary:my-app'],
        ),
    },
    policies=aws.iam.EKSRolePoliciesArgs(
        vpn_cni=aws.iam.EKSVPNCNIPolicyArgs(
            attach=True,
            enable_ipv4=True,
        ),
    ),
)

pulumi.export('role_for_service_account_eks', role_for_service_account_eks)

# User
user = aws.iam.User(
    'user',
    name='pulumipus',
    force_destroy=True,
)

pulumi.export('user', user)
