// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Vm Cluster Network resource in Oracle Cloud Infrastructure Database service.
//
// Creates the VM cluster network. Applies to Exadata Cloud@Customer instances only.
// To create a cloud VM cluster in an Exadata Cloud Service instance, use the [CreateCloudVmCluster ](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/CloudVmCluster/CreateCloudVmCluster) operation.
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
// 		_, err := database.NewVmClusterNetwork(ctx, "testVmClusterNetwork", &database.VmClusterNetworkArgs{
// 			CompartmentId:           pulumi.Any(_var.Compartment_id),
// 			DisplayName:             pulumi.Any(_var.Vm_cluster_network_display_name),
// 			ExadataInfrastructureId: pulumi.Any(oci_database_exadata_infrastructure.Test_exadata_infrastructure.Id),
// 			Scans: database.VmClusterNetworkScanArray{
// 				&database.VmClusterNetworkScanArgs{
// 					Hostname: pulumi.Any(_var.Vm_cluster_network_scans_hostname),
// 					Ips:      pulumi.Any(_var.Vm_cluster_network_scans_ips),
// 					Port:     pulumi.Any(_var.Vm_cluster_network_scans_port),
// 				},
// 			},
// 			VmNetworks: database.VmClusterNetworkVmNetworkArray{
// 				&database.VmClusterNetworkVmNetworkArgs{
// 					DomainName:  pulumi.Any(_var.Vm_cluster_network_vm_networks_domain_name),
// 					Gateway:     pulumi.Any(_var.Vm_cluster_network_vm_networks_gateway),
// 					Netmask:     pulumi.Any(_var.Vm_cluster_network_vm_networks_netmask),
// 					NetworkType: pulumi.Any(_var.Vm_cluster_network_vm_networks_network_type),
// 					Nodes: database.VmClusterNetworkVmNetworkNodeArray{
// 						&database.VmClusterNetworkVmNetworkNodeArgs{
// 							Hostname:    pulumi.Any(_var.Vm_cluster_network_vm_networks_nodes_hostname),
// 							Ip:          pulumi.Any(_var.Vm_cluster_network_vm_networks_nodes_ip),
// 							Vip:         pulumi.Any(_var.Vm_cluster_network_vm_networks_nodes_vip),
// 							VipHostname: pulumi.Any(_var.Vm_cluster_network_vm_networks_nodes_vip_hostname),
// 						},
// 					},
// 					VlanId: pulumi.Any(_var.Vm_cluster_network_vm_networks_vlan_id),
// 				},
// 			},
// 			DefinedTags: pulumi.Any(_var.Vm_cluster_network_defined_tags),
// 			Dns:         pulumi.Any(_var.Vm_cluster_network_dns),
// 			FreeformTags: pulumi.AnyMap{
// 				"Department": pulumi.Any("Finance"),
// 			},
// 			Ntps:                     pulumi.Any(_var.Vm_cluster_network_ntp),
// 			ValidateVmClusterNetwork: pulumi.Any(_var.Vm_cluster_network_validate_vm_cluster_network),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// VmClusterNetworks can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:database/vmClusterNetwork:VmClusterNetwork test_vm_cluster_network "exadataInfrastructures/{exadataInfrastructureId}/vmClusterNetworks/{vmClusterNetworkId}"
// ```
type VmClusterNetwork struct {
	pulumi.CustomResourceState

	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// The user-friendly name for the Exadata Cloud@Customer VM cluster network. The name does not need to be unique.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// (Updatable) The list of DNS server IP addresses. Maximum of 3 allowed.
	Dns pulumi.StringArrayOutput `pulumi:"dns"`
	// The Exadata infrastructure [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInfrastructureId pulumi.StringOutput `pulumi:"exadataInfrastructureId"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// Additional information about the current lifecycle state.
	LifecycleDetails pulumi.StringOutput `pulumi:"lifecycleDetails"`
	// (Updatable) The list of NTP server IP addresses. Maximum of 3 allowed.
	Ntps pulumi.StringArrayOutput `pulumi:"ntps"`
	// (Updatable) The SCAN details.
	Scans VmClusterNetworkScanArrayOutput `pulumi:"scans"`
	// The current state of the VM cluster network.
	State pulumi.StringOutput `pulumi:"state"`
	// The date and time when the VM cluster network was created.
	TimeCreated              pulumi.StringOutput  `pulumi:"timeCreated"`
	ValidateVmClusterNetwork pulumi.BoolPtrOutput `pulumi:"validateVmClusterNetwork"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated VM Cluster.
	VmClusterId pulumi.StringOutput `pulumi:"vmClusterId"`
	// (Updatable) Details of the client and backup networks.
	VmNetworks VmClusterNetworkVmNetworkArrayOutput `pulumi:"vmNetworks"`
}

// NewVmClusterNetwork registers a new resource with the given unique name, arguments, and options.
func NewVmClusterNetwork(ctx *pulumi.Context,
	name string, args *VmClusterNetworkArgs, opts ...pulumi.ResourceOption) (*VmClusterNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ExadataInfrastructureId == nil {
		return nil, errors.New("invalid value for required argument 'ExadataInfrastructureId'")
	}
	if args.Scans == nil {
		return nil, errors.New("invalid value for required argument 'Scans'")
	}
	if args.VmNetworks == nil {
		return nil, errors.New("invalid value for required argument 'VmNetworks'")
	}
	var resource VmClusterNetwork
	err := ctx.RegisterResource("oci:database/vmClusterNetwork:VmClusterNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVmClusterNetwork gets an existing VmClusterNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVmClusterNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VmClusterNetworkState, opts ...pulumi.ResourceOption) (*VmClusterNetwork, error) {
	var resource VmClusterNetwork
	err := ctx.ReadResource("oci:database/vmClusterNetwork:VmClusterNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VmClusterNetwork resources.
type vmClusterNetworkState struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// The user-friendly name for the Exadata Cloud@Customer VM cluster network. The name does not need to be unique.
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) The list of DNS server IP addresses. Maximum of 3 allowed.
	Dns []string `pulumi:"dns"`
	// The Exadata infrastructure [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInfrastructureId *string `pulumi:"exadataInfrastructureId"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// Additional information about the current lifecycle state.
	LifecycleDetails *string `pulumi:"lifecycleDetails"`
	// (Updatable) The list of NTP server IP addresses. Maximum of 3 allowed.
	Ntps []string `pulumi:"ntps"`
	// (Updatable) The SCAN details.
	Scans []VmClusterNetworkScan `pulumi:"scans"`
	// The current state of the VM cluster network.
	State *string `pulumi:"state"`
	// The date and time when the VM cluster network was created.
	TimeCreated              *string `pulumi:"timeCreated"`
	ValidateVmClusterNetwork *bool   `pulumi:"validateVmClusterNetwork"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated VM Cluster.
	VmClusterId *string `pulumi:"vmClusterId"`
	// (Updatable) Details of the client and backup networks.
	VmNetworks []VmClusterNetworkVmNetwork `pulumi:"vmNetworks"`
}

type VmClusterNetworkState struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags pulumi.MapInput
	// The user-friendly name for the Exadata Cloud@Customer VM cluster network. The name does not need to be unique.
	DisplayName pulumi.StringPtrInput
	// (Updatable) The list of DNS server IP addresses. Maximum of 3 allowed.
	Dns pulumi.StringArrayInput
	// The Exadata infrastructure [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInfrastructureId pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// Additional information about the current lifecycle state.
	LifecycleDetails pulumi.StringPtrInput
	// (Updatable) The list of NTP server IP addresses. Maximum of 3 allowed.
	Ntps pulumi.StringArrayInput
	// (Updatable) The SCAN details.
	Scans VmClusterNetworkScanArrayInput
	// The current state of the VM cluster network.
	State pulumi.StringPtrInput
	// The date and time when the VM cluster network was created.
	TimeCreated              pulumi.StringPtrInput
	ValidateVmClusterNetwork pulumi.BoolPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated VM Cluster.
	VmClusterId pulumi.StringPtrInput
	// (Updatable) Details of the client and backup networks.
	VmNetworks VmClusterNetworkVmNetworkArrayInput
}

func (VmClusterNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmClusterNetworkState)(nil)).Elem()
}

type vmClusterNetworkArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// The user-friendly name for the Exadata Cloud@Customer VM cluster network. The name does not need to be unique.
	DisplayName string `pulumi:"displayName"`
	// (Updatable) The list of DNS server IP addresses. Maximum of 3 allowed.
	Dns []string `pulumi:"dns"`
	// The Exadata infrastructure [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInfrastructureId string `pulumi:"exadataInfrastructureId"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable) The list of NTP server IP addresses. Maximum of 3 allowed.
	Ntps []string `pulumi:"ntps"`
	// (Updatable) The SCAN details.
	Scans                    []VmClusterNetworkScan `pulumi:"scans"`
	ValidateVmClusterNetwork *bool                  `pulumi:"validateVmClusterNetwork"`
	// (Updatable) Details of the client and backup networks.
	VmNetworks []VmClusterNetworkVmNetwork `pulumi:"vmNetworks"`
}

// The set of arguments for constructing a VmClusterNetwork resource.
type VmClusterNetworkArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId pulumi.StringInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags pulumi.MapInput
	// The user-friendly name for the Exadata Cloud@Customer VM cluster network. The name does not need to be unique.
	DisplayName pulumi.StringInput
	// (Updatable) The list of DNS server IP addresses. Maximum of 3 allowed.
	Dns pulumi.StringArrayInput
	// The Exadata infrastructure [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInfrastructureId pulumi.StringInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// (Updatable) The list of NTP server IP addresses. Maximum of 3 allowed.
	Ntps pulumi.StringArrayInput
	// (Updatable) The SCAN details.
	Scans                    VmClusterNetworkScanArrayInput
	ValidateVmClusterNetwork pulumi.BoolPtrInput
	// (Updatable) Details of the client and backup networks.
	VmNetworks VmClusterNetworkVmNetworkArrayInput
}

func (VmClusterNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmClusterNetworkArgs)(nil)).Elem()
}

type VmClusterNetworkInput interface {
	pulumi.Input

	ToVmClusterNetworkOutput() VmClusterNetworkOutput
	ToVmClusterNetworkOutputWithContext(ctx context.Context) VmClusterNetworkOutput
}

func (*VmClusterNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((*VmClusterNetwork)(nil))
}

func (i *VmClusterNetwork) ToVmClusterNetworkOutput() VmClusterNetworkOutput {
	return i.ToVmClusterNetworkOutputWithContext(context.Background())
}

func (i *VmClusterNetwork) ToVmClusterNetworkOutputWithContext(ctx context.Context) VmClusterNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmClusterNetworkOutput)
}

func (i *VmClusterNetwork) ToVmClusterNetworkPtrOutput() VmClusterNetworkPtrOutput {
	return i.ToVmClusterNetworkPtrOutputWithContext(context.Background())
}

func (i *VmClusterNetwork) ToVmClusterNetworkPtrOutputWithContext(ctx context.Context) VmClusterNetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmClusterNetworkPtrOutput)
}

type VmClusterNetworkPtrInput interface {
	pulumi.Input

	ToVmClusterNetworkPtrOutput() VmClusterNetworkPtrOutput
	ToVmClusterNetworkPtrOutputWithContext(ctx context.Context) VmClusterNetworkPtrOutput
}

type vmClusterNetworkPtrType VmClusterNetworkArgs

func (*vmClusterNetworkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VmClusterNetwork)(nil))
}

func (i *vmClusterNetworkPtrType) ToVmClusterNetworkPtrOutput() VmClusterNetworkPtrOutput {
	return i.ToVmClusterNetworkPtrOutputWithContext(context.Background())
}

func (i *vmClusterNetworkPtrType) ToVmClusterNetworkPtrOutputWithContext(ctx context.Context) VmClusterNetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmClusterNetworkPtrOutput)
}

// VmClusterNetworkArrayInput is an input type that accepts VmClusterNetworkArray and VmClusterNetworkArrayOutput values.
// You can construct a concrete instance of `VmClusterNetworkArrayInput` via:
//
//          VmClusterNetworkArray{ VmClusterNetworkArgs{...} }
type VmClusterNetworkArrayInput interface {
	pulumi.Input

	ToVmClusterNetworkArrayOutput() VmClusterNetworkArrayOutput
	ToVmClusterNetworkArrayOutputWithContext(context.Context) VmClusterNetworkArrayOutput
}

type VmClusterNetworkArray []VmClusterNetworkInput

func (VmClusterNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VmClusterNetwork)(nil)).Elem()
}

func (i VmClusterNetworkArray) ToVmClusterNetworkArrayOutput() VmClusterNetworkArrayOutput {
	return i.ToVmClusterNetworkArrayOutputWithContext(context.Background())
}

func (i VmClusterNetworkArray) ToVmClusterNetworkArrayOutputWithContext(ctx context.Context) VmClusterNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmClusterNetworkArrayOutput)
}

// VmClusterNetworkMapInput is an input type that accepts VmClusterNetworkMap and VmClusterNetworkMapOutput values.
// You can construct a concrete instance of `VmClusterNetworkMapInput` via:
//
//          VmClusterNetworkMap{ "key": VmClusterNetworkArgs{...} }
type VmClusterNetworkMapInput interface {
	pulumi.Input

	ToVmClusterNetworkMapOutput() VmClusterNetworkMapOutput
	ToVmClusterNetworkMapOutputWithContext(context.Context) VmClusterNetworkMapOutput
}

type VmClusterNetworkMap map[string]VmClusterNetworkInput

func (VmClusterNetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VmClusterNetwork)(nil)).Elem()
}

func (i VmClusterNetworkMap) ToVmClusterNetworkMapOutput() VmClusterNetworkMapOutput {
	return i.ToVmClusterNetworkMapOutputWithContext(context.Background())
}

func (i VmClusterNetworkMap) ToVmClusterNetworkMapOutputWithContext(ctx context.Context) VmClusterNetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmClusterNetworkMapOutput)
}

type VmClusterNetworkOutput struct {
	*pulumi.OutputState
}

func (VmClusterNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmClusterNetwork)(nil))
}

func (o VmClusterNetworkOutput) ToVmClusterNetworkOutput() VmClusterNetworkOutput {
	return o
}

func (o VmClusterNetworkOutput) ToVmClusterNetworkOutputWithContext(ctx context.Context) VmClusterNetworkOutput {
	return o
}

func (o VmClusterNetworkOutput) ToVmClusterNetworkPtrOutput() VmClusterNetworkPtrOutput {
	return o.ToVmClusterNetworkPtrOutputWithContext(context.Background())
}

func (o VmClusterNetworkOutput) ToVmClusterNetworkPtrOutputWithContext(ctx context.Context) VmClusterNetworkPtrOutput {
	return o.ApplyT(func(v VmClusterNetwork) *VmClusterNetwork {
		return &v
	}).(VmClusterNetworkPtrOutput)
}

type VmClusterNetworkPtrOutput struct {
	*pulumi.OutputState
}

func (VmClusterNetworkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VmClusterNetwork)(nil))
}

func (o VmClusterNetworkPtrOutput) ToVmClusterNetworkPtrOutput() VmClusterNetworkPtrOutput {
	return o
}

func (o VmClusterNetworkPtrOutput) ToVmClusterNetworkPtrOutputWithContext(ctx context.Context) VmClusterNetworkPtrOutput {
	return o
}

type VmClusterNetworkArrayOutput struct{ *pulumi.OutputState }

func (VmClusterNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmClusterNetwork)(nil))
}

func (o VmClusterNetworkArrayOutput) ToVmClusterNetworkArrayOutput() VmClusterNetworkArrayOutput {
	return o
}

func (o VmClusterNetworkArrayOutput) ToVmClusterNetworkArrayOutputWithContext(ctx context.Context) VmClusterNetworkArrayOutput {
	return o
}

func (o VmClusterNetworkArrayOutput) Index(i pulumi.IntInput) VmClusterNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VmClusterNetwork {
		return vs[0].([]VmClusterNetwork)[vs[1].(int)]
	}).(VmClusterNetworkOutput)
}

type VmClusterNetworkMapOutput struct{ *pulumi.OutputState }

func (VmClusterNetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VmClusterNetwork)(nil))
}

func (o VmClusterNetworkMapOutput) ToVmClusterNetworkMapOutput() VmClusterNetworkMapOutput {
	return o
}

func (o VmClusterNetworkMapOutput) ToVmClusterNetworkMapOutputWithContext(ctx context.Context) VmClusterNetworkMapOutput {
	return o
}

func (o VmClusterNetworkMapOutput) MapIndex(k pulumi.StringInput) VmClusterNetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VmClusterNetwork {
		return vs[0].(map[string]VmClusterNetwork)[vs[1].(string)]
	}).(VmClusterNetworkOutput)
}

func init() {
	pulumi.RegisterOutputType(VmClusterNetworkOutput{})
	pulumi.RegisterOutputType(VmClusterNetworkPtrOutput{})
	pulumi.RegisterOutputType(VmClusterNetworkArrayOutput{})
	pulumi.RegisterOutputType(VmClusterNetworkMapOutput{})
}
