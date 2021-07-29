// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Batch.V1
{

    /// <summary>
    /// JobStatus represents the current state of a Job.
    /// </summary>
    [OutputType]
    public sealed class JobStatus
    {
        /// <summary>
        /// The number of actively running pods.
        /// </summary>
        public readonly int Active;
        /// <summary>
        /// CompletedIndexes holds the completed indexes when .spec.completionMode = "Indexed" in a text format. The indexes are represented as decimal integers separated by commas. The numbers are listed in increasing order. Three or more consecutive numbers are compressed and represented by the first and last element of the series, separated by a hyphen. For example, if the completed indexes are 1, 3, 4, 5 and 7, they are represented as "1,3-5,7".
        /// </summary>
        public readonly string CompletedIndexes;
        /// <summary>
        /// Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC. The completion time is only set when the job finishes successfully.
        /// </summary>
        public readonly string CompletionTime;
        /// <summary>
        /// The latest available observations of an object's current state. When a Job fails, one of the conditions will have type "Failed" and status true. When a Job is suspended, one of the conditions will have type "Suspended" and status true; when the Job is resumed, the status of this condition will become false. When a Job is completed, one of the conditions will have type "Complete" and status true. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Batch.V1.JobCondition> Conditions;
        /// <summary>
        /// The number of pods which reached phase Failed.
        /// </summary>
        public readonly int Failed;
        /// <summary>
        /// Represents time when the job controller started processing a job. When a Job is created in the suspended state, this field is not set until the first time it is resumed. This field is reset every time a Job is resumed from suspension. It is represented in RFC3339 form and is in UTC.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// The number of pods which reached phase Succeeded.
        /// </summary>
        public readonly int Succeeded;

        [OutputConstructor]
        private JobStatus(
            int active,

            string completedIndexes,

            string completionTime,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Batch.V1.JobCondition> conditions,

            int failed,

            string startTime,

            int succeeded)
        {
            Active = active;
            CompletedIndexes = completedIndexes;
            CompletionTime = completionTime;
            Conditions = conditions;
            Failed = failed;
            StartTime = startTime;
            Succeeded = succeeded;
        }
    }
}
