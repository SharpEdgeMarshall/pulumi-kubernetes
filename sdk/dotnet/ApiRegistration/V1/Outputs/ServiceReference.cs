// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.ApiRegistration.V1
{

    /// <summary>
    /// ServiceReference holds a reference to Service.legacy.k8s.io
    /// </summary>
    [OutputType]
    public sealed class ServiceReference
    {
        /// <summary>
        /// Name is the name of the service
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Namespace is the namespace of the service
        /// </summary>
        public readonly string Namespace;
        /// <summary>
        /// If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
        /// </summary>
        public readonly int Port;

        [OutputConstructor]
        private ServiceReference(
            string name,

            string @namespace,

            int port)
        {
            Name = name;
            Namespace = @namespace;
            Port = port;
        }
    }
}
