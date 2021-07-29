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
    /// Represents a StorageOS persistent volume resource.
    /// </summary>
    [OutputType]
    public sealed class StorageOSPersistentVolumeSource
    {
        /// <summary>
        /// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
        /// </summary>
        public readonly string FsType;
        /// <summary>
        /// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
        /// </summary>
        public readonly bool ReadOnly;
        /// <summary>
        /// SecretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.ObjectReference SecretRef;
        /// <summary>
        /// VolumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.
        /// </summary>
        public readonly string VolumeName;
        /// <summary>
        /// VolumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to "default" if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.
        /// </summary>
        public readonly string VolumeNamespace;

        [OutputConstructor]
        private StorageOSPersistentVolumeSource(
            string fsType,

            bool readOnly,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.ObjectReference secretRef,

            string volumeName,

            string volumeNamespace)
        {
            FsType = fsType;
            ReadOnly = readOnly;
            SecretRef = secretRef;
            VolumeName = volumeName;
            VolumeNamespace = volumeNamespace;
        }
    }
}
