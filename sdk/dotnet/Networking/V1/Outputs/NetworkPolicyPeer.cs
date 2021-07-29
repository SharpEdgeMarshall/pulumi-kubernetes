// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Networking.V1
{

    /// <summary>
    /// NetworkPolicyPeer describes a peer to allow traffic to/from. Only certain combinations of fields are allowed
    /// </summary>
    [OutputType]
    public sealed class NetworkPolicyPeer
    {
        /// <summary>
        /// IPBlock defines policy on a particular IPBlock. If this field is set then neither of the other fields can be.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Networking.V1.IPBlock IpBlock;
        /// <summary>
        /// Selects Namespaces using cluster-scoped labels. This field follows standard label selector semantics; if present but empty, it selects all namespaces.
        /// 
        /// If PodSelector is also set, then the NetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects all Pods in the Namespaces selected by NamespaceSelector.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector NamespaceSelector;
        /// <summary>
        /// This is a label selector which selects Pods. This field follows standard label selector semantics; if present but empty, it selects all pods.
        /// 
        /// If NamespaceSelector is also set, then the NetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects the Pods matching PodSelector in the policy's own Namespace.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector PodSelector;

        [OutputConstructor]
        private NetworkPolicyPeer(
            Pulumi.Kubernetes.Types.Outputs.Networking.V1.IPBlock ipBlock,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector namespaceSelector,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector podSelector)
        {
            IpBlock = ipBlock;
            NamespaceSelector = namespaceSelector;
            PodSelector = podSelector;
        }
    }
}
