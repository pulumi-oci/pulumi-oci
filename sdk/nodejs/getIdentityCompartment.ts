// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides details about a specific Compartment resource in Oracle Cloud Infrastructure Identity service.
 *
 * Gets the specified compartment's information.
 *
 * This operation does not return a list of all the resources inside the compartment. There is no single
 * API operation that does that. Compartments can contain multiple types of resources (instances, block
 * storage volumes, etc.). To find out what's in a compartment, you must call the "List" operation for
 * each resource type and specify the compartment's OCID as a query parameter in the request. For example,
 * call the [ListInstances](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Instance/ListInstances) operation in the Cloud Compute
 * Service or the [ListVolumes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Volume/ListVolumes) operation in Cloud Block Storage.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testCompartment = oci.GetIdentityCompartment({
 *     id: _var.compartment_id,
 * });
 * ```
 */
export function getIdentityCompartment(args: GetIdentityCompartmentArgs, opts?: pulumi.InvokeOptions): Promise<GetIdentityCompartmentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getIdentityCompartment:GetIdentityCompartment", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking GetIdentityCompartment.
 */
export interface GetIdentityCompartmentArgs {
    /**
     * The OCID of the compartment.
     */
    id: string;
}

/**
 * A collection of values returned by GetIdentityCompartment.
 */
export interface GetIdentityCompartmentResult {
    /**
     * The OCID of the parent compartment containing the compartment.
     */
    readonly compartmentId: string;
    /**
     * Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
     */
    readonly definedTags: {[key: string]: any};
    /**
     * The description you assign to the compartment. Does not have to be unique, and it's changeable.
     */
    readonly description: string;
    /**
     * Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
     */
    readonly freeformTags: {[key: string]: any};
    /**
     * The OCID of the compartment.
     */
    readonly id: string;
    /**
     * The detailed status of INACTIVE lifecycleState.
     */
    readonly inactiveState: string;
    /**
     * Indicates whether or not the compartment is accessible for the user making the request. Returns true when the user has INSPECT permissions directly on a resource in the compartment or indirectly (permissions can be on a resource in a subcompartment).
     */
    readonly isAccessible: boolean;
    /**
     * The name you assign to the compartment during creation. The name must be unique across all compartments in the parent. Avoid entering confidential information.
     */
    readonly name: string;
    /**
     * The compartment's current state.
     */
    readonly state: string;
    /**
     * Date and time the compartment was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`
     */
    readonly timeCreated: string;
}