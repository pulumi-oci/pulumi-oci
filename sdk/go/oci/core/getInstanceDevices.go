// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Instance Devices in Oracle Cloud Infrastructure Core service.
//
// Gets a list of all the devices for given instance. You can optionally filter results by device availability.
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
// 		opt0 := _var.Instance_device_is_available
// 		opt1 := _var.Instance_device_name
// 		_, err := core.GetInstanceDevices(ctx, &core.GetInstanceDevicesArgs{
// 			InstanceId:  oci_core_instance.Test_instance.Id,
// 			IsAvailable: &opt0,
// 			Name:        &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetInstanceDevices(ctx *pulumi.Context, args *GetInstanceDevicesArgs, opts ...pulumi.InvokeOption) (*GetInstanceDevicesResult, error) {
	var rv GetInstanceDevicesResult
	err := ctx.Invoke("oci:core/getInstanceDevices:getInstanceDevices", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceDevices.
type GetInstanceDevicesArgs struct {
	Filters []GetInstanceDevicesFilter `pulumi:"filters"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.
	InstanceId string `pulumi:"instanceId"`
	// A filter to return only available devices or only used devices.
	IsAvailable *bool `pulumi:"isAvailable"`
	// A filter to return only devices that match the given name exactly.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getInstanceDevices.
type GetInstanceDevicesResult struct {
	// The list of devices.
	Devices []GetInstanceDevicesDevice `pulumi:"devices"`
	Filters []GetInstanceDevicesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// The flag denoting whether device is available.
	IsAvailable *bool `pulumi:"isAvailable"`
	// The device name.
	Name *string `pulumi:"name"`
}
