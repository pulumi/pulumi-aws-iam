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
    /// This resource helps you manage an Iam Account's Alias and Password Policy. If your IAM Account Alias was previously
    /// set (either via the AWS console or when AWS created your Account) you will see an error like the below:
    /// 
    /// If you want to manage you Alias using Pulumi you will need to import this resource.
    /// 
    /// ## Example Usage
    /// ## Account
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
    ///         var account = new Account("account", new AccountArgs
    ///         {
    ///             AccountAlias = "cool-alias",
    ///             PasswordPolicy=new AccountPasswordPolicyArgs
    ///             {
    ///                 MinimumLength = 37,
    ///                 RequireNumbers = false,
    ///                 AllowUsersToChange = true,
    ///                 HardExpiry = true,
    ///                 RequireSymbols = true,
    ///                 RequireLowercaseCharacters = true,
    ///                 RequireUppercaseCharacters = true,
    ///             }
    /// 
    ///         });
    /// 
    ///         this.Account = Output.Create&lt;Account&gt;(account);
    ///     }
    /// 
    ///     [Output]
    ///     public Output&lt;Account&gt; Account { get; set; }
    /// }
    /// ```
    /// {{ /example }}
    /// </summary>
    [AwsIamResourceType("aws-iam:index:Account")]
    public partial class Account : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// The AWS ARN associated with the calling entity.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The AWS Account ID number of the account that owns or contains the calling entity.
        /// </summary>
        [Output("id")]
        public Output<string> Id { get; private set; } = null!;

        /// <summary>
        /// Indicates whether passwords in the account expire. Returns true if max password
        /// age contains a value greater than 0. Returns false if it is 0 or not present.
        /// </summary>
        [Output("passwordPolicyExpirePasswords")]
        public Output<bool> PasswordPolicyExpirePasswords { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the calling entity.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a Account resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Account(string name, AccountArgs args, ComponentResourceOptions? options = null)
            : base("aws-iam:index:Account", name, args ?? new AccountArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class AccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AWS IAM account alias for this account.
        /// </summary>
        [Input("accountAlias", required: true)]
        public Input<string> AccountAlias { get; set; } = null!;

        /// <summary>
        /// Options to specify complexity requirements and mandatory rotation periods for your IAM users' passwords. If
        /// left empty the default AWS password policy will be applied.
        /// </summary>
        [Input("passwordPolicy", required: true)]
        public Input<Inputs.AccountPasswordPolicyArgs> PasswordPolicy { get; set; } = null!;

        public AccountArgs()
        {
        }
        public static new AccountArgs Empty => new AccountArgs();
    }
}
