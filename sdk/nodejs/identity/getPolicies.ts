// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the list of Policies in Oracle Cloud Infrastructure Identity service.
 *
 * Lists the policies in the specified compartment (either the tenancy or another of your compartments).
 * See [Where to Get the Tenancy's OCID and User's OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/apisigningkey.htm#five).
 *
 * To determine which policies apply to a particular group or compartment, you must view the individual
 * statements inside all your policies. There isn't a way to automatically obtain that information via the API.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testPolicies = oci.identity.getPolicies({
 *     compartmentId: _var.tenancy_ocid,
 *     name: _var.policy_name,
 *     state: _var.policy_state,
 * });
 * ```
 */
export function getPolicies(args: GetPoliciesArgs, opts?: pulumi.InvokeOptions): Promise<GetPoliciesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:identity/getPolicies:getPolicies", {
        "compartmentId": args.compartmentId,
        "filters": args.filters,
        "name": args.name,
        "state": args.state,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicies.
 */
export interface GetPoliciesArgs {
    /**
     * The OCID of the compartment (remember that the tenancy is simply the root compartment).
     */
    compartmentId: string;
    filters?: inputs.identity.GetPoliciesFilter[];
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
 * A collection of values returned by getPolicies.
 */
export interface GetPoliciesResult {
    /**
     * The OCID of the compartment containing the policy (either the tenancy or another compartment).
     */
    readonly compartmentId: string;
    readonly filters?: outputs.identity.GetPoliciesFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name you assign to the policy during creation. The name must be unique across all policies in the tenancy and cannot be changed.
     */
    readonly name?: string;
    /**
     * The list of policies.
     */
    readonly policies: outputs.identity.GetPoliciesPolicy[];
    /**
     * The policy's current state.
     */
    readonly state?: string;
}