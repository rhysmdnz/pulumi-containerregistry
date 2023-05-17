# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ResourceArgs', 'Resource']

@pulumi.input_type
class ResourceArgs:
    def __init__(__self__, *,
                 image_tarball: pulumi.Input[str],
                 image_tarball_hash: pulumi.Input[str],
                 remote_tag: pulumi.Input[str]):
        """
        The set of arguments for constructing a Resource resource.
        :param pulumi.Input[str] image_tarball: Image tarball thing.
        :param pulumi.Input[str] image_tarball_hash: Hash of the image tarball.
        :param pulumi.Input[str] remote_tag: The tag to save the image to.
        """
        pulumi.set(__self__, "image_tarball", image_tarball)
        pulumi.set(__self__, "image_tarball_hash", image_tarball_hash)
        pulumi.set(__self__, "remote_tag", remote_tag)

    @property
    @pulumi.getter(name="imageTarball")
    def image_tarball(self) -> pulumi.Input[str]:
        """
        Image tarball thing.
        """
        return pulumi.get(self, "image_tarball")

    @image_tarball.setter
    def image_tarball(self, value: pulumi.Input[str]):
        pulumi.set(self, "image_tarball", value)

    @property
    @pulumi.getter(name="imageTarballHash")
    def image_tarball_hash(self) -> pulumi.Input[str]:
        """
        Hash of the image tarball.
        """
        return pulumi.get(self, "image_tarball_hash")

    @image_tarball_hash.setter
    def image_tarball_hash(self, value: pulumi.Input[str]):
        pulumi.set(self, "image_tarball_hash", value)

    @property
    @pulumi.getter(name="remoteTag")
    def remote_tag(self) -> pulumi.Input[str]:
        """
        The tag to save the image to.
        """
        return pulumi.get(self, "remote_tag")

    @remote_tag.setter
    def remote_tag(self, value: pulumi.Input[str]):
        pulumi.set(self, "remote_tag", value)


@pulumi.input_type
class _ResourceState:
    def __init__(__self__, *,
                 image_tarball: Optional[pulumi.Input[str]] = None,
                 image_tarball_hash: Optional[pulumi.Input[str]] = None,
                 remote_tag: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Resource resources.
        :param pulumi.Input[str] image_tarball: Image tarball thing.
        :param pulumi.Input[str] image_tarball_hash: Hash of the image tarball.
        :param pulumi.Input[str] remote_tag: The tag to save the image to.
        """
        if image_tarball is not None:
            pulumi.set(__self__, "image_tarball", image_tarball)
        if image_tarball_hash is not None:
            pulumi.set(__self__, "image_tarball_hash", image_tarball_hash)
        if remote_tag is not None:
            pulumi.set(__self__, "remote_tag", remote_tag)

    @property
    @pulumi.getter(name="imageTarball")
    def image_tarball(self) -> Optional[pulumi.Input[str]]:
        """
        Image tarball thing.
        """
        return pulumi.get(self, "image_tarball")

    @image_tarball.setter
    def image_tarball(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_tarball", value)

    @property
    @pulumi.getter(name="imageTarballHash")
    def image_tarball_hash(self) -> Optional[pulumi.Input[str]]:
        """
        Hash of the image tarball.
        """
        return pulumi.get(self, "image_tarball_hash")

    @image_tarball_hash.setter
    def image_tarball_hash(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_tarball_hash", value)

    @property
    @pulumi.getter(name="remoteTag")
    def remote_tag(self) -> Optional[pulumi.Input[str]]:
        """
        The tag to save the image to.
        """
        return pulumi.get(self, "remote_tag")

    @remote_tag.setter
    def remote_tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_tag", value)


class Resource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 image_tarball: Optional[pulumi.Input[str]] = None,
                 image_tarball_hash: Optional[pulumi.Input[str]] = None,
                 remote_tag: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Resource resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] image_tarball: Image tarball thing.
        :param pulumi.Input[str] image_tarball_hash: Hash of the image tarball.
        :param pulumi.Input[str] remote_tag: The tag to save the image to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResourceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Resource resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ResourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 image_tarball: Optional[pulumi.Input[str]] = None,
                 image_tarball_hash: Optional[pulumi.Input[str]] = None,
                 remote_tag: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResourceArgs.__new__(ResourceArgs)

            if image_tarball is None and not opts.urn:
                raise TypeError("Missing required property 'image_tarball'")
            __props__.__dict__["image_tarball"] = image_tarball
            if image_tarball_hash is None and not opts.urn:
                raise TypeError("Missing required property 'image_tarball_hash'")
            __props__.__dict__["image_tarball_hash"] = image_tarball_hash
            if remote_tag is None and not opts.urn:
                raise TypeError("Missing required property 'remote_tag'")
            __props__.__dict__["remote_tag"] = remote_tag
        super(Resource, __self__).__init__(
            'containerregistry:index/resource:Resource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            image_tarball: Optional[pulumi.Input[str]] = None,
            image_tarball_hash: Optional[pulumi.Input[str]] = None,
            remote_tag: Optional[pulumi.Input[str]] = None) -> 'Resource':
        """
        Get an existing Resource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] image_tarball: Image tarball thing.
        :param pulumi.Input[str] image_tarball_hash: Hash of the image tarball.
        :param pulumi.Input[str] remote_tag: The tag to save the image to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ResourceState.__new__(_ResourceState)

        __props__.__dict__["image_tarball"] = image_tarball
        __props__.__dict__["image_tarball_hash"] = image_tarball_hash
        __props__.__dict__["remote_tag"] = remote_tag
        return Resource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="imageTarball")
    def image_tarball(self) -> pulumi.Output[str]:
        """
        Image tarball thing.
        """
        return pulumi.get(self, "image_tarball")

    @property
    @pulumi.getter(name="imageTarballHash")
    def image_tarball_hash(self) -> pulumi.Output[str]:
        """
        Hash of the image tarball.
        """
        return pulumi.get(self, "image_tarball_hash")

    @property
    @pulumi.getter(name="remoteTag")
    def remote_tag(self) -> pulumi.Output[str]:
        """
        The tag to save the image to.
        """
        return pulumi.get(self, "remote_tag")

