// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./podDisruptionBudget";
export * from "./podDisruptionBudgetList";

// Import resources to register:
import { PodDisruptionBudget } from "./podDisruptionBudget";
import { PodDisruptionBudgetList } from "./podDisruptionBudgetList";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "kubernetes:policy/v1:PodDisruptionBudget":
                return new PodDisruptionBudget(name, <any>undefined, { urn })
            case "kubernetes:policy/v1:PodDisruptionBudgetList":
                return new PodDisruptionBudgetList(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("kubernetes", "policy/v1", _module)