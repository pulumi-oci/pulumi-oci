// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Image Shapes in Oracle Cloud Infrastructure Core service.
//
// Lists the compatible shapes for the specified image.
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
// 		_, err := core.GetImageShapes(ctx, &core.GetImageShapesArgs{
// 			ImageId: oci_core_image.Test_image.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetImageShapes(ctx *pulumi.Context, args *GetImageShapesArgs, opts ...pulumi.InvokeOption) (*GetImageShapesResult, error) {
	var rv GetImageShapesResult
	err := ctx.Invoke("oci:core/getImageShapes:getImageShapes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImageShapes.
type GetImageShapesArgs struct {
	Filters []GetImageShapesFilter `pulumi:"filters"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the image.
	ImageId string `pulumi:"imageId"`
}

// A collection of values returned by getImageShapes.
type GetImageShapesResult struct {
	Filters []GetImageShapesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ImageId string `pulumi:"imageId"`
	// The list of image_shape_compatibilities.
	ImageShapeCompatibilities []GetImageShapesImageShapeCompatibility `pulumi:"imageShapeCompatibilities"`
}
