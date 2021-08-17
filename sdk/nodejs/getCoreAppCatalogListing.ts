// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides details about a specific App Catalog Listing resource in Oracle Cloud Infrastructure Core service.
 *
 * Gets the specified listing.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testAppCatalogListing = oci.GetCoreAppCatalogListing({
 *     listingId: data.oci_core_app_catalog_listing.test_listing.id,
 * });
 * ```
 */
export function getCoreAppCatalogListing(args: GetCoreAppCatalogListingArgs, opts?: pulumi.InvokeOptions): Promise<GetCoreAppCatalogListingResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getCoreAppCatalogListing:GetCoreAppCatalogListing", {
        "listingId": args.listingId,
    }, opts);
}

/**
 * A collection of arguments for invoking GetCoreAppCatalogListing.
 */
export interface GetCoreAppCatalogListingArgs {
    /**
     * The OCID of the listing.
     */
    listingId: string;
}

/**
 * A collection of values returned by GetCoreAppCatalogListing.
 */
export interface GetCoreAppCatalogListingResult {
    /**
     * Listing's contact URL.
     */
    readonly contactUrl: string;
    /**
     * Description of the listing.
     */
    readonly description: string;
    /**
     * The display name of the listing.
     */
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * the region free ocid of the listing resource.
     */
    readonly listingId: string;
    /**
     * Publisher's logo URL.
     */
    readonly publisherLogoUrl: string;
    /**
     * The name of the publisher who published this listing.
     */
    readonly publisherName: string;
    /**
     * The short summary for the listing.
     */
    readonly summary: string;
    /**
     * Date and time the listing was published, in [RFC3339](https://tools.ietf.org/html/rfc3339) format. Example: `2018-03-20T12:32:53.532Z`
     */
    readonly timePublished: string;
}