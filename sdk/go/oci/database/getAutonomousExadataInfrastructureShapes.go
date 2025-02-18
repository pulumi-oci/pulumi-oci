// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Autonomous Exadata Infrastructure Shapes in Oracle Cloud Infrastructure Database service.
//
// Gets a list of the shapes that can be used to launch a new Autonomous Exadata Infrastructure resource. The shape determines resources to allocate (CPU cores, memory and storage).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/database"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := database.GetAutonomousExadataInfrastructureShapes(ctx, &database.GetAutonomousExadataInfrastructureShapesArgs{
// 			AvailabilityDomain: _var.Autonomous_exadata_infrastructure_shape_availability_domain,
// 			CompartmentId:      _var.Compartment_id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetAutonomousExadataInfrastructureShapes(ctx *pulumi.Context, args *GetAutonomousExadataInfrastructureShapesArgs, opts ...pulumi.InvokeOption) (*GetAutonomousExadataInfrastructureShapesResult, error) {
	var rv GetAutonomousExadataInfrastructureShapesResult
	err := ctx.Invoke("oci:database/getAutonomousExadataInfrastructureShapes:getAutonomousExadataInfrastructureShapes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAutonomousExadataInfrastructureShapes.
type GetAutonomousExadataInfrastructureShapesArgs struct {
	// The name of the Availability Domain.
	AvailabilityDomain string `pulumi:"availabilityDomain"`
	// The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId string                                           `pulumi:"compartmentId"`
	Filters       []GetAutonomousExadataInfrastructureShapesFilter `pulumi:"filters"`
}

// A collection of values returned by getAutonomousExadataInfrastructureShapes.
type GetAutonomousExadataInfrastructureShapesResult struct {
	// The list of autonomous_exadata_infrastructure_shapes.
	AutonomousExadataInfrastructureShapes []GetAutonomousExadataInfrastructureShapesAutonomousExadataInfrastructureShape `pulumi:"autonomousExadataInfrastructureShapes"`
	AvailabilityDomain                    string                                                                         `pulumi:"availabilityDomain"`
	CompartmentId                         string                                                                         `pulumi:"compartmentId"`
	Filters                               []GetAutonomousExadataInfrastructureShapesFilter                               `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
