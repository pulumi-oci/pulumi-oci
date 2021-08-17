// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the list of Http Probe Results in Oracle Cloud Infrastructure Health Checks service.
 *
 * Gets the HTTP probe results for the specified probe or monitor, where
 * the `probeConfigurationId` is the OCID of either a monitor or an
 * on-demand probe.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testHttpProbeResults = oci.GetHealthChecksHttpProbeResults({
 *     probeConfigurationId: oci_health_checks_probe_configuration.test_probe_configuration.id,
 *     startTimeGreaterThanOrEqualTo: _var.http_probe_result_start_time_greater_than_or_equal_to,
 *     startTimeLessThanOrEqualTo: _var.http_probe_result_start_time_less_than_or_equal_to,
 *     target: _var.http_probe_result_target,
 * });
 * ```
 */
export function getHealthChecksHttpProbeResults(args: GetHealthChecksHttpProbeResultsArgs, opts?: pulumi.InvokeOptions): Promise<GetHealthChecksHttpProbeResultsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getHealthChecksHttpProbeResults:GetHealthChecksHttpProbeResults", {
        "filters": args.filters,
        "probeConfigurationId": args.probeConfigurationId,
        "startTimeGreaterThanOrEqualTo": args.startTimeGreaterThanOrEqualTo,
        "startTimeLessThanOrEqualTo": args.startTimeLessThanOrEqualTo,
        "target": args.target,
    }, opts);
}

/**
 * A collection of arguments for invoking GetHealthChecksHttpProbeResults.
 */
export interface GetHealthChecksHttpProbeResultsArgs {
    filters?: inputs.GetHealthChecksHttpProbeResultsFilter[];
    /**
     * The OCID of a monitor or on-demand probe.
     */
    probeConfigurationId: string;
    /**
     * Returns results with a `startTime` equal to or greater than the specified value.
     */
    startTimeGreaterThanOrEqualTo?: number;
    /**
     * Returns results with a `startTime` equal to or less than the specified value.
     */
    startTimeLessThanOrEqualTo?: number;
    /**
     * Filters results that match the `target`.
     */
    target?: string;
}

/**
 * A collection of values returned by GetHealthChecksHttpProbeResults.
 */
export interface GetHealthChecksHttpProbeResultsResult {
    readonly filters?: outputs.GetHealthChecksHttpProbeResultsFilter[];
    /**
     * The list of http_probe_results.
     */
    readonly httpProbeResults: outputs.GetHealthChecksHttpProbeResultsHttpProbeResult[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The OCID of the monitor or on-demand probe responsible for creating this result.
     */
    readonly probeConfigurationId: string;
    readonly startTimeGreaterThanOrEqualTo?: number;
    readonly startTimeLessThanOrEqualTo?: number;
    /**
     * The target hostname or IP address of the probe.
     */
    readonly target?: string;
}