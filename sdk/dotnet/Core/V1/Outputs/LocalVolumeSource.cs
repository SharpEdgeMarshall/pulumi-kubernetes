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
    /// Local represents directly-attached storage with node affinity (Beta feature)
    /// </summary>
    [OutputType]
    public sealed class LocalVolumeSource
    {
        /// <summary>
        /// Filesystem type to mount. It applies only when the Path is a block device. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default value is to auto-select a fileystem if unspecified.
        /// </summary>
        public readonly string FsType;
        /// <summary>
        /// The full path to the volume on the node. It can be either a directory or block device (disk, partition, ...).
        /// </summary>
        public readonly string Path;

        [OutputConstructor]
        private LocalVolumeSource(
            string fsType,

            string path)
        {
            FsType = fsType;
            Path = path;
        }
    }
}
