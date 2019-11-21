# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [session/service.proto](#session/service.proto)
 - Services
    - [Session](#session.Session)
    
  - Messages
    - [CreateRequest](#session.CreateRequest)
    - [CreateResponse](#session.CreateResponse)
    

- [session/types.proto](#session/types.proto)

  - Messages
    - [Token](#session.Token)
    - [VerificationHeader](#session.VerificationHeader)
    

- [Scalar Value Types](#scalar-value-types)



<a name="session/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## session/service.proto




<a name="session.Session"></a>

### Service "session.Session"


```
rpc Create(stream CreateRequest) returns (stream CreateResponse);

```

#### Method Create

Create is a method that used to open a trusted session to manipulate
an object. In order to put or delete object client have to obtain session
token with trusted node. Trusted node will modify client's object
(add missing headers, checksums, homomorphic hash) and sign id with
session key. Session is established during 4-step handshake in one gRPC stream

- First client stream message SHOULD BE type of `CreateRequest_Init`.
- First server stream message SHOULD BE type of `CreateResponse_Unsigned`.
- Second client stream message SHOULD BE type of `CreateRequest_Signed`.
- Second server stream message SHOULD BE type of `CreateResponse_Result`.

| Name | Input | Output |
| ---- | ----- | ------ |
| Create | [CreateRequest](#session.CreateRequest) | [CreateResponse](#session.CreateResponse) |
 <!-- end services -->


<a name="session.CreateRequest"></a>

### Message CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Init | [Token](#session.Token) |  | Init is a message to initialize session opening. Carry: owner of manipulation object; ID of manipulation object; token lifetime bounds. |
| Signed | [Token](#session.Token) |  | Signed Init message response (Unsigned) from server with user private key |
| Meta | [service.RequestMetaHeader](#service.RequestMetaHeader) |  | RequestMetaHeader contains information about request meta headers (should be embedded into message) |
| Verify | [service.RequestVerificationHeader](#service.RequestVerificationHeader) |  | RequestVerificationHeader is a set of signatures of every NeoFS Node that processed request (should be embedded into message) |


<a name="session.CreateResponse"></a>

### Message CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Unsigned | [Token](#session.Token) |  | Unsigned token with token ID and session public key generated on server side |
| Result | [Token](#session.Token) |  | Result is a resulting token which can be used for object placing through an trusted intermediary |

 <!-- end messages -->

 <!-- end enums -->



<a name="session/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## session/types.proto


 <!-- end services -->


<a name="session.Token"></a>

### Message Token
User token granting rights for object manipulation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Header | [VerificationHeader](#session.VerificationHeader) |  | Header carries verification data of session key |
| OwnerID | [bytes](#bytes) |  | OwnerID is an owner of manipulation object |
| FirstEpoch | [uint64](#uint64) |  | FirstEpoch is an initial epoch of token lifetime |
| LastEpoch | [uint64](#uint64) |  | LastEpoch is a last epoch of token lifetime |
| ObjectID | [bytes](#bytes) | repeated | ObjectID is an object identifier of manipulation object |
| Signature | [bytes](#bytes) |  | Signature is a token signature, signed by owner of manipulation object |
| ID | [bytes](#bytes) |  | ID is a token identifier. valid UUIDv4 represented in bytes |


<a name="session.VerificationHeader"></a>

### Message VerificationHeader



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PublicKey | [bytes](#bytes) |  | PublicKey is a session public key |
| KeySignature | [bytes](#bytes) |  | KeySignature is a session public key signature. Signed by trusted side |

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

