// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Beta1
{

    /// <summary>
    /// FlowSchemaStatus represents the current state of a FlowSchema.
    /// </summary>
    [OutputType]
    public sealed class FlowSchemaStatus
    {
        /// <summary>
        /// `conditions` is a list of the current states of FlowSchema.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Beta1.FlowSchemaCondition> Conditions;

        [OutputConstructor]
        private FlowSchemaStatus(ImmutableArray<Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Beta1.FlowSchemaCondition> conditions)
        {
            Conditions = conditions;
        }
    }
}
