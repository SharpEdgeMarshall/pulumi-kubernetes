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
    /// The node this Taint is attached to has the "effect" on any pod that does not tolerate the Taint.
    /// </summary>
    [OutputType]
    public sealed class Taint
    {
        /// <summary>
        /// Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.
        /// </summary>
        public readonly string Effect;
        /// <summary>
        /// Required. The taint key to be applied to a node.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.
        /// </summary>
        public readonly string TimeAdded;
        /// <summary>
        /// The taint value corresponding to the taint key.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private Taint(
            string effect,

            string key,

            string timeAdded,

            string value)
        {
            Effect = effect;
            Key = key;
            TimeAdded = timeAdded;
            Value = value;
        }
    }
}
