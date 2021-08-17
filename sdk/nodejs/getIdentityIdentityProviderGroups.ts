// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the list of Identity Provider Groups in Oracle Cloud Infrastructure Identity service.
 *
 * Lists the identity provider groups.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testIdentityProviderGroups = oci.GetIdentityIdentityProviderGroups({
 *     identityProviderId: oci_identity_identity_provider.test_identity_provider.id,
 *     name: _var.identity_provider_group_name,
 *     state: _var.identity_provider_group_state,
 * });
 * ```
 */
export function getIdentityIdentityProviderGroups(args: GetIdentityIdentityProviderGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetIdentityIdentityProviderGroupsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getIdentityIdentityProviderGroups:GetIdentityIdentityProviderGroups", {
        "filters": args.filters,
        "identityProviderId": args.identityProviderId,
        "name": args.name,
        "state": args.state,
    }, opts);
}

/**
 * A collection of arguments for invoking GetIdentityIdentityProviderGroups.
 */
export interface GetIdentityIdentityProviderGroupsArgs {
    filters?: inputs.GetIdentityIdentityProviderGroupsFilter[];
    /**
     * The OCID of the identity provider.
     */
    identityProviderId: string;
    /**
     * A filter to only return resources that match the given name exactly.
     */
    name?: string;
    /**
     * A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
     */
    state?: string;
}

/**
 * A collection of values returned by GetIdentityIdentityProviderGroups.
 */
export interface GetIdentityIdentityProviderGroupsResult {
    readonly filters?: outputs.GetIdentityIdentityProviderGroupsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of identity_provider_groups.
     */
    readonly identityProviderGroups: outputs.GetIdentityIdentityProviderGroupsIdentityProviderGroup[];
    /**
     * The OCID of the `IdentityProvider` this group belongs to.
     */
    readonly identityProviderId: string;
    /**
     * Display name of the group
     */
    readonly name?: string;
    readonly state?: string;
}