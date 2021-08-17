// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the list of Cpes in Oracle Cloud Infrastructure Core service.
 *
 * Lists the customer-premises equipment objects (CPEs) in the specified compartment.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testCpes = oci.GetCoreCpes({
 *     compartmentId: _var.compartment_id,
 * });
 * ```
 */
export function getCoreCpes(args: GetCoreCpesArgs, opts?: pulumi.InvokeOptions): Promise<GetCoreCpesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getCoreCpes:GetCoreCpes", {
        "compartmentId": args.compartmentId,
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking GetCoreCpes.
 */
export interface GetCoreCpesArgs {
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
     */
    compartmentId: string;
    filters?: inputs.GetCoreCpesFilter[];
}

/**
 * A collection of values returned by GetCoreCpes.
 */
export interface GetCoreCpesResult {
    /**
     * The OCID of the compartment containing the CPE.
     */
    readonly compartmentId: string;
    /**
     * The list of cpes.
     */
    readonly cpes: outputs.GetCoreCpesCpe[];
    readonly filters?: outputs.GetCoreCpesFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}