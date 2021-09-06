// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides details about a specific Vault Usage resource in Oracle Cloud Infrastructure Kms service.
 *
 * Gets the count of keys and key versions in the specified vault to calculate usage against service limits.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testVaultUsage = oci.kms.getVaultUsage({
 *     vaultId: oci_kms_vault.test_vault.id,
 * });
 * ```
 */
export function getVaultUsage(args: GetVaultUsageArgs, opts?: pulumi.InvokeOptions): Promise<GetVaultUsageResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:kms/getVaultUsage:getVaultUsage", {
        "vaultId": args.vaultId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVaultUsage.
 */
export interface GetVaultUsageArgs {
    /**
     * The OCID of the vault.
     */
    vaultId: string;
}

/**
 * A collection of values returned by getVaultUsage.
 */
export interface GetVaultUsageResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The number of keys in this vault, across all compartments, excluding keys in a `DELETED` state.
     */
    readonly keyCount: number;
    /**
     * The number of key versions in this vault, across all compartments, excluding key versions in a `DELETED` state.
     */
    readonly keyVersionCount: number;
    /**
     * The number of keys in this vault that persist on the server, across all compartments, excluding keys in a `DELETED` state.
     */
    readonly softwareKeyCount: number;
    /**
     * The number of key versions in this vault that persist on the server, across all compartments, excluding key versions in a `DELETED` state.
     */
    readonly softwareKeyVersionCount: number;
    readonly vaultId: string;
}