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
    /// SupplementalGroupsStrategyOptions defines the strategy type and options used to create the strategy. Deprecated: use SupplementalGroupsStrategyOptions from policy API Group instead.
    /// </summary>
    [OutputType]
    public sealed class SupplementalGroupsStrategyOptions
    {
        /// <summary>
        /// ranges are the allowed ranges of supplemental groups.  If you would like to force a single supplemental group then supply a single range with the same start and end. Required for MustRunAs.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1.IDRange> Ranges;
        /// <summary>
        /// rule is the strategy that will dictate what supplemental groups is used in the SecurityContext.
        /// </summary>
        public readonly string Rule;

        [OutputConstructor]
        private SupplementalGroupsStrategyOptions(
            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Extensions.V1Beta1.IDRange> ranges,

            string rule)
        {
            Ranges = ranges;
            Rule = rule;
        }
    }
}
