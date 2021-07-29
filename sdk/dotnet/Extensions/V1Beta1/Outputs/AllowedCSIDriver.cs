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
    /// AllowedCSIDriver represents a single inline CSI Driver that is allowed to be used.
    /// </summary>
    [OutputType]
    public sealed class AllowedCSIDriver
    {
        /// <summary>
        /// Name is the registered name of the CSI driver
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private AllowedCSIDriver(string name)
        {
            Name = name;
        }
    }
}
