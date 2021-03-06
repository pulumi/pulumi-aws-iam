// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.awsiam;

import com.pulumi.awsiam.inputs.RoleArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AssumableRoleWithOIDCArgs extends com.pulumi.resources.ResourceArgs {

    public static final AssumableRoleWithOIDCArgs Empty = new AssumableRoleWithOIDCArgs();

    /**
     * The AWS account ID where the OIDC provider lives, leave empty to use the account for the AWS provider.
     * 
     */
    @Import(name="awsAccountId")
    private @Nullable Output<String> awsAccountId;

    /**
     * @return The AWS account ID where the OIDC provider lives, leave empty to use the account for the AWS provider.
     * 
     */
    public Optional<Output<String>> awsAccountId() {
        return Optional.ofNullable(this.awsAccountId);
    }

    /**
     * Whether policies should be detached from this role when destroying.
     * 
     */
    @Import(name="forceDetachPolicies")
    private @Nullable Output<Boolean> forceDetachPolicies;

    /**
     * @return Whether policies should be detached from this role when destroying.
     * 
     */
    public Optional<Output<Boolean>> forceDetachPolicies() {
        return Optional.ofNullable(this.forceDetachPolicies);
    }

    /**
     * Maximum CLI/API session duration in seconds between 3600 and 43200.
     * 
     */
    @Import(name="maxSessionDuration")
    private @Nullable Output<Integer> maxSessionDuration;

    /**
     * @return Maximum CLI/API session duration in seconds between 3600 and 43200.
     * 
     */
    public Optional<Output<Integer>> maxSessionDuration() {
        return Optional.ofNullable(this.maxSessionDuration);
    }

    /**
     * The audience to be added to the role policy. Set to sts.amazonaws.com for cross-account assumable role. Leave empty otherwise.
     * 
     */
    @Import(name="oidcFullyQualifiedAudiences")
    private @Nullable Output<List<String>> oidcFullyQualifiedAudiences;

    /**
     * @return The audience to be added to the role policy. Set to sts.amazonaws.com for cross-account assumable role. Leave empty otherwise.
     * 
     */
    public Optional<Output<List<String>>> oidcFullyQualifiedAudiences() {
        return Optional.ofNullable(this.oidcFullyQualifiedAudiences);
    }

    /**
     * The fully qualified OIDC subjects to be added to the role policy.
     * 
     */
    @Import(name="oidcFullyQualifiedSubjects")
    private @Nullable Output<List<String>> oidcFullyQualifiedSubjects;

    /**
     * @return The fully qualified OIDC subjects to be added to the role policy.
     * 
     */
    public Optional<Output<List<String>>> oidcFullyQualifiedSubjects() {
        return Optional.ofNullable(this.oidcFullyQualifiedSubjects);
    }

    /**
     * The OIDC subject using wildcards to be added to the role policy.
     * 
     */
    @Import(name="oidcSubjectsWithWildcards")
    private @Nullable Output<List<String>> oidcSubjectsWithWildcards;

    /**
     * @return The OIDC subject using wildcards to be added to the role policy.
     * 
     */
    public Optional<Output<List<String>>> oidcSubjectsWithWildcards() {
        return Optional.ofNullable(this.oidcSubjectsWithWildcards);
    }

    /**
     * List of URLs of the OIDC Providers.
     * 
     */
    @Import(name="providerUrls")
    private @Nullable Output<List<String>> providerUrls;

    /**
     * @return List of URLs of the OIDC Providers.
     * 
     */
    public Optional<Output<List<String>>> providerUrls() {
        return Optional.ofNullable(this.providerUrls);
    }

    /**
     * The IAM role.
     * 
     */
    @Import(name="role")
    private @Nullable Output<RoleArgs> role;

    /**
     * @return The IAM role.
     * 
     */
    public Optional<Output<RoleArgs>> role() {
        return Optional.ofNullable(this.role);
    }

    /**
     * A map of tags to add.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to add.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private AssumableRoleWithOIDCArgs() {}

    private AssumableRoleWithOIDCArgs(AssumableRoleWithOIDCArgs $) {
        this.awsAccountId = $.awsAccountId;
        this.forceDetachPolicies = $.forceDetachPolicies;
        this.maxSessionDuration = $.maxSessionDuration;
        this.oidcFullyQualifiedAudiences = $.oidcFullyQualifiedAudiences;
        this.oidcFullyQualifiedSubjects = $.oidcFullyQualifiedSubjects;
        this.oidcSubjectsWithWildcards = $.oidcSubjectsWithWildcards;
        this.providerUrls = $.providerUrls;
        this.role = $.role;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AssumableRoleWithOIDCArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AssumableRoleWithOIDCArgs $;

        public Builder() {
            $ = new AssumableRoleWithOIDCArgs();
        }

        public Builder(AssumableRoleWithOIDCArgs defaults) {
            $ = new AssumableRoleWithOIDCArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param awsAccountId The AWS account ID where the OIDC provider lives, leave empty to use the account for the AWS provider.
         * 
         * @return builder
         * 
         */
        public Builder awsAccountId(@Nullable Output<String> awsAccountId) {
            $.awsAccountId = awsAccountId;
            return this;
        }

        /**
         * @param awsAccountId The AWS account ID where the OIDC provider lives, leave empty to use the account for the AWS provider.
         * 
         * @return builder
         * 
         */
        public Builder awsAccountId(String awsAccountId) {
            return awsAccountId(Output.of(awsAccountId));
        }

        /**
         * @param forceDetachPolicies Whether policies should be detached from this role when destroying.
         * 
         * @return builder
         * 
         */
        public Builder forceDetachPolicies(@Nullable Output<Boolean> forceDetachPolicies) {
            $.forceDetachPolicies = forceDetachPolicies;
            return this;
        }

        /**
         * @param forceDetachPolicies Whether policies should be detached from this role when destroying.
         * 
         * @return builder
         * 
         */
        public Builder forceDetachPolicies(Boolean forceDetachPolicies) {
            return forceDetachPolicies(Output.of(forceDetachPolicies));
        }

        /**
         * @param maxSessionDuration Maximum CLI/API session duration in seconds between 3600 and 43200.
         * 
         * @return builder
         * 
         */
        public Builder maxSessionDuration(@Nullable Output<Integer> maxSessionDuration) {
            $.maxSessionDuration = maxSessionDuration;
            return this;
        }

        /**
         * @param maxSessionDuration Maximum CLI/API session duration in seconds between 3600 and 43200.
         * 
         * @return builder
         * 
         */
        public Builder maxSessionDuration(Integer maxSessionDuration) {
            return maxSessionDuration(Output.of(maxSessionDuration));
        }

        /**
         * @param oidcFullyQualifiedAudiences The audience to be added to the role policy. Set to sts.amazonaws.com for cross-account assumable role. Leave empty otherwise.
         * 
         * @return builder
         * 
         */
        public Builder oidcFullyQualifiedAudiences(@Nullable Output<List<String>> oidcFullyQualifiedAudiences) {
            $.oidcFullyQualifiedAudiences = oidcFullyQualifiedAudiences;
            return this;
        }

        /**
         * @param oidcFullyQualifiedAudiences The audience to be added to the role policy. Set to sts.amazonaws.com for cross-account assumable role. Leave empty otherwise.
         * 
         * @return builder
         * 
         */
        public Builder oidcFullyQualifiedAudiences(List<String> oidcFullyQualifiedAudiences) {
            return oidcFullyQualifiedAudiences(Output.of(oidcFullyQualifiedAudiences));
        }

        /**
         * @param oidcFullyQualifiedAudiences The audience to be added to the role policy. Set to sts.amazonaws.com for cross-account assumable role. Leave empty otherwise.
         * 
         * @return builder
         * 
         */
        public Builder oidcFullyQualifiedAudiences(String... oidcFullyQualifiedAudiences) {
            return oidcFullyQualifiedAudiences(List.of(oidcFullyQualifiedAudiences));
        }

        /**
         * @param oidcFullyQualifiedSubjects The fully qualified OIDC subjects to be added to the role policy.
         * 
         * @return builder
         * 
         */
        public Builder oidcFullyQualifiedSubjects(@Nullable Output<List<String>> oidcFullyQualifiedSubjects) {
            $.oidcFullyQualifiedSubjects = oidcFullyQualifiedSubjects;
            return this;
        }

        /**
         * @param oidcFullyQualifiedSubjects The fully qualified OIDC subjects to be added to the role policy.
         * 
         * @return builder
         * 
         */
        public Builder oidcFullyQualifiedSubjects(List<String> oidcFullyQualifiedSubjects) {
            return oidcFullyQualifiedSubjects(Output.of(oidcFullyQualifiedSubjects));
        }

        /**
         * @param oidcFullyQualifiedSubjects The fully qualified OIDC subjects to be added to the role policy.
         * 
         * @return builder
         * 
         */
        public Builder oidcFullyQualifiedSubjects(String... oidcFullyQualifiedSubjects) {
            return oidcFullyQualifiedSubjects(List.of(oidcFullyQualifiedSubjects));
        }

        /**
         * @param oidcSubjectsWithWildcards The OIDC subject using wildcards to be added to the role policy.
         * 
         * @return builder
         * 
         */
        public Builder oidcSubjectsWithWildcards(@Nullable Output<List<String>> oidcSubjectsWithWildcards) {
            $.oidcSubjectsWithWildcards = oidcSubjectsWithWildcards;
            return this;
        }

        /**
         * @param oidcSubjectsWithWildcards The OIDC subject using wildcards to be added to the role policy.
         * 
         * @return builder
         * 
         */
        public Builder oidcSubjectsWithWildcards(List<String> oidcSubjectsWithWildcards) {
            return oidcSubjectsWithWildcards(Output.of(oidcSubjectsWithWildcards));
        }

        /**
         * @param oidcSubjectsWithWildcards The OIDC subject using wildcards to be added to the role policy.
         * 
         * @return builder
         * 
         */
        public Builder oidcSubjectsWithWildcards(String... oidcSubjectsWithWildcards) {
            return oidcSubjectsWithWildcards(List.of(oidcSubjectsWithWildcards));
        }

        /**
         * @param providerUrls List of URLs of the OIDC Providers.
         * 
         * @return builder
         * 
         */
        public Builder providerUrls(@Nullable Output<List<String>> providerUrls) {
            $.providerUrls = providerUrls;
            return this;
        }

        /**
         * @param providerUrls List of URLs of the OIDC Providers.
         * 
         * @return builder
         * 
         */
        public Builder providerUrls(List<String> providerUrls) {
            return providerUrls(Output.of(providerUrls));
        }

        /**
         * @param providerUrls List of URLs of the OIDC Providers.
         * 
         * @return builder
         * 
         */
        public Builder providerUrls(String... providerUrls) {
            return providerUrls(List.of(providerUrls));
        }

        /**
         * @param role The IAM role.
         * 
         * @return builder
         * 
         */
        public Builder role(@Nullable Output<RoleArgs> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The IAM role.
         * 
         * @return builder
         * 
         */
        public Builder role(RoleArgs role) {
            return role(Output.of(role));
        }

        /**
         * @param tags A map of tags to add.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to add.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public AssumableRoleWithOIDCArgs build() {
            $.awsAccountId = Codegen.stringProp("awsAccountId").output().arg($.awsAccountId).def("").getNullable();
            $.forceDetachPolicies = Codegen.booleanProp("forceDetachPolicies").output().arg($.forceDetachPolicies).def(false).getNullable();
            $.maxSessionDuration = Codegen.integerProp("maxSessionDuration").output().arg($.maxSessionDuration).def(3600).getNullable();
            return $;
        }
    }

}
