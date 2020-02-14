# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [service/meta.proto](#service/meta.proto)

  - Messages
    - [RequestMetaHeader](#service.RequestMetaHeader)
    - [ResponseMetaHeader](#service.ResponseMetaHeader)
    

- [service/verify.proto](#service/verify.proto)

  - Messages
    - [RequestVerificationHeader](#service.RequestVerificationHeader)
    - [RequestVerificationHeader.Sign](#service.RequestVerificationHeader.Sign)
    - [RequestVerificationHeader.Signature](#service.RequestVerificationHeader.Signature)
    

- [service/verify_test.proto](#service/verify_test.proto)

  - Messages
    - [TestRequest](#service.TestRequest)
    

- [Scalar Value Types](#scalar-value-types)



<a name="service/meta.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## service/meta.proto


 <!-- end services -->


<a name="service.RequestMetaHeader"></a>

### Message RequestMetaHeader
RequestMetaHeader contains information about request meta headers
(should be embedded into message)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TTL | [uint32](#uint32) |  | TTL must be larger than zero, it decreased in every NeoFS Node |
| Epoch | [uint64](#uint64) |  | Epoch for user can be empty, because node sets epoch to the actual value |
| Version | [uint32](#uint32) |  | Version defines protocol version TODO: not used for now, should be implemented in future |


<a name="service.ResponseMetaHeader"></a>

### Message ResponseMetaHeader
ResponseMetaHeader contains meta information based on request processing by server
(should be embedded into message)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Epoch | [uint64](#uint64) |  | Current NeoFS epoch on server |
| Version | [uint32](#uint32) |  | Version defines protocol version TODO: not used for now, should be implemented in future |

 <!-- end messages -->

 <!-- end enums -->



<a name="service/verify.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## service/verify.proto


 <!-- end services -->


<a name="service.RequestVerificationHeader"></a>

### Message RequestVerificationHeader
RequestVerificationHeader is a set of signatures of every NeoFS Node that processed request
(should be embedded into message).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Signatures | [RequestVerificationHeader.Signature](#service.RequestVerificationHeader.Signature) | repeated | Signatures is a set of signatures of every passed NeoFS Node |


<a name="service.RequestVerificationHeader.Sign"></a>

### Message RequestVerificationHeader.Sign



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Sign | [bytes](#bytes) |  | Sign is signature of the request or session key. |
| Peer | [bytes](#bytes) |  | Peer is compressed public key used for signature. |


<a name="service.RequestVerificationHeader.Signature"></a>

### Message RequestVerificationHeader.Signature



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Sign | [RequestVerificationHeader.Sign](#service.RequestVerificationHeader.Sign) |  | Sign is a signature and public key of the request. |
| Origin | [RequestVerificationHeader.Sign](#service.RequestVerificationHeader.Sign) |  | Origin used for requests, when trusted node changes it and re-sign with session key. If session key used for signature request, then Origin should contain public key of user and signed session key. |

 <!-- end messages -->

 <!-- end enums -->



<a name="service/verify_test.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## service/verify_test.proto


 <!-- end services -->


<a name="service.TestRequest"></a>

### Message TestRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| IntField | [int32](#int32) |  |  |
| StringField | [string](#string) |  |  |
| BytesField | [bytes](#bytes) |  |  |
| CustomField | [bytes](#bytes) |  |  |
| Meta | [RequestMetaHeader](#service.RequestMetaHeader) |  |  |
| Header | [RequestVerificationHeader](#service.RequestVerificationHeader) |  |  |

 <!-- end messages -->

 <!-- end enums -->



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

