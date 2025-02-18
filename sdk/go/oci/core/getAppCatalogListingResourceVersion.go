// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific App Catalog Listing Resource Version resource in Oracle Cloud Infrastructure Core service.
//
// Gets the specified listing resource version.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := core.GetAppCatalogListingResourceVersion(ctx, &core.GetAppCatalogListingResourceVersionArgs{
// 			ListingId:       data.Oci_core_app_catalog_listing.Test_listing.Id,
// 			ResourceVersion: _var.App_catalog_listing_resource_version_resource_version,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetAppCatalogListingResourceVersion(ctx *pulumi.Context, args *GetAppCatalogListingResourceVersionArgs, opts ...pulumi.InvokeOption) (*GetAppCatalogListingResourceVersionResult, error) {
	var rv GetAppCatalogListingResourceVersionResult
	err := ctx.Invoke("oci:core/getAppCatalogListingResourceVersion:getAppCatalogListingResourceVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppCatalogListingResourceVersion.
type GetAppCatalogListingResourceVersionArgs struct {
	// The OCID of the listing.
	ListingId string `pulumi:"listingId"`
	// Listing Resource Version.
	ResourceVersion string `pulumi:"resourceVersion"`
}

// A collection of values returned by getAppCatalogListingResourceVersion.
type GetAppCatalogListingResourceVersionResult struct {
	// List of accessible ports for instances launched with this listing resource version.
	AccessiblePorts []int `pulumi:"accessiblePorts"`
	// Allowed actions for the listing resource.
	AllowedActions []string `pulumi:"allowedActions"`
	// List of regions that this listing resource version is available.
	AvailableRegions []string `pulumi:"availableRegions"`
	// Array of shapes compatible with this resource.
	CompatibleShapes []string `pulumi:"compatibleShapes"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The OCID of the listing this resource version belongs to.
	ListingId string `pulumi:"listingId"`
	// OCID of the listing resource.
	ListingResourceId string `pulumi:"listingResourceId"`
	// Resource Version.
	ListingResourceVersion string `pulumi:"listingResourceVersion"`
	ResourceVersion        string `pulumi:"resourceVersion"`
	// Date and time the listing resource version was published, in [RFC3339](https://tools.ietf.org/html/rfc3339) format. Example: `2018-03-20T12:32:53.532Z`
	TimePublished string `pulumi:"timePublished"`
}
