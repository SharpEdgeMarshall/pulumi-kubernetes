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

// ClusterRoleList is a collection of ClusterRoles
type ClusterRoleList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Items is a list of ClusterRoles
	Items ClusterRoleTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata.
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewClusterRoleList registers a new resource with the given unique name, arguments, and options.
func NewClusterRoleList(ctx *pulumi.Context,
	name string, args *ClusterRoleListArgs, opts ...pulumi.ResourceOption) (*ClusterRoleList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("rbac.authorization.k8s.io/v1")
	args.Kind = pulumi.StringPtr("ClusterRoleList")
	var resource ClusterRoleList
	err := ctx.RegisterResource("kubernetes:rbac.authorization.k8s.io/v1:ClusterRoleList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterRoleList gets an existing ClusterRoleList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterRoleList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterRoleListState, opts ...pulumi.ResourceOption) (*ClusterRoleList, error) {
	var resource ClusterRoleList
	err := ctx.ReadResource("kubernetes:rbac.authorization.k8s.io/v1:ClusterRoleList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterRoleList resources.
type clusterRoleListState struct {
}

type ClusterRoleListState struct {
}

func (ClusterRoleListState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterRoleListState)(nil)).Elem()
}

type clusterRoleListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of ClusterRoles
	Items []ClusterRoleType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ClusterRoleList resource.
type ClusterRoleListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is a list of ClusterRoles
	Items ClusterRoleTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ListMetaPtrInput
}

func (ClusterRoleListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterRoleListArgs)(nil)).Elem()
}

type ClusterRoleListInput interface {
	pulumi.Input

	ToClusterRoleListOutput() ClusterRoleListOutput
	ToClusterRoleListOutputWithContext(ctx context.Context) ClusterRoleListOutput
}

func (*ClusterRoleList) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterRoleList)(nil))
}

func (i *ClusterRoleList) ToClusterRoleListOutput() ClusterRoleListOutput {
	return i.ToClusterRoleListOutputWithContext(context.Background())
}

func (i *ClusterRoleList) ToClusterRoleListOutputWithContext(ctx context.Context) ClusterRoleListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleListOutput)
}

func (i *ClusterRoleList) ToClusterRoleListPtrOutput() ClusterRoleListPtrOutput {
	return i.ToClusterRoleListPtrOutputWithContext(context.Background())
}

func (i *ClusterRoleList) ToClusterRoleListPtrOutputWithContext(ctx context.Context) ClusterRoleListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleListPtrOutput)
}

type ClusterRoleListPtrInput interface {
	pulumi.Input

	ToClusterRoleListPtrOutput() ClusterRoleListPtrOutput
	ToClusterRoleListPtrOutputWithContext(ctx context.Context) ClusterRoleListPtrOutput
}

type clusterRoleListPtrType ClusterRoleListArgs

func (*clusterRoleListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterRoleList)(nil))
}

func (i *clusterRoleListPtrType) ToClusterRoleListPtrOutput() ClusterRoleListPtrOutput {
	return i.ToClusterRoleListPtrOutputWithContext(context.Background())
}

func (i *clusterRoleListPtrType) ToClusterRoleListPtrOutputWithContext(ctx context.Context) ClusterRoleListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleListPtrOutput)
}

// ClusterRoleListArrayInput is an input type that accepts ClusterRoleListArray and ClusterRoleListArrayOutput values.
// You can construct a concrete instance of `ClusterRoleListArrayInput` via:
//
//          ClusterRoleListArray{ ClusterRoleListArgs{...} }
type ClusterRoleListArrayInput interface {
	pulumi.Input

	ToClusterRoleListArrayOutput() ClusterRoleListArrayOutput
	ToClusterRoleListArrayOutputWithContext(context.Context) ClusterRoleListArrayOutput
}

type ClusterRoleListArray []ClusterRoleListInput

func (ClusterRoleListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterRoleList)(nil)).Elem()
}

func (i ClusterRoleListArray) ToClusterRoleListArrayOutput() ClusterRoleListArrayOutput {
	return i.ToClusterRoleListArrayOutputWithContext(context.Background())
}

func (i ClusterRoleListArray) ToClusterRoleListArrayOutputWithContext(ctx context.Context) ClusterRoleListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleListArrayOutput)
}

// ClusterRoleListMapInput is an input type that accepts ClusterRoleListMap and ClusterRoleListMapOutput values.
// You can construct a concrete instance of `ClusterRoleListMapInput` via:
//
//          ClusterRoleListMap{ "key": ClusterRoleListArgs{...} }
type ClusterRoleListMapInput interface {
	pulumi.Input

	ToClusterRoleListMapOutput() ClusterRoleListMapOutput
	ToClusterRoleListMapOutputWithContext(context.Context) ClusterRoleListMapOutput
}

type ClusterRoleListMap map[string]ClusterRoleListInput

func (ClusterRoleListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterRoleList)(nil)).Elem()
}

func (i ClusterRoleListMap) ToClusterRoleListMapOutput() ClusterRoleListMapOutput {
	return i.ToClusterRoleListMapOutputWithContext(context.Background())
}

func (i ClusterRoleListMap) ToClusterRoleListMapOutputWithContext(ctx context.Context) ClusterRoleListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleListMapOutput)
}

type ClusterRoleListOutput struct {
	*pulumi.OutputState
}

func (ClusterRoleListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterRoleList)(nil))
}

func (o ClusterRoleListOutput) ToClusterRoleListOutput() ClusterRoleListOutput {
	return o
}

func (o ClusterRoleListOutput) ToClusterRoleListOutputWithContext(ctx context.Context) ClusterRoleListOutput {
	return o
}

func (o ClusterRoleListOutput) ToClusterRoleListPtrOutput() ClusterRoleListPtrOutput {
	return o.ToClusterRoleListPtrOutputWithContext(context.Background())
}

func (o ClusterRoleListOutput) ToClusterRoleListPtrOutputWithContext(ctx context.Context) ClusterRoleListPtrOutput {
	return o.ApplyT(func(v ClusterRoleList) *ClusterRoleList {
		return &v
	}).(ClusterRoleListPtrOutput)
}

type ClusterRoleListPtrOutput struct {
	*pulumi.OutputState
}

func (ClusterRoleListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterRoleList)(nil))
}

func (o ClusterRoleListPtrOutput) ToClusterRoleListPtrOutput() ClusterRoleListPtrOutput {
	return o
}

func (o ClusterRoleListPtrOutput) ToClusterRoleListPtrOutputWithContext(ctx context.Context) ClusterRoleListPtrOutput {
	return o
}

type ClusterRoleListArrayOutput struct{ *pulumi.OutputState }

func (ClusterRoleListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterRoleList)(nil))
}

func (o ClusterRoleListArrayOutput) ToClusterRoleListArrayOutput() ClusterRoleListArrayOutput {
	return o
}

func (o ClusterRoleListArrayOutput) ToClusterRoleListArrayOutputWithContext(ctx context.Context) ClusterRoleListArrayOutput {
	return o
}

func (o ClusterRoleListArrayOutput) Index(i pulumi.IntInput) ClusterRoleListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterRoleList {
		return vs[0].([]ClusterRoleList)[vs[1].(int)]
	}).(ClusterRoleListOutput)
}

type ClusterRoleListMapOutput struct{ *pulumi.OutputState }

func (ClusterRoleListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ClusterRoleList)(nil))
}

func (o ClusterRoleListMapOutput) ToClusterRoleListMapOutput() ClusterRoleListMapOutput {
	return o
}

func (o ClusterRoleListMapOutput) ToClusterRoleListMapOutputWithContext(ctx context.Context) ClusterRoleListMapOutput {
	return o
}

func (o ClusterRoleListMapOutput) MapIndex(k pulumi.StringInput) ClusterRoleListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ClusterRoleList {
		return vs[0].(map[string]ClusterRoleList)[vs[1].(string)]
	}).(ClusterRoleListOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterRoleListOutput{})
	pulumi.RegisterOutputType(ClusterRoleListPtrOutput{})
	pulumi.RegisterOutputType(ClusterRoleListArrayOutput{})
	pulumi.RegisterOutputType(ClusterRoleListMapOutput{})
}
