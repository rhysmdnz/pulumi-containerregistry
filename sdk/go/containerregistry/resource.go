// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerregistry

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Resource struct {
	pulumi.CustomResourceState

	// Hash of the image tarball.
	ImageTarballHash pulumi.StringOutput `pulumi:"imageTarballHash"`
	// Image tarball thing.
	Image_tarball pulumi.AssetOrArchiveOutput `pulumi:"image_tarball"`
	// The tag to save the image to.
	RemoteTag pulumi.StringOutput `pulumi:"remoteTag"`
}

// NewResource registers a new resource with the given unique name, arguments, and options.
func NewResource(ctx *pulumi.Context,
	name string, args *ResourceArgs, opts ...pulumi.ResourceOption) (*Resource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImageTarballHash == nil {
		return nil, errors.New("invalid value for required argument 'ImageTarballHash'")
	}
	if args.Image_tarball == nil {
		return nil, errors.New("invalid value for required argument 'Image_tarball'")
	}
	if args.RemoteTag == nil {
		return nil, errors.New("invalid value for required argument 'RemoteTag'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Resource
	err := ctx.RegisterResource("containerregistry:index/resource:Resource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResource gets an existing Resource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceState, opts ...pulumi.ResourceOption) (*Resource, error) {
	var resource Resource
	err := ctx.ReadResource("containerregistry:index/resource:Resource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Resource resources.
type resourceState struct {
	// Hash of the image tarball.
	ImageTarballHash *string `pulumi:"imageTarballHash"`
	// Image tarball thing.
	Image_tarball pulumi.AssetOrArchive `pulumi:"image_tarball"`
	// The tag to save the image to.
	RemoteTag *string `pulumi:"remoteTag"`
}

type ResourceState struct {
	// Hash of the image tarball.
	ImageTarballHash pulumi.StringPtrInput
	// Image tarball thing.
	Image_tarball pulumi.AssetOrArchiveInput
	// The tag to save the image to.
	RemoteTag pulumi.StringPtrInput
}

func (ResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceState)(nil)).Elem()
}

type resourceArgs struct {
	// Hash of the image tarball.
	ImageTarballHash string `pulumi:"imageTarballHash"`
	// Image tarball thing.
	Image_tarball pulumi.AssetOrArchive `pulumi:"image_tarball"`
	// The tag to save the image to.
	RemoteTag string `pulumi:"remoteTag"`
}

// The set of arguments for constructing a Resource resource.
type ResourceArgs struct {
	// Hash of the image tarball.
	ImageTarballHash pulumi.StringInput
	// Image tarball thing.
	Image_tarball pulumi.AssetOrArchiveInput
	// The tag to save the image to.
	RemoteTag pulumi.StringInput
}

func (ResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceArgs)(nil)).Elem()
}

type ResourceInput interface {
	pulumi.Input

	ToResourceOutput() ResourceOutput
	ToResourceOutputWithContext(ctx context.Context) ResourceOutput
}

func (*Resource) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (i *Resource) ToResourceOutput() ResourceOutput {
	return i.ToResourceOutputWithContext(context.Background())
}

func (i *Resource) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceOutput)
}

// ResourceArrayInput is an input type that accepts ResourceArray and ResourceArrayOutput values.
// You can construct a concrete instance of `ResourceArrayInput` via:
//
//	ResourceArray{ ResourceArgs{...} }
type ResourceArrayInput interface {
	pulumi.Input

	ToResourceArrayOutput() ResourceArrayOutput
	ToResourceArrayOutputWithContext(context.Context) ResourceArrayOutput
}

type ResourceArray []ResourceInput

func (ResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Resource)(nil)).Elem()
}

func (i ResourceArray) ToResourceArrayOutput() ResourceArrayOutput {
	return i.ToResourceArrayOutputWithContext(context.Background())
}

func (i ResourceArray) ToResourceArrayOutputWithContext(ctx context.Context) ResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceArrayOutput)
}

// ResourceMapInput is an input type that accepts ResourceMap and ResourceMapOutput values.
// You can construct a concrete instance of `ResourceMapInput` via:
//
//	ResourceMap{ "key": ResourceArgs{...} }
type ResourceMapInput interface {
	pulumi.Input

	ToResourceMapOutput() ResourceMapOutput
	ToResourceMapOutputWithContext(context.Context) ResourceMapOutput
}

type ResourceMap map[string]ResourceInput

func (ResourceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Resource)(nil)).Elem()
}

func (i ResourceMap) ToResourceMapOutput() ResourceMapOutput {
	return i.ToResourceMapOutputWithContext(context.Background())
}

func (i ResourceMap) ToResourceMapOutputWithContext(ctx context.Context) ResourceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceMapOutput)
}

type ResourceOutput struct{ *pulumi.OutputState }

func (ResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (o ResourceOutput) ToResourceOutput() ResourceOutput {
	return o
}

func (o ResourceOutput) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return o
}

// Hash of the image tarball.
func (o ResourceOutput) ImageTarballHash() pulumi.StringOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringOutput { return v.ImageTarballHash }).(pulumi.StringOutput)
}

// Image tarball thing.
func (o ResourceOutput) Image_tarball() pulumi.AssetOrArchiveOutput {
	return o.ApplyT(func(v *Resource) pulumi.AssetOrArchiveOutput { return v.Image_tarball }).(pulumi.AssetOrArchiveOutput)
}

// The tag to save the image to.
func (o ResourceOutput) RemoteTag() pulumi.StringOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringOutput { return v.RemoteTag }).(pulumi.StringOutput)
}

type ResourceArrayOutput struct{ *pulumi.OutputState }

func (ResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Resource)(nil)).Elem()
}

func (o ResourceArrayOutput) ToResourceArrayOutput() ResourceArrayOutput {
	return o
}

func (o ResourceArrayOutput) ToResourceArrayOutputWithContext(ctx context.Context) ResourceArrayOutput {
	return o
}

func (o ResourceArrayOutput) Index(i pulumi.IntInput) ResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Resource {
		return vs[0].([]*Resource)[vs[1].(int)]
	}).(ResourceOutput)
}

type ResourceMapOutput struct{ *pulumi.OutputState }

func (ResourceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Resource)(nil)).Elem()
}

func (o ResourceMapOutput) ToResourceMapOutput() ResourceMapOutput {
	return o
}

func (o ResourceMapOutput) ToResourceMapOutputWithContext(ctx context.Context) ResourceMapOutput {
	return o
}

func (o ResourceMapOutput) MapIndex(k pulumi.StringInput) ResourceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Resource {
		return vs[0].(map[string]*Resource)[vs[1].(string)]
	}).(ResourceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceInput)(nil)).Elem(), &Resource{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceArrayInput)(nil)).Elem(), ResourceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceMapInput)(nil)).Elem(), ResourceMap{})
	pulumi.RegisterOutputType(ResourceOutput{})
	pulumi.RegisterOutputType(ResourceArrayOutput{})
	pulumi.RegisterOutputType(ResourceMapOutput{})
}
