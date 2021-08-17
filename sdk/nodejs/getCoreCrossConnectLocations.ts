// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the list of Cross Connect Locations in Oracle Cloud Infrastructure Core service.
 *
 * Lists the available FastConnect locations for cross-connect installation. You need
 * this information so you can specify your desired location when you create a cross-connect.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testCrossConnectLocations = oci.GetCoreCrossConnectLocations({
 *     compartmentId: _var.compartment_id,
 * });
 * ```
 */
export function getCoreCrossConnectLocations(args: GetCoreCrossConnectLocationsArgs, opts?: pulumi.InvokeOptions): Promise<GetCoreCrossConnectLocationsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getCoreCrossConnectLocations:GetCoreCrossConnectLocations", {
        "compartmentId": args.compartmentId,
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking GetCoreCrossConnectLocations.
 */
export interface GetCoreCrossConnectLocationsArgs {
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
     */
    compartmentId: string;
    filters?: inputs.GetCoreCrossConnectLocationsFilter[];
}

/**
 * A collection of values returned by GetCoreCrossConnectLocations.
 */
export interface GetCoreCrossConnectLocationsResult {
    readonly compartmentId: string;
    /**
     * The list of cross_connect_locations.
     */
    readonly crossConnectLocations: outputs.GetCoreCrossConnectLocationsCrossConnectLocation[];
    readonly filters?: outputs.GetCoreCrossConnectLocationsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}