// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the list of Db Node Console Connections in Oracle Cloud Infrastructure Database service.
 *
 * Lists the console connections for the specified database node.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testDbNodeConsoleConnections = oci.GetDatabaseDbNodeConsoleConnections({
 *     dbNodeId: oci_database_db_node.test_db_node.id,
 * });
 * ```
 */
export function getDatabaseDbNodeConsoleConnections(args: GetDatabaseDbNodeConsoleConnectionsArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseDbNodeConsoleConnectionsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getDatabaseDbNodeConsoleConnections:GetDatabaseDbNodeConsoleConnections", {
        "dbNodeId": args.dbNodeId,
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking GetDatabaseDbNodeConsoleConnections.
 */
export interface GetDatabaseDbNodeConsoleConnectionsArgs {
    /**
     * The database node [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
     */
    dbNodeId: string;
    filters?: inputs.GetDatabaseDbNodeConsoleConnectionsFilter[];
}

/**
 * A collection of values returned by GetDatabaseDbNodeConsoleConnections.
 */
export interface GetDatabaseDbNodeConsoleConnectionsResult {
    /**
     * The list of console_connections.
     */
    readonly consoleConnections: outputs.GetDatabaseDbNodeConsoleConnectionsConsoleConnection[];
    /**
     * The OCID of the database node.
     */
    readonly dbNodeId: string;
    readonly filters?: outputs.GetDatabaseDbNodeConsoleConnectionsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}