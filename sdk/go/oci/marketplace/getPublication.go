// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package marketplace

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Publication resource in Oracle Cloud Infrastructure Marketplace service.
//
// Get details of a publication
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/marketplace"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := marketplace.LookupPublication(ctx, &marketplace.LookupPublicationArgs{
// 			PublicationId: oci_marketplace_publication.Test_publication.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupPublication(ctx *pulumi.Context, args *LookupPublicationArgs, opts ...pulumi.InvokeOption) (*LookupPublicationResult, error) {
	var rv LookupPublicationResult
	err := ctx.Invoke("oci:marketplace/getPublication:getPublication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPublication.
type LookupPublicationArgs struct {
	// The unique identifier for the listing.
	PublicationId string `pulumi:"publicationId"`
}

// A collection of values returned by getPublication.
type LookupPublicationResult struct {
	// The Compartment id where the listings exists
	CompartmentId string `pulumi:"compartmentId"`
	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The model for upload data for images and icons.
	Icon GetPublicationIcon `pulumi:"icon"`
	// The unique identifier for the listing in Marketplace.
	Id                      string `pulumi:"id"`
	IsAgreementAcknowledged bool   `pulumi:"isAgreementAcknowledged"`
	// In which catalog the listing should exist.
	ListingType string `pulumi:"listingType"`
	// A long description of the listing.
	LongDescription string `pulumi:"longDescription"`
	// name of the operating system
	Name           string                       `pulumi:"name"`
	PackageDetails GetPublicationPackageDetails `pulumi:"packageDetails"`
	// The listing's package type.
	PackageType   string `pulumi:"packageType"`
	PublicationId string `pulumi:"publicationId"`
	// A short description of the listing.
	ShortDescription string `pulumi:"shortDescription"`
	// The state of the listing in its lifecycle
	State string `pulumi:"state"`
	// Contact information to use to get support from the publisher for the listing.
	SupportContacts []GetPublicationSupportContact `pulumi:"supportContacts"`
	// List of operating systems supprted.
	SupportedOperatingSystems []GetPublicationSupportedOperatingSystem `pulumi:"supportedOperatingSystems"`
	// The date and time this publication was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated string `pulumi:"timeCreated"`
}
