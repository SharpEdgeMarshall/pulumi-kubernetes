// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Apps.V1
{

    /// <summary>
    /// DeploymentStatus is the most recently observed status of the Deployment.
    /// </summary>
    [OutputType]
    public sealed class DeploymentStatus
    {
        /// <summary>
        /// Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.
        /// </summary>
        public readonly int AvailableReplicas;
        /// <summary>
        /// Count of hash collisions for the Deployment. The Deployment controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ReplicaSet.
        /// </summary>
        public readonly int CollisionCount;
        /// <summary>
        /// Represents the latest available observations of a deployment's current state.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Apps.V1.DeploymentCondition> Conditions;
        /// <summary>
        /// The generation observed by the deployment controller.
        /// </summary>
        public readonly int ObservedGeneration;
        /// <summary>
        /// Total number of ready pods targeted by this deployment.
        /// </summary>
        public readonly int ReadyReplicas;
        /// <summary>
        /// Total number of non-terminated pods targeted by this deployment (their labels match the selector).
        /// </summary>
        public readonly int Replicas;
        /// <summary>
        /// Total number of unavailable pods targeted by this deployment. This is the total number of pods that are still required for the deployment to have 100% available capacity. They may either be pods that are running but not yet available or pods that still have not been created.
        /// </summary>
        public readonly int UnavailableReplicas;
        /// <summary>
        /// Total number of non-terminated pods targeted by this deployment that have the desired template spec.
        /// </summary>
        public readonly int UpdatedReplicas;

        [OutputConstructor]
        private DeploymentStatus(
            int availableReplicas,

            int collisionCount,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Apps.V1.DeploymentCondition> conditions,

            int observedGeneration,

            int readyReplicas,

            int replicas,

            int unavailableReplicas,

            int updatedReplicas)
        {
            AvailableReplicas = availableReplicas;
            CollisionCount = collisionCount;
            Conditions = conditions;
            ObservedGeneration = observedGeneration;
            ReadyReplicas = readyReplicas;
            Replicas = replicas;
            UnavailableReplicas = unavailableReplicas;
            UpdatedReplicas = updatedReplicas;
        }
    }
}
