// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ObjectstorageNamespaceMetadata extends pulumi.CustomResource {
    /**
     * Get an existing ObjectstorageNamespaceMetadata resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ObjectstorageNamespaceMetadataState, opts?: pulumi.CustomResourceOptions): ObjectstorageNamespaceMetadata {
        return new ObjectstorageNamespaceMetadata(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'oci:index/objectstorageNamespaceMetadata:ObjectstorageNamespaceMetadata';

    /**
     * Returns true if the given object is an instance of ObjectstorageNamespaceMetadata.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ObjectstorageNamespaceMetadata {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ObjectstorageNamespaceMetadata.__pulumiType;
    }

    public readonly defaultS3compartmentId!: pulumi.Output<string>;
    public readonly defaultSwiftCompartmentId!: pulumi.Output<string>;
    public readonly namespace!: pulumi.Output<string>;

    /**
     * Create a ObjectstorageNamespaceMetadata resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ObjectstorageNamespaceMetadataArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ObjectstorageNamespaceMetadataArgs | ObjectstorageNamespaceMetadataState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ObjectstorageNamespaceMetadataState | undefined;
            inputs["defaultS3compartmentId"] = state ? state.defaultS3compartmentId : undefined;
            inputs["defaultSwiftCompartmentId"] = state ? state.defaultSwiftCompartmentId : undefined;
            inputs["namespace"] = state ? state.namespace : undefined;
        } else {
            const args = argsOrState as ObjectstorageNamespaceMetadataArgs | undefined;
            if ((!args || args.namespace === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespace'");
            }
            inputs["defaultS3compartmentId"] = args ? args.defaultS3compartmentId : undefined;
            inputs["defaultSwiftCompartmentId"] = args ? args.defaultSwiftCompartmentId : undefined;
            inputs["namespace"] = args ? args.namespace : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ObjectstorageNamespaceMetadata.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ObjectstorageNamespaceMetadata resources.
 */
export interface ObjectstorageNamespaceMetadataState {
    defaultS3compartmentId?: pulumi.Input<string>;
    defaultSwiftCompartmentId?: pulumi.Input<string>;
    namespace?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ObjectstorageNamespaceMetadata resource.
 */
export interface ObjectstorageNamespaceMetadataArgs {
    defaultS3compartmentId?: pulumi.Input<string>;
    defaultSwiftCompartmentId?: pulumi.Input<string>;
    namespace: pulumi.Input<string>;
}