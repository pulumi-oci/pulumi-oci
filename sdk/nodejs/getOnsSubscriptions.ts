// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the list of Subscriptions in Oracle Cloud Infrastructure Notifications service.
 *
 * Lists the subscriptions in the specified compartment or topic.
 *
 * Transactions Per Minute (TPM) per-tenancy limit for this operation: 60.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testSubscriptions = oci.GetOnsSubscriptions({
 *     compartmentId: _var.compartment_id,
 *     topicId: oci_ons_notification_topic.test_notification_topic.id,
 * });
 * ```
 */
export function getOnsSubscriptions(args: GetOnsSubscriptionsArgs, opts?: pulumi.InvokeOptions): Promise<GetOnsSubscriptionsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getOnsSubscriptions:GetOnsSubscriptions", {
        "compartmentId": args.compartmentId,
        "filters": args.filters,
        "topicId": args.topicId,
    }, opts);
}

/**
 * A collection of arguments for invoking GetOnsSubscriptions.
 */
export interface GetOnsSubscriptionsArgs {
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
     */
    compartmentId: string;
    filters?: inputs.GetOnsSubscriptionsFilter[];
    /**
     * Return all subscriptions that are subscribed to the given topic OCID. Either this query parameter or the compartmentId query parameter must be set.
     */
    topicId?: string;
}

/**
 * A collection of values returned by GetOnsSubscriptions.
 */
export interface GetOnsSubscriptionsResult {
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for the subscription.
     */
    readonly compartmentId: string;
    readonly filters?: outputs.GetOnsSubscriptionsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of subscriptions.
     */
    readonly subscriptions: outputs.GetOnsSubscriptionsSubscription[];
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated topic.
     */
    readonly topicId?: string;
}