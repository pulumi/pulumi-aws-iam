// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.awsiam.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * The VPC CNI IAM policy to the role.
 * 
 */
public final class EKSVPNCNIPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final EKSVPNCNIPolicyArgs Empty = new EKSVPNCNIPolicyArgs();

    /**
     * Determines whether to attach the VPC CNI IAM policy to the role.
     * 
     */
    @Import(name="attach", required=true)
    private Output<Boolean> attach;

    /**
     * @return Determines whether to attach the VPC CNI IAM policy to the role.
     * 
     */
    public Output<Boolean> attach() {
        return this.attach;
    }

    /**
     * Determines whether to enable IPv4 permissions for VPC CNI policy.
     * 
     */
    @Import(name="enableIpv4")
    private @Nullable Output<Boolean> enableIpv4;

    /**
     * @return Determines whether to enable IPv4 permissions for VPC CNI policy.
     * 
     */
    public Optional<Output<Boolean>> enableIpv4() {
        return Optional.ofNullable(this.enableIpv4);
    }

    /**
     * Determines whether to enable IPv6 permissions for VPC CNI policy.
     * 
     */
    @Import(name="enableIpv6")
    private @Nullable Output<Boolean> enableIpv6;

    /**
     * @return Determines whether to enable IPv6 permissions for VPC CNI policy.
     * 
     */
    public Optional<Output<Boolean>> enableIpv6() {
        return Optional.ofNullable(this.enableIpv6);
    }

    private EKSVPNCNIPolicyArgs() {}

    private EKSVPNCNIPolicyArgs(EKSVPNCNIPolicyArgs $) {
        this.attach = $.attach;
        this.enableIpv4 = $.enableIpv4;
        this.enableIpv6 = $.enableIpv6;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EKSVPNCNIPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EKSVPNCNIPolicyArgs $;

        public Builder() {
            $ = new EKSVPNCNIPolicyArgs();
        }

        public Builder(EKSVPNCNIPolicyArgs defaults) {
            $ = new EKSVPNCNIPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param attach Determines whether to attach the VPC CNI IAM policy to the role.
         * 
         * @return builder
         * 
         */
        public Builder attach(Output<Boolean> attach) {
            $.attach = attach;
            return this;
        }

        /**
         * @param attach Determines whether to attach the VPC CNI IAM policy to the role.
         * 
         * @return builder
         * 
         */
        public Builder attach(Boolean attach) {
            return attach(Output.of(attach));
        }

        /**
         * @param enableIpv4 Determines whether to enable IPv4 permissions for VPC CNI policy.
         * 
         * @return builder
         * 
         */
        public Builder enableIpv4(@Nullable Output<Boolean> enableIpv4) {
            $.enableIpv4 = enableIpv4;
            return this;
        }

        /**
         * @param enableIpv4 Determines whether to enable IPv4 permissions for VPC CNI policy.
         * 
         * @return builder
         * 
         */
        public Builder enableIpv4(Boolean enableIpv4) {
            return enableIpv4(Output.of(enableIpv4));
        }

        /**
         * @param enableIpv6 Determines whether to enable IPv6 permissions for VPC CNI policy.
         * 
         * @return builder
         * 
         */
        public Builder enableIpv6(@Nullable Output<Boolean> enableIpv6) {
            $.enableIpv6 = enableIpv6;
            return this;
        }

        /**
         * @param enableIpv6 Determines whether to enable IPv6 permissions for VPC CNI policy.
         * 
         * @return builder
         * 
         */
        public Builder enableIpv6(Boolean enableIpv6) {
            return enableIpv6(Output.of(enableIpv6));
        }

        public EKSVPNCNIPolicyArgs build() {
            $.attach = Objects.requireNonNull($.attach, "expected parameter 'attach' to be non-null");
            return $;
        }
    }

}
