# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._inputs import *

__all__ = ['AssumableRoleWithOIDCArgs', 'AssumableRoleWithOIDC']

@pulumi.input_type
class AssumableRoleWithOIDCArgs:
    def __init__(__self__, *,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 force_detach_policies: Optional[pulumi.Input[bool]] = None,
                 max_session_duration: Optional[pulumi.Input[int]] = None,
                 oidc_fully_qualified_audiences: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 oidc_fully_qualified_subjects: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 oidc_subjects_with_wildcards: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 provider_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input['RoleArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a AssumableRoleWithOIDC resource.
        :param pulumi.Input[str] aws_account_id: The AWS account ID where the OIDC provider lives, leave empty to use the account for the AWS provider.
        :param pulumi.Input[bool] force_detach_policies: Whether policies should be detached from this role when destroying.
        :param pulumi.Input[int] max_session_duration: Maximum CLI/API session duration in seconds between 3600 and 43200.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] oidc_fully_qualified_audiences: The audience to be added to the role policy. Set to sts.amazonaws.com for cross-account assumable role. Leave empty otherwise.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] oidc_fully_qualified_subjects: The fully qualified OIDC subjects to be added to the role policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] oidc_subjects_with_wildcards: The OIDC subject using wildcards to be added to the role policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] provider_urls: List of URLs of the OIDC Providers.
        :param pulumi.Input['RoleArgs'] role: The IAM role.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to add.
        """
        if aws_account_id is None:
            aws_account_id = ''
        if aws_account_id is not None:
            pulumi.set(__self__, "aws_account_id", aws_account_id)
        if force_detach_policies is None:
            force_detach_policies = False
        if force_detach_policies is not None:
            pulumi.set(__self__, "force_detach_policies", force_detach_policies)
        if max_session_duration is None:
            max_session_duration = 3600
        if max_session_duration is not None:
            pulumi.set(__self__, "max_session_duration", max_session_duration)
        if oidc_fully_qualified_audiences is not None:
            pulumi.set(__self__, "oidc_fully_qualified_audiences", oidc_fully_qualified_audiences)
        if oidc_fully_qualified_subjects is not None:
            pulumi.set(__self__, "oidc_fully_qualified_subjects", oidc_fully_qualified_subjects)
        if oidc_subjects_with_wildcards is not None:
            pulumi.set(__self__, "oidc_subjects_with_wildcards", oidc_subjects_with_wildcards)
        if provider_urls is not None:
            pulumi.set(__self__, "provider_urls", provider_urls)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS account ID where the OIDC provider lives, leave empty to use the account for the AWS provider.
        """
        return pulumi.get(self, "aws_account_id")

    @aws_account_id.setter
    def aws_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_account_id", value)

    @property
    @pulumi.getter(name="forceDetachPolicies")
    def force_detach_policies(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether policies should be detached from this role when destroying.
        """
        return pulumi.get(self, "force_detach_policies")

    @force_detach_policies.setter
    def force_detach_policies(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_detach_policies", value)

    @property
    @pulumi.getter(name="maxSessionDuration")
    def max_session_duration(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum CLI/API session duration in seconds between 3600 and 43200.
        """
        return pulumi.get(self, "max_session_duration")

    @max_session_duration.setter
    def max_session_duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_session_duration", value)

    @property
    @pulumi.getter(name="oidcFullyQualifiedAudiences")
    def oidc_fully_qualified_audiences(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The audience to be added to the role policy. Set to sts.amazonaws.com for cross-account assumable role. Leave empty otherwise.
        """
        return pulumi.get(self, "oidc_fully_qualified_audiences")

    @oidc_fully_qualified_audiences.setter
    def oidc_fully_qualified_audiences(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "oidc_fully_qualified_audiences", value)

    @property
    @pulumi.getter(name="oidcFullyQualifiedSubjects")
    def oidc_fully_qualified_subjects(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The fully qualified OIDC subjects to be added to the role policy.
        """
        return pulumi.get(self, "oidc_fully_qualified_subjects")

    @oidc_fully_qualified_subjects.setter
    def oidc_fully_qualified_subjects(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "oidc_fully_qualified_subjects", value)

    @property
    @pulumi.getter(name="oidcSubjectsWithWildcards")
    def oidc_subjects_with_wildcards(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The OIDC subject using wildcards to be added to the role policy.
        """
        return pulumi.get(self, "oidc_subjects_with_wildcards")

    @oidc_subjects_with_wildcards.setter
    def oidc_subjects_with_wildcards(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "oidc_subjects_with_wildcards", value)

    @property
    @pulumi.getter(name="providerUrls")
    def provider_urls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of URLs of the OIDC Providers.
        """
        return pulumi.get(self, "provider_urls")

    @provider_urls.setter
    def provider_urls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "provider_urls", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input['RoleArgs']]:
        """
        The IAM role.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input['RoleArgs']]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to add.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class AssumableRoleWithOIDC(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 force_detach_policies: Optional[pulumi.Input[bool]] = None,
                 max_session_duration: Optional[pulumi.Input[int]] = None,
                 oidc_fully_qualified_audiences: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 oidc_fully_qualified_subjects: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 oidc_subjects_with_wildcards: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 provider_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[pulumi.InputType['RoleArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        This resources helps you create a single IAM role which can be assume by trusted
        resources using OpenID Connect Federated Users.

        ## Example Usage
        ## Assumable Role With OIDC

        ```python
        import pulumi
        import pulumi_aws_iam as iam

        assumable_role_with_oidc = iam.AssumableRoleWithOIDC(
            'assumable_role_with_oidc',
            role=iam.RoleArgs(
                name='oidc-role',
                policy_arns=['arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy']
            ),
            tags={
                'Role': 'oidc-role',
            },
            provider_urls=['oidc.eks.eu-west-1.amazonaws.com/id/BA9E170D464AF7B92084EF72A69B9DC8']
        )

        pulumi.export('assumable_role_with_oidc', assumable_role_with_oidc)
        ```
        {{ /example }}

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_account_id: The AWS account ID where the OIDC provider lives, leave empty to use the account for the AWS provider.
        :param pulumi.Input[bool] force_detach_policies: Whether policies should be detached from this role when destroying.
        :param pulumi.Input[int] max_session_duration: Maximum CLI/API session duration in seconds between 3600 and 43200.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] oidc_fully_qualified_audiences: The audience to be added to the role policy. Set to sts.amazonaws.com for cross-account assumable role. Leave empty otherwise.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] oidc_fully_qualified_subjects: The fully qualified OIDC subjects to be added to the role policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] oidc_subjects_with_wildcards: The OIDC subject using wildcards to be added to the role policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] provider_urls: List of URLs of the OIDC Providers.
        :param pulumi.Input[pulumi.InputType['RoleArgs']] role: The IAM role.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to add.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AssumableRoleWithOIDCArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resources helps you create a single IAM role which can be assume by trusted
        resources using OpenID Connect Federated Users.

        ## Example Usage
        ## Assumable Role With OIDC

        ```python
        import pulumi
        import pulumi_aws_iam as iam

        assumable_role_with_oidc = iam.AssumableRoleWithOIDC(
            'assumable_role_with_oidc',
            role=iam.RoleArgs(
                name='oidc-role',
                policy_arns=['arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy']
            ),
            tags={
                'Role': 'oidc-role',
            },
            provider_urls=['oidc.eks.eu-west-1.amazonaws.com/id/BA9E170D464AF7B92084EF72A69B9DC8']
        )

        pulumi.export('assumable_role_with_oidc', assumable_role_with_oidc)
        ```
        {{ /example }}

        :param str resource_name: The name of the resource.
        :param AssumableRoleWithOIDCArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AssumableRoleWithOIDCArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 force_detach_policies: Optional[pulumi.Input[bool]] = None,
                 max_session_duration: Optional[pulumi.Input[int]] = None,
                 oidc_fully_qualified_audiences: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 oidc_fully_qualified_subjects: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 oidc_subjects_with_wildcards: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 provider_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[pulumi.InputType['RoleArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AssumableRoleWithOIDCArgs.__new__(AssumableRoleWithOIDCArgs)

            if aws_account_id is None:
                aws_account_id = ''
            __props__.__dict__["aws_account_id"] = aws_account_id
            if force_detach_policies is None:
                force_detach_policies = False
            __props__.__dict__["force_detach_policies"] = force_detach_policies
            if max_session_duration is None:
                max_session_duration = 3600
            __props__.__dict__["max_session_duration"] = max_session_duration
            __props__.__dict__["oidc_fully_qualified_audiences"] = oidc_fully_qualified_audiences
            __props__.__dict__["oidc_fully_qualified_subjects"] = oidc_fully_qualified_subjects
            __props__.__dict__["oidc_subjects_with_wildcards"] = oidc_subjects_with_wildcards
            __props__.__dict__["provider_urls"] = provider_urls
            __props__.__dict__["role"] = role
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["path"] = None
            __props__.__dict__["unique_id"] = None
        super(AssumableRoleWithOIDC, __self__).__init__(
            'aws-iam:index:AssumableRoleWithOIDC',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        ARN of IAM role.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of IAM role.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        Path of IAM role.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="uniqueId")
    def unique_id(self) -> pulumi.Output[str]:
        """
        Unique ID of IAM role.
        """
        return pulumi.get(self, "unique_id")

