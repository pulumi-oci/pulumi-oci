// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Autonomous Database Instance Wallet Management resource in Oracle Cloud Infrastructure Database service.
//
// Gets the wallet details for the specified Autonomous Database.
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
// 		_, err := database.LookupAutonomousDatabaseInstanceWalletManagement(ctx, &database.LookupAutonomousDatabaseInstanceWalletManagementArgs{
// 			AutonomousDatabaseId: oci_database_autonomous_database.Test_autonomous_database.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupAutonomousDatabaseInstanceWalletManagement(ctx *pulumi.Context, args *LookupAutonomousDatabaseInstanceWalletManagementArgs, opts ...pulumi.InvokeOption) (*LookupAutonomousDatabaseInstanceWalletManagementResult, error) {
	var rv LookupAutonomousDatabaseInstanceWalletManagementResult
	err := ctx.Invoke("oci:database/getAutonomousDatabaseInstanceWalletManagement:getAutonomousDatabaseInstanceWalletManagement", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAutonomousDatabaseInstanceWalletManagement.
type LookupAutonomousDatabaseInstanceWalletManagementArgs struct {
	// The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	AutonomousDatabaseId string `pulumi:"autonomousDatabaseId"`
}

// A collection of values returned by getAutonomousDatabaseInstanceWalletManagement.
type LookupAutonomousDatabaseInstanceWalletManagementResult struct {
	// The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	AutonomousDatabaseId string `pulumi:"autonomousDatabaseId"`
	Id                   string `pulumi:"id"`
	// Indicates whether to rotate the wallet or not. If `false`, the wallet will not be rotated. The default is `false`.
	ShouldRotate bool `pulumi:"shouldRotate"`
	// The current lifecycle state of the Autonomous Database wallet.
	State string `pulumi:"state"`
	// The date and time the wallet was last rotated.
	TimeRotated string `pulumi:"timeRotated"`
}
