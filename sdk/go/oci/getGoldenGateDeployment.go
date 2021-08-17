// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Deployment resource in Oracle Cloud Infrastructure Golden Gate service.
//
// Retrieves a deployment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := oci.GetGoldenGateDeployment(ctx, &GetGoldenGateDeploymentArgs{
// 			DeploymentId: oci_golden_gate_deployment.Test_deployment.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupGoldenGateDeployment(ctx *pulumi.Context, args *LookupGoldenGateDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupGoldenGateDeploymentResult, error) {
	var rv LookupGoldenGateDeploymentResult
	err := ctx.Invoke("oci:index/getGoldenGateDeployment:GetGoldenGateDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetGoldenGateDeployment.
type LookupGoldenGateDeploymentArgs struct {
	// A unique Deployment identifier.
	DeploymentId string `pulumi:"deploymentId"`
}

// A collection of values returned by GetGoldenGateDeployment.
type LookupGoldenGateDeploymentResult struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId string `pulumi:"compartmentId"`
	// The Minimum number of OCPUs to be made available for this Deployment.
	CpuCoreCount int `pulumi:"cpuCoreCount"`
	// Tags defined for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup being referenced.
	DeploymentBackupId string `pulumi:"deploymentBackupId"`
	DeploymentId       string `pulumi:"deploymentId"`
	// The deployment type.
	DeploymentType string `pulumi:"deploymentType"`
	// The URL of a resource.
	DeploymentUrl string `pulumi:"deploymentUrl"`
	// Metadata about this specific object.
	Description string `pulumi:"description"`
	// An object's Display Name.
	DisplayName string `pulumi:"displayName"`
	// A three-label Fully Qualified Domain Name (FQDN) for a resource.
	Fqdn string `pulumi:"fqdn"`
	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced.
	Id string `pulumi:"id"`
	// Indicates if auto scaling is enabled for the Deployment's CPU core count.
	IsAutoScalingEnabled bool `pulumi:"isAutoScalingEnabled"`
	// True if all of the aggregate resources are working correctly.
	IsHealthy bool `pulumi:"isHealthy"`
	// Indicates if the resource is the the latest available version.
	IsLatestVersion bool `pulumi:"isLatestVersion"`
	// True if this object is publicly available.
	IsPublic bool `pulumi:"isPublic"`
	// The Oracle license model that applies to a Deployment.
	LicenseModel string `pulumi:"licenseModel"`
	// Describes the object's current state in detail. For example, it can be used to provide actionable information for a resource in a Failed state.
	LifecycleDetails string `pulumi:"lifecycleDetails"`
	// An array of [Network Security Group](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/networksecuritygroups.htm) OCIDs used to define network access for a deployment.
	NsgIds []string `pulumi:"nsgIds"`
	// Deployment Data for an OggDeployment
	OggData GetGoldenGateDeploymentOggData `pulumi:"oggData"`
	// The private IP address in the customer's VCN representing the access point for the associated endpoint service in the GoldenGate service VCN.
	PrivateIpAddress string `pulumi:"privateIpAddress"`
	// The public IP address representing the access point for the Deployment.
	PublicIpAddress string `pulumi:"publicIpAddress"`
	// Possible lifecycle states.
	State string `pulumi:"state"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet being referenced.
	SubnetId string `pulumi:"subnetId"`
	// The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]interface{} `pulumi:"systemTags"`
	// The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeCreated string `pulumi:"timeCreated"`
	// The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated string `pulumi:"timeUpdated"`
}