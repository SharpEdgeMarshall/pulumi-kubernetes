// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Storage.V1Beta1
{

    /// <summary>
    /// CSINodeSpec holds information about the specification of all CSI drivers installed on a node
    /// </summary>
    [OutputType]
    public sealed class CSINodeSpec
    {
        /// <summary>
        /// drivers is a list of information of all CSI Drivers existing on a node. If all drivers in the list are uninstalled, this can become empty.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Storage.V1Beta1.CSINodeDriver> Drivers;

        [OutputConstructor]
        private CSINodeSpec(ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Storage.V1Beta1.CSINodeDriver> drivers)
        {
            Drivers = drivers;
        }
    }
}
