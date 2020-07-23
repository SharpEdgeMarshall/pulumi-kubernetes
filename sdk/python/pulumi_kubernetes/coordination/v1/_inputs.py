# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from ... import meta as _meta

__all__ = [
    'LeaseArgs',
    'LeaseSpecArgs',
]

@pulumi.input_type
class LeaseArgs:
    def __init__(__self__, *,
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None,
                 spec: Optional[pulumi.Input['LeaseSpecArgs']] = None):
        """
        Lease defines a lease concept.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input['_meta.v1.ObjectMetaArgs'] metadata: More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input['LeaseSpecArgs'] spec: Specification of the Lease. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        pulumi.set(__self__, "apiVersion", 'coordination.k8s.io/v1')
        pulumi.set(__self__, "kind", 'Lease')
        pulumi.set(__self__, "metadata", metadata)
        pulumi.set(__self__, "spec", spec)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[pulumi.Input[str]]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        ...

    @api_version.setter
    def api_version(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        ...

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]:
        """
        More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        ...

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]):
        ...

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['LeaseSpecArgs']]:
        """
        Specification of the Lease. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        """
        ...

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['LeaseSpecArgs']]):
        ...


@pulumi.input_type
class LeaseSpecArgs:
    def __init__(__self__, *,
                 acquire_time: Optional[pulumi.Input[str]] = None,
                 holder_identity: Optional[pulumi.Input[str]] = None,
                 lease_duration_seconds: Optional[pulumi.Input[float]] = None,
                 lease_transitions: Optional[pulumi.Input[float]] = None,
                 renew_time: Optional[pulumi.Input[str]] = None):
        """
        LeaseSpec is a specification of a Lease.
        :param pulumi.Input[str] acquire_time: acquireTime is a time when the current lease was acquired.
        :param pulumi.Input[str] holder_identity: holderIdentity contains the identity of the holder of a current lease.
        :param pulumi.Input[float] lease_duration_seconds: leaseDurationSeconds is a duration that candidates for a lease need to wait to force acquire it. This is measure against time of last observed RenewTime.
        :param pulumi.Input[float] lease_transitions: leaseTransitions is the number of transitions of a lease between holders.
        :param pulumi.Input[str] renew_time: renewTime is a time when the current holder of a lease has last updated the lease.
        """
        pulumi.set(__self__, "acquireTime", acquire_time)
        pulumi.set(__self__, "holderIdentity", holder_identity)
        pulumi.set(__self__, "leaseDurationSeconds", lease_duration_seconds)
        pulumi.set(__self__, "leaseTransitions", lease_transitions)
        pulumi.set(__self__, "renewTime", renew_time)

    @property
    @pulumi.getter(name="acquireTime")
    def acquire_time(self) -> Optional[pulumi.Input[str]]:
        """
        acquireTime is a time when the current lease was acquired.
        """
        ...

    @acquire_time.setter
    def acquire_time(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter(name="holderIdentity")
    def holder_identity(self) -> Optional[pulumi.Input[str]]:
        """
        holderIdentity contains the identity of the holder of a current lease.
        """
        ...

    @holder_identity.setter
    def holder_identity(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter(name="leaseDurationSeconds")
    def lease_duration_seconds(self) -> Optional[pulumi.Input[float]]:
        """
        leaseDurationSeconds is a duration that candidates for a lease need to wait to force acquire it. This is measure against time of last observed RenewTime.
        """
        ...

    @lease_duration_seconds.setter
    def lease_duration_seconds(self, value: Optional[pulumi.Input[float]]):
        ...

    @property
    @pulumi.getter(name="leaseTransitions")
    def lease_transitions(self) -> Optional[pulumi.Input[float]]:
        """
        leaseTransitions is the number of transitions of a lease between holders.
        """
        ...

    @lease_transitions.setter
    def lease_transitions(self, value: Optional[pulumi.Input[float]]):
        ...

    @property
    @pulumi.getter(name="renewTime")
    def renew_time(self) -> Optional[pulumi.Input[str]]:
        """
        renewTime is a time when the current holder of a lease has last updated the lease.
        """
        ...

    @renew_time.setter
    def renew_time(self, value: Optional[pulumi.Input[str]]):
        ...

