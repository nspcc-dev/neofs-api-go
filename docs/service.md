# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [service/meta.proto](#service/meta.proto)

  - Messages
    - [RequestExtendedHeader](#service.RequestExtendedHeader)
    - [RequestExtendedHeader.KV](#service.RequestExtendedHeader.KV)
    - [RequestMetaHeader](#service.RequestMetaHeader)
    - [ResponseMetaHeader](#service.ResponseMetaHeader)
    

- [service/verify.proto](#service/verify.proto)

  - Messages
    - [BearerTokenMsg](#service.BearerTokenMsg)
    - [BearerTokenMsg.Info](#service.BearerTokenMsg.Info)
    - [RequestVerificationHeader](#service.RequestVerificationHeader)
    - [RequestVerificationHeader.Signature](#service.RequestVerificationHeader.Signature)
    - [Token](#service.Token)
    - [Token.Info](#service.Token.Info)
    - [TokenLifetime](#service.TokenLifetime)
    

- [service/verify_test.proto](#service/verify_test.proto)

  - Messages
    - [TestRequest](#service.TestRequest)
    

- [Scalar Value Types](#scalar-value-types)



<a name="service/meta.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## service/meta.proto


 <!-- end services -->


<a name="service.RequestExtendedHeader"></a>

### Message RequestExtendedHeader
RequestExtendedHeader contains extended headers of request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Headers | [RequestExtendedHeader.KV](#service.RequestExtendedHeader.KV) | repeated | Headers carries list of key-value headers |


<a name="service.RequestExtendedHeader.KV"></a>

### Message RequestExtendedHeader.KV
KV contains string key-value pair


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| K | [string](#string) |  | K carries extended header key |
| V | [string](#string) |  | V carries extended header value |


<a name="service.RequestMetaHeader"></a>

### Message RequestMetaHeader
RequestMetaHeader contains information about request meta headers
(should be embedded into message)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TTL | [uint32](#uint32) |  | TTL must be larger than zero, it decreased in every NeoFS Node |
| Epoch | [uint64](#uint64) |  | Epoch for user can be empty, because node sets epoch to the actual value |
| Version | [uint32](#uint32) |  | Version defines protocol version TODO: not used for now, should be implemented in future |
| Raw | [bool](#bool) |  | Raw determines whether the request is raw or not |
| ExtendedHeader | [RequestExtendedHeader](#service.RequestExtendedHeader) |  | ExtendedHeader carries extended headers of the request |


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


<a name="service.BearerTokenMsg"></a>

### Message BearerTokenMsg
BearerTokenMsg carries information about request ACL rules with limited lifetime


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TokenInfo | [BearerTokenMsg.Info](#service.BearerTokenMsg.Info) |  | TokenInfo is a grouped information about token |
| OwnerKey | [bytes](#bytes) |  | OwnerKey is a public key of the token owner |
| Signature | [bytes](#bytes) |  | Signature is a signature of token information |


<a name="service.BearerTokenMsg.Info"></a>

### Message BearerTokenMsg.Info



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ACLRules | [bytes](#bytes) |  | ACLRules carries a binary representation of the table of extended ACL rules |
| OwnerID | [bytes](#bytes) |  | OwnerID is an owner of token |
| ValidUntil | [uint64](#uint64) |  | ValidUntil carries a last epoch of token lifetime |


<a name="service.RequestVerificationHeader"></a>

### Message RequestVerificationHeader
RequestVerificationHeader is a set of signatures of every NeoFS Node that processed request
(should be embedded into message).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Signatures | [RequestVerificationHeader.Signature](#service.RequestVerificationHeader.Signature) | repeated | Signatures is a set of signatures of every passed NeoFS Node |
| Token | [Token](#service.Token) |  | Token is a token of the session within which the request is sent |
| Bearer | [BearerTokenMsg](#service.BearerTokenMsg) |  | Bearer is a Bearer token of the request |


<a name="service.RequestVerificationHeader.Signature"></a>

### Message RequestVerificationHeader.Signature



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Sign | [bytes](#bytes) |  | Sign is signature of the request or session key. |
| Peer | [bytes](#bytes) |  | Peer is compressed public key used for signature. |


<a name="service.Token"></a>

### Message Token
User token granting rights for object manipulation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TokenInfo | [Token.Info](#service.Token.Info) |  | TokenInfo is a grouped information about token |
| Signature | [bytes](#bytes) |  | Signature is a signature of session token information |


<a name="service.Token.Info"></a>

### Message Token.Info



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [bytes](#bytes) |  | ID is a token identifier. valid UUIDv4 represented in bytes |
| OwnerID | [bytes](#bytes) |  | OwnerID is an owner of manipulation object |
| verb | [Token.Info.Verb](#service.Token.Info.Verb) |  | Verb is a type of request for which the token is issued |
| Address | [refs.Address](#refs.Address) |  | Address is an object address for which token is issued |
| Lifetime | [TokenLifetime](#service.TokenLifetime) |  | Lifetime is a lifetime of the session |
| SessionKey | [bytes](#bytes) |  | SessionKey is a public key of session key |
| OwnerKey | [bytes](#bytes) |  | OwnerKey is a public key of the token owner |


<a name="service.TokenLifetime"></a>

### Message TokenLifetime
TokenLifetime carries a group of lifetime parameters of the token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Created | [uint64](#uint64) |  | Created carries an initial epoch of token lifetime |
| ValidUntil | [uint64](#uint64) |  | ValidUntil carries a last epoch of token lifetime |

 <!-- end messages -->


<a name="service.Token.Info.Verb"></a>

### Token.Info.Verb
Verb is an enumeration of session request types

| Name | Number | Description |
| ---- | ------ | ----------- |
| Put | 0 | Put refers to object.Put RPC call |
| Get | 1 | Get refers to object.Get RPC call |
| Head | 2 | Head refers to object.Head RPC call |
| Search | 3 | Search refers to object.Search RPC call |
| Delete | 4 | Delete refers to object.Delete RPC call |
| Range | 5 | Range refers to object.GetRange RPC call |
| RangeHash | 6 | RangeHash refers to object.GetRangeHash RPC call |


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

