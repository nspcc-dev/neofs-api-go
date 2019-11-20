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
Open a trusted session to manipulate an object. In order to put or
delete object client have to obtain session token with trusted node.
Trusted node will modify client's object (add missing headers, checksums,
homomorphic hash) and sign id with session key. Session is established
during 4-step handshake in one gRPC stream

```
rpc Create(stream CreateRequest) returns (stream CreateResponse);

```

#### Method Create



| Name | Input | Output |
| ---- | ----- | ------ |
| Create | [CreateRequest](#session.CreateRequest) | [CreateResponse](#session.CreateResponse) |
 <!-- end services -->


<a name="session.CreateRequest"></a>

### Message CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Init | [Token](#session.Token) |  | Message to init session opening. Carry: owner of manipulation object; ID of manipulation object; token lifetime bounds. |
| Signed | [Token](#session.Token) |  | Signed Init message response (Unsigned) from server with user private key |


<a name="session.CreateResponse"></a>

### Message CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Unsigned | [Token](#session.Token) |  | Unsigned token with token ID and session public key generated on server side |
| Result | [Token](#session.Token) |  | Resulting token which can be used for object placing through an trusted intermediary |

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
| OwnerID | [bytes](#bytes) |  | Owner of manipulation object |
| FirstEpoch | [uint64](#uint64) |  | Initial epoch of token lifetime |
| LastEpoch | [uint64](#uint64) |  | Last epoch of token lifetime |
| ObjectID | [bytes](#bytes) | repeated | ID of manipulation object |
| Signature | [bytes](#bytes) |  | Token signature. Signed by owner of manipulation object |
| ID | [bytes](#bytes) |  | Token ID (UUID) |


<a name="session.VerificationHeader"></a>

### Message VerificationHeader



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PublicKey | [bytes](#bytes) |  | Session public key |
| KeySignature | [bytes](#bytes) |  | Session public key signature. Signed by trusted side |

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

