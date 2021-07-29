// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2alpha1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CronJob represents the configuration of a single cron job.
type CronJob struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Specification of the desired behavior of a cron job, including the schedule. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec CronJobSpecPtrOutput `pulumi:"spec"`
	// Current status of a cron job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status CronJobStatusPtrOutput `pulumi:"status"`
}

// NewCronJob registers a new resource with the given unique name, arguments, and options.
func NewCronJob(ctx *pulumi.Context,
	name string, args *CronJobArgs, opts ...pulumi.ResourceOption) (*CronJob, error) {
	if args == nil {
		args = &CronJobArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("batch/v2alpha1")
	args.Kind = pulumi.StringPtr("CronJob")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:batch/v1:CronJob"),
		},
		{
			Type: pulumi.String("kubernetes:batch/v1beta1:CronJob"),
		},
	})
	opts = append(opts, aliases)
	var resource CronJob
	err := ctx.RegisterResource("kubernetes:batch/v2alpha1:CronJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCronJob gets an existing CronJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCronJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CronJobState, opts ...pulumi.ResourceOption) (*CronJob, error) {
	var resource CronJob
	err := ctx.ReadResource("kubernetes:batch/v2alpha1:CronJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CronJob resources.
type cronJobState struct {
}

type CronJobState struct {
}

func (CronJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*cronJobState)(nil)).Elem()
}

type cronJobArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Specification of the desired behavior of a cron job, including the schedule. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *CronJobSpec `pulumi:"spec"`
}

// The set of arguments for constructing a CronJob resource.
type CronJobArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Specification of the desired behavior of a cron job, including the schedule. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec CronJobSpecPtrInput
}

func (CronJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cronJobArgs)(nil)).Elem()
}

type CronJobInput interface {
	pulumi.Input

	ToCronJobOutput() CronJobOutput
	ToCronJobOutputWithContext(ctx context.Context) CronJobOutput
}

func (*CronJob) ElementType() reflect.Type {
	return reflect.TypeOf((*CronJob)(nil))
}

func (i *CronJob) ToCronJobOutput() CronJobOutput {
	return i.ToCronJobOutputWithContext(context.Background())
}

func (i *CronJob) ToCronJobOutputWithContext(ctx context.Context) CronJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobOutput)
}

func (i *CronJob) ToCronJobPtrOutput() CronJobPtrOutput {
	return i.ToCronJobPtrOutputWithContext(context.Background())
}

func (i *CronJob) ToCronJobPtrOutputWithContext(ctx context.Context) CronJobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobPtrOutput)
}

type CronJobPtrInput interface {
	pulumi.Input

	ToCronJobPtrOutput() CronJobPtrOutput
	ToCronJobPtrOutputWithContext(ctx context.Context) CronJobPtrOutput
}

type cronJobPtrType CronJobArgs

func (*cronJobPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CronJob)(nil))
}

func (i *cronJobPtrType) ToCronJobPtrOutput() CronJobPtrOutput {
	return i.ToCronJobPtrOutputWithContext(context.Background())
}

func (i *cronJobPtrType) ToCronJobPtrOutputWithContext(ctx context.Context) CronJobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobPtrOutput)
}

// CronJobArrayInput is an input type that accepts CronJobArray and CronJobArrayOutput values.
// You can construct a concrete instance of `CronJobArrayInput` via:
//
//          CronJobArray{ CronJobArgs{...} }
type CronJobArrayInput interface {
	pulumi.Input

	ToCronJobArrayOutput() CronJobArrayOutput
	ToCronJobArrayOutputWithContext(context.Context) CronJobArrayOutput
}

type CronJobArray []CronJobInput

func (CronJobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CronJob)(nil)).Elem()
}

func (i CronJobArray) ToCronJobArrayOutput() CronJobArrayOutput {
	return i.ToCronJobArrayOutputWithContext(context.Background())
}

func (i CronJobArray) ToCronJobArrayOutputWithContext(ctx context.Context) CronJobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobArrayOutput)
}

// CronJobMapInput is an input type that accepts CronJobMap and CronJobMapOutput values.
// You can construct a concrete instance of `CronJobMapInput` via:
//
//          CronJobMap{ "key": CronJobArgs{...} }
type CronJobMapInput interface {
	pulumi.Input

	ToCronJobMapOutput() CronJobMapOutput
	ToCronJobMapOutputWithContext(context.Context) CronJobMapOutput
}

type CronJobMap map[string]CronJobInput

func (CronJobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CronJob)(nil)).Elem()
}

func (i CronJobMap) ToCronJobMapOutput() CronJobMapOutput {
	return i.ToCronJobMapOutputWithContext(context.Background())
}

func (i CronJobMap) ToCronJobMapOutputWithContext(ctx context.Context) CronJobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobMapOutput)
}

type CronJobOutput struct {
	*pulumi.OutputState
}

func (CronJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CronJob)(nil))
}

func (o CronJobOutput) ToCronJobOutput() CronJobOutput {
	return o
}

func (o CronJobOutput) ToCronJobOutputWithContext(ctx context.Context) CronJobOutput {
	return o
}

func (o CronJobOutput) ToCronJobPtrOutput() CronJobPtrOutput {
	return o.ToCronJobPtrOutputWithContext(context.Background())
}

func (o CronJobOutput) ToCronJobPtrOutputWithContext(ctx context.Context) CronJobPtrOutput {
	return o.ApplyT(func(v CronJob) *CronJob {
		return &v
	}).(CronJobPtrOutput)
}

type CronJobPtrOutput struct {
	*pulumi.OutputState
}

func (CronJobPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CronJob)(nil))
}

func (o CronJobPtrOutput) ToCronJobPtrOutput() CronJobPtrOutput {
	return o
}

func (o CronJobPtrOutput) ToCronJobPtrOutputWithContext(ctx context.Context) CronJobPtrOutput {
	return o
}

type CronJobArrayOutput struct{ *pulumi.OutputState }

func (CronJobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CronJob)(nil))
}

func (o CronJobArrayOutput) ToCronJobArrayOutput() CronJobArrayOutput {
	return o
}

func (o CronJobArrayOutput) ToCronJobArrayOutputWithContext(ctx context.Context) CronJobArrayOutput {
	return o
}

func (o CronJobArrayOutput) Index(i pulumi.IntInput) CronJobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CronJob {
		return vs[0].([]CronJob)[vs[1].(int)]
	}).(CronJobOutput)
}

type CronJobMapOutput struct{ *pulumi.OutputState }

func (CronJobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CronJob)(nil))
}

func (o CronJobMapOutput) ToCronJobMapOutput() CronJobMapOutput {
	return o
}

func (o CronJobMapOutput) ToCronJobMapOutputWithContext(ctx context.Context) CronJobMapOutput {
	return o
}

func (o CronJobMapOutput) MapIndex(k pulumi.StringInput) CronJobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CronJob {
		return vs[0].(map[string]CronJob)[vs[1].(string)]
	}).(CronJobOutput)
}

func init() {
	pulumi.RegisterOutputType(CronJobOutput{})
	pulumi.RegisterOutputType(CronJobPtrOutput{})
	pulumi.RegisterOutputType(CronJobArrayOutput{})
	pulumi.RegisterOutputType(CronJobMapOutput{})
}
