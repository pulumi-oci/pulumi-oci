// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Private Application Packages in Oracle Cloud Infrastructure Service Catalog service.
//
// Lists the packages in the specified private application.
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
// 		opt0 := _var.Private_application_package_display_name
// 		opt1 := oci_service_catalog_private_application_package.Test_private_application_package.Id
// 		_, err := servicecatalog.GetPrivateApplicationPackages(ctx, &servicecatalog.GetPrivateApplicationPackagesArgs{
// 			PrivateApplicationId:        oci_service_catalog_private_application.Test_private_application.Id,
// 			DisplayName:                 &opt0,
// 			PackageTypes:                _var.Private_application_package_package_type,
// 			PrivateApplicationPackageId: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetPrivateApplicationPackages(ctx *pulumi.Context, args *GetPrivateApplicationPackagesArgs, opts ...pulumi.InvokeOption) (*GetPrivateApplicationPackagesResult, error) {
	var rv GetPrivateApplicationPackagesResult
	err := ctx.Invoke("oci:servicecatalog/getPrivateApplicationPackages:getPrivateApplicationPackages", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrivateApplicationPackages.
type GetPrivateApplicationPackagesArgs struct {
	// Exact match name filter.
	DisplayName *string                               `pulumi:"displayName"`
	Filters     []GetPrivateApplicationPackagesFilter `pulumi:"filters"`
	// Name of the package type. If multiple package types are provided, then any resource with one or more matching package types will be returned.
	PackageTypes []string `pulumi:"packageTypes"`
	// The unique identifier for the private application.
	PrivateApplicationId string `pulumi:"privateApplicationId"`
	// The unique identifier for the private application package.
	PrivateApplicationPackageId *string `pulumi:"privateApplicationPackageId"`
}

// A collection of values returned by getPrivateApplicationPackages.
type GetPrivateApplicationPackagesResult struct {
	// The display name of the package.
	DisplayName *string                               `pulumi:"displayName"`
	Filters     []GetPrivateApplicationPackagesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The specified package's type.
	PackageTypes []string `pulumi:"packageTypes"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private application where the package is hosted.
	PrivateApplicationId string `pulumi:"privateApplicationId"`
	// The list of private_application_package_collection.
	PrivateApplicationPackageCollections []GetPrivateApplicationPackagesPrivateApplicationPackageCollection `pulumi:"privateApplicationPackageCollections"`
	PrivateApplicationPackageId          *string                                                            `pulumi:"privateApplicationPackageId"`
}