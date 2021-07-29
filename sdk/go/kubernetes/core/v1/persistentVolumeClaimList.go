// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// PersistentVolumeClaimList is a list of PersistentVolumeClaim items.
type PersistentVolumeClaimList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// A list of persistent volume claims. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Items PersistentVolumeClaimTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewPersistentVolumeClaimList registers a new resource with the given unique name, arguments, and options.
func NewPersistentVolumeClaimList(ctx *pulumi.Context,
	name string, args *PersistentVolumeClaimListArgs, opts ...pulumi.ResourceOption) (*PersistentVolumeClaimList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("PersistentVolumeClaimList")
	var resource PersistentVolumeClaimList
	err := ctx.RegisterResource("kubernetes:core/v1:PersistentVolumeClaimList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPersistentVolumeClaimList gets an existing PersistentVolumeClaimList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPersistentVolumeClaimList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PersistentVolumeClaimListState, opts ...pulumi.ResourceOption) (*PersistentVolumeClaimList, error) {
	var resource PersistentVolumeClaimList
	err := ctx.ReadResource("kubernetes:core/v1:PersistentVolumeClaimList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PersistentVolumeClaimList resources.
type persistentVolumeClaimListState struct {
}

type PersistentVolumeClaimListState struct {
}

func (PersistentVolumeClaimListState) ElementType() reflect.Type {
	return reflect.TypeOf((*persistentVolumeClaimListState)(nil)).Elem()
}

type persistentVolumeClaimListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// A list of persistent volume claims. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Items []PersistentVolumeClaimType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a PersistentVolumeClaimList resource.
type PersistentVolumeClaimListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// A list of persistent volume claims. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Items PersistentVolumeClaimTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (PersistentVolumeClaimListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*persistentVolumeClaimListArgs)(nil)).Elem()
}

type PersistentVolumeClaimListInput interface {
	pulumi.Input

	ToPersistentVolumeClaimListOutput() PersistentVolumeClaimListOutput
	ToPersistentVolumeClaimListOutputWithContext(ctx context.Context) PersistentVolumeClaimListOutput
}

func (*PersistentVolumeClaimList) ElementType() reflect.Type {
	return reflect.TypeOf((*PersistentVolumeClaimList)(nil))
}

func (i *PersistentVolumeClaimList) ToPersistentVolumeClaimListOutput() PersistentVolumeClaimListOutput {
	return i.ToPersistentVolumeClaimListOutputWithContext(context.Background())
}

func (i *PersistentVolumeClaimList) ToPersistentVolumeClaimListOutputWithContext(ctx context.Context) PersistentVolumeClaimListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumeClaimListOutput)
}

func (i *PersistentVolumeClaimList) ToPersistentVolumeClaimListPtrOutput() PersistentVolumeClaimListPtrOutput {
	return i.ToPersistentVolumeClaimListPtrOutputWithContext(context.Background())
}

func (i *PersistentVolumeClaimList) ToPersistentVolumeClaimListPtrOutputWithContext(ctx context.Context) PersistentVolumeClaimListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumeClaimListPtrOutput)
}

type PersistentVolumeClaimListPtrInput interface {
	pulumi.Input

	ToPersistentVolumeClaimListPtrOutput() PersistentVolumeClaimListPtrOutput
	ToPersistentVolumeClaimListPtrOutputWithContext(ctx context.Context) PersistentVolumeClaimListPtrOutput
}

type persistentVolumeClaimListPtrType PersistentVolumeClaimListArgs

func (*persistentVolumeClaimListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistentVolumeClaimList)(nil))
}

func (i *persistentVolumeClaimListPtrType) ToPersistentVolumeClaimListPtrOutput() PersistentVolumeClaimListPtrOutput {
	return i.ToPersistentVolumeClaimListPtrOutputWithContext(context.Background())
}

func (i *persistentVolumeClaimListPtrType) ToPersistentVolumeClaimListPtrOutputWithContext(ctx context.Context) PersistentVolumeClaimListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumeClaimListPtrOutput)
}

// PersistentVolumeClaimListArrayInput is an input type that accepts PersistentVolumeClaimListArray and PersistentVolumeClaimListArrayOutput values.
// You can construct a concrete instance of `PersistentVolumeClaimListArrayInput` via:
//
//          PersistentVolumeClaimListArray{ PersistentVolumeClaimListArgs{...} }
type PersistentVolumeClaimListArrayInput interface {
	pulumi.Input

	ToPersistentVolumeClaimListArrayOutput() PersistentVolumeClaimListArrayOutput
	ToPersistentVolumeClaimListArrayOutputWithContext(context.Context) PersistentVolumeClaimListArrayOutput
}

type PersistentVolumeClaimListArray []PersistentVolumeClaimListInput

func (PersistentVolumeClaimListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PersistentVolumeClaimList)(nil)).Elem()
}

func (i PersistentVolumeClaimListArray) ToPersistentVolumeClaimListArrayOutput() PersistentVolumeClaimListArrayOutput {
	return i.ToPersistentVolumeClaimListArrayOutputWithContext(context.Background())
}

func (i PersistentVolumeClaimListArray) ToPersistentVolumeClaimListArrayOutputWithContext(ctx context.Context) PersistentVolumeClaimListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumeClaimListArrayOutput)
}

// PersistentVolumeClaimListMapInput is an input type that accepts PersistentVolumeClaimListMap and PersistentVolumeClaimListMapOutput values.
// You can construct a concrete instance of `PersistentVolumeClaimListMapInput` via:
//
//          PersistentVolumeClaimListMap{ "key": PersistentVolumeClaimListArgs{...} }
type PersistentVolumeClaimListMapInput interface {
	pulumi.Input

	ToPersistentVolumeClaimListMapOutput() PersistentVolumeClaimListMapOutput
	ToPersistentVolumeClaimListMapOutputWithContext(context.Context) PersistentVolumeClaimListMapOutput
}

type PersistentVolumeClaimListMap map[string]PersistentVolumeClaimListInput

func (PersistentVolumeClaimListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PersistentVolumeClaimList)(nil)).Elem()
}

func (i PersistentVolumeClaimListMap) ToPersistentVolumeClaimListMapOutput() PersistentVolumeClaimListMapOutput {
	return i.ToPersistentVolumeClaimListMapOutputWithContext(context.Background())
}

func (i PersistentVolumeClaimListMap) ToPersistentVolumeClaimListMapOutputWithContext(ctx context.Context) PersistentVolumeClaimListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentVolumeClaimListMapOutput)
}

type PersistentVolumeClaimListOutput struct {
	*pulumi.OutputState
}

func (PersistentVolumeClaimListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PersistentVolumeClaimList)(nil))
}

func (o PersistentVolumeClaimListOutput) ToPersistentVolumeClaimListOutput() PersistentVolumeClaimListOutput {
	return o
}

func (o PersistentVolumeClaimListOutput) ToPersistentVolumeClaimListOutputWithContext(ctx context.Context) PersistentVolumeClaimListOutput {
	return o
}

func (o PersistentVolumeClaimListOutput) ToPersistentVolumeClaimListPtrOutput() PersistentVolumeClaimListPtrOutput {
	return o.ToPersistentVolumeClaimListPtrOutputWithContext(context.Background())
}

func (o PersistentVolumeClaimListOutput) ToPersistentVolumeClaimListPtrOutputWithContext(ctx context.Context) PersistentVolumeClaimListPtrOutput {
	return o.ApplyT(func(v PersistentVolumeClaimList) *PersistentVolumeClaimList {
		return &v
	}).(PersistentVolumeClaimListPtrOutput)
}

type PersistentVolumeClaimListPtrOutput struct {
	*pulumi.OutputState
}

func (PersistentVolumeClaimListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistentVolumeClaimList)(nil))
}

func (o PersistentVolumeClaimListPtrOutput) ToPersistentVolumeClaimListPtrOutput() PersistentVolumeClaimListPtrOutput {
	return o
}

func (o PersistentVolumeClaimListPtrOutput) ToPersistentVolumeClaimListPtrOutputWithContext(ctx context.Context) PersistentVolumeClaimListPtrOutput {
	return o
}

type PersistentVolumeClaimListArrayOutput struct{ *pulumi.OutputState }

func (PersistentVolumeClaimListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PersistentVolumeClaimList)(nil))
}

func (o PersistentVolumeClaimListArrayOutput) ToPersistentVolumeClaimListArrayOutput() PersistentVolumeClaimListArrayOutput {
	return o
}

func (o PersistentVolumeClaimListArrayOutput) ToPersistentVolumeClaimListArrayOutputWithContext(ctx context.Context) PersistentVolumeClaimListArrayOutput {
	return o
}

func (o PersistentVolumeClaimListArrayOutput) Index(i pulumi.IntInput) PersistentVolumeClaimListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PersistentVolumeClaimList {
		return vs[0].([]PersistentVolumeClaimList)[vs[1].(int)]
	}).(PersistentVolumeClaimListOutput)
}

type PersistentVolumeClaimListMapOutput struct{ *pulumi.OutputState }

func (PersistentVolumeClaimListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PersistentVolumeClaimList)(nil))
}

func (o PersistentVolumeClaimListMapOutput) ToPersistentVolumeClaimListMapOutput() PersistentVolumeClaimListMapOutput {
	return o
}

func (o PersistentVolumeClaimListMapOutput) ToPersistentVolumeClaimListMapOutputWithContext(ctx context.Context) PersistentVolumeClaimListMapOutput {
	return o
}

func (o PersistentVolumeClaimListMapOutput) MapIndex(k pulumi.StringInput) PersistentVolumeClaimListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PersistentVolumeClaimList {
		return vs[0].(map[string]PersistentVolumeClaimList)[vs[1].(string)]
	}).(PersistentVolumeClaimListOutput)
}

func init() {
	pulumi.RegisterOutputType(PersistentVolumeClaimListOutput{})
	pulumi.RegisterOutputType(PersistentVolumeClaimListPtrOutput{})
	pulumi.RegisterOutputType(PersistentVolumeClaimListArrayOutput{})
	pulumi.RegisterOutputType(PersistentVolumeClaimListMapOutput{})
}
