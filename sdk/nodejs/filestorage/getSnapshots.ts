// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the list of Snapshots in Oracle Cloud Infrastructure File Storage service.
 *
 * Lists snapshots of the specified file system.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testSnapshots = oci.filestorage.getSnapshots({
 *     fileSystemId: oci_file_storage_file_system.test_file_system.id,
 *     id: _var.snapshot_id,
 *     state: _var.snapshot_state,
 * });
 * ```
 */
export function getSnapshots(args: GetSnapshotsArgs, opts?: pulumi.InvokeOptions): Promise<GetSnapshotsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:filestorage/getSnapshots:getSnapshots", {
        "fileSystemId": args.fileSystemId,
        "filters": args.filters,
        "id": args.id,
        "state": args.state,
    }, opts);
}

/**
 * A collection of arguments for invoking getSnapshots.
 */
export interface GetSnapshotsArgs {
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system.
     */
    fileSystemId: string;
    filters?: inputs.filestorage.GetSnapshotsFilter[];
    /**
     * Filter results by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resouce type.
     */
    id?: string;
    /**
     * Filter results by the specified lifecycle state. Must be a valid state for the resource type.
     */
    state?: string;
}

/**
 * A collection of values returned by getSnapshots.
 */
export interface GetSnapshotsResult {
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system from which the snapshot was created.
     */
    readonly fileSystemId: string;
    readonly filters?: outputs.filestorage.GetSnapshotsFilter[];
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the snapshot.
     */
    readonly id?: string;
    /**
     * The list of snapshots.
     */
    readonly snapshots: outputs.filestorage.GetSnapshotsSnapshot[];
    /**
     * The current state of the snapshot.
     */
    readonly state?: string;
}