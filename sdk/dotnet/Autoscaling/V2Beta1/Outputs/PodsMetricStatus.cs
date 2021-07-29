// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta1
{

    /// <summary>
    /// PodsMetricStatus indicates the current value of a metric describing each pod in the current scale target (for example, transactions-processed-per-second).
    /// </summary>
    [OutputType]
    public sealed class PodsMetricStatus
    {
        /// <summary>
        /// currentAverageValue is the current value of the average of the metric across all relevant pods (as a quantity)
        /// </summary>
        public readonly string CurrentAverageValue;
        /// <summary>
        /// metricName is the name of the metric in question
        /// </summary>
        public readonly string MetricName;
        /// <summary>
        /// selector is the string-encoded form of a standard kubernetes label selector for the given metric When set in the PodsMetricSource, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector Selector;

        [OutputConstructor]
        private PodsMetricStatus(
            string currentAverageValue,

            string metricName,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector selector)
        {
            CurrentAverageValue = currentAverageValue;
            MetricName = metricName;
            Selector = selector;
        }
    }
}
