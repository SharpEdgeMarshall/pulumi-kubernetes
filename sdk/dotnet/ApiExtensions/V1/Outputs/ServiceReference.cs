// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1
{

    /// <summary>
    /// ServiceReference holds a reference to Service.legacy.k8s.io
    /// </summary>
    [OutputType]
    public sealed class ServiceReference
    {
        /// <summary>
        /// name is the name of the service. Required
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// namespace is the namespace of the service. Required
        /// </summary>
        public readonly string Namespace;
        /// <summary>
        /// path is an optional URL path at which the webhook will be contacted.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// port is an optional service port at which the webhook will be contacted. `port` should be a valid port number (1-65535, inclusive). Defaults to 443 for backward compatibility.
        /// </summary>
        public readonly int Port;

        [OutputConstructor]
        private ServiceReference(
            string name,

            string @namespace,

            string path,

            int port)
        {
            Name = name;
            Namespace = @namespace;
            Path = path;
            Port = port;
        }
    }
}
