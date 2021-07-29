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
    /// DeploymentSpec is the specification of the desired behavior of the Deployment.
    /// </summary>
    [OutputType]
    public sealed class DeploymentSpec
    {
        /// <summary>
        /// Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
        /// </summary>
        public readonly int MinReadySeconds;
        /// <summary>
        /// Indicates that the deployment is paused.
        /// </summary>
        public readonly bool Paused;
        /// <summary>
        /// The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s.
        /// </summary>
        public readonly int ProgressDeadlineSeconds;
        /// <summary>
        /// Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
        /// </summary>
        public readonly int Replicas;
        /// <summary>
        /// The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
        /// </summary>
        public readonly int RevisionHistoryLimit;
        /// <summary>
        /// Label selector for pods. Existing ReplicaSets whose pods are selected by this will be the ones affected by this deployment. It must match the pod template's labels.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector Selector;
        /// <summary>
        /// The deployment strategy to use to replace existing pods with new ones.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2.DeploymentStrategy Strategy;
        /// <summary>
        /// Template describes the pods that will be created.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.PodTemplateSpec Template;

        [OutputConstructor]
        private DeploymentSpec(
            int minReadySeconds,

            bool paused,

            int progressDeadlineSeconds,

            int replicas,

            int revisionHistoryLimit,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector selector,

            Pulumi.Kubernetes.Types.Outputs.Apps.V1Beta2.DeploymentStrategy strategy,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.PodTemplateSpec template)
        {
            MinReadySeconds = minReadySeconds;
            Paused = paused;
            ProgressDeadlineSeconds = progressDeadlineSeconds;
            Replicas = replicas;
            RevisionHistoryLimit = revisionHistoryLimit;
            Selector = selector;
            Strategy = strategy;
            Template = template;
        }
    }
}
