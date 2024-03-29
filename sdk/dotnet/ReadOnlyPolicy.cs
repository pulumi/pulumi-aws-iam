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
    /// This resource helps you create an IAM read-only policy for the services you specify. The default AWS
    /// read-only policies may not include services you need or may contain services you do not need access to.
    /// This resource helps ensure your read-only policy has permissions to exactly what you specify.
    /// 
    /// ## Example Usage
    /// ## RDS and Dynamo Read Only Policy
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
    ///         var readOnlyPolicy = new ReadOnlyPolicy("read-only-policy", new ReadOnlyPolicyArgs
    ///         {
    ///             Name = "example",
    ///             Path = "/",
    ///             Description = "My example read only policy",
    ///             AllowedServices = {"rds", "dynamodb"},
    ///         });
    /// 
    ///         this.ReadOnlyPolicy = Output.Create&lt;ReadOnlyPolicy&gt;(readOnlyPolicy);
    ///     }
    /// 
    ///     [Output]
    ///     public Output&lt;ReadOnlyPolicy&gt; ReadOnlyPolicy { get; set; }
    /// }
    /// ```
    /// {{ /example }}
    /// </summary>
    [AwsIamResourceType("aws-iam:index:ReadOnlyPolicy")]
    public partial class ReadOnlyPolicy : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// The ARN assigned by AWS to this policy.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The description of the policy.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The policy's ID.
        /// </summary>
        [Output("id")]
        public Output<string> Id { get; private set; } = null!;

        /// <summary>
        /// The name of the policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The path of the policy in IAM.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// The policy document.
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;

        /// <summary>
        /// Policy document as json. Useful if you need document but do not want to create IAM policy itself. For example for SSO Permission Set inline policies.
        /// </summary>
        [Output("policyJson")]
        public Output<string> PolicyJson { get; private set; } = null!;


        /// <summary>
        /// Create a ReadOnlyPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReadOnlyPolicy(string name, ReadOnlyPolicyArgs args, ComponentResourceOptions? options = null)
            : base("aws-iam:index:ReadOnlyPolicy", name, args ?? new ReadOnlyPolicyArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class ReadOnlyPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// JSON policy document if you want to add custom actions.
        /// </summary>
        [Input("additionalPolicyJson")]
        public Input<string>? AdditionalPolicyJson { get; set; }

        /// <summary>
        /// Allows StartQuery/StopQuery/FilterLogEvents CloudWatch actions.
        /// </summary>
        [Input("allowCloudwatchLogsQuery")]
        public Input<bool>? AllowCloudwatchLogsQuery { get; set; }

        /// <summary>
        /// Allows GetCallerIdentity/GetSessionToken/GetAccessKeyInfo sts actions.
        /// </summary>
        [Input("allowPredefinedStsActions")]
        public Input<bool>? AllowPredefinedStsActions { get; set; }

        /// <summary>
        /// Allows List/Get/Describe/View actions for services used when browsing AWS console (e.g. resource-groups, tag, health services).
        /// </summary>
        [Input("allowWebConsoleServices")]
        public Input<bool>? AllowWebConsoleServices { get; set; }

        [Input("allowedServices")]
        private InputList<string>? _allowedServices;

        /// <summary>
        /// List of services to allow Get/List/Describe/View options. Service name should be the same as corresponding service IAM prefix. See what it is for each service here https://docs.aws.amazon.com/service-authorization/latest/reference/reference_policies_actions-resources-contextkeys.html.
        /// </summary>
        public InputList<string> AllowedServices
        {
            get => _allowedServices ?? (_allowedServices = new InputList<string>());
            set => _allowedServices = value;
        }

        /// <summary>
        /// The description of the policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the policy.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The path of the policy in IAM.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

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

        [Input("webConsoleServices")]
        private InputList<string>? _webConsoleServices;

        /// <summary>
        /// List of web console services to allow.
        /// </summary>
        public InputList<string> WebConsoleServices
        {
            get => _webConsoleServices ?? (_webConsoleServices = new InputList<string>());
            set => _webConsoleServices = value;
        }

        public ReadOnlyPolicyArgs()
        {
            AdditionalPolicyJson = "{}";
            AllowCloudwatchLogsQuery = true;
            AllowPredefinedStsActions = true;
            AllowWebConsoleServices = true;
            Description = "IAM Policy";
            Path = "/";
        }
        public static new ReadOnlyPolicyArgs Empty => new ReadOnlyPolicyArgs();
    }
}
