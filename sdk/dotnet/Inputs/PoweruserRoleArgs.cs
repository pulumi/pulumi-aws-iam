// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsIam.Inputs
{

    /// <summary>
    /// The poweruser role.
    /// </summary>
    public sealed class PoweruserRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IAM role with poweruser access.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Path of poweruser IAM role.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Permissions boundary ARN to use for poweruser role.
        /// </summary>
        [Input("permissionsBoundaryArn")]
        public Input<string>? PermissionsBoundaryArn { get; set; }

        [Input("policyArns")]
        private InputList<string>? _policyArns;

        /// <summary>
        /// List of policy ARNs to use for poweruser role.
        /// </summary>
        public InputList<string> PolicyArns
        {
            get => _policyArns ?? (_policyArns = new InputList<string>());
            set => _policyArns = value;
        }

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

        public PoweruserRoleArgs()
        {
            Name = "poweruser";
            Path = "/";
            PermissionsBoundaryArn = "";
        }
        public static new PoweruserRoleArgs Empty => new PoweruserRoleArgs();
    }
}
