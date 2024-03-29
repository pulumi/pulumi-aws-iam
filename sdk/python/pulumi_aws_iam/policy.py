# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PolicyArgs', 'Policy']

@pulumi.input_type
class PolicyArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 policy_document: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Policy resource.
        :param pulumi.Input[str] name: The name of the policy.
        :param pulumi.Input[str] policy_document: The policy document.
        :param pulumi.Input[str] description: The description of the policy.
        :param pulumi.Input[str] path: The path of the policy in IAM.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to add.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "policy_document", policy_document)
        if description is None:
            description = 'IAM Policy'
        if description is not None:
            pulumi.set(__self__, "description", description)
        if path is None:
            path = '/'
        if path is not None:
            pulumi.set(__self__, "path", path)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="policyDocument")
    def policy_document(self) -> pulumi.Input[str]:
        """
        The policy document.
        """
        return pulumi.get(self, "policy_document")

    @policy_document.setter
    def policy_document(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_document", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        The path of the policy in IAM.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

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


class Policy(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 policy_document: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        This resource helps you create an IAM policy.

        ## Example Usage
        ## Policy

        ```python
        import json
        import pulumi
        import pulumi_aws_iam as iam

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
        ```
        {{ /example }}

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the policy.
        :param pulumi.Input[str] name: The name of the policy.
        :param pulumi.Input[str] path: The path of the policy in IAM.
        :param pulumi.Input[str] policy_document: The policy document.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to add.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource helps you create an IAM policy.

        ## Example Usage
        ## Policy

        ```python
        import json
        import pulumi
        import pulumi_aws_iam as iam

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
        ```
        {{ /example }}

        :param str resource_name: The name of the resource.
        :param PolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 policy_document: Optional[pulumi.Input[str]] = None,
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
            __props__ = PolicyArgs.__new__(PolicyArgs)

            if description is None:
                description = 'IAM Policy'
            __props__.__dict__["description"] = description
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            if path is None:
                path = '/'
            __props__.__dict__["path"] = path
            if policy_document is None and not opts.urn:
                raise TypeError("Missing required property 'policy_document'")
            __props__.__dict__["policy_document"] = policy_document
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["id"] = None
        super(Policy, __self__).__init__(
            'aws-iam:index:Policy',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN assigned by AWS to this policy.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> pulumi.Output[str]:
        """
        The policy's ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the policy.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        The path of the policy in IAM.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="policyDocument")
    def policy_document(self) -> pulumi.Output[str]:
        """
        The policy document.
        """
        return pulumi.get(self, "policy_document")

