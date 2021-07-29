// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Rbac.V1Alpha1
{

    /// <summary>
    /// Subject contains a reference to the object or user identities a role binding applies to.  This can either hold a direct API object reference, or a value for non-objects such as user and group names.
    /// </summary>
    [OutputType]
    public sealed class Subject
    {
        /// <summary>
        /// APIVersion holds the API group and version of the referenced subject. Defaults to "v1" for ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io/v1alpha1" for User and Group subjects.
        /// </summary>
        public readonly string ApiVersion;
        /// <summary>
        /// Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount". If the Authorizer does not recognized the kind value, the Authorizer should report an error.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the object being referenced.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Namespace of the referenced object.  If the object kind is non-namespace, such as "User" or "Group", and this value is not empty the Authorizer should report an error.
        /// </summary>
        public readonly string Namespace;

        [OutputConstructor]
        private Subject(
            string apiVersion,

            string kind,

            string name,

            string @namespace)
        {
            ApiVersion = apiVersion;
            Kind = kind;
            Name = name;
            Namespace = @namespace;
        }
    }
}
