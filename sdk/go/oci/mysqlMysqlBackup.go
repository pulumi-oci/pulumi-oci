// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Mysql Backup resource in Oracle Cloud Infrastructure MySQL Database service.
//
// Create a backup of a DB System.
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
// 		_, err := oci.NewMysqlMysqlBackup(ctx, "testMysqlBackup", &oci.MysqlMysqlBackupArgs{
// 			DbSystemId: pulumi.Any(oci_mysql_mysql_db_system.Test_db_system.Id),
// 			BackupType: pulumi.Any(_var.Mysql_backup_backup_type),
// 			DefinedTags: pulumi.AnyMap{
// 				"foo-namespace.bar-key": pulumi.Any("value"),
// 			},
// 			Description: pulumi.Any(_var.Mysql_backup_description),
// 			DisplayName: pulumi.Any(_var.Mysql_backup_display_name),
// 			FreeformTags: pulumi.AnyMap{
// 				"bar-key": pulumi.Any("value"),
// 			},
// 			RetentionInDays: pulumi.Any(_var.Mysql_backup_retention_in_days),
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
// MysqlBackups can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:index/mysqlMysqlBackup:MysqlMysqlBackup test_mysql_backup "id"
// ```
type MysqlMysqlBackup struct {
	pulumi.CustomResourceState

	// The size of the backup in base-2 (IEC) gibibytes. (GiB).
	BackupSizeInGbs pulumi.IntOutput `pulumi:"backupSizeInGbs"`
	// The type of backup.
	BackupType pulumi.StringOutput `pulumi:"backupType"`
	// (Updatable) The OCID of the compartment.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// If the backup was created automatically, or by a manual request.
	CreationType pulumi.StringOutput `pulumi:"creationType"`
	// Initial size of the data volume in GiBs that will be created and attached.
	DataStorageSizeInGb pulumi.IntOutput `pulumi:"dataStorageSizeInGb"`
	// The OCID of the DB System the Backup is associated with.
	DbSystemId pulumi.StringOutput `pulumi:"dbSystemId"`
	// Snapshot of the DbSystem details at the time of the backup
	DbSystemSnapshot MysqlMysqlBackupDbSystemSnapshotOutput `pulumi:"dbSystemSnapshot"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// (Updatable) A user-supplied description for the backup.
	Description pulumi.StringOutput `pulumi:"description"`
	// (Updatable) A user-supplied display name for the backup.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// Additional information about the current lifecycleState.
	LifecycleDetails pulumi.StringOutput `pulumi:"lifecycleDetails"`
	// The MySQL server version of the DB System used for backup.
	MysqlVersion pulumi.StringOutput `pulumi:"mysqlVersion"`
	// (Updatable) Number of days to retain this backup.
	RetentionInDays pulumi.IntOutput `pulumi:"retentionInDays"`
	// The shape of the DB System instance used for backup.
	ShapeName pulumi.StringOutput `pulumi:"shapeName"`
	// The state of the backup.
	State pulumi.StringOutput `pulumi:"state"`
	// The time the backup record was created.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// The time at which the backup was updated.
	TimeUpdated pulumi.StringOutput `pulumi:"timeUpdated"`
}

// NewMysqlMysqlBackup registers a new resource with the given unique name, arguments, and options.
func NewMysqlMysqlBackup(ctx *pulumi.Context,
	name string, args *MysqlMysqlBackupArgs, opts ...pulumi.ResourceOption) (*MysqlMysqlBackup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbSystemId == nil {
		return nil, errors.New("invalid value for required argument 'DbSystemId'")
	}
	var resource MysqlMysqlBackup
	err := ctx.RegisterResource("oci:index/mysqlMysqlBackup:MysqlMysqlBackup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMysqlMysqlBackup gets an existing MysqlMysqlBackup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMysqlMysqlBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MysqlMysqlBackupState, opts ...pulumi.ResourceOption) (*MysqlMysqlBackup, error) {
	var resource MysqlMysqlBackup
	err := ctx.ReadResource("oci:index/mysqlMysqlBackup:MysqlMysqlBackup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MysqlMysqlBackup resources.
type mysqlMysqlBackupState struct {
	// The size of the backup in base-2 (IEC) gibibytes. (GiB).
	BackupSizeInGbs *int `pulumi:"backupSizeInGbs"`
	// The type of backup.
	BackupType *string `pulumi:"backupType"`
	// (Updatable) The OCID of the compartment.
	CompartmentId *string `pulumi:"compartmentId"`
	// If the backup was created automatically, or by a manual request.
	CreationType *string `pulumi:"creationType"`
	// Initial size of the data volume in GiBs that will be created and attached.
	DataStorageSizeInGb *int `pulumi:"dataStorageSizeInGb"`
	// The OCID of the DB System the Backup is associated with.
	DbSystemId *string `pulumi:"dbSystemId"`
	// Snapshot of the DbSystem details at the time of the backup
	DbSystemSnapshot *MysqlMysqlBackupDbSystemSnapshot `pulumi:"dbSystemSnapshot"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A user-supplied description for the backup.
	Description *string `pulumi:"description"`
	// (Updatable) A user-supplied display name for the backup.
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// Additional information about the current lifecycleState.
	LifecycleDetails *string `pulumi:"lifecycleDetails"`
	// The MySQL server version of the DB System used for backup.
	MysqlVersion *string `pulumi:"mysqlVersion"`
	// (Updatable) Number of days to retain this backup.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// The shape of the DB System instance used for backup.
	ShapeName *string `pulumi:"shapeName"`
	// The state of the backup.
	State *string `pulumi:"state"`
	// The time the backup record was created.
	TimeCreated *string `pulumi:"timeCreated"`
	// The time at which the backup was updated.
	TimeUpdated *string `pulumi:"timeUpdated"`
}

type MysqlMysqlBackupState struct {
	// The size of the backup in base-2 (IEC) gibibytes. (GiB).
	BackupSizeInGbs pulumi.IntPtrInput
	// The type of backup.
	BackupType pulumi.StringPtrInput
	// (Updatable) The OCID of the compartment.
	CompartmentId pulumi.StringPtrInput
	// If the backup was created automatically, or by a manual request.
	CreationType pulumi.StringPtrInput
	// Initial size of the data volume in GiBs that will be created and attached.
	DataStorageSizeInGb pulumi.IntPtrInput
	// The OCID of the DB System the Backup is associated with.
	DbSystemId pulumi.StringPtrInput
	// Snapshot of the DbSystem details at the time of the backup
	DbSystemSnapshot MysqlMysqlBackupDbSystemSnapshotPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A user-supplied description for the backup.
	Description pulumi.StringPtrInput
	// (Updatable) A user-supplied display name for the backup.
	DisplayName pulumi.StringPtrInput
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags pulumi.MapInput
	// Additional information about the current lifecycleState.
	LifecycleDetails pulumi.StringPtrInput
	// The MySQL server version of the DB System used for backup.
	MysqlVersion pulumi.StringPtrInput
	// (Updatable) Number of days to retain this backup.
	RetentionInDays pulumi.IntPtrInput
	// The shape of the DB System instance used for backup.
	ShapeName pulumi.StringPtrInput
	// The state of the backup.
	State pulumi.StringPtrInput
	// The time the backup record was created.
	TimeCreated pulumi.StringPtrInput
	// The time at which the backup was updated.
	TimeUpdated pulumi.StringPtrInput
}

func (MysqlMysqlBackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*mysqlMysqlBackupState)(nil)).Elem()
}

type mysqlMysqlBackupArgs struct {
	// The type of backup.
	BackupType *string `pulumi:"backupType"`
	// (Updatable) The OCID of the compartment.
	CompartmentId *string `pulumi:"compartmentId"`
	// The OCID of the DB System the Backup is associated with.
	DbSystemId string `pulumi:"dbSystemId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A user-supplied description for the backup.
	Description *string `pulumi:"description"`
	// (Updatable) A user-supplied display name for the backup.
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable) Number of days to retain this backup.
	RetentionInDays *int `pulumi:"retentionInDays"`
}

// The set of arguments for constructing a MysqlMysqlBackup resource.
type MysqlMysqlBackupArgs struct {
	// The type of backup.
	BackupType pulumi.StringPtrInput
	// (Updatable) The OCID of the compartment.
	CompartmentId pulumi.StringPtrInput
	// The OCID of the DB System the Backup is associated with.
	DbSystemId pulumi.StringInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A user-supplied description for the backup.
	Description pulumi.StringPtrInput
	// (Updatable) A user-supplied display name for the backup.
	DisplayName pulumi.StringPtrInput
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags pulumi.MapInput
	// (Updatable) Number of days to retain this backup.
	RetentionInDays pulumi.IntPtrInput
}

func (MysqlMysqlBackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mysqlMysqlBackupArgs)(nil)).Elem()
}

type MysqlMysqlBackupInput interface {
	pulumi.Input

	ToMysqlMysqlBackupOutput() MysqlMysqlBackupOutput
	ToMysqlMysqlBackupOutputWithContext(ctx context.Context) MysqlMysqlBackupOutput
}

func (*MysqlMysqlBackup) ElementType() reflect.Type {
	return reflect.TypeOf((*MysqlMysqlBackup)(nil))
}

func (i *MysqlMysqlBackup) ToMysqlMysqlBackupOutput() MysqlMysqlBackupOutput {
	return i.ToMysqlMysqlBackupOutputWithContext(context.Background())
}

func (i *MysqlMysqlBackup) ToMysqlMysqlBackupOutputWithContext(ctx context.Context) MysqlMysqlBackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MysqlMysqlBackupOutput)
}

func (i *MysqlMysqlBackup) ToMysqlMysqlBackupPtrOutput() MysqlMysqlBackupPtrOutput {
	return i.ToMysqlMysqlBackupPtrOutputWithContext(context.Background())
}

func (i *MysqlMysqlBackup) ToMysqlMysqlBackupPtrOutputWithContext(ctx context.Context) MysqlMysqlBackupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MysqlMysqlBackupPtrOutput)
}

type MysqlMysqlBackupPtrInput interface {
	pulumi.Input

	ToMysqlMysqlBackupPtrOutput() MysqlMysqlBackupPtrOutput
	ToMysqlMysqlBackupPtrOutputWithContext(ctx context.Context) MysqlMysqlBackupPtrOutput
}

type mysqlMysqlBackupPtrType MysqlMysqlBackupArgs

func (*mysqlMysqlBackupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MysqlMysqlBackup)(nil))
}

func (i *mysqlMysqlBackupPtrType) ToMysqlMysqlBackupPtrOutput() MysqlMysqlBackupPtrOutput {
	return i.ToMysqlMysqlBackupPtrOutputWithContext(context.Background())
}

func (i *mysqlMysqlBackupPtrType) ToMysqlMysqlBackupPtrOutputWithContext(ctx context.Context) MysqlMysqlBackupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MysqlMysqlBackupPtrOutput)
}

// MysqlMysqlBackupArrayInput is an input type that accepts MysqlMysqlBackupArray and MysqlMysqlBackupArrayOutput values.
// You can construct a concrete instance of `MysqlMysqlBackupArrayInput` via:
//
//          MysqlMysqlBackupArray{ MysqlMysqlBackupArgs{...} }
type MysqlMysqlBackupArrayInput interface {
	pulumi.Input

	ToMysqlMysqlBackupArrayOutput() MysqlMysqlBackupArrayOutput
	ToMysqlMysqlBackupArrayOutputWithContext(context.Context) MysqlMysqlBackupArrayOutput
}

type MysqlMysqlBackupArray []MysqlMysqlBackupInput

func (MysqlMysqlBackupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MysqlMysqlBackup)(nil)).Elem()
}

func (i MysqlMysqlBackupArray) ToMysqlMysqlBackupArrayOutput() MysqlMysqlBackupArrayOutput {
	return i.ToMysqlMysqlBackupArrayOutputWithContext(context.Background())
}

func (i MysqlMysqlBackupArray) ToMysqlMysqlBackupArrayOutputWithContext(ctx context.Context) MysqlMysqlBackupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MysqlMysqlBackupArrayOutput)
}

// MysqlMysqlBackupMapInput is an input type that accepts MysqlMysqlBackupMap and MysqlMysqlBackupMapOutput values.
// You can construct a concrete instance of `MysqlMysqlBackupMapInput` via:
//
//          MysqlMysqlBackupMap{ "key": MysqlMysqlBackupArgs{...} }
type MysqlMysqlBackupMapInput interface {
	pulumi.Input

	ToMysqlMysqlBackupMapOutput() MysqlMysqlBackupMapOutput
	ToMysqlMysqlBackupMapOutputWithContext(context.Context) MysqlMysqlBackupMapOutput
}

type MysqlMysqlBackupMap map[string]MysqlMysqlBackupInput

func (MysqlMysqlBackupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MysqlMysqlBackup)(nil)).Elem()
}

func (i MysqlMysqlBackupMap) ToMysqlMysqlBackupMapOutput() MysqlMysqlBackupMapOutput {
	return i.ToMysqlMysqlBackupMapOutputWithContext(context.Background())
}

func (i MysqlMysqlBackupMap) ToMysqlMysqlBackupMapOutputWithContext(ctx context.Context) MysqlMysqlBackupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MysqlMysqlBackupMapOutput)
}

type MysqlMysqlBackupOutput struct {
	*pulumi.OutputState
}

func (MysqlMysqlBackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MysqlMysqlBackup)(nil))
}

func (o MysqlMysqlBackupOutput) ToMysqlMysqlBackupOutput() MysqlMysqlBackupOutput {
	return o
}

func (o MysqlMysqlBackupOutput) ToMysqlMysqlBackupOutputWithContext(ctx context.Context) MysqlMysqlBackupOutput {
	return o
}

func (o MysqlMysqlBackupOutput) ToMysqlMysqlBackupPtrOutput() MysqlMysqlBackupPtrOutput {
	return o.ToMysqlMysqlBackupPtrOutputWithContext(context.Background())
}

func (o MysqlMysqlBackupOutput) ToMysqlMysqlBackupPtrOutputWithContext(ctx context.Context) MysqlMysqlBackupPtrOutput {
	return o.ApplyT(func(v MysqlMysqlBackup) *MysqlMysqlBackup {
		return &v
	}).(MysqlMysqlBackupPtrOutput)
}

type MysqlMysqlBackupPtrOutput struct {
	*pulumi.OutputState
}

func (MysqlMysqlBackupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MysqlMysqlBackup)(nil))
}

func (o MysqlMysqlBackupPtrOutput) ToMysqlMysqlBackupPtrOutput() MysqlMysqlBackupPtrOutput {
	return o
}

func (o MysqlMysqlBackupPtrOutput) ToMysqlMysqlBackupPtrOutputWithContext(ctx context.Context) MysqlMysqlBackupPtrOutput {
	return o
}

type MysqlMysqlBackupArrayOutput struct{ *pulumi.OutputState }

func (MysqlMysqlBackupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MysqlMysqlBackup)(nil))
}

func (o MysqlMysqlBackupArrayOutput) ToMysqlMysqlBackupArrayOutput() MysqlMysqlBackupArrayOutput {
	return o
}

func (o MysqlMysqlBackupArrayOutput) ToMysqlMysqlBackupArrayOutputWithContext(ctx context.Context) MysqlMysqlBackupArrayOutput {
	return o
}

func (o MysqlMysqlBackupArrayOutput) Index(i pulumi.IntInput) MysqlMysqlBackupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MysqlMysqlBackup {
		return vs[0].([]MysqlMysqlBackup)[vs[1].(int)]
	}).(MysqlMysqlBackupOutput)
}

type MysqlMysqlBackupMapOutput struct{ *pulumi.OutputState }

func (MysqlMysqlBackupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MysqlMysqlBackup)(nil))
}

func (o MysqlMysqlBackupMapOutput) ToMysqlMysqlBackupMapOutput() MysqlMysqlBackupMapOutput {
	return o
}

func (o MysqlMysqlBackupMapOutput) ToMysqlMysqlBackupMapOutputWithContext(ctx context.Context) MysqlMysqlBackupMapOutput {
	return o
}

func (o MysqlMysqlBackupMapOutput) MapIndex(k pulumi.StringInput) MysqlMysqlBackupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MysqlMysqlBackup {
		return vs[0].(map[string]MysqlMysqlBackup)[vs[1].(string)]
	}).(MysqlMysqlBackupOutput)
}

func init() {
	pulumi.RegisterOutputType(MysqlMysqlBackupOutput{})
	pulumi.RegisterOutputType(MysqlMysqlBackupPtrOutput{})
	pulumi.RegisterOutputType(MysqlMysqlBackupArrayOutput{})
	pulumi.RegisterOutputType(MysqlMysqlBackupMapOutput{})
}