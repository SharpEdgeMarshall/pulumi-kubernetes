// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Core.V1
{

    /// <summary>
    /// ResourceQuota sets aggregate quota restrictions enforced per namespace
    /// </summary>
    [OutputType]
    public sealed class ResourceQuota
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
        /// Spec defines the desired quota. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.ResourceQuotaSpec Spec;
        /// <summary>
        /// Status defines the actual enforced quota and its current usage. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.ResourceQuotaStatus Status;

        [OutputConstructor]
        private ResourceQuota(
            string apiVersion,

            string kind,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta metadata,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.ResourceQuotaSpec spec,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.ResourceQuotaStatus status)
        {
            ApiVersion = apiVersion;
            Kind = kind;
            Metadata = metadata;
            Spec = spec;
            Status = status;
        }
    }
}
