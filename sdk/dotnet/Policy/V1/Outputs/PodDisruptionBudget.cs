// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Policy.V1
{

    /// <summary>
    /// PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods
    /// </summary>
    [OutputType]
    public sealed class PodDisruptionBudget
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        public readonly string ApiVersion;
        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta Metadata;
        /// <summary>
        /// Specification of the desired behavior of the PodDisruptionBudget.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Policy.V1.PodDisruptionBudgetSpec Spec;
        /// <summary>
        /// Most recently observed status of the PodDisruptionBudget.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Policy.V1.PodDisruptionBudgetStatus Status;

        [OutputConstructor]
        private PodDisruptionBudget(
            string apiVersion,

            string kind,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta metadata,

            Pulumi.Kubernetes.Types.Outputs.Policy.V1.PodDisruptionBudgetSpec spec,

            Pulumi.Kubernetes.Types.Outputs.Policy.V1.PodDisruptionBudgetStatus status)
        {
            ApiVersion = apiVersion;
            Kind = kind;
            Metadata = metadata;
            Spec = spec;
            Status = status;
        }
    }
}
