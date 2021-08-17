// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the list of Fault Domains in Oracle Cloud Infrastructure Identity service.
 *
 * Lists the Fault Domains in your tenancy. Specify the OCID of either the tenancy or another
 * of your compartments as the value for the compartment ID (remember that the tenancy is simply the root compartment).
 * See [Where to Get the Tenancy's OCID and User's OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/apisigningkey.htm#five).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testFaultDomains = oci.GetIdentityFaultDomains({
 *     availabilityDomain: _var.fault_domain_availability_domain,
 *     compartmentId: _var.compartment_id,
 * });
 * ```
 */
export function getIdentityFaultDomains(args: GetIdentityFaultDomainsArgs, opts?: pulumi.InvokeOptions): Promise<GetIdentityFaultDomainsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getIdentityFaultDomains:GetIdentityFaultDomains", {
        "availabilityDomain": args.availabilityDomain,
        "compartmentId": args.compartmentId,
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking GetIdentityFaultDomains.
 */
export interface GetIdentityFaultDomainsArgs {
    /**
     * The name of the availibilityDomain.
     */
    availabilityDomain: string;
    /**
     * The OCID of the compartment (remember that the tenancy is simply the root compartment).
     */
    compartmentId: string;
    filters?: inputs.GetIdentityFaultDomainsFilter[];
}

/**
 * A collection of values returned by GetIdentityFaultDomains.
 */
export interface GetIdentityFaultDomainsResult {
    /**
     * The name of the availabilityDomain where the Fault Domain belongs.
     */
    readonly availabilityDomain: string;
    /**
     * The OCID of the compartment. Currently only tenancy (root) compartment can be provided.
     */
    readonly compartmentId: string;
    /**
     * The list of fault_domains.
     */
    readonly faultDomains: outputs.GetIdentityFaultDomainsFaultDomain[];
    readonly filters?: outputs.GetIdentityFaultDomainsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}