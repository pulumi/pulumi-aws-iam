# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['GroupWithPoliciesArgs', 'GroupWithPolicies']

@pulumi.input_type
class GroupWithPoliciesArgs:
    def __init__(__self__, *,
                 group_users: pulumi.Input[Sequence[pulumi.Input[str]]],
                 name: pulumi.Input[str],
                 attach_iam_self_management_policy: Optional[pulumi.Input[bool]] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 custom_group_policies: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]]] = None,
                 custom_group_policy_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 iam_self_management_policy_name_prefix: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a GroupWithPolicies resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] group_users: List of IAM users to have in an IAM group which can assume the role.
        :param pulumi.Input[str] name: Name of IAM group.
        :param pulumi.Input[bool] attach_iam_self_management_policy: Whether to attach IAM policy which allows IAM users to manage their credentials and MFA.
        :param pulumi.Input[str] aws_account_id: AWS account id to use inside IAM policies. If empty, current AWS account ID will be used.
        :param pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]] custom_group_policies: List of maps of inline IAM policies to attach to IAM group. Should have `name` and `policy` keys in each element.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] custom_group_policy_arns: List of IAM policies ARNs to attach to IAM group.
        :param pulumi.Input[str] iam_self_management_policy_name_prefix: Name prefix for IAM policy to create with IAM self-management permissions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to add.
        """
        pulumi.set(__self__, "group_users", group_users)
        if name is None:
            name = ''
        pulumi.set(__self__, "name", name)
        if attach_iam_self_management_policy is None:
            attach_iam_self_management_policy = True
        if attach_iam_self_management_policy is not None:
            pulumi.set(__self__, "attach_iam_self_management_policy", attach_iam_self_management_policy)
        if aws_account_id is None:
            aws_account_id = ''
        if aws_account_id is not None:
            pulumi.set(__self__, "aws_account_id", aws_account_id)
        if custom_group_policies is not None:
            pulumi.set(__self__, "custom_group_policies", custom_group_policies)
        if custom_group_policy_arns is not None:
            pulumi.set(__self__, "custom_group_policy_arns", custom_group_policy_arns)
        if iam_self_management_policy_name_prefix is None:
            iam_self_management_policy_name_prefix = 'IAMSelfManagement-'
        if iam_self_management_policy_name_prefix is not None:
            pulumi.set(__self__, "iam_self_management_policy_name_prefix", iam_self_management_policy_name_prefix)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="groupUsers")
    def group_users(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of IAM users to have in an IAM group which can assume the role.
        """
        return pulumi.get(self, "group_users")

    @group_users.setter
    def group_users(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "group_users", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of IAM group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="attachIamSelfManagementPolicy")
    def attach_iam_self_management_policy(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to attach IAM policy which allows IAM users to manage their credentials and MFA.
        """
        return pulumi.get(self, "attach_iam_self_management_policy")

    @attach_iam_self_management_policy.setter
    def attach_iam_self_management_policy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "attach_iam_self_management_policy", value)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        AWS account id to use inside IAM policies. If empty, current AWS account ID will be used.
        """
        return pulumi.get(self, "aws_account_id")

    @aws_account_id.setter
    def aws_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_account_id", value)

    @property
    @pulumi.getter(name="customGroupPolicies")
    def custom_group_policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]]]:
        """
        List of maps of inline IAM policies to attach to IAM group. Should have `name` and `policy` keys in each element.
        """
        return pulumi.get(self, "custom_group_policies")

    @custom_group_policies.setter
    def custom_group_policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]]]):
        pulumi.set(self, "custom_group_policies", value)

    @property
    @pulumi.getter(name="customGroupPolicyArns")
    def custom_group_policy_arns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of IAM policies ARNs to attach to IAM group.
        """
        return pulumi.get(self, "custom_group_policy_arns")

    @custom_group_policy_arns.setter
    def custom_group_policy_arns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "custom_group_policy_arns", value)

    @property
    @pulumi.getter(name="iamSelfManagementPolicyNamePrefix")
    def iam_self_management_policy_name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Name prefix for IAM policy to create with IAM self-management permissions.
        """
        return pulumi.get(self, "iam_self_management_policy_name_prefix")

    @iam_self_management_policy_name_prefix.setter
    def iam_self_management_policy_name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iam_self_management_policy_name_prefix", value)

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


class GroupWithPolicies(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attach_iam_self_management_policy: Optional[pulumi.Input[bool]] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 custom_group_policies: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]]] = None,
                 custom_group_policy_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 group_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 iam_self_management_policy_name_prefix: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        This resources allows you to create an IAM group with specified IAM policies,
        and then add specified users into your created group.

        ## Example Usage
        ## Group With Policies

        ```python
        import pulumi
        import pulumi_aws_iam as iam

        group_with_policies = iam.GroupWithPolicies(
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
        ```
        {{ /example }}

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] attach_iam_self_management_policy: Whether to attach IAM policy which allows IAM users to manage their credentials and MFA.
        :param pulumi.Input[str] aws_account_id: AWS account id to use inside IAM policies. If empty, current AWS account ID will be used.
        :param pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]] custom_group_policies: List of maps of inline IAM policies to attach to IAM group. Should have `name` and `policy` keys in each element.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] custom_group_policy_arns: List of IAM policies ARNs to attach to IAM group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] group_users: List of IAM users to have in an IAM group which can assume the role.
        :param pulumi.Input[str] iam_self_management_policy_name_prefix: Name prefix for IAM policy to create with IAM self-management permissions.
        :param pulumi.Input[str] name: Name of IAM group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to add.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupWithPoliciesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resources allows you to create an IAM group with specified IAM policies,
        and then add specified users into your created group.

        ## Example Usage
        ## Group With Policies

        ```python
        import pulumi
        import pulumi_aws_iam as iam

        group_with_policies = iam.GroupWithPolicies(
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
        ```
        {{ /example }}

        :param str resource_name: The name of the resource.
        :param GroupWithPoliciesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupWithPoliciesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attach_iam_self_management_policy: Optional[pulumi.Input[bool]] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 custom_group_policies: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]]] = None,
                 custom_group_policy_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 group_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 iam_self_management_policy_name_prefix: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
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
            __props__ = GroupWithPoliciesArgs.__new__(GroupWithPoliciesArgs)

            if attach_iam_self_management_policy is None:
                attach_iam_self_management_policy = True
            __props__.__dict__["attach_iam_self_management_policy"] = attach_iam_self_management_policy
            if aws_account_id is None:
                aws_account_id = ''
            __props__.__dict__["aws_account_id"] = aws_account_id
            __props__.__dict__["custom_group_policies"] = custom_group_policies
            __props__.__dict__["custom_group_policy_arns"] = custom_group_policy_arns
            if group_users is None and not opts.urn:
                raise TypeError("Missing required property 'group_users'")
            __props__.__dict__["group_users"] = group_users
            if iam_self_management_policy_name_prefix is None:
                iam_self_management_policy_name_prefix = 'IAMSelfManagement-'
            __props__.__dict__["iam_self_management_policy_name_prefix"] = iam_self_management_policy_name_prefix
            if name is None:
                name = ''
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["group_arn"] = None
            __props__.__dict__["group_name"] = None
        super(GroupWithPolicies, __self__).__init__(
            'aws-iam:index:GroupWithPolicies',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> pulumi.Output[str]:
        """
        IAM AWS account id.
        """
        return pulumi.get(self, "aws_account_id")

    @property
    @pulumi.getter(name="groupArn")
    def group_arn(self) -> pulumi.Output[str]:
        """
        IAM group arn.
        """
        return pulumi.get(self, "group_arn")

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Output[str]:
        """
        IAM group name.
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter(name="groupUsers")
    def group_users(self) -> pulumi.Output[Sequence[str]]:
        """
        List of IAM users in IAM group
        """
        return pulumi.get(self, "group_users")

