// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the list of Profiles in Oracle Cloud Infrastructure Optimizer service.
 *
 * Lists the existing profiles.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testProfiles = oci.GetOptimizerProfiles({
 *     compartmentId: _var.compartment_id,
 *     name: _var.profile_name,
 *     state: _var.profile_state,
 * });
 * ```
 */
export function getOptimizerProfiles(args: GetOptimizerProfilesArgs, opts?: pulumi.InvokeOptions): Promise<GetOptimizerProfilesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getOptimizerProfiles:GetOptimizerProfiles", {
        "compartmentId": args.compartmentId,
        "filters": args.filters,
        "name": args.name,
        "state": args.state,
    }, opts);
}

/**
 * A collection of arguments for invoking GetOptimizerProfiles.
 */
export interface GetOptimizerProfilesArgs {
    /**
     * The OCID of the compartment.
     */
    compartmentId: string;
    filters?: inputs.GetOptimizerProfilesFilter[];
    /**
     * Optional. A filter that returns results that match the name specified.
     */
    name?: string;
    /**
     * A filter that returns results that match the lifecycle state specified.
     */
    state?: string;
}

/**
 * A collection of values returned by GetOptimizerProfiles.
 */
export interface GetOptimizerProfilesResult {
    /**
     * The OCID of the tenancy. The tenancy is the root compartment.
     */
    readonly compartmentId: string;
    readonly filters?: outputs.GetOptimizerProfilesFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name assigned to the profile. Avoid entering confidential information.
     */
    readonly name?: string;
    /**
     * The list of profile_collection.
     */
    readonly profileCollections: outputs.GetOptimizerProfilesProfileCollection[];
    /**
     * The profile's current state.
     */
    readonly state?: string;
}