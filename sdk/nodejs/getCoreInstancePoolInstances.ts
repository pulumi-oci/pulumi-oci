// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the list of Instance Pool Instances in Oracle Cloud Infrastructure Core service.
 *
 * List the instances in the specified instance pool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testInstancePoolInstances = oci.GetCoreInstancePoolInstances({
 *     compartmentId: _var.compartment_id,
 *     instancePoolId: oci_core_instance_pool.test_instance_pool.id,
 *     displayName: _var.instance_pool_instance_display_name,
 * });
 * ```
 */
export function getCoreInstancePoolInstances(args: GetCoreInstancePoolInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetCoreInstancePoolInstancesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getCoreInstancePoolInstances:GetCoreInstancePoolInstances", {
        "compartmentId": args.compartmentId,
        "displayName": args.displayName,
        "filters": args.filters,
        "instancePoolId": args.instancePoolId,
    }, opts);
}

/**
 * A collection of arguments for invoking GetCoreInstancePoolInstances.
 */
export interface GetCoreInstancePoolInstancesArgs {
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
     */
    compartmentId: string;
    /**
     * A filter to return only resources that match the given display name exactly.
     */
    displayName?: string;
    filters?: inputs.GetCoreInstancePoolInstancesFilter[];
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance pool.
     */
    instancePoolId: string;
}

/**
 * A collection of values returned by GetCoreInstancePoolInstances.
 */
export interface GetCoreInstancePoolInstancesResult {
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the instance.
     */
    readonly compartmentId: string;
    /**
     * The user-friendly name. Does not have to be unique.
     */
    readonly displayName?: string;
    readonly filters?: outputs.GetCoreInstancePoolInstancesFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instancePoolId: string;
    /**
     * The list of instances.
     */
    readonly instances: outputs.GetCoreInstancePoolInstancesInstance[];
}