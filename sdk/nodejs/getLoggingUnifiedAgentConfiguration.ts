// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides details about a specific Unified Agent Configuration resource in Oracle Cloud Infrastructure Logging service.
 *
 * Get the unified agent configuration for an ID.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testUnifiedAgentConfiguration = oci.GetLoggingUnifiedAgentConfiguration({
 *     unifiedAgentConfigurationId: oci_logging_unified_agent_configuration.test_unified_agent_configuration.id,
 * });
 * ```
 */
export function getLoggingUnifiedAgentConfiguration(args: GetLoggingUnifiedAgentConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetLoggingUnifiedAgentConfigurationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getLoggingUnifiedAgentConfiguration:GetLoggingUnifiedAgentConfiguration", {
        "unifiedAgentConfigurationId": args.unifiedAgentConfigurationId,
    }, opts);
}

/**
 * A collection of arguments for invoking GetLoggingUnifiedAgentConfiguration.
 */
export interface GetLoggingUnifiedAgentConfigurationArgs {
    /**
     * The OCID of the Unified Agent configuration.
     */
    unifiedAgentConfigurationId: string;
}

/**
 * A collection of values returned by GetLoggingUnifiedAgentConfiguration.
 */
export interface GetLoggingUnifiedAgentConfigurationResult {
    /**
     * The OCID of the compartment that the resource belongs to.
     */
    readonly compartmentId: string;
    /**
     * State of unified agent service configuration.
     */
    readonly configurationState: string;
    /**
     * Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
     */
    readonly definedTags: {[key: string]: any};
    /**
     * Description for this resource.
     */
    readonly description: string;
    /**
     * The user-friendly display name. This must be unique within the enclosing resource, and it's changeable. Avoid entering confidential information.
     */
    readonly displayName: string;
    /**
     * Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
     */
    readonly freeformTags: {[key: string]: any};
    /**
     * Groups using the configuration.
     */
    readonly groupAssociation: outputs.GetLoggingUnifiedAgentConfigurationGroupAssociation;
    /**
     * The OCID of the resource.
     */
    readonly id: string;
    /**
     * Whether or not this resource is currently enabled.
     */
    readonly isEnabled: boolean;
    /**
     * Top level Unified Agent service configuration object.
     */
    readonly serviceConfiguration: outputs.GetLoggingUnifiedAgentConfigurationServiceConfiguration;
    /**
     * The pipeline state.
     */
    readonly state: string;
    /**
     * Time the resource was created.
     */
    readonly timeCreated: string;
    /**
     * Time the resource was last modified.
     */
    readonly timeLastModified: string;
    readonly unifiedAgentConfigurationId: string;
}