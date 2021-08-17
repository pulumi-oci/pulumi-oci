// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides details about a specific Configuration resource in Oracle Cloud Infrastructure Metering Computation service.
 *
 * Returns the configurations list for the UI drop-down list.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testConfiguration = oci.GetMeteringComputationConfiguration({
 *     tenantId: oci_metering_computation_tenant.test_tenant.id,
 * });
 * ```
 */
export function getMeteringComputationConfiguration(args: GetMeteringComputationConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetMeteringComputationConfigurationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getMeteringComputationConfiguration:GetMeteringComputationConfiguration", {
        "tenantId": args.tenantId,
    }, opts);
}

/**
 * A collection of arguments for invoking GetMeteringComputationConfiguration.
 */
export interface GetMeteringComputationConfigurationArgs {
    /**
     * tenant id
     */
    tenantId: string;
}

/**
 * A collection of values returned by GetMeteringComputationConfiguration.
 */
export interface GetMeteringComputationConfigurationResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of available configurations.
     */
    readonly items: outputs.GetMeteringComputationConfigurationItem[];
    readonly tenantId: string;
}