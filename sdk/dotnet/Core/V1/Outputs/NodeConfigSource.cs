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
    /// NodeConfigSource specifies a source of node configuration. Exactly one subfield (excluding metadata) must be non-nil.
    /// </summary>
    [OutputType]
    public sealed class NodeConfigSource
    {
        /// <summary>
        /// ConfigMap is a reference to a Node's ConfigMap
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.ConfigMapNodeConfigSource ConfigMap;

        [OutputConstructor]
        private NodeConfigSource(Pulumi.Kubernetes.Types.Outputs.Core.V1.ConfigMapNodeConfigSource configMap)
        {
            ConfigMap = configMap;
        }
    }
}
