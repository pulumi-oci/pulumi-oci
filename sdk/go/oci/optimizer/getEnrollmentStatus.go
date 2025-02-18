// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package optimizer

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Enrollment Status resource in Oracle Cloud Infrastructure Optimizer service.
//
// Gets the Cloud Advisor enrollment status.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/optimizer"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := optimizer.LookupEnrollmentStatus(ctx, &optimizer.LookupEnrollmentStatusArgs{
// 			EnrollmentStatusId: oci_optimizer_enrollment_status.Test_enrollment_status.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupEnrollmentStatus(ctx *pulumi.Context, args *LookupEnrollmentStatusArgs, opts ...pulumi.InvokeOption) (*LookupEnrollmentStatusResult, error) {
	var rv LookupEnrollmentStatusResult
	err := ctx.Invoke("oci:optimizer/getEnrollmentStatus:getEnrollmentStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEnrollmentStatus.
type LookupEnrollmentStatusArgs struct {
	// The unique OCID associated with the enrollment status.
	EnrollmentStatusId string `pulumi:"enrollmentStatusId"`
}

// A collection of values returned by getEnrollmentStatus.
type LookupEnrollmentStatusResult struct {
	// The OCID of the compartment.
	CompartmentId      string `pulumi:"compartmentId"`
	EnrollmentStatusId string `pulumi:"enrollmentStatusId"`
	// The OCID of the enrollment status.
	Id string `pulumi:"id"`
	// The enrollment status' current state.
	State string `pulumi:"state"`
	// The current Cloud Advisor enrollment status.
	Status string `pulumi:"status"`
	// The reason for the enrollment status of the tenancy.
	StatusReason string `pulumi:"statusReason"`
	// The date and time the enrollment status was created, in the format defined by RFC3339.
	TimeCreated string `pulumi:"timeCreated"`
	// The date and time the enrollment status was last updated, in the format defined by RFC3339.
	TimeUpdated string `pulumi:"timeUpdated"`
}
