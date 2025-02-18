// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of App Catalog Listing Resource Versions in Oracle Cloud Infrastructure Core service.
//
// Gets all resource versions for a particular listing.
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
// 		_, err := core.GetAppCatalogListingResourceVersions(ctx, &core.GetAppCatalogListingResourceVersionsArgs{
// 			ListingId: data.Oci_core_app_catalog_listing.Test_listing.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetAppCatalogListingResourceVersions(ctx *pulumi.Context, args *GetAppCatalogListingResourceVersionsArgs, opts ...pulumi.InvokeOption) (*GetAppCatalogListingResourceVersionsResult, error) {
	var rv GetAppCatalogListingResourceVersionsResult
	err := ctx.Invoke("oci:core/getAppCatalogListingResourceVersions:getAppCatalogListingResourceVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAppCatalogListingResourceVersions.
type GetAppCatalogListingResourceVersionsArgs struct {
	Filters []GetAppCatalogListingResourceVersionsFilter `pulumi:"filters"`
	// The OCID of the listing.
	ListingId string `pulumi:"listingId"`
}

// A collection of values returned by getAppCatalogListingResourceVersions.
type GetAppCatalogListingResourceVersionsResult struct {
	// The list of app_catalog_listing_resource_versions.
	AppCatalogListingResourceVersions []GetAppCatalogListingResourceVersionsAppCatalogListingResourceVersion `pulumi:"appCatalogListingResourceVersions"`
	Filters                           []GetAppCatalogListingResourceVersionsFilter                           `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The OCID of the listing this resource version belongs to.
	ListingId string `pulumi:"listingId"`
}
