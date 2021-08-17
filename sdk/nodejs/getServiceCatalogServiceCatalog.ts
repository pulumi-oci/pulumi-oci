// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides details about a specific Service Catalog resource in Oracle Cloud Infrastructure Service Catalog service.
 *
 * Gets detailed information about the service catalog including name, compartmentId
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testServiceCatalog = oci.GetServiceCatalogServiceCatalog({
 *     serviceCatalogId: oci_service_catalog_service_catalog.test_service_catalog.id,
 * });
 * ```
 */
export function getServiceCatalogServiceCatalog(args: GetServiceCatalogServiceCatalogArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceCatalogServiceCatalogResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getServiceCatalogServiceCatalog:GetServiceCatalogServiceCatalog", {
        "serviceCatalogId": args.serviceCatalogId,
    }, opts);
}

/**
 * A collection of arguments for invoking GetServiceCatalogServiceCatalog.
 */
export interface GetServiceCatalogServiceCatalogArgs {
    /**
     * The unique identifier for the service catalog.
     */
    serviceCatalogId: string;
}

/**
 * A collection of values returned by GetServiceCatalogServiceCatalog.
 */
export interface GetServiceCatalogServiceCatalogResult {
    /**
     * The Compartment id where the service catalog exists
     */
    readonly compartmentId: string;
    /**
     * Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
     */
    readonly definedTags: {[key: string]: any};
    /**
     * The name of the service catalog.
     */
    readonly displayName: string;
    /**
     * Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
     */
    readonly freeformTags: {[key: string]: any};
    /**
     * The unique identifier for the Service catalog.
     */
    readonly id: string;
    readonly serviceCatalogId: string;
    /**
     * The lifecycle state of the service catalog.
     */
    readonly state: string;
    /**
     * The date and time the service catalog was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2021-05-26T21:10:29.600Z`
     */
    readonly timeCreated: string;
    /**
     * The date and time the service catalog was last modified, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2021-12-10T05:10:29.721Z`
     */
    readonly timeUpdated: string;
}