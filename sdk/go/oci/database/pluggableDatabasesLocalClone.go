// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Pluggable Databases Local Clone resource in Oracle Cloud Infrastructure Database service. Although pluggable databases(PDB) belong to a container database(CDB), there is no change to the parent(CDB) as a result of this operation.
//
// Clones and starts a pluggable database (PDB) in the same database (CDB) as the source PDB. The source PDB must be in the `READ_WRITE` openMode to perform the clone operation.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/database"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := database.NewPluggableDatabasesLocalClone(ctx, "testPluggableDatabasesLocalClone", &database.PluggableDatabasesLocalCloneArgs{
// 			ClonedPdbName:           pulumi.Any(_var.Pluggable_databases_local_clone_cloned_pdb_name),
// 			PdbAdminPassword:        pulumi.Any(_var.Pluggable_databases_local_clone_pdb_admin_password),
// 			PluggableDatabaseId:     pulumi.Any(oci_database_pluggable_database.Test_pluggable_database.Id),
// 			TargetTdeWalletPassword: pulumi.Any(_var.Pluggable_databases_local_clone_target_tde_wallet_password),
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
// Import is not supported for this resource.
type PluggableDatabasesLocalClone struct {
	pulumi.CustomResourceState

	// The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
	ClonedPdbName pulumi.StringOutput `pulumi:"clonedPdbName"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// Connection strings to connect to an Oracle Pluggable Database.
	ConnectionStrings PluggableDatabasesLocalCloneConnectionStringsOutput `pulumi:"connectionStrings"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CDB.
	ContainerDatabaseId pulumi.StringOutput `pulumi:"containerDatabaseId"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// The restricted mode of the pluggable database. If a pluggable database is opened in restricted mode, the user needs both create a session and have restricted session privileges to connect to it.
	IsRestricted pulumi.BoolOutput `pulumi:"isRestricted"`
	// Detailed message for the lifecycle state.
	LifecycleDetails pulumi.StringOutput `pulumi:"lifecycleDetails"`
	// The mode that pluggable database is in. Open mode can only be changed to READ_ONLY or MIGRATE directly from the backend (within the Oracle Database software).
	OpenMode pulumi.StringOutput `pulumi:"openMode"`
	// A strong password for PDB Admin of the newly cloned PDB. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
	PdbAdminPassword pulumi.StringOutput `pulumi:"pdbAdminPassword"`
	// The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
	PdbName pulumi.StringOutput `pulumi:"pdbName"`
	// The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	PluggableDatabaseId pulumi.StringOutput `pulumi:"pluggableDatabaseId"`
	// The current state of the pluggable database.
	State pulumi.StringOutput `pulumi:"state"`
	// The existing TDE wallet password of the target CDB.
	TargetTdeWalletPassword pulumi.StringOutput `pulumi:"targetTdeWalletPassword"`
	// The date and time the pluggable database was created.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
}

// NewPluggableDatabasesLocalClone registers a new resource with the given unique name, arguments, and options.
func NewPluggableDatabasesLocalClone(ctx *pulumi.Context,
	name string, args *PluggableDatabasesLocalCloneArgs, opts ...pulumi.ResourceOption) (*PluggableDatabasesLocalClone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClonedPdbName == nil {
		return nil, errors.New("invalid value for required argument 'ClonedPdbName'")
	}
	if args.PdbAdminPassword == nil {
		return nil, errors.New("invalid value for required argument 'PdbAdminPassword'")
	}
	if args.PluggableDatabaseId == nil {
		return nil, errors.New("invalid value for required argument 'PluggableDatabaseId'")
	}
	if args.TargetTdeWalletPassword == nil {
		return nil, errors.New("invalid value for required argument 'TargetTdeWalletPassword'")
	}
	var resource PluggableDatabasesLocalClone
	err := ctx.RegisterResource("oci:database/pluggableDatabasesLocalClone:PluggableDatabasesLocalClone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPluggableDatabasesLocalClone gets an existing PluggableDatabasesLocalClone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPluggableDatabasesLocalClone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PluggableDatabasesLocalCloneState, opts ...pulumi.ResourceOption) (*PluggableDatabasesLocalClone, error) {
	var resource PluggableDatabasesLocalClone
	err := ctx.ReadResource("oci:database/pluggableDatabasesLocalClone:PluggableDatabasesLocalClone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PluggableDatabasesLocalClone resources.
type pluggableDatabasesLocalCloneState struct {
	// The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
	ClonedPdbName *string `pulumi:"clonedPdbName"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `pulumi:"compartmentId"`
	// Connection strings to connect to an Oracle Pluggable Database.
	ConnectionStrings *PluggableDatabasesLocalCloneConnectionStrings `pulumi:"connectionStrings"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CDB.
	ContainerDatabaseId *string `pulumi:"containerDatabaseId"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The restricted mode of the pluggable database. If a pluggable database is opened in restricted mode, the user needs both create a session and have restricted session privileges to connect to it.
	IsRestricted *bool `pulumi:"isRestricted"`
	// Detailed message for the lifecycle state.
	LifecycleDetails *string `pulumi:"lifecycleDetails"`
	// The mode that pluggable database is in. Open mode can only be changed to READ_ONLY or MIGRATE directly from the backend (within the Oracle Database software).
	OpenMode *string `pulumi:"openMode"`
	// A strong password for PDB Admin of the newly cloned PDB. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
	PdbAdminPassword *string `pulumi:"pdbAdminPassword"`
	// The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
	PdbName *string `pulumi:"pdbName"`
	// The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	PluggableDatabaseId *string `pulumi:"pluggableDatabaseId"`
	// The current state of the pluggable database.
	State *string `pulumi:"state"`
	// The existing TDE wallet password of the target CDB.
	TargetTdeWalletPassword *string `pulumi:"targetTdeWalletPassword"`
	// The date and time the pluggable database was created.
	TimeCreated *string `pulumi:"timeCreated"`
}

type PluggableDatabasesLocalCloneState struct {
	// The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
	ClonedPdbName pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId pulumi.StringPtrInput
	// Connection strings to connect to an Oracle Pluggable Database.
	ConnectionStrings PluggableDatabasesLocalCloneConnectionStringsPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CDB.
	ContainerDatabaseId pulumi.StringPtrInput
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags pulumi.MapInput
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// The restricted mode of the pluggable database. If a pluggable database is opened in restricted mode, the user needs both create a session and have restricted session privileges to connect to it.
	IsRestricted pulumi.BoolPtrInput
	// Detailed message for the lifecycle state.
	LifecycleDetails pulumi.StringPtrInput
	// The mode that pluggable database is in. Open mode can only be changed to READ_ONLY or MIGRATE directly from the backend (within the Oracle Database software).
	OpenMode pulumi.StringPtrInput
	// A strong password for PDB Admin of the newly cloned PDB. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
	PdbAdminPassword pulumi.StringPtrInput
	// The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
	PdbName pulumi.StringPtrInput
	// The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	PluggableDatabaseId pulumi.StringPtrInput
	// The current state of the pluggable database.
	State pulumi.StringPtrInput
	// The existing TDE wallet password of the target CDB.
	TargetTdeWalletPassword pulumi.StringPtrInput
	// The date and time the pluggable database was created.
	TimeCreated pulumi.StringPtrInput
}

func (PluggableDatabasesLocalCloneState) ElementType() reflect.Type {
	return reflect.TypeOf((*pluggableDatabasesLocalCloneState)(nil)).Elem()
}

type pluggableDatabasesLocalCloneArgs struct {
	// The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
	ClonedPdbName string `pulumi:"clonedPdbName"`
	// A strong password for PDB Admin of the newly cloned PDB. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
	PdbAdminPassword string `pulumi:"pdbAdminPassword"`
	// The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	PluggableDatabaseId string `pulumi:"pluggableDatabaseId"`
	// The existing TDE wallet password of the target CDB.
	TargetTdeWalletPassword string `pulumi:"targetTdeWalletPassword"`
}

// The set of arguments for constructing a PluggableDatabasesLocalClone resource.
type PluggableDatabasesLocalCloneArgs struct {
	// The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
	ClonedPdbName pulumi.StringInput
	// A strong password for PDB Admin of the newly cloned PDB. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
	PdbAdminPassword pulumi.StringInput
	// The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	PluggableDatabaseId pulumi.StringInput
	// The existing TDE wallet password of the target CDB.
	TargetTdeWalletPassword pulumi.StringInput
}

func (PluggableDatabasesLocalCloneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pluggableDatabasesLocalCloneArgs)(nil)).Elem()
}

type PluggableDatabasesLocalCloneInput interface {
	pulumi.Input

	ToPluggableDatabasesLocalCloneOutput() PluggableDatabasesLocalCloneOutput
	ToPluggableDatabasesLocalCloneOutputWithContext(ctx context.Context) PluggableDatabasesLocalCloneOutput
}

func (*PluggableDatabasesLocalClone) ElementType() reflect.Type {
	return reflect.TypeOf((*PluggableDatabasesLocalClone)(nil))
}

func (i *PluggableDatabasesLocalClone) ToPluggableDatabasesLocalCloneOutput() PluggableDatabasesLocalCloneOutput {
	return i.ToPluggableDatabasesLocalCloneOutputWithContext(context.Background())
}

func (i *PluggableDatabasesLocalClone) ToPluggableDatabasesLocalCloneOutputWithContext(ctx context.Context) PluggableDatabasesLocalCloneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PluggableDatabasesLocalCloneOutput)
}

func (i *PluggableDatabasesLocalClone) ToPluggableDatabasesLocalClonePtrOutput() PluggableDatabasesLocalClonePtrOutput {
	return i.ToPluggableDatabasesLocalClonePtrOutputWithContext(context.Background())
}

func (i *PluggableDatabasesLocalClone) ToPluggableDatabasesLocalClonePtrOutputWithContext(ctx context.Context) PluggableDatabasesLocalClonePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PluggableDatabasesLocalClonePtrOutput)
}

type PluggableDatabasesLocalClonePtrInput interface {
	pulumi.Input

	ToPluggableDatabasesLocalClonePtrOutput() PluggableDatabasesLocalClonePtrOutput
	ToPluggableDatabasesLocalClonePtrOutputWithContext(ctx context.Context) PluggableDatabasesLocalClonePtrOutput
}

type pluggableDatabasesLocalClonePtrType PluggableDatabasesLocalCloneArgs

func (*pluggableDatabasesLocalClonePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PluggableDatabasesLocalClone)(nil))
}

func (i *pluggableDatabasesLocalClonePtrType) ToPluggableDatabasesLocalClonePtrOutput() PluggableDatabasesLocalClonePtrOutput {
	return i.ToPluggableDatabasesLocalClonePtrOutputWithContext(context.Background())
}

func (i *pluggableDatabasesLocalClonePtrType) ToPluggableDatabasesLocalClonePtrOutputWithContext(ctx context.Context) PluggableDatabasesLocalClonePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PluggableDatabasesLocalClonePtrOutput)
}

// PluggableDatabasesLocalCloneArrayInput is an input type that accepts PluggableDatabasesLocalCloneArray and PluggableDatabasesLocalCloneArrayOutput values.
// You can construct a concrete instance of `PluggableDatabasesLocalCloneArrayInput` via:
//
//          PluggableDatabasesLocalCloneArray{ PluggableDatabasesLocalCloneArgs{...} }
type PluggableDatabasesLocalCloneArrayInput interface {
	pulumi.Input

	ToPluggableDatabasesLocalCloneArrayOutput() PluggableDatabasesLocalCloneArrayOutput
	ToPluggableDatabasesLocalCloneArrayOutputWithContext(context.Context) PluggableDatabasesLocalCloneArrayOutput
}

type PluggableDatabasesLocalCloneArray []PluggableDatabasesLocalCloneInput

func (PluggableDatabasesLocalCloneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PluggableDatabasesLocalClone)(nil)).Elem()
}

func (i PluggableDatabasesLocalCloneArray) ToPluggableDatabasesLocalCloneArrayOutput() PluggableDatabasesLocalCloneArrayOutput {
	return i.ToPluggableDatabasesLocalCloneArrayOutputWithContext(context.Background())
}

func (i PluggableDatabasesLocalCloneArray) ToPluggableDatabasesLocalCloneArrayOutputWithContext(ctx context.Context) PluggableDatabasesLocalCloneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PluggableDatabasesLocalCloneArrayOutput)
}

// PluggableDatabasesLocalCloneMapInput is an input type that accepts PluggableDatabasesLocalCloneMap and PluggableDatabasesLocalCloneMapOutput values.
// You can construct a concrete instance of `PluggableDatabasesLocalCloneMapInput` via:
//
//          PluggableDatabasesLocalCloneMap{ "key": PluggableDatabasesLocalCloneArgs{...} }
type PluggableDatabasesLocalCloneMapInput interface {
	pulumi.Input

	ToPluggableDatabasesLocalCloneMapOutput() PluggableDatabasesLocalCloneMapOutput
	ToPluggableDatabasesLocalCloneMapOutputWithContext(context.Context) PluggableDatabasesLocalCloneMapOutput
}

type PluggableDatabasesLocalCloneMap map[string]PluggableDatabasesLocalCloneInput

func (PluggableDatabasesLocalCloneMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PluggableDatabasesLocalClone)(nil)).Elem()
}

func (i PluggableDatabasesLocalCloneMap) ToPluggableDatabasesLocalCloneMapOutput() PluggableDatabasesLocalCloneMapOutput {
	return i.ToPluggableDatabasesLocalCloneMapOutputWithContext(context.Background())
}

func (i PluggableDatabasesLocalCloneMap) ToPluggableDatabasesLocalCloneMapOutputWithContext(ctx context.Context) PluggableDatabasesLocalCloneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PluggableDatabasesLocalCloneMapOutput)
}

type PluggableDatabasesLocalCloneOutput struct {
	*pulumi.OutputState
}

func (PluggableDatabasesLocalCloneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PluggableDatabasesLocalClone)(nil))
}

func (o PluggableDatabasesLocalCloneOutput) ToPluggableDatabasesLocalCloneOutput() PluggableDatabasesLocalCloneOutput {
	return o
}

func (o PluggableDatabasesLocalCloneOutput) ToPluggableDatabasesLocalCloneOutputWithContext(ctx context.Context) PluggableDatabasesLocalCloneOutput {
	return o
}

func (o PluggableDatabasesLocalCloneOutput) ToPluggableDatabasesLocalClonePtrOutput() PluggableDatabasesLocalClonePtrOutput {
	return o.ToPluggableDatabasesLocalClonePtrOutputWithContext(context.Background())
}

func (o PluggableDatabasesLocalCloneOutput) ToPluggableDatabasesLocalClonePtrOutputWithContext(ctx context.Context) PluggableDatabasesLocalClonePtrOutput {
	return o.ApplyT(func(v PluggableDatabasesLocalClone) *PluggableDatabasesLocalClone {
		return &v
	}).(PluggableDatabasesLocalClonePtrOutput)
}

type PluggableDatabasesLocalClonePtrOutput struct {
	*pulumi.OutputState
}

func (PluggableDatabasesLocalClonePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PluggableDatabasesLocalClone)(nil))
}

func (o PluggableDatabasesLocalClonePtrOutput) ToPluggableDatabasesLocalClonePtrOutput() PluggableDatabasesLocalClonePtrOutput {
	return o
}

func (o PluggableDatabasesLocalClonePtrOutput) ToPluggableDatabasesLocalClonePtrOutputWithContext(ctx context.Context) PluggableDatabasesLocalClonePtrOutput {
	return o
}

type PluggableDatabasesLocalCloneArrayOutput struct{ *pulumi.OutputState }

func (PluggableDatabasesLocalCloneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PluggableDatabasesLocalClone)(nil))
}

func (o PluggableDatabasesLocalCloneArrayOutput) ToPluggableDatabasesLocalCloneArrayOutput() PluggableDatabasesLocalCloneArrayOutput {
	return o
}

func (o PluggableDatabasesLocalCloneArrayOutput) ToPluggableDatabasesLocalCloneArrayOutputWithContext(ctx context.Context) PluggableDatabasesLocalCloneArrayOutput {
	return o
}

func (o PluggableDatabasesLocalCloneArrayOutput) Index(i pulumi.IntInput) PluggableDatabasesLocalCloneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PluggableDatabasesLocalClone {
		return vs[0].([]PluggableDatabasesLocalClone)[vs[1].(int)]
	}).(PluggableDatabasesLocalCloneOutput)
}

type PluggableDatabasesLocalCloneMapOutput struct{ *pulumi.OutputState }

func (PluggableDatabasesLocalCloneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PluggableDatabasesLocalClone)(nil))
}

func (o PluggableDatabasesLocalCloneMapOutput) ToPluggableDatabasesLocalCloneMapOutput() PluggableDatabasesLocalCloneMapOutput {
	return o
}

func (o PluggableDatabasesLocalCloneMapOutput) ToPluggableDatabasesLocalCloneMapOutputWithContext(ctx context.Context) PluggableDatabasesLocalCloneMapOutput {
	return o
}

func (o PluggableDatabasesLocalCloneMapOutput) MapIndex(k pulumi.StringInput) PluggableDatabasesLocalCloneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PluggableDatabasesLocalClone {
		return vs[0].(map[string]PluggableDatabasesLocalClone)[vs[1].(string)]
	}).(PluggableDatabasesLocalCloneOutput)
}

func init() {
	pulumi.RegisterOutputType(PluggableDatabasesLocalCloneOutput{})
	pulumi.RegisterOutputType(PluggableDatabasesLocalClonePtrOutput{})
	pulumi.RegisterOutputType(PluggableDatabasesLocalCloneArrayOutput{})
	pulumi.RegisterOutputType(PluggableDatabasesLocalCloneMapOutput{})
}