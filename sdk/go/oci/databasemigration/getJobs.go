// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databasemigration

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Jobs in Oracle Cloud Infrastructure Database Migration service.
//
// List all the names of the Migration jobs associated to the specified
// migration site.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/databasemigration"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Job_display_name
// 		opt1 := _var.Job_state
// 		_, err := databasemigration.GetJobs(ctx, &databasemigration.GetJobsArgs{
// 			MigrationId: oci_database_migration_migration.Test_migration.Id,
// 			DisplayName: &opt0,
// 			State:       &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetJobs(ctx *pulumi.Context, args *GetJobsArgs, opts ...pulumi.InvokeOption) (*GetJobsResult, error) {
	var rv GetJobsResult
	err := ctx.Invoke("oci:databasemigration/getJobs:getJobs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getJobs.
type GetJobsArgs struct {
	// A filter to return only resources that match the entire display name given.
	DisplayName *string         `pulumi:"displayName"`
	Filters     []GetJobsFilter `pulumi:"filters"`
	// The ID of the migration in which to list resources.
	MigrationId string `pulumi:"migrationId"`
	// The lifecycle state of the Migration Job.
	State *string `pulumi:"state"`
}

// A collection of values returned by getJobs.
type GetJobsResult struct {
	// Name of the job.
	DisplayName *string         `pulumi:"displayName"`
	Filters     []GetJobsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of job_collection.
	JobCollections []GetJobsJobCollection `pulumi:"jobCollections"`
	// The OCID of the Migration that this job belongs to.
	MigrationId string `pulumi:"migrationId"`
	// The current state of the migration job.
	State *string `pulumi:"state"`
}