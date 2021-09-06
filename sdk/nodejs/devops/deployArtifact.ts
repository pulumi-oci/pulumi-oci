// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This resource provides the Deploy Artifact resource in Oracle Cloud Infrastructure Devops service.
 *
 * Creates a new deployment artifact.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testDeployArtifact = new oci.devops.DeployArtifact("testDeployArtifact", {
 *     argumentSubstitutionMode: _var.deploy_artifact_argument_substitution_mode,
 *     deployArtifactSource: {
 *         deployArtifactSourceType: _var.deploy_artifact_deploy_artifact_source_deploy_artifact_source_type,
 *         base64encodedContent: _var.deploy_artifact_deploy_artifact_source_base64encoded_content,
 *         deployArtifactPath: _var.deploy_artifact_deploy_artifact_source_deploy_artifact_path,
 *         deployArtifactVersion: _var.deploy_artifact_deploy_artifact_source_deploy_artifact_version,
 *         imageDigest: _var.deploy_artifact_deploy_artifact_source_image_digest,
 *         imageUri: _var.deploy_artifact_deploy_artifact_source_image_uri,
 *         repositoryId: oci_artifacts_repository.test_repository.id,
 *     },
 *     deployArtifactType: _var.deploy_artifact_deploy_artifact_type,
 *     projectId: oci_devops_project.test_project.id,
 *     definedTags: {
 *         "foo-namespace.bar-key": "value",
 *     },
 *     description: _var.deploy_artifact_description,
 *     displayName: _var.deploy_artifact_display_name,
 *     freeformTags: {
 *         "bar-key": "value",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * DeployArtifacts can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import oci:devops/deployArtifact:DeployArtifact test_deploy_artifact "id"
 * ```
 */
export class DeployArtifact extends pulumi.CustomResource {
    /**
     * Get an existing DeployArtifact resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeployArtifactState, opts?: pulumi.CustomResourceOptions): DeployArtifact {
        return new DeployArtifact(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'oci:devops/deployArtifact:DeployArtifact';

    /**
     * Returns true if the given object is an instance of DeployArtifact.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeployArtifact {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeployArtifact.__pulumiType;
    }

    /**
     * (Updatable) Mode for artifact parameter substitution.
     */
    public readonly argumentSubstitutionMode!: pulumi.Output<string>;
    /**
     * The OCID of a compartment.
     */
    public /*out*/ readonly compartmentId!: pulumi.Output<string>;
    /**
     * (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
     */
    public readonly definedTags!: pulumi.Output<{[key: string]: any}>;
    /**
     * (Updatable) Specifies source of an artifact.
     */
    public readonly deployArtifactSource!: pulumi.Output<outputs.devops.DeployArtifactDeployArtifactSource>;
    /**
     * (Updatable) Type of the deployment artifact.
     */
    public readonly deployArtifactType!: pulumi.Output<string>;
    /**
     * (Updatable) Optional description about the deployment artifact.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * (Updatable) Deployment artifact display name. Avoid entering confidential information.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
     */
    public readonly freeformTags!: pulumi.Output<{[key: string]: any}>;
    /**
     * A detailed message describing the current state. For example, can be used to provide actionable information for a resource in Failed state.
     */
    public /*out*/ readonly lifecycleDetails!: pulumi.Output<string>;
    /**
     * The OCID of a project.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Current state of the deployment artifact.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
     */
    public /*out*/ readonly systemTags!: pulumi.Output<{[key: string]: any}>;
    /**
     * Time the deployment artifact was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
     */
    public /*out*/ readonly timeCreated!: pulumi.Output<string>;
    /**
     * Time the deployment artifact was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
     */
    public /*out*/ readonly timeUpdated!: pulumi.Output<string>;

    /**
     * Create a DeployArtifact resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeployArtifactArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeployArtifactArgs | DeployArtifactState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeployArtifactState | undefined;
            inputs["argumentSubstitutionMode"] = state ? state.argumentSubstitutionMode : undefined;
            inputs["compartmentId"] = state ? state.compartmentId : undefined;
            inputs["definedTags"] = state ? state.definedTags : undefined;
            inputs["deployArtifactSource"] = state ? state.deployArtifactSource : undefined;
            inputs["deployArtifactType"] = state ? state.deployArtifactType : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["freeformTags"] = state ? state.freeformTags : undefined;
            inputs["lifecycleDetails"] = state ? state.lifecycleDetails : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["systemTags"] = state ? state.systemTags : undefined;
            inputs["timeCreated"] = state ? state.timeCreated : undefined;
            inputs["timeUpdated"] = state ? state.timeUpdated : undefined;
        } else {
            const args = argsOrState as DeployArtifactArgs | undefined;
            if ((!args || args.argumentSubstitutionMode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'argumentSubstitutionMode'");
            }
            if ((!args || args.deployArtifactSource === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deployArtifactSource'");
            }
            if ((!args || args.deployArtifactType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deployArtifactType'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            inputs["argumentSubstitutionMode"] = args ? args.argumentSubstitutionMode : undefined;
            inputs["definedTags"] = args ? args.definedTags : undefined;
            inputs["deployArtifactSource"] = args ? args.deployArtifactSource : undefined;
            inputs["deployArtifactType"] = args ? args.deployArtifactType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["freeformTags"] = args ? args.freeformTags : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["compartmentId"] = undefined /*out*/;
            inputs["lifecycleDetails"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["systemTags"] = undefined /*out*/;
            inputs["timeCreated"] = undefined /*out*/;
            inputs["timeUpdated"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DeployArtifact.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeployArtifact resources.
 */
export interface DeployArtifactState {
    /**
     * (Updatable) Mode for artifact parameter substitution.
     */
    argumentSubstitutionMode?: pulumi.Input<string>;
    /**
     * The OCID of a compartment.
     */
    compartmentId?: pulumi.Input<string>;
    /**
     * (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
     */
    definedTags?: pulumi.Input<{[key: string]: any}>;
    /**
     * (Updatable) Specifies source of an artifact.
     */
    deployArtifactSource?: pulumi.Input<inputs.devops.DeployArtifactDeployArtifactSource>;
    /**
     * (Updatable) Type of the deployment artifact.
     */
    deployArtifactType?: pulumi.Input<string>;
    /**
     * (Updatable) Optional description about the deployment artifact.
     */
    description?: pulumi.Input<string>;
    /**
     * (Updatable) Deployment artifact display name. Avoid entering confidential information.
     */
    displayName?: pulumi.Input<string>;
    /**
     * (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
     */
    freeformTags?: pulumi.Input<{[key: string]: any}>;
    /**
     * A detailed message describing the current state. For example, can be used to provide actionable information for a resource in Failed state.
     */
    lifecycleDetails?: pulumi.Input<string>;
    /**
     * The OCID of a project.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Current state of the deployment artifact.
     */
    state?: pulumi.Input<string>;
    /**
     * Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
     */
    systemTags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Time the deployment artifact was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
     */
    timeCreated?: pulumi.Input<string>;
    /**
     * Time the deployment artifact was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
     */
    timeUpdated?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DeployArtifact resource.
 */
export interface DeployArtifactArgs {
    /**
     * (Updatable) Mode for artifact parameter substitution.
     */
    argumentSubstitutionMode: pulumi.Input<string>;
    /**
     * (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
     */
    definedTags?: pulumi.Input<{[key: string]: any}>;
    /**
     * (Updatable) Specifies source of an artifact.
     */
    deployArtifactSource: pulumi.Input<inputs.devops.DeployArtifactDeployArtifactSource>;
    /**
     * (Updatable) Type of the deployment artifact.
     */
    deployArtifactType: pulumi.Input<string>;
    /**
     * (Updatable) Optional description about the deployment artifact.
     */
    description?: pulumi.Input<string>;
    /**
     * (Updatable) Deployment artifact display name. Avoid entering confidential information.
     */
    displayName?: pulumi.Input<string>;
    /**
     * (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
     */
    freeformTags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The OCID of a project.
     */
    projectId: pulumi.Input<string>;
}