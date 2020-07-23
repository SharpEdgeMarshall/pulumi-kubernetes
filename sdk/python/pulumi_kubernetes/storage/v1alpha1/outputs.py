# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ... import core as _core
from ... import meta as _meta

__all__ = [
    'VolumeAttachment',
    'VolumeAttachmentSource',
    'VolumeAttachmentSpec',
    'VolumeAttachmentStatus',
    'VolumeError',
]

@pulumi.output_type
class VolumeAttachment(dict):
    """
    VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node.

    VolumeAttachment objects are non-namespaced.
    """
    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[str]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        ...

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        ...

    @property
    @pulumi.getter
    def metadata(self) -> Optional['_meta.v1.outputs.ObjectMeta']:
        """
        Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        ...

    @property
    @pulumi.getter
    def spec(self) -> 'outputs.VolumeAttachmentSpec':
        """
        Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
        """
        ...

    @property
    @pulumi.getter
    def status(self) -> Optional['outputs.VolumeAttachmentStatus']:
        """
        Status of the VolumeAttachment request. Populated by the entity completing the attach or detach operation, i.e. the external-attacher.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class VolumeAttachmentSource(dict):
    """
    VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
    """
    @property
    @pulumi.getter(name="inlineVolumeSpec")
    def inline_volume_spec(self) -> Optional['_core.v1.outputs.PersistentVolumeSpec']:
        """
        inlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's inline VolumeSource to a PersistentVolumeSpec. This field is alpha-level and is only honored by servers that enabled the CSIMigration feature.
        """
        ...

    @property
    @pulumi.getter(name="persistentVolumeName")
    def persistent_volume_name(self) -> Optional[str]:
        """
        Name of the persistent volume to attach.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class VolumeAttachmentSpec(dict):
    """
    VolumeAttachmentSpec is the specification of a VolumeAttachment request.
    """
    @property
    @pulumi.getter
    def attacher(self) -> str:
        """
        Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
        """
        ...

    @property
    @pulumi.getter(name="nodeName")
    def node_name(self) -> str:
        """
        The node that the volume should be attached to.
        """
        ...

    @property
    @pulumi.getter
    def source(self) -> 'outputs.VolumeAttachmentSource':
        """
        Source represents the volume that should be attached.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class VolumeAttachmentStatus(dict):
    """
    VolumeAttachmentStatus is the status of a VolumeAttachment request.
    """
    @property
    @pulumi.getter(name="attachError")
    def attach_error(self) -> Optional['outputs.VolumeError']:
        """
        The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        """
        ...

    @property
    @pulumi.getter
    def attached(self) -> bool:
        """
        Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        """
        ...

    @property
    @pulumi.getter(name="attachmentMetadata")
    def attachment_metadata(self) -> Optional[Mapping[str, str]]:
        """
        Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        """
        ...

    @property
    @pulumi.getter(name="detachError")
    def detach_error(self) -> Optional['outputs.VolumeError']:
        """
        The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class VolumeError(dict):
    """
    VolumeError captures an error encountered during a volume operation.
    """
    @property
    @pulumi.getter
    def message(self) -> Optional[str]:
        """
        String detailing the error encountered during Attach or Detach operation. This string maybe logged, so it should not contain sensitive information.
        """
        ...

    @property
    @pulumi.getter
    def time(self) -> Optional[str]:
        """
        Time the error was encountered.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

