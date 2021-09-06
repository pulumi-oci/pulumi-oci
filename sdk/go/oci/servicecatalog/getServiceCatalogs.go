// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Service Catalogs in Oracle Cloud Infrastructure Service Catalog service.
//
// Lists all the service catalogs in the given compartment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/servicecatalog"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Service_catalog_display_name
// 		opt1 := oci_service_catalog_service_catalog.Test_service_catalog.Id
// 		_, err := servicecatalog.GetServiceCatalogs(ctx, &servicecatalog.GetServiceCatalogsArgs{
// 			CompartmentId:    _var.Compartment_id,
// 			DisplayName:      &opt0,
// 			ServiceCatalogId: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetServiceCatalogs(ctx *pulumi.Context, args *GetServiceCatalogsArgs, opts ...pulumi.InvokeOption) (*GetServiceCatalogsResult, error) {
	var rv GetServiceCatalogsResult
	err := ctx.Invoke("oci:servicecatalog/getServiceCatalogs:getServiceCatalogs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceCatalogs.
type GetServiceCatalogsArgs struct {
	// The unique identifier for the compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// Exact match name filter.
	DisplayName *string                    `pulumi:"displayName"`
	Filters     []GetServiceCatalogsFilter `pulumi:"filters"`
	// The unique identifier for the service catalog.
	ServiceCatalogId *string `pulumi:"serviceCatalogId"`
}

// A collection of values returned by getServiceCatalogs.
type GetServiceCatalogsResult struct {
	// The Compartment id where the service catalog exists
	CompartmentId string `pulumi:"compartmentId"`
	// The name of the service catalog.
	DisplayName *string                    `pulumi:"displayName"`
	Filters     []GetServiceCatalogsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of service_catalog_collection.
	ServiceCatalogCollections []GetServiceCatalogsServiceCatalogCollection `pulumi:"serviceCatalogCollections"`
	ServiceCatalogId          *string                                      `pulumi:"serviceCatalogId"`
}