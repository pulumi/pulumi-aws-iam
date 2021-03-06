// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./account";
export * from "./assumableRole";
export * from "./assumableRoleWithOIDC";
export * from "./assumableRoleWithSAML";
export * from "./assumableRoles";
export * from "./assumableRolesWithSAML";
export * from "./eksrole";
export * from "./groupWithAssumableRolesPolicy";
export * from "./groupWithPolicies";
export * from "./policy";
export * from "./provider";
export * from "./readOnlyPolicy";
export * from "./roleForServiceAccountsEks";
export * from "./user";

// Export sub-modules:
import * as types from "./types";

export {
    types,
};

// Import resources to register:
import { Account } from "./account";
import { AssumableRole } from "./assumableRole";
import { AssumableRoleWithOIDC } from "./assumableRoleWithOIDC";
import { AssumableRoleWithSAML } from "./assumableRoleWithSAML";
import { AssumableRoles } from "./assumableRoles";
import { AssumableRolesWithSAML } from "./assumableRolesWithSAML";
import { EKSRole } from "./eksrole";
import { GroupWithAssumableRolesPolicy } from "./groupWithAssumableRolesPolicy";
import { GroupWithPolicies } from "./groupWithPolicies";
import { Policy } from "./policy";
import { ReadOnlyPolicy } from "./readOnlyPolicy";
import { RoleForServiceAccountsEks } from "./roleForServiceAccountsEks";
import { User } from "./user";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-iam:index:Account":
                return new Account(name, <any>undefined, { urn })
            case "aws-iam:index:AssumableRole":
                return new AssumableRole(name, <any>undefined, { urn })
            case "aws-iam:index:AssumableRoleWithOIDC":
                return new AssumableRoleWithOIDC(name, <any>undefined, { urn })
            case "aws-iam:index:AssumableRoleWithSAML":
                return new AssumableRoleWithSAML(name, <any>undefined, { urn })
            case "aws-iam:index:AssumableRoles":
                return new AssumableRoles(name, <any>undefined, { urn })
            case "aws-iam:index:AssumableRolesWithSAML":
                return new AssumableRolesWithSAML(name, <any>undefined, { urn })
            case "aws-iam:index:EKSRole":
                return new EKSRole(name, <any>undefined, { urn })
            case "aws-iam:index:GroupWithAssumableRolesPolicy":
                return new GroupWithAssumableRolesPolicy(name, <any>undefined, { urn })
            case "aws-iam:index:GroupWithPolicies":
                return new GroupWithPolicies(name, <any>undefined, { urn })
            case "aws-iam:index:Policy":
                return new Policy(name, <any>undefined, { urn })
            case "aws-iam:index:ReadOnlyPolicy":
                return new ReadOnlyPolicy(name, <any>undefined, { urn })
            case "aws-iam:index:RoleForServiceAccountsEks":
                return new RoleForServiceAccountsEks(name, <any>undefined, { urn })
            case "aws-iam:index:User":
                return new User(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-iam", "index", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("aws-iam", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:aws-iam") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
