// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsIam
{
    /// <summary>
    /// This resources allows you to create an IAM group with specified IAM policies,
    /// and then add specified users into your created group.
    /// 
    /// ## Example Usage
    /// ## Group With Policies
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Pulumi.AwsIam;
    /// using Pulumi.AwsIam.Inputs;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var groupWithPolicies = new GroupWithPolicies("group-with-policies", new GroupWithPoliciesArgs
    ///         {
    ///             Name = "superadmins",
    ///             GroupUsers = {"user1", "user2"},
    ///             AttachIamSelfManagementPolicy = true,
    ///             CustomGroupPolicyArns = {"arn:aws:iam::aws:policy/AdministratorAccess"},
    ///             CustomGroupPolicies = new InputList&lt;ImmutableDictionary&lt;string, string&gt;&gt;
    ///             {
    ///                 ImmutableDictionary.Create&lt;string, string&gt;()
    ///                     .Add("name", "AllowS3Listing")
    ///                     .Add("policy", "{}"),
    ///             },
    ///         });
    /// 
    ///         this.GroupWithPolicies = Output.Create&lt;GroupWithPolicies&gt;(groupWithPolicies);
    ///     }
    /// 
    ///     [Output]
    ///     public Output&lt;GroupWithPolicies&gt; GroupWithPolicies { get; set; }
    /// }
    /// ```
    /// {{ /example }}
    /// </summary>
    [AwsIamResourceType("aws-iam:index:GroupWithPolicies")]
    public partial class GroupWithPolicies : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// IAM AWS account id.
        /// </summary>
        [Output("awsAccountId")]
        public Output<string> AwsAccountId { get; private set; } = null!;

        /// <summary>
        /// IAM group arn.
        /// </summary>
        [Output("groupArn")]
        public Output<string> GroupArn { get; private set; } = null!;

        /// <summary>
        /// IAM group name.
        /// </summary>
        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;

        /// <summary>
        /// List of IAM users in IAM group
        /// </summary>
        [Output("groupUsers")]
        public Output<ImmutableArray<string>> GroupUsers { get; private set; } = null!;


        /// <summary>
        /// Create a GroupWithPolicies resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupWithPolicies(string name, GroupWithPoliciesArgs args, ComponentResourceOptions? options = null)
            : base("aws-iam:index:GroupWithPolicies", name, args ?? new GroupWithPoliciesArgs(), MakeResourceOptions(options, ""), remote: true)
        {
        }

        private static ComponentResourceOptions MakeResourceOptions(ComponentResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ComponentResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = ComponentResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class GroupWithPoliciesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to attach IAM policy which allows IAM users to manage their credentials and MFA.
        /// </summary>
        [Input("attachIamSelfManagementPolicy")]
        public Input<bool>? AttachIamSelfManagementPolicy { get; set; }

        /// <summary>
        /// AWS account id to use inside IAM policies. If empty, current AWS account ID will be used.
        /// </summary>
        [Input("awsAccountId")]
        public Input<string>? AwsAccountId { get; set; }

        [Input("customGroupPolicies")]
        private InputList<ImmutableDictionary<string, string>>? _customGroupPolicies;

        /// <summary>
        /// List of maps of inline IAM policies to attach to IAM group. Should have `name` and `policy` keys in each element.
        /// </summary>
        public InputList<ImmutableDictionary<string, string>> CustomGroupPolicies
        {
            get => _customGroupPolicies ?? (_customGroupPolicies = new InputList<ImmutableDictionary<string, string>>());
            set => _customGroupPolicies = value;
        }

        [Input("customGroupPolicyArns")]
        private InputList<string>? _customGroupPolicyArns;

        /// <summary>
        /// List of IAM policies ARNs to attach to IAM group.
        /// </summary>
        public InputList<string> CustomGroupPolicyArns
        {
            get => _customGroupPolicyArns ?? (_customGroupPolicyArns = new InputList<string>());
            set => _customGroupPolicyArns = value;
        }

        [Input("groupUsers", required: true)]
        private InputList<string>? _groupUsers;

        /// <summary>
        /// List of IAM users to have in an IAM group which can assume the role.
        /// </summary>
        public InputList<string> GroupUsers
        {
            get => _groupUsers ?? (_groupUsers = new InputList<string>());
            set => _groupUsers = value;
        }

        /// <summary>
        /// Name prefix for IAM policy to create with IAM self-management permissions.
        /// </summary>
        [Input("iamSelfManagementPolicyNamePrefix")]
        public Input<string>? IamSelfManagementPolicyNamePrefix { get; set; }

        /// <summary>
        /// Name of IAM group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to add.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GroupWithPoliciesArgs()
        {
            AttachIamSelfManagementPolicy = true;
            AwsAccountId = "";
            IamSelfManagementPolicyNamePrefix = "IAMSelfManagement-";
            Name = "";
        }
        public static new GroupWithPoliciesArgs Empty => new GroupWithPoliciesArgs();
    }
}
