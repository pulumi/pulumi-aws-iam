// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.awsiam.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * The readonly role.
 * 
 */
public final class ReadonlyRoleWithMFAArgs extends com.pulumi.resources.ResourceArgs {

    public static final ReadonlyRoleWithMFAArgs Empty = new ReadonlyRoleWithMFAArgs();

    /**
     * IAM role with readonly access.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return IAM role with readonly access.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Path of readonly IAM role.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return Path of readonly IAM role.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * Permissions boundary ARN to use for readonly role.
     * 
     */
    @Import(name="permissionsBoundaryArn")
    private @Nullable Output<String> permissionsBoundaryArn;

    /**
     * @return Permissions boundary ARN to use for readonly role.
     * 
     */
    public Optional<Output<String>> permissionsBoundaryArn() {
        return Optional.ofNullable(this.permissionsBoundaryArn);
    }

    /**
     * List of policy ARNs to use for readonly role.
     * 
     */
    @Import(name="policyArns")
    private @Nullable Output<List<String>> policyArns;

    /**
     * @return List of policy ARNs to use for readonly role.
     * 
     */
    public Optional<Output<List<String>>> policyArns() {
        return Optional.ofNullable(this.policyArns);
    }

    /**
     * Whether admin role requires MFA.
     * 
     */
    @Import(name="requiresMfa")
    private @Nullable Output<Boolean> requiresMfa;

    /**
     * @return Whether admin role requires MFA.
     * 
     */
    public Optional<Output<Boolean>> requiresMfa() {
        return Optional.ofNullable(this.requiresMfa);
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

    private ReadonlyRoleWithMFAArgs() {}

    private ReadonlyRoleWithMFAArgs(ReadonlyRoleWithMFAArgs $) {
        this.name = $.name;
        this.path = $.path;
        this.permissionsBoundaryArn = $.permissionsBoundaryArn;
        this.policyArns = $.policyArns;
        this.requiresMfa = $.requiresMfa;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ReadonlyRoleWithMFAArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReadonlyRoleWithMFAArgs $;

        public Builder() {
            $ = new ReadonlyRoleWithMFAArgs();
        }

        public Builder(ReadonlyRoleWithMFAArgs defaults) {
            $ = new ReadonlyRoleWithMFAArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name IAM role with readonly access.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name IAM role with readonly access.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param path Path of readonly IAM role.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path Path of readonly IAM role.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param permissionsBoundaryArn Permissions boundary ARN to use for readonly role.
         * 
         * @return builder
         * 
         */
        public Builder permissionsBoundaryArn(@Nullable Output<String> permissionsBoundaryArn) {
            $.permissionsBoundaryArn = permissionsBoundaryArn;
            return this;
        }

        /**
         * @param permissionsBoundaryArn Permissions boundary ARN to use for readonly role.
         * 
         * @return builder
         * 
         */
        public Builder permissionsBoundaryArn(String permissionsBoundaryArn) {
            return permissionsBoundaryArn(Output.of(permissionsBoundaryArn));
        }

        /**
         * @param policyArns List of policy ARNs to use for readonly role.
         * 
         * @return builder
         * 
         */
        public Builder policyArns(@Nullable Output<List<String>> policyArns) {
            $.policyArns = policyArns;
            return this;
        }

        /**
         * @param policyArns List of policy ARNs to use for readonly role.
         * 
         * @return builder
         * 
         */
        public Builder policyArns(List<String> policyArns) {
            return policyArns(Output.of(policyArns));
        }

        /**
         * @param policyArns List of policy ARNs to use for readonly role.
         * 
         * @return builder
         * 
         */
        public Builder policyArns(String... policyArns) {
            return policyArns(List.of(policyArns));
        }

        /**
         * @param requiresMfa Whether admin role requires MFA.
         * 
         * @return builder
         * 
         */
        public Builder requiresMfa(@Nullable Output<Boolean> requiresMfa) {
            $.requiresMfa = requiresMfa;
            return this;
        }

        /**
         * @param requiresMfa Whether admin role requires MFA.
         * 
         * @return builder
         * 
         */
        public Builder requiresMfa(Boolean requiresMfa) {
            return requiresMfa(Output.of(requiresMfa));
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

        public ReadonlyRoleWithMFAArgs build() {
            $.name = Codegen.stringProp("name").output().arg($.name).def("readonly").getNullable();
            $.path = Codegen.stringProp("path").output().arg($.path).def("/").getNullable();
            $.permissionsBoundaryArn = Codegen.stringProp("permissionsBoundaryArn").output().arg($.permissionsBoundaryArn).def("").getNullable();
            $.requiresMfa = Codegen.booleanProp("requiresMfa").output().arg($.requiresMfa).def(true).getNullable();
            return $;
        }
    }

}
