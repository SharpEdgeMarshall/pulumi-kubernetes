// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deployment enables declarative updates for Pods and ReplicaSets.
//
// This resource waits until its status is ready before registering success
// for create/update, and populating output properties from the current state of the resource.
// The following conditions are used to determine whether the resource creation has
// succeeded or failed:
//
// 1. The Deployment has begun to be updated by the Deployment controller. If the current
//    generation of the Deployment is > 1, then this means that the current generation must
//    be different from the generation reported by the last outputs.
// 2. There exists a ReplicaSet whose revision is equal to the current revision of the
//    Deployment.
// 3. The Deployment's '.status.conditions' has a status of type 'Available' whose 'status'
//    member is set to 'True'.
// 4. If the Deployment has generation > 1, then '.status.conditions' has a status of type
//    'Progressing', whose 'status' member is set to 'True', and whose 'reason' is
//    'NewReplicaSetAvailable'. For generation <= 1, this status field does not exist,
//    because it doesn't do a rollout (i.e., it simply creates the Deployment and
//    corresponding ReplicaSet), and therefore there is no rollout to mark as 'Progressing'.
//
// If the Deployment has not reached a Ready state after 10 minutes, it will
// time out and mark the resource update as Failed. You can override the default timeout value
// by setting the 'customTimeouts' option on the resource.
//
// Deprecated: extensions/v1beta1/Deployment is deprecated by apps/v1/Deployment and not supported by Kubernetes v1.16+ clusters.
type Deployment struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object metadata.
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Specification of the desired behavior of the Deployment.
	Spec DeploymentSpecPtrOutput `pulumi:"spec"`
	// Most recently observed status of the Deployment.
	Status DeploymentStatusPtrOutput `pulumi:"status"`
}

// NewDeployment registers a new resource with the given unique name, arguments, and options.
func NewDeployment(ctx *pulumi.Context,
	name string, args *DeploymentArgs, opts ...pulumi.ResourceOption) (*Deployment, error) {
	if args == nil {
		args = &DeploymentArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("extensions/v1beta1")
	args.Kind = pulumi.StringPtr("Deployment")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:apps/v1:Deployment"),
		},
		{
			Type: pulumi.String("kubernetes:apps/v1beta1:Deployment"),
		},
		{
			Type: pulumi.String("kubernetes:apps/v1beta2:Deployment"),
		},
	})
	opts = append(opts, aliases)
	var resource Deployment
	err := ctx.RegisterResource("kubernetes:extensions/v1beta1:Deployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeployment gets an existing Deployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentState, opts ...pulumi.ResourceOption) (*Deployment, error) {
	var resource Deployment
	err := ctx.ReadResource("kubernetes:extensions/v1beta1:Deployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Deployment resources.
type deploymentState struct {
}

type DeploymentState struct {
}

func (DeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentState)(nil)).Elem()
}

type deploymentArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Specification of the desired behavior of the Deployment.
	Spec *DeploymentSpec `pulumi:"spec"`
}

// The set of arguments for constructing a Deployment resource.
type DeploymentArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata.
	Metadata metav1.ObjectMetaPtrInput
	// Specification of the desired behavior of the Deployment.
	Spec DeploymentSpecPtrInput
}

func (DeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentArgs)(nil)).Elem()
}

type DeploymentInput interface {
	pulumi.Input

	ToDeploymentOutput() DeploymentOutput
	ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput
}

func (*Deployment) ElementType() reflect.Type {
	return reflect.TypeOf((*Deployment)(nil))
}

func (i *Deployment) ToDeploymentOutput() DeploymentOutput {
	return i.ToDeploymentOutputWithContext(context.Background())
}

func (i *Deployment) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentOutput)
}

func (i *Deployment) ToDeploymentPtrOutput() DeploymentPtrOutput {
	return i.ToDeploymentPtrOutputWithContext(context.Background())
}

func (i *Deployment) ToDeploymentPtrOutputWithContext(ctx context.Context) DeploymentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPtrOutput)
}

type DeploymentPtrInput interface {
	pulumi.Input

	ToDeploymentPtrOutput() DeploymentPtrOutput
	ToDeploymentPtrOutputWithContext(ctx context.Context) DeploymentPtrOutput
}

type deploymentPtrType DeploymentArgs

func (*deploymentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil))
}

func (i *deploymentPtrType) ToDeploymentPtrOutput() DeploymentPtrOutput {
	return i.ToDeploymentPtrOutputWithContext(context.Background())
}

func (i *deploymentPtrType) ToDeploymentPtrOutputWithContext(ctx context.Context) DeploymentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPtrOutput)
}

// DeploymentArrayInput is an input type that accepts DeploymentArray and DeploymentArrayOutput values.
// You can construct a concrete instance of `DeploymentArrayInput` via:
//
//          DeploymentArray{ DeploymentArgs{...} }
type DeploymentArrayInput interface {
	pulumi.Input

	ToDeploymentArrayOutput() DeploymentArrayOutput
	ToDeploymentArrayOutputWithContext(context.Context) DeploymentArrayOutput
}

type DeploymentArray []DeploymentInput

func (DeploymentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Deployment)(nil)).Elem()
}

func (i DeploymentArray) ToDeploymentArrayOutput() DeploymentArrayOutput {
	return i.ToDeploymentArrayOutputWithContext(context.Background())
}

func (i DeploymentArray) ToDeploymentArrayOutputWithContext(ctx context.Context) DeploymentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentArrayOutput)
}

// DeploymentMapInput is an input type that accepts DeploymentMap and DeploymentMapOutput values.
// You can construct a concrete instance of `DeploymentMapInput` via:
//
//          DeploymentMap{ "key": DeploymentArgs{...} }
type DeploymentMapInput interface {
	pulumi.Input

	ToDeploymentMapOutput() DeploymentMapOutput
	ToDeploymentMapOutputWithContext(context.Context) DeploymentMapOutput
}

type DeploymentMap map[string]DeploymentInput

func (DeploymentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Deployment)(nil)).Elem()
}

func (i DeploymentMap) ToDeploymentMapOutput() DeploymentMapOutput {
	return i.ToDeploymentMapOutputWithContext(context.Background())
}

func (i DeploymentMap) ToDeploymentMapOutputWithContext(ctx context.Context) DeploymentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentMapOutput)
}

type DeploymentOutput struct {
	*pulumi.OutputState
}

func (DeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Deployment)(nil))
}

func (o DeploymentOutput) ToDeploymentOutput() DeploymentOutput {
	return o
}

func (o DeploymentOutput) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return o
}

func (o DeploymentOutput) ToDeploymentPtrOutput() DeploymentPtrOutput {
	return o.ToDeploymentPtrOutputWithContext(context.Background())
}

func (o DeploymentOutput) ToDeploymentPtrOutputWithContext(ctx context.Context) DeploymentPtrOutput {
	return o.ApplyT(func(v Deployment) *Deployment {
		return &v
	}).(DeploymentPtrOutput)
}

type DeploymentPtrOutput struct {
	*pulumi.OutputState
}

func (DeploymentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil))
}

func (o DeploymentPtrOutput) ToDeploymentPtrOutput() DeploymentPtrOutput {
	return o
}

func (o DeploymentPtrOutput) ToDeploymentPtrOutputWithContext(ctx context.Context) DeploymentPtrOutput {
	return o
}

type DeploymentArrayOutput struct{ *pulumi.OutputState }

func (DeploymentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Deployment)(nil))
}

func (o DeploymentArrayOutput) ToDeploymentArrayOutput() DeploymentArrayOutput {
	return o
}

func (o DeploymentArrayOutput) ToDeploymentArrayOutputWithContext(ctx context.Context) DeploymentArrayOutput {
	return o
}

func (o DeploymentArrayOutput) Index(i pulumi.IntInput) DeploymentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Deployment {
		return vs[0].([]Deployment)[vs[1].(int)]
	}).(DeploymentOutput)
}

type DeploymentMapOutput struct{ *pulumi.OutputState }

func (DeploymentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Deployment)(nil))
}

func (o DeploymentMapOutput) ToDeploymentMapOutput() DeploymentMapOutput {
	return o
}

func (o DeploymentMapOutput) ToDeploymentMapOutputWithContext(ctx context.Context) DeploymentMapOutput {
	return o
}

func (o DeploymentMapOutput) MapIndex(k pulumi.StringInput) DeploymentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Deployment {
		return vs[0].(map[string]Deployment)[vs[1].(string)]
	}).(DeploymentOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentOutput{})
	pulumi.RegisterOutputType(DeploymentPtrOutput{})
	pulumi.RegisterOutputType(DeploymentArrayOutput{})
	pulumi.RegisterOutputType(DeploymentMapOutput{})
}
