// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Discovery.V1
{

    /// <summary>
    /// EndpointPort represents a Port used by an EndpointSlice
    /// </summary>
    [OutputType]
    public sealed class EndpointPort
    {
        /// <summary>
        /// The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and http://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.
        /// </summary>
        public readonly string AppProtocol;
        /// <summary>
        /// The name of this port. All ports in an EndpointSlice must have a unique name. If the EndpointSlice is dervied from a Kubernetes service, this corresponds to the Service.ports[].name. Name must either be an empty string or pass DNS_LABEL validation: * must be no more than 63 characters long. * must consist of lower case alphanumeric characters or '-'. * must start and end with an alphanumeric character. Default is empty string.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The port number of the endpoint. If this is not specified, ports are not restricted and must be interpreted in the context of the specific consumer.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
        /// </summary>
        public readonly string Protocol;

        [OutputConstructor]
        private EndpointPort(
            string appProtocol,

            string name,

            int port,

            string protocol)
        {
            AppProtocol = appProtocol;
            Name = name;
            Port = port;
            Protocol = protocol;
        }
    }
}
