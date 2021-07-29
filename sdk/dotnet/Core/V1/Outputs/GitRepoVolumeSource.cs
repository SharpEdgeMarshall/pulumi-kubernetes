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
    /// Represents a volume that is populated with the contents of a git repository. Git repo volumes do not support ownership management. Git repo volumes support SELinux relabeling.
    /// 
    /// DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.
    /// </summary>
    [OutputType]
    public sealed class GitRepoVolumeSource
    {
        /// <summary>
        /// Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.
        /// </summary>
        public readonly string Directory;
        /// <summary>
        /// Repository URL
        /// </summary>
        public readonly string Repository;
        /// <summary>
        /// Commit hash for the specified revision.
        /// </summary>
        public readonly string Revision;

        [OutputConstructor]
        private GitRepoVolumeSource(
            string directory,

            string repository,

            string revision)
        {
            Directory = directory;
            Repository = repository;
            Revision = revision;
        }
    }
}
