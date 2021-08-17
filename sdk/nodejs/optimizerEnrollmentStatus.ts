// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource provides the Enrollment Status resource in Oracle Cloud Infrastructure Optimizer service.
 *
 * Updates the enrollment status of the tenancy.
 *
 * ## Import
 *
 * EnrollmentStatus can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import oci:index/optimizerEnrollmentStatus:OptimizerEnrollmentStatus test_enrollment_status "id"
 * ```
 */
export class OptimizerEnrollmentStatus extends pulumi.CustomResource {
    /**
     * Get an existing OptimizerEnrollmentStatus resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OptimizerEnrollmentStatusState, opts?: pulumi.CustomResourceOptions): OptimizerEnrollmentStatus {
        return new OptimizerEnrollmentStatus(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'oci:index/optimizerEnrollmentStatus:OptimizerEnrollmentStatus';

    /**
     * Returns true if the given object is an instance of OptimizerEnrollmentStatus.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OptimizerEnrollmentStatus {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OptimizerEnrollmentStatus.__pulumiType;
    }

    /**
     * The OCID of the compartment.
     */
    public /*out*/ readonly compartmentId!: pulumi.Output<string>;
    /**
     * The unique OCID associated with the enrollment status.
     */
    public readonly enrollmentStatusId!: pulumi.Output<string>;
    /**
     * The enrollment status' current state.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * (Updatable) The Cloud Advisor enrollment status.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * The reason for the enrollment status of the tenancy.
     */
    public /*out*/ readonly statusReason!: pulumi.Output<string>;
    /**
     * The date and time the enrollment status was created, in the format defined by RFC3339.
     */
    public /*out*/ readonly timeCreated!: pulumi.Output<string>;
    /**
     * The date and time the enrollment status was last updated, in the format defined by RFC3339.
     */
    public /*out*/ readonly timeUpdated!: pulumi.Output<string>;

    /**
     * Create a OptimizerEnrollmentStatus resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OptimizerEnrollmentStatusArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OptimizerEnrollmentStatusArgs | OptimizerEnrollmentStatusState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OptimizerEnrollmentStatusState | undefined;
            inputs["compartmentId"] = state ? state.compartmentId : undefined;
            inputs["enrollmentStatusId"] = state ? state.enrollmentStatusId : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["statusReason"] = state ? state.statusReason : undefined;
            inputs["timeCreated"] = state ? state.timeCreated : undefined;
            inputs["timeUpdated"] = state ? state.timeUpdated : undefined;
        } else {
            const args = argsOrState as OptimizerEnrollmentStatusArgs | undefined;
            if ((!args || args.enrollmentStatusId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enrollmentStatusId'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            inputs["enrollmentStatusId"] = args ? args.enrollmentStatusId : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["compartmentId"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["statusReason"] = undefined /*out*/;
            inputs["timeCreated"] = undefined /*out*/;
            inputs["timeUpdated"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(OptimizerEnrollmentStatus.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OptimizerEnrollmentStatus resources.
 */
export interface OptimizerEnrollmentStatusState {
    /**
     * The OCID of the compartment.
     */
    compartmentId?: pulumi.Input<string>;
    /**
     * The unique OCID associated with the enrollment status.
     */
    enrollmentStatusId?: pulumi.Input<string>;
    /**
     * The enrollment status' current state.
     */
    state?: pulumi.Input<string>;
    /**
     * (Updatable) The Cloud Advisor enrollment status.
     */
    status?: pulumi.Input<string>;
    /**
     * The reason for the enrollment status of the tenancy.
     */
    statusReason?: pulumi.Input<string>;
    /**
     * The date and time the enrollment status was created, in the format defined by RFC3339.
     */
    timeCreated?: pulumi.Input<string>;
    /**
     * The date and time the enrollment status was last updated, in the format defined by RFC3339.
     */
    timeUpdated?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OptimizerEnrollmentStatus resource.
 */
export interface OptimizerEnrollmentStatusArgs {
    /**
     * The unique OCID associated with the enrollment status.
     */
    enrollmentStatusId: pulumi.Input<string>;
    /**
     * (Updatable) The Cloud Advisor enrollment status.
     */
    status: pulumi.Input<string>;
}