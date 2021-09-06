// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./alertRule";
export * from "./budget";
export * from "./getAlertRule";
export * from "./getAlertRules";
export * from "./getBudget";
export * from "./getBudgets";

// Import resources to register:
import { AlertRule } from "./alertRule";
import { Budget } from "./budget";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "oci:budget/alertRule:AlertRule":
                return new AlertRule(name, <any>undefined, { urn })
            case "oci:budget/budget:Budget":
                return new Budget(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("oci", "budget/alertRule", _module)
pulumi.runtime.registerResourceModule("oci", "budget/budget", _module)