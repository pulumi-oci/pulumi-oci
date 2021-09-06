// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Import is not supported for this resource.
 */
export class AutonomousDatabaseWallet extends pulumi.CustomResource {
    /**
     * Get an existing AutonomousDatabaseWallet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AutonomousDatabaseWalletState, opts?: pulumi.CustomResourceOptions): AutonomousDatabaseWallet {
        return new AutonomousDatabaseWallet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'oci:database/autonomousDatabaseWallet:AutonomousDatabaseWallet';

    /**
     * Returns true if the given object is an instance of AutonomousDatabaseWallet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AutonomousDatabaseWallet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AutonomousDatabaseWallet.__pulumiType;
    }

    /**
     * The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
     */
    public readonly autonomousDatabaseId!: pulumi.Output<string>;
    public readonly base64EncodeContent!: pulumi.Output<boolean | undefined>;
    /**
     * content of the downloaded zipped wallet for the Autonomous Database. If `base64EncodeContent` is set to `true`, then this content will be base64 encoded.
     */
    public /*out*/ readonly content!: pulumi.Output<string>;
    /**
     * The type of wallet to generate.
     */
    public readonly generateType!: pulumi.Output<string | undefined>;
    /**
     * The password to encrypt the keys inside the wallet. The password must be at least 8 characters long and must include at least 1 letter and either 1 numeric character or 1 special character.
     */
    public readonly password!: pulumi.Output<string>;

    /**
     * Create a AutonomousDatabaseWallet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AutonomousDatabaseWalletArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AutonomousDatabaseWalletArgs | AutonomousDatabaseWalletState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AutonomousDatabaseWalletState | undefined;
            inputs["autonomousDatabaseId"] = state ? state.autonomousDatabaseId : undefined;
            inputs["base64EncodeContent"] = state ? state.base64EncodeContent : undefined;
            inputs["content"] = state ? state.content : undefined;
            inputs["generateType"] = state ? state.generateType : undefined;
            inputs["password"] = state ? state.password : undefined;
        } else {
            const args = argsOrState as AutonomousDatabaseWalletArgs | undefined;
            if ((!args || args.autonomousDatabaseId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autonomousDatabaseId'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            inputs["autonomousDatabaseId"] = args ? args.autonomousDatabaseId : undefined;
            inputs["base64EncodeContent"] = args ? args.base64EncodeContent : undefined;
            inputs["generateType"] = args ? args.generateType : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["content"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AutonomousDatabaseWallet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AutonomousDatabaseWallet resources.
 */
export interface AutonomousDatabaseWalletState {
    /**
     * The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
     */
    autonomousDatabaseId?: pulumi.Input<string>;
    base64EncodeContent?: pulumi.Input<boolean>;
    /**
     * content of the downloaded zipped wallet for the Autonomous Database. If `base64EncodeContent` is set to `true`, then this content will be base64 encoded.
     */
    content?: pulumi.Input<string>;
    /**
     * The type of wallet to generate.
     */
    generateType?: pulumi.Input<string>;
    /**
     * The password to encrypt the keys inside the wallet. The password must be at least 8 characters long and must include at least 1 letter and either 1 numeric character or 1 special character.
     */
    password?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AutonomousDatabaseWallet resource.
 */
export interface AutonomousDatabaseWalletArgs {
    /**
     * The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
     */
    autonomousDatabaseId: pulumi.Input<string>;
    base64EncodeContent?: pulumi.Input<boolean>;
    /**
     * The type of wallet to generate.
     */
    generateType?: pulumi.Input<string>;
    /**
     * The password to encrypt the keys inside the wallet. The password must be at least 8 characters long and must include at least 1 letter and either 1 numeric character or 1 special character.
     */
    password: pulumi.Input<string>;
}