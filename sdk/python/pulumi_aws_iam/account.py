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

__all__ = ['AccountArgs', 'Account']

@pulumi.input_type
class AccountArgs:
    def __init__(__self__, *,
                 account_alias: pulumi.Input[str],
                 password_policy: pulumi.Input['AccountPasswordPolicyArgs']):
        """
        The set of arguments for constructing a Account resource.
        :param pulumi.Input[str] account_alias: AWS IAM account alias for this account.
        :param pulumi.Input['AccountPasswordPolicyArgs'] password_policy: Options to specify complexity requirements and mandatory rotation periods for your IAM users' passwords. If
               left empty the default AWS password policy will be applied.
        """
        pulumi.set(__self__, "account_alias", account_alias)
        pulumi.set(__self__, "password_policy", password_policy)

    @property
    @pulumi.getter(name="accountAlias")
    def account_alias(self) -> pulumi.Input[str]:
        """
        AWS IAM account alias for this account.
        """
        return pulumi.get(self, "account_alias")

    @account_alias.setter
    def account_alias(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_alias", value)

    @property
    @pulumi.getter(name="passwordPolicy")
    def password_policy(self) -> pulumi.Input['AccountPasswordPolicyArgs']:
        """
        Options to specify complexity requirements and mandatory rotation periods for your IAM users' passwords. If
        left empty the default AWS password policy will be applied.
        """
        return pulumi.get(self, "password_policy")

    @password_policy.setter
    def password_policy(self, value: pulumi.Input['AccountPasswordPolicyArgs']):
        pulumi.set(self, "password_policy", value)


class Account(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_alias: Optional[pulumi.Input[str]] = None,
                 password_policy: Optional[pulumi.Input[pulumi.InputType['AccountPasswordPolicyArgs']]] = None,
                 __props__=None):
        """
        This resource helps you manage an Iam Account's Alias and Password Policy. If your IAM Account Alias was previously
        set (either via the AWS console or when AWS created your Account) you will see an error like the below:

        If you want to manage you Alias using Pulumi you will need to import this resource.

        ## Example Usage
        ## Account

        ```python
        import pulumi
        import pulumi_aws_iam as iam

        account = iam.Account(
            'account',
            account_alias='cool-alias',
            password_policy=iam.AccountPasswordPolicyArgs(
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
        ```
        {{ /example }}

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_alias: AWS IAM account alias for this account.
        :param pulumi.Input[pulumi.InputType['AccountPasswordPolicyArgs']] password_policy: Options to specify complexity requirements and mandatory rotation periods for your IAM users' passwords. If
               left empty the default AWS password policy will be applied.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource helps you manage an Iam Account's Alias and Password Policy. If your IAM Account Alias was previously
        set (either via the AWS console or when AWS created your Account) you will see an error like the below:

        If you want to manage you Alias using Pulumi you will need to import this resource.

        ## Example Usage
        ## Account

        ```python
        import pulumi
        import pulumi_aws_iam as iam

        account = iam.Account(
            'account',
            account_alias='cool-alias',
            password_policy=iam.AccountPasswordPolicyArgs(
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
        ```
        {{ /example }}

        :param str resource_name: The name of the resource.
        :param AccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_alias: Optional[pulumi.Input[str]] = None,
                 password_policy: Optional[pulumi.Input[pulumi.InputType['AccountPasswordPolicyArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccountArgs.__new__(AccountArgs)

            if account_alias is None and not opts.urn:
                raise TypeError("Missing required property 'account_alias'")
            __props__.__dict__["account_alias"] = account_alias
            if password_policy is None and not opts.urn:
                raise TypeError("Missing required property 'password_policy'")
            __props__.__dict__["password_policy"] = password_policy
            __props__.__dict__["arn"] = None
            __props__.__dict__["id"] = None
            __props__.__dict__["password_policy_expire_passwords"] = None
            __props__.__dict__["user_id"] = None
        super(Account, __self__).__init__(
            'aws-iam:index:Account',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The AWS ARN associated with the calling entity.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def id(self) -> pulumi.Output[str]:
        """
        The AWS Account ID number of the account that owns or contains the calling entity.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="passwordPolicyExpirePasswords")
    def password_policy_expire_passwords(self) -> pulumi.Output[bool]:
        """
        Indicates whether passwords in the account expire. Returns true if max password
        age contains a value greater than 0. Returns false if it is 0 or not present.
        """
        return pulumi.get(self, "password_policy_expire_passwords")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        The unique identifier of the calling entity.
        """
        return pulumi.get(self, "user_id")

