// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Dedicated Vm Host Shapes in Oracle Cloud Infrastructure Core service.
//
// Lists the shapes that can be used to launch a dedicated virtual machine host within the specified compartment.
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
// 		opt0 := _var.Dedicated_vm_host_shape_availability_domain
// 		opt1 := _var.Dedicated_vm_host_shape_instance_shape_name
// 		_, err := core.GetDedicatedVmHostShapes(ctx, &core.GetDedicatedVmHostShapesArgs{
// 			CompartmentId:      _var.Compartment_id,
// 			AvailabilityDomain: &opt0,
// 			InstanceShapeName:  &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDedicatedVmHostShapes(ctx *pulumi.Context, args *GetDedicatedVmHostShapesArgs, opts ...pulumi.InvokeOption) (*GetDedicatedVmHostShapesResult, error) {
	var rv GetDedicatedVmHostShapesResult
	err := ctx.Invoke("oci:core/getDedicatedVmHostShapes:getDedicatedVmHostShapes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDedicatedVmHostShapes.
type GetDedicatedVmHostShapesArgs struct {
	// The name of the availability domain.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `pulumi:"availabilityDomain"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId string                           `pulumi:"compartmentId"`
	Filters       []GetDedicatedVmHostShapesFilter `pulumi:"filters"`
	// The name for the instance's shape.
	InstanceShapeName *string `pulumi:"instanceShapeName"`
}

// A collection of values returned by getDedicatedVmHostShapes.
type GetDedicatedVmHostShapesResult struct {
	// The shape's availability domain.
	AvailabilityDomain *string `pulumi:"availabilityDomain"`
	CompartmentId      string  `pulumi:"compartmentId"`
	// The list of dedicated_vm_host_shapes.
	DedicatedVmHostShapes []GetDedicatedVmHostShapesDedicatedVmHostShape `pulumi:"dedicatedVmHostShapes"`
	Filters               []GetDedicatedVmHostShapesFilter               `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id                string  `pulumi:"id"`
	InstanceShapeName *string `pulumi:"instanceShapeName"`
}
