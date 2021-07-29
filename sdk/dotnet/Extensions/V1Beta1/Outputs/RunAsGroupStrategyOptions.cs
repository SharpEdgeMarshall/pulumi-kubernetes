// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1
{

    /// <summary>
    /// RunAsGroupStrategyOptions defines the strategy type and any options used to create the strategy. Deprecated: use RunAsGroupStrategyOptions from policy API Group instead.
    /// </summary>
    [OutputType]
    public sealed class RunAsGroupStrategyOptions
    {
        /// <summary>
        /// ranges are the allowed ranges of gids that may be used. If you would like to force a single gid then supply a single range with the same start and end. Required for MustRunAs.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1.IDRange> Ranges;
        /// <summary>
        /// rule is the strategy that will dictate the allowable RunAsGroup values that may be set.
        /// </summary>
        public readonly string Rule;

        [OutputConstructor]
        private RunAsGroupStrategyOptions(
            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1.IDRange> ranges,

            string rule)
        {
            Ranges = ranges;
            Rule = rule;
        }
    }
}
