# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'InstanceBootDiskInitializeParams',
]

@pulumi.output_type
class InstanceBootDiskInitializeParams(dict):
    def __init__(__self__, *,
                 image: Optional[str] = None):
        """
        :param str image: The image from which to initialize this disk. This can be
               one of: the image's `self_link`, `projects/{project}/global/images/{image}`,
               `projects/{project}/global/images/family/{family}`, `global/images/{image}`,
               `global/images/family/{family}`, `family/{family}`, `{project}/{family}`,
               `{project}/{image}`, `{family}`, or `{image}`. If referred by family, the
               images names must include the family name. If they don't, use the
               [gcp.compute.Image data source](https://www.terraform.io/docs/providers/google/d/compute_image.html).
               For instance, the image `centos-6-v20180104` includes its family name `centos-6`.
               These images can be referred by family name here.
        """
        InstanceBootDiskInitializeParams._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            image=image,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             image: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'image' in kwargs:
            image = kwargs['image']

        if image is not None:
            _setter("image", image)

    @property
    @pulumi.getter
    def image(self) -> Optional[str]:
        """
        The image from which to initialize this disk. This can be
        one of: the image's `self_link`, `projects/{project}/global/images/{image}`,
        `projects/{project}/global/images/family/{family}`, `global/images/{image}`,
        `global/images/family/{family}`, `family/{family}`, `{project}/{family}`,
        `{project}/{image}`, `{family}`, or `{image}`. If referred by family, the
        images names must include the family name. If they don't, use the
        [gcp.compute.Image data source](https://www.terraform.io/docs/providers/google/d/compute_image.html).
        For instance, the image `centos-6-v20180104` includes its family name `centos-6`.
        These images can be referred by family name here.
        """
        return pulumi.get(self, "image")

