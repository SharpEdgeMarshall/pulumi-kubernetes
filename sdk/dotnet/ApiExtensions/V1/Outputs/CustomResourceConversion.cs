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
    /// CustomResourceConversion describes how to convert different versions of a CR.
    /// </summary>
    [OutputType]
    public sealed class CustomResourceConversion
    {
        /// <summary>
        /// strategy specifies how custom resources are converted between versions. Allowed values are: - `None`: The converter only change the apiVersion and would not touch any other field in the custom resource. - `Webhook`: API Server will call to an external webhook to do the conversion. Additional information
        ///   is needed for this option. This requires spec.preserveUnknownFields to be false, and spec.conversion.webhook to be set.
        /// </summary>
        public readonly string Strategy;
        /// <summary>
        /// webhook describes how to call the conversion webhook. Required when `strategy` is set to `Webhook`.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1.WebhookConversion Webhook;

        [OutputConstructor]
        private CustomResourceConversion(
            string strategy,

            Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1.WebhookConversion webhook)
        {
            Strategy = strategy;
            Webhook = webhook;
        }
    }
}
