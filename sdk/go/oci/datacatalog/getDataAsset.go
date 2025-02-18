// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datacatalog

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Data Asset resource in Oracle Cloud Infrastructure Data Catalog service.
//
// Gets a specific data asset for the given key within a data catalog.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/datacatalog"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := datacatalog.LookupDataAsset(ctx, &datacatalog.LookupDataAssetArgs{
// 			CatalogId:    oci_datacatalog_catalog.Test_catalog.Id,
// 			DataAssetKey: _var.Data_asset_data_asset_key,
// 			Fields:       _var.Data_asset_fields,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupDataAsset(ctx *pulumi.Context, args *LookupDataAssetArgs, opts ...pulumi.InvokeOption) (*LookupDataAssetResult, error) {
	var rv LookupDataAssetResult
	err := ctx.Invoke("oci:datacatalog/getDataAsset:getDataAsset", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDataAsset.
type LookupDataAssetArgs struct {
	// Unique catalog identifier.
	CatalogId string `pulumi:"catalogId"`
	// Unique data asset key.
	DataAssetKey string `pulumi:"dataAssetKey"`
	// Specifies the fields to return in a data asset response.
	Fields []string `pulumi:"fields"`
}

// A collection of values returned by getDataAsset.
type LookupDataAssetResult struct {
	// The data catalog's OCID.
	CatalogId string `pulumi:"catalogId"`
	// OCID of the user who created the data asset.
	CreatedById  string `pulumi:"createdById"`
	DataAssetKey string `pulumi:"dataAssetKey"`
	// Detailed description of the data asset.
	Description string `pulumi:"description"`
	// A user-friendly display name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName string `pulumi:"displayName"`
	// External URI that can be used to reference the object. Format will differ based on the type of object.
	ExternalKey string   `pulumi:"externalKey"`
	Fields      []string `pulumi:"fields"`
	Id          string   `pulumi:"id"`
	// Unique data asset key that is immutable.
	Key string `pulumi:"key"`
	// A map of maps that contains the properties which are specific to the asset type. Each data asset type definition defines it's set of required and optional properties. The map keys are category names and the values are maps of property name to property value. Every property is contained inside of a category. Most data assets have required properties within the "default" category. Example: `{"properties": { "default": { "host": "host1", "port": "1521", "database": "orcl"}}}`
	Properties map[string]interface{} `pulumi:"properties"`
	// The current state of the data asset.
	State string `pulumi:"state"`
	// The date and time the data asset was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2019-03-25T21:10:29.600Z`
	TimeCreated string `pulumi:"timeCreated"`
	// The last time that any change was made to the data asset. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated string `pulumi:"timeUpdated"`
	// The key of the object type. Type key's can be found via the '/types' endpoint.
	TypeKey string `pulumi:"typeKey"`
	// OCID of the user who last modified the data asset.
	UpdatedById string `pulumi:"updatedById"`
	// URI to the data asset instance in the API.
	Uri string `pulumi:"uri"`
}
