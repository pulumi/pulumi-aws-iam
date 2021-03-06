// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsIam.Outputs
{

    [OutputType]
    public sealed class KeybaseOutput
    {
        /// <summary>
        /// Decrypt user password command.
        /// </summary>
        public readonly string? PasswordDecryptCommand;
        /// <summary>
        /// Encrypted password
        /// </summary>
        public readonly string? PasswordPgpMessage;
        /// <summary>
        /// Decrypt access secret key command.
        /// </summary>
        public readonly string? SecretKeyDecryptCommand;
        /// <summary>
        /// Encrypted access secret key.
        /// </summary>
        public readonly string? SecretKeyPgpMessage;

        [OutputConstructor]
        private KeybaseOutput(
            string? passwordDecryptCommand,

            string? passwordPgpMessage,

            string? secretKeyDecryptCommand,

            string? secretKeyPgpMessage)
        {
            PasswordDecryptCommand = passwordDecryptCommand;
            PasswordPgpMessage = passwordPgpMessage;
            SecretKeyDecryptCommand = secretKeyDecryptCommand;
            SecretKeyPgpMessage = secretKeyPgpMessage;
        }
    }
}
