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
    /// This resources helps you create an IAM User, Login Profile, and Access Key. Additionally you
    /// can optionally upload an IAM SSH User Public Key.
    /// 
    /// ## Example Usage
    /// ### User
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
    ///         var user = new User("user", new UserArgs
    ///         {
    ///             Name = "pulumipus",
    ///             ForceDestroy = true,
    ///             PgpKey = "keybase:test",
    ///             PasswordResetRequired = false,
    ///         });
    /// 
    ///         this.User = Output.Create&lt;User&gt;(user);
    ///     }
    /// 
    ///     [Output]
    ///     public Output&lt;User&gt; User { get; set; }
    /// }
    /// ```
    /// {{ /example }}
    /// </summary>
    [AwsIamResourceType("aws-iam:index:User")]
    public partial class User : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// The IAM access key.
        /// </summary>
        [Output("accessKey")]
        public Output<Outputs.AccessKeyOutput> AccessKey { get; private set; } = null!;

        [Output("keybase")]
        public Output<Outputs.KeybaseOutput> Keybase { get; private set; } = null!;

        /// <summary>
        /// PGP key used to encrypt sensitive data for this user (if empty - secrets are not encrypted).
        /// </summary>
        [Output("pgpKey")]
        public Output<string> PgpKey { get; private set; } = null!;

        /// <summary>
        /// The IAM user.
        /// </summary>
        [Output("userInfo")]
        public Output<Outputs.UserOutput> UserInfo { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, ComponentResourceOptions? options = null)
            : base("aws-iam:index:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class UserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When destroying this user, destroy even if it has non-Pulumi-managed IAM access keys, login profile or MFA devices. Without forceDestroy a user with non-Pulumi-managed access keys and login profile will fail to be destroyed.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// Desired name for the IAM user.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The length of the generated password
        /// </summary>
        [Input("passwordLength")]
        public Input<int>? PasswordLength { get; set; }

        /// <summary>
        /// Whether the user should be forced to reset the generated password on first login.
        /// </summary>
        [Input("passwordResetRequired")]
        public Input<bool>? PasswordResetRequired { get; set; }

        /// <summary>
        /// Desired path for the IAM user.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The ARN of the policy that is used to set the permissions boundary for the user.
        /// </summary>
        [Input("permissionsBoundary")]
        public Input<string>? PermissionsBoundary { get; set; }

        /// <summary>
        /// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:username`. Used to encrypt password and access key.
        /// </summary>
        [Input("pgpKey")]
        public Input<string>? PgpKey { get; set; }

        /// <summary>
        /// Specifies the public key encoding format to use in the response. To retrieve the public key in ssh-rsa format, use SSH. To retrieve the public key in PEM format, use PEM.
        /// </summary>
        [Input("sshKeyEncoding")]
        public Input<string>? SshKeyEncoding { get; set; }

        /// <summary>
        /// The SSH public key. The public key must be encoded in ssh-rsa format or PEM format.
        /// </summary>
        [Input("sshPublicKey")]
        public Input<string>? SshPublicKey { get; set; }

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

        /// <summary>
        /// Whether to upload a public ssh key to the IAM user.
        /// </summary>
        [Input("uploadIamUserSshKey")]
        public Input<bool>? UploadIamUserSshKey { get; set; }

        public UserArgs()
        {
            Path = "/";
            SshKeyEncoding = "SSH";
        }
        public static new UserArgs Empty => new UserArgs();
    }
}
