// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { AccountArgs } from "./account";
export type Account = import("./account").Account;
export const Account: typeof import("./account").Account = null as any;
utilities.lazyLoad(exports, ["Account"], () => require("./account"));

export { AssumableRoleArgs } from "./assumableRole";
export type AssumableRole = import("./assumableRole").AssumableRole;
export const AssumableRole: typeof import("./assumableRole").AssumableRole = null as any;
utilities.lazyLoad(exports, ["AssumableRole"], () => require("./assumableRole"));

export { AssumableRoleWithOIDCArgs } from "./assumableRoleWithOIDC";
export type AssumableRoleWithOIDC = import("./assumableRoleWithOIDC").AssumableRoleWithOIDC;
export const AssumableRoleWithOIDC: typeof import("./assumableRoleWithOIDC").AssumableRoleWithOIDC = null as any;
utilities.lazyLoad(exports, ["AssumableRoleWithOIDC"], () => require("./assumableRoleWithOIDC"));

export { AssumableRoleWithSAMLArgs } from "./assumableRoleWithSAML";
export type AssumableRoleWithSAML = import("./assumableRoleWithSAML").AssumableRoleWithSAML;
export const AssumableRoleWithSAML: typeof import("./assumableRoleWithSAML").AssumableRoleWithSAML = null as any;
utilities.lazyLoad(exports, ["AssumableRoleWithSAML"], () => require("./assumableRoleWithSAML"));

export { AssumableRolesArgs } from "./assumableRoles";
export type AssumableRoles = import("./assumableRoles").AssumableRoles;
export const AssumableRoles: typeof import("./assumableRoles").AssumableRoles = null as any;
utilities.lazyLoad(exports, ["AssumableRoles"], () => require("./assumableRoles"));

export { AssumableRolesWithSAMLArgs } from "./assumableRolesWithSAML";
export type AssumableRolesWithSAML = import("./assumableRolesWithSAML").AssumableRolesWithSAML;
export const AssumableRolesWithSAML: typeof import("./assumableRolesWithSAML").AssumableRolesWithSAML = null as any;
utilities.lazyLoad(exports, ["AssumableRolesWithSAML"], () => require("./assumableRolesWithSAML"));

export { EKSRoleArgs } from "./eksrole";
export type EKSRole = import("./eksrole").EKSRole;
export const EKSRole: typeof import("./eksrole").EKSRole = null as any;
utilities.lazyLoad(exports, ["EKSRole"], () => require("./eksrole"));

export { GroupWithAssumableRolesPolicyArgs } from "./groupWithAssumableRolesPolicy";
export type GroupWithAssumableRolesPolicy = import("./groupWithAssumableRolesPolicy").GroupWithAssumableRolesPolicy;
export const GroupWithAssumableRolesPolicy: typeof import("./groupWithAssumableRolesPolicy").GroupWithAssumableRolesPolicy = null as any;
utilities.lazyLoad(exports, ["GroupWithAssumableRolesPolicy"], () => require("./groupWithAssumableRolesPolicy"));

export { GroupWithPoliciesArgs } from "./groupWithPolicies";
export type GroupWithPolicies = import("./groupWithPolicies").GroupWithPolicies;
export const GroupWithPolicies: typeof import("./groupWithPolicies").GroupWithPolicies = null as any;
utilities.lazyLoad(exports, ["GroupWithPolicies"], () => require("./groupWithPolicies"));

export { PolicyArgs } from "./policy";
export type Policy = import("./policy").Policy;
export const Policy: typeof import("./policy").Policy = null as any;
utilities.lazyLoad(exports, ["Policy"], () => require("./policy"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { ReadOnlyPolicyArgs } from "./readOnlyPolicy";
export type ReadOnlyPolicy = import("./readOnlyPolicy").ReadOnlyPolicy;
export const ReadOnlyPolicy: typeof import("./readOnlyPolicy").ReadOnlyPolicy = null as any;
utilities.lazyLoad(exports, ["ReadOnlyPolicy"], () => require("./readOnlyPolicy"));

export { RoleForServiceAccountsEksArgs } from "./roleForServiceAccountsEks";
export type RoleForServiceAccountsEks = import("./roleForServiceAccountsEks").RoleForServiceAccountsEks;
export const RoleForServiceAccountsEks: typeof import("./roleForServiceAccountsEks").RoleForServiceAccountsEks = null as any;
utilities.lazyLoad(exports, ["RoleForServiceAccountsEks"], () => require("./roleForServiceAccountsEks"));

export { UserArgs } from "./user";
export type User = import("./user").User;
export const User: typeof import("./user").User = null as any;
utilities.lazyLoad(exports, ["User"], () => require("./user"));


// Export sub-modules:
import * as types from "./types";

export {
    types,
};

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
pulumi.runtime.registerResourcePackage("aws-iam", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:aws-iam") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
