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
    /// The Karpenter Controller policy to the role.
    /// </summary>
    public sealed class EKSKarpenterControllerPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines whether to attach the Karpenter Controller policy to the role.
        /// </summary>
        [Input("attach", required: true)]
        public Input<bool> Attach { get; set; } = null!;

        /// <summary>
        /// Cluster ID where the Karpenter controller is provisioned/managing.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("nodeIamRoleArns")]
        private InputList<string>? _nodeIamRoleArns;

        /// <summary>
        /// List of node IAM role ARNs Karpenter can use to launch nodes. If not provided,
        /// the default ARN "*" will be applied.
        /// </summary>
        public InputList<string> NodeIamRoleArns
        {
            get => _nodeIamRoleArns ?? (_nodeIamRoleArns = new InputList<string>());
            set => _nodeIamRoleArns = value;
        }

        [Input("ssmParameterArns")]
        private InputList<string>? _ssmParameterArns;

        /// <summary>
        /// List of SSM Parameter ARNs that contain AMI IDs launched by Karpenter. If not provided,
        /// the default ARN "arn:aws:ssm:*:*:parameter/aws/service/*" will be applied.
        /// </summary>
        public InputList<string> SsmParameterArns
        {
            get => _ssmParameterArns ?? (_ssmParameterArns = new InputList<string>());
            set => _ssmParameterArns = value;
        }

        /// <summary>
        /// Account ID of where the subnets Karpenter will utilize resides. Used when subnets are shared from another account.
        /// </summary>
        [Input("subnetAccountId")]
        public Input<string>? SubnetAccountId { get; set; }

        /// <summary>
        /// Tag key (`{key = value}`) applied to resources launched by Karpenter through the Karpenter provisioner.
        /// </summary>
        [Input("tagKey")]
        public Input<string>? TagKey { get; set; }

        public EKSKarpenterControllerPolicyArgs()
        {
            ClusterId = "*";
            TagKey = "karpenter.sh/discovery";
        }
        public static new EKSKarpenterControllerPolicyArgs Empty => new EKSKarpenterControllerPolicyArgs();
    }
}
