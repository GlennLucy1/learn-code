from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class DetectRequest(_message.Message):
    __slots__ = ("b64_img",)
    B64_IMG_FIELD_NUMBER: _ClassVar[int]
    b64_img: str
    def __init__(self, b64_img: _Optional[str] = ...) -> None: ...

class DetectResponse(_message.Message):
    __slots__ = ("response",)
    RESPONSE_FIELD_NUMBER: _ClassVar[int]
    response: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, response: _Optional[_Iterable[str]] = ...) -> None: ...
