// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides details about a specific Cluster Option resource in Oracle Cloud Infrastructure Container Engine service.
 *
 * Get options available for clusters.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testClusterOption = oci.GetContainerengineClusterOption({
 *     clusterOptionId: oci_containerengine_cluster_option.test_cluster_option.id,
 *     compartmentId: _var.compartment_id,
 * });
 * ```
 */
export function getContainerengineClusterOption(args: GetContainerengineClusterOptionArgs, opts?: pulumi.InvokeOptions): Promise<GetContainerengineClusterOptionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getContainerengineClusterOption:GetContainerengineClusterOption", {
        "clusterOptionId": args.clusterOptionId,
        "compartmentId": args.compartmentId,
    }, opts);
}

/**
 * A collection of arguments for invoking GetContainerengineClusterOption.
 */
export interface GetContainerengineClusterOptionArgs {
    /**
     * The id of the option set to retrieve. Use "all" get all options, or use a cluster ID to get options specific to the provided cluster.
     */
    clusterOptionId: string;
    /**
     * The OCID of the compartment.
     */
    compartmentId?: string;
}

/**
 * A collection of values returned by GetContainerengineClusterOption.
 */
export interface GetContainerengineClusterOptionResult {
    readonly clusterOptionId: string;
    readonly compartmentId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Available Kubernetes versions.
     */
    readonly kubernetesVersions: string[];
}