// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2
{

    /// <summary>
    /// Deployment enables declarative updates for Pods and ReplicaSets.
    /// 
    /// This resource waits until its status is ready before registering success
    /// for create/update, and populating output properties from the current state of the resource.
    /// The following conditions are used to determine whether the resource creation has
    /// succeeded or failed:
    /// 
    /// 1. The Deployment has begun to be updated by the Deployment controller. If the current
    ///    generation of the Deployment is &gt; 1, then this means that the current generation must
    ///    be different from the generation reported by the last outputs.
    /// 2. There exists a ReplicaSet whose revision is equal to the current revision of the
    ///    Deployment.
    /// 3. The Deployment's '.status.conditions' has a status of type 'Available' whose 'status'
    ///    member is set to 'True'.
    /// 4. If the Deployment has generation &gt; 1, then '.status.conditions' has a status of type
    ///    'Progressing', whose 'status' member is set to 'True', and whose 'reason' is
    ///    'NewReplicaSetAvailable'. For generation &lt;= 1, this status field does not exist,
    ///    because it doesn't do a rollout (i.e., it simply creates the Deployment and
    ///    corresponding ReplicaSet), and therefore there is no rollout to mark as 'Progressing'.
    /// 
    /// If the Deployment has not reached a Ready state after 10 minutes, it will
    /// time out and mark the resource update as Failed. You can override the default timeout value
    /// by setting the 'customTimeouts' option on the resource.
    /// </summary>
    [OutputType]
    public sealed class Deployment
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
        /// Standard object metadata.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta Metadata;
        /// <summary>
        /// Specification of the desired behavior of the Deployment.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2.DeploymentSpec Spec;
        /// <summary>
        /// Most recently observed status of the Deployment.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2.DeploymentStatus Status;

        [OutputConstructor]
        private Deployment(
            string apiVersion,

            string kind,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta metadata,

            Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2.DeploymentSpec spec,

            Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2.DeploymentStatus status)
        {
            ApiVersion = apiVersion;
            Kind = kind;
            Metadata = metadata;
            Spec = spec;
            Status = status;
        }
    }
}
