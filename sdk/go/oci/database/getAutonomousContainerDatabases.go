// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Autonomous Container Databases in Oracle Cloud Infrastructure Database service.
//
// Gets a list of the Autonomous Container Databases in the specified compartment.
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
// 		opt0 := oci_database_autonomous_exadata_infrastructure.Test_autonomous_exadata_infrastructure.Id
// 		opt1 := oci_database_autonomous_vm_cluster.Test_autonomous_vm_cluster.Id
// 		opt2 := _var.Autonomous_container_database_availability_domain
// 		opt3 := _var.Autonomous_container_database_display_name
// 		opt4 := _var.Autonomous_container_database_infrastructure_type
// 		opt5 := _var.Autonomous_container_database_service_level_agreement_type
// 		opt6 := _var.Autonomous_container_database_state
// 		_, err := database.GetAutonomousContainerDatabases(ctx, &database.GetAutonomousContainerDatabasesArgs{
// 			CompartmentId:                     _var.Compartment_id,
// 			AutonomousExadataInfrastructureId: &opt0,
// 			AutonomousVmClusterId:             &opt1,
// 			AvailabilityDomain:                &opt2,
// 			DisplayName:                       &opt3,
// 			InfrastructureType:                &opt4,
// 			ServiceLevelAgreementType:         &opt5,
// 			State:                             &opt6,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetAutonomousContainerDatabases(ctx *pulumi.Context, args *GetAutonomousContainerDatabasesArgs, opts ...pulumi.InvokeOption) (*GetAutonomousContainerDatabasesResult, error) {
	var rv GetAutonomousContainerDatabasesResult
	err := ctx.Invoke("oci:database/getAutonomousContainerDatabases:getAutonomousContainerDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAutonomousContainerDatabases.
type GetAutonomousContainerDatabasesArgs struct {
	// The Autonomous Exadata Infrastructure [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	AutonomousExadataInfrastructureId *string `pulumi:"autonomousExadataInfrastructureId"`
	// The Autonomous VM Cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	AutonomousVmClusterId *string `pulumi:"autonomousVmClusterId"`
	// A filter to return only resources that match the given availability domain exactly.
	AvailabilityDomain *string `pulumi:"availabilityDomain"`
	// The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId string `pulumi:"compartmentId"`
	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string                                 `pulumi:"displayName"`
	Filters     []GetAutonomousContainerDatabasesFilter `pulumi:"filters"`
	// A filter to return only resources that match the given Infrastructure Type.
	InfrastructureType *string `pulumi:"infrastructureType"`
	// A filter to return only resources that match the given service level agreement type exactly.
	ServiceLevelAgreementType *string `pulumi:"serviceLevelAgreementType"`
	// A filter to return only resources that match the given lifecycle state exactly.
	State *string `pulumi:"state"`
}

// A collection of values returned by getAutonomousContainerDatabases.
type GetAutonomousContainerDatabasesResult struct {
	// The list of autonomous_container_databases.
	AutonomousContainerDatabases []GetAutonomousContainerDatabasesAutonomousContainerDatabase `pulumi:"autonomousContainerDatabases"`
	// The OCID of the Autonomous Exadata Infrastructure.
	AutonomousExadataInfrastructureId *string `pulumi:"autonomousExadataInfrastructureId"`
	// The OCID of the Autonomous VM Cluster.
	AutonomousVmClusterId *string `pulumi:"autonomousVmClusterId"`
	// The availability domain of the Autonomous Container Database.
	AvailabilityDomain *string `pulumi:"availabilityDomain"`
	// The OCID of the compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// The user-provided name for the Autonomous Container Database.
	DisplayName *string                                 `pulumi:"displayName"`
	Filters     []GetAutonomousContainerDatabasesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The infrastructure type this resource belongs to.
	InfrastructureType *string `pulumi:"infrastructureType"`
	// The service level agreement type of the container database. The default is STANDARD.
	ServiceLevelAgreementType *string `pulumi:"serviceLevelAgreementType"`
	// The current state of the Autonomous Container Database.
	State *string `pulumi:"state"`
}
