// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the list of Remote Peering Connections in Oracle Cloud Infrastructure Core service.
 *
 * Lists the remote peering connections (RPCs) for the specified DRG and compartment
 * (the RPC's compartment).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testRemotePeeringConnections = oci.GetCoreRemotePeeringConnections({
 *     compartmentId: _var.compartment_id,
 *     drgId: oci_core_drg.test_drg.id,
 * });
 * ```
 */
export function getCoreRemotePeeringConnections(args: GetCoreRemotePeeringConnectionsArgs, opts?: pulumi.InvokeOptions): Promise<GetCoreRemotePeeringConnectionsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getCoreRemotePeeringConnections:GetCoreRemotePeeringConnections", {
        "compartmentId": args.compartmentId,
        "drgId": args.drgId,
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking GetCoreRemotePeeringConnections.
 */
export interface GetCoreRemotePeeringConnectionsArgs {
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
     */
    compartmentId: string;
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
     */
    drgId?: string;
    filters?: inputs.GetCoreRemotePeeringConnectionsFilter[];
}

/**
 * A collection of values returned by GetCoreRemotePeeringConnections.
 */
export interface GetCoreRemotePeeringConnectionsResult {
    /**
     * The OCID of the compartment that contains the RPC.
     */
    readonly compartmentId: string;
    /**
     * The OCID of the DRG that this RPC belongs to.
     */
    readonly drgId?: string;
    readonly filters?: outputs.GetCoreRemotePeeringConnectionsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of remote_peering_connections.
     */
    readonly remotePeeringConnections: outputs.GetCoreRemotePeeringConnectionsRemotePeeringConnection[];
}