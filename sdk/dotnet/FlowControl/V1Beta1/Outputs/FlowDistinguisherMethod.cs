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
    /// FlowDistinguisherMethod specifies the method of a flow distinguisher.
    /// </summary>
    [OutputType]
    public sealed class FlowDistinguisherMethod
    {
        /// <summary>
        /// `type` is the type of flow distinguisher method The supported types are "ByUser" and "ByNamespace". Required.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private FlowDistinguisherMethod(string type)
        {
            Type = type;
        }
    }
}
