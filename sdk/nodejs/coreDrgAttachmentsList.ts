// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This resource provides the Drg Attachments List resource in Oracle Cloud Infrastructure Core service.
 *
 * Returns a complete list of DRG attachments that belong to a particular DRG.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testDrgAttachmentsList = new oci.CoreDrgAttachmentsList("testDrgAttachmentsList", {
 *     drgId: oci_core_drg.test_drg.id,
 *     attachmentType: _var.drg_attachments_list_attachment_type,
 *     isCrossTenancy: _var.drg_attachments_list_is_cross_tenancy,
 * });
 * ```
 *
 * ## Import
 *
 * Import is not supported for this resource.
 */
export class CoreDrgAttachmentsList extends pulumi.CustomResource {
    /**
     * Get an existing CoreDrgAttachmentsList resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CoreDrgAttachmentsListState, opts?: pulumi.CustomResourceOptions): CoreDrgAttachmentsList {
        return new CoreDrgAttachmentsList(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'oci:index/coreDrgAttachmentsList:CoreDrgAttachmentsList';

    /**
     * Returns true if the given object is an instance of CoreDrgAttachmentsList.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CoreDrgAttachmentsList {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CoreDrgAttachmentsList.__pulumiType;
    }

    /**
     * The type for the network resource attached to the DRG.
     */
    public readonly attachmentType!: pulumi.Output<string | undefined>;
    /**
     * The list of drg_attachments.
     */
    public /*out*/ readonly drgAllAttachments!: pulumi.Output<outputs.CoreDrgAttachmentsListDrgAllAttachment[]>;
    /**
     * The [[OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
     */
    public readonly drgId!: pulumi.Output<string>;
    /**
     * Whether the DRG attachment lives in a different tenancy than the DRG.
     */
    public readonly isCrossTenancy!: pulumi.Output<boolean | undefined>;

    /**
     * Create a CoreDrgAttachmentsList resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CoreDrgAttachmentsListArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CoreDrgAttachmentsListArgs | CoreDrgAttachmentsListState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CoreDrgAttachmentsListState | undefined;
            inputs["attachmentType"] = state ? state.attachmentType : undefined;
            inputs["drgAllAttachments"] = state ? state.drgAllAttachments : undefined;
            inputs["drgId"] = state ? state.drgId : undefined;
            inputs["isCrossTenancy"] = state ? state.isCrossTenancy : undefined;
        } else {
            const args = argsOrState as CoreDrgAttachmentsListArgs | undefined;
            if ((!args || args.drgId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'drgId'");
            }
            inputs["attachmentType"] = args ? args.attachmentType : undefined;
            inputs["drgId"] = args ? args.drgId : undefined;
            inputs["isCrossTenancy"] = args ? args.isCrossTenancy : undefined;
            inputs["drgAllAttachments"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CoreDrgAttachmentsList.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CoreDrgAttachmentsList resources.
 */
export interface CoreDrgAttachmentsListState {
    /**
     * The type for the network resource attached to the DRG.
     */
    attachmentType?: pulumi.Input<string>;
    /**
     * The list of drg_attachments.
     */
    drgAllAttachments?: pulumi.Input<pulumi.Input<inputs.CoreDrgAttachmentsListDrgAllAttachment>[]>;
    /**
     * The [[OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
     */
    drgId?: pulumi.Input<string>;
    /**
     * Whether the DRG attachment lives in a different tenancy than the DRG.
     */
    isCrossTenancy?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a CoreDrgAttachmentsList resource.
 */
export interface CoreDrgAttachmentsListArgs {
    /**
     * The type for the network resource attached to the DRG.
     */
    attachmentType?: pulumi.Input<string>;
    /**
     * The [[OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
     */
    drgId: pulumi.Input<string>;
    /**
     * Whether the DRG attachment lives in a different tenancy than the DRG.
     */
    isCrossTenancy?: pulumi.Input<boolean>;
}