// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CronJobList is a collection of cron jobs.
type CronJobList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// items is the list of CronJobs.
	Items CronJobTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewCronJobList registers a new resource with the given unique name, arguments, and options.
func NewCronJobList(ctx *pulumi.Context,
	name string, args *CronJobListArgs, opts ...pulumi.ResourceOption) (*CronJobList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("batch/v1beta1")
	args.Kind = pulumi.StringPtr("CronJobList")
	var resource CronJobList
	err := ctx.RegisterResource("kubernetes:batch/v1beta1:CronJobList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCronJobList gets an existing CronJobList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCronJobList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CronJobListState, opts ...pulumi.ResourceOption) (*CronJobList, error) {
	var resource CronJobList
	err := ctx.ReadResource("kubernetes:batch/v1beta1:CronJobList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CronJobList resources.
type cronJobListState struct {
}

type CronJobListState struct {
}

func (CronJobListState) ElementType() reflect.Type {
	return reflect.TypeOf((*cronJobListState)(nil)).Elem()
}

type cronJobListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items is the list of CronJobs.
	Items []CronJobType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a CronJobList resource.
type CronJobListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items is the list of CronJobs.
	Items CronJobTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (CronJobListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cronJobListArgs)(nil)).Elem()
}

type CronJobListInput interface {
	pulumi.Input

	ToCronJobListOutput() CronJobListOutput
	ToCronJobListOutputWithContext(ctx context.Context) CronJobListOutput
}

func (*CronJobList) ElementType() reflect.Type {
	return reflect.TypeOf((*CronJobList)(nil))
}

func (i *CronJobList) ToCronJobListOutput() CronJobListOutput {
	return i.ToCronJobListOutputWithContext(context.Background())
}

func (i *CronJobList) ToCronJobListOutputWithContext(ctx context.Context) CronJobListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobListOutput)
}

func (i *CronJobList) ToCronJobListPtrOutput() CronJobListPtrOutput {
	return i.ToCronJobListPtrOutputWithContext(context.Background())
}

func (i *CronJobList) ToCronJobListPtrOutputWithContext(ctx context.Context) CronJobListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobListPtrOutput)
}

type CronJobListPtrInput interface {
	pulumi.Input

	ToCronJobListPtrOutput() CronJobListPtrOutput
	ToCronJobListPtrOutputWithContext(ctx context.Context) CronJobListPtrOutput
}

type cronJobListPtrType CronJobListArgs

func (*cronJobListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CronJobList)(nil))
}

func (i *cronJobListPtrType) ToCronJobListPtrOutput() CronJobListPtrOutput {
	return i.ToCronJobListPtrOutputWithContext(context.Background())
}

func (i *cronJobListPtrType) ToCronJobListPtrOutputWithContext(ctx context.Context) CronJobListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobListPtrOutput)
}

// CronJobListArrayInput is an input type that accepts CronJobListArray and CronJobListArrayOutput values.
// You can construct a concrete instance of `CronJobListArrayInput` via:
//
//          CronJobListArray{ CronJobListArgs{...} }
type CronJobListArrayInput interface {
	pulumi.Input

	ToCronJobListArrayOutput() CronJobListArrayOutput
	ToCronJobListArrayOutputWithContext(context.Context) CronJobListArrayOutput
}

type CronJobListArray []CronJobListInput

func (CronJobListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CronJobList)(nil)).Elem()
}

func (i CronJobListArray) ToCronJobListArrayOutput() CronJobListArrayOutput {
	return i.ToCronJobListArrayOutputWithContext(context.Background())
}

func (i CronJobListArray) ToCronJobListArrayOutputWithContext(ctx context.Context) CronJobListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobListArrayOutput)
}

// CronJobListMapInput is an input type that accepts CronJobListMap and CronJobListMapOutput values.
// You can construct a concrete instance of `CronJobListMapInput` via:
//
//          CronJobListMap{ "key": CronJobListArgs{...} }
type CronJobListMapInput interface {
	pulumi.Input

	ToCronJobListMapOutput() CronJobListMapOutput
	ToCronJobListMapOutputWithContext(context.Context) CronJobListMapOutput
}

type CronJobListMap map[string]CronJobListInput

func (CronJobListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CronJobList)(nil)).Elem()
}

func (i CronJobListMap) ToCronJobListMapOutput() CronJobListMapOutput {
	return i.ToCronJobListMapOutputWithContext(context.Background())
}

func (i CronJobListMap) ToCronJobListMapOutputWithContext(ctx context.Context) CronJobListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobListMapOutput)
}

type CronJobListOutput struct {
	*pulumi.OutputState
}

func (CronJobListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CronJobList)(nil))
}

func (o CronJobListOutput) ToCronJobListOutput() CronJobListOutput {
	return o
}

func (o CronJobListOutput) ToCronJobListOutputWithContext(ctx context.Context) CronJobListOutput {
	return o
}

func (o CronJobListOutput) ToCronJobListPtrOutput() CronJobListPtrOutput {
	return o.ToCronJobListPtrOutputWithContext(context.Background())
}

func (o CronJobListOutput) ToCronJobListPtrOutputWithContext(ctx context.Context) CronJobListPtrOutput {
	return o.ApplyT(func(v CronJobList) *CronJobList {
		return &v
	}).(CronJobListPtrOutput)
}

type CronJobListPtrOutput struct {
	*pulumi.OutputState
}

func (CronJobListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CronJobList)(nil))
}

func (o CronJobListPtrOutput) ToCronJobListPtrOutput() CronJobListPtrOutput {
	return o
}

func (o CronJobListPtrOutput) ToCronJobListPtrOutputWithContext(ctx context.Context) CronJobListPtrOutput {
	return o
}

type CronJobListArrayOutput struct{ *pulumi.OutputState }

func (CronJobListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CronJobList)(nil))
}

func (o CronJobListArrayOutput) ToCronJobListArrayOutput() CronJobListArrayOutput {
	return o
}

func (o CronJobListArrayOutput) ToCronJobListArrayOutputWithContext(ctx context.Context) CronJobListArrayOutput {
	return o
}

func (o CronJobListArrayOutput) Index(i pulumi.IntInput) CronJobListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CronJobList {
		return vs[0].([]CronJobList)[vs[1].(int)]
	}).(CronJobListOutput)
}

type CronJobListMapOutput struct{ *pulumi.OutputState }

func (CronJobListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CronJobList)(nil))
}

func (o CronJobListMapOutput) ToCronJobListMapOutput() CronJobListMapOutput {
	return o
}

func (o CronJobListMapOutput) ToCronJobListMapOutputWithContext(ctx context.Context) CronJobListMapOutput {
	return o
}

func (o CronJobListMapOutput) MapIndex(k pulumi.StringInput) CronJobListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CronJobList {
		return vs[0].(map[string]CronJobList)[vs[1].(string)]
	}).(CronJobListOutput)
}

func init() {
	pulumi.RegisterOutputType(CronJobListOutput{})
	pulumi.RegisterOutputType(CronJobListPtrOutput{})
	pulumi.RegisterOutputType(CronJobListArrayOutput{})
	pulumi.RegisterOutputType(CronJobListMapOutput{})
}
