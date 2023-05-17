// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Resource extends pulumi.CustomResource {
    /**
     * Get an existing Resource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourceState, opts?: pulumi.CustomResourceOptions): Resource {
        return new Resource(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'containerregistry:index/resource:Resource';

    /**
     * Returns true if the given object is an instance of Resource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Resource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Resource.__pulumiType;
    }

    /**
     * Hash of the image tarball.
     */
    public readonly imageTarballHash!: pulumi.Output<string>;
    /**
     * Image tarball thing.
     */
    public readonly image_tarball!: pulumi.Output<pulumi.asset.Asset | pulumi.asset.Archive>;
    /**
     * The tag to save the image to.
     */
    public readonly remoteTag!: pulumi.Output<string>;

    /**
     * Create a Resource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourceArgs | ResourceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResourceState | undefined;
            resourceInputs["imageTarballHash"] = state ? state.imageTarballHash : undefined;
            resourceInputs["image_tarball"] = state ? state.image_tarball : undefined;
            resourceInputs["remoteTag"] = state ? state.remoteTag : undefined;
        } else {
            const args = argsOrState as ResourceArgs | undefined;
            if ((!args || args.imageTarballHash === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageTarballHash'");
            }
            if ((!args || args.image_tarball === undefined) && !opts.urn) {
                throw new Error("Missing required property 'image_tarball'");
            }
            if ((!args || args.remoteTag === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remoteTag'");
            }
            resourceInputs["imageTarballHash"] = args ? args.imageTarballHash : undefined;
            resourceInputs["image_tarball"] = args ? args.image_tarball : undefined;
            resourceInputs["remoteTag"] = args ? args.remoteTag : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Resource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Resource resources.
 */
export interface ResourceState {
    /**
     * Hash of the image tarball.
     */
    imageTarballHash?: pulumi.Input<string>;
    /**
     * Image tarball thing.
     */
    image_tarball?: pulumi.Input<pulumi.asset.Asset | pulumi.asset.Archive>;
    /**
     * The tag to save the image to.
     */
    remoteTag?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Resource resource.
 */
export interface ResourceArgs {
    /**
     * Hash of the image tarball.
     */
    imageTarballHash: pulumi.Input<string>;
    /**
     * Image tarball thing.
     */
    image_tarball: pulumi.Input<pulumi.asset.Asset | pulumi.asset.Archive>;
    /**
     * The tag to save the image to.
     */
    remoteTag: pulumi.Input<string>;
}
