# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [container/service.proto](#container/service.proto)
 - Services
    - [Service](#container.Service)
    
  - Messages
    - [DeleteRequest](#container.DeleteRequest)
    - [DeleteResponse](#container.DeleteResponse)
    - [ExtendedACLKey](#container.ExtendedACLKey)
    - [ExtendedACLValue](#container.ExtendedACLValue)
    - [GetExtendedACLRequest](#container.GetExtendedACLRequest)
    - [GetExtendedACLResponse](#container.GetExtendedACLResponse)
    - [GetRequest](#container.GetRequest)
    - [GetResponse](#container.GetResponse)
    - [ListRequest](#container.ListRequest)
    - [ListResponse](#container.ListResponse)
    - [PutRequest](#container.PutRequest)
    - [PutResponse](#container.PutResponse)
    - [SetExtendedACLRequest](#container.SetExtendedACLRequest)
    - [SetExtendedACLResponse](#container.SetExtendedACLResponse)
    

- [container/types.proto](#container/types.proto)

  - Messages
    - [Container](#container.Container)
    

- [Scalar Value Types](#scalar-value-types)



<a name="container/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## container/service.proto




<a name="container.Service"></a>

### Service "container.Service"
Container service provides API for manipulating with the container.

```
rpc Put(PutRequest) returns (PutResponse);
rpc Delete(DeleteRequest) returns (DeleteResponse);
rpc Get(GetRequest) returns (GetResponse);
rpc List(ListRequest) returns (ListResponse);
rpc SetExtendedACL(SetExtendedACLRequest) returns (SetExtendedACLResponse);
rpc GetExtendedACL(GetExtendedACLRequest) returns (GetExtendedACLResponse);

```

#### Method Put

Put request proposes container to the inner ring nodes. They will
accept new container if user has enough deposit. All containers
are accepted by the consensus, therefore it is asynchronous process.

| Name | Input | Output |
| ---- | ----- | ------ |
| Put | [PutRequest](#container.PutRequest) | [PutResponse](#container.PutResponse) |
#### Method Delete

Delete container removes it from the inner ring container storage. It
also asynchronous process done by consensus.

| Name | Input | Output |
| ---- | ----- | ------ |
| Delete | [DeleteRequest](#container.DeleteRequest) | [DeleteResponse](#container.DeleteResponse) |
#### Method Get

Get container returns container instance

| Name | Input | Output |
| ---- | ----- | ------ |
| Get | [GetRequest](#container.GetRequest) | [GetResponse](#container.GetResponse) |
#### Method List

List returns all user's containers

| Name | Input | Output |
| ---- | ----- | ------ |
| List | [ListRequest](#container.ListRequest) | [ListResponse](#container.ListResponse) |
#### Method SetExtendedACL

SetExtendedACL changes extended ACL rules of the container

| Name | Input | Output |
| ---- | ----- | ------ |
| SetExtendedACL | [SetExtendedACLRequest](#container.SetExtendedACLRequest) | [SetExtendedACLResponse](#container.SetExtendedACLResponse) |
#### Method GetExtendedACL

GetExtendedACL returns extended ACL rules of the container

| Name | Input | Output |
| ---- | ----- | ------ |
| GetExtendedACL | [GetExtendedACLRequest](#container.GetExtendedACLRequest) | [GetExtendedACLResponse](#container.GetExtendedACLResponse) |
 <!-- end services -->


<a name="container.DeleteRequest"></a>

### Message DeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CID | [bytes](#bytes) |  | CID (container id) is a SHA256 hash of the container structure |
| Meta | [service.RequestMetaHeader](#service.RequestMetaHeader) |  | RequestMetaHeader contains information about request meta headers (should be embedded into message) |
| Verify | [service.RequestVerificationHeader](#service.RequestVerificationHeader) |  | RequestVerificationHeader is a set of signatures of every NeoFS Node that processed request (should be embedded into message) |


<a name="container.DeleteResponse"></a>

### Message DeleteResponse
DeleteResponse is empty because delete operation is asynchronous and done
via consensus in inner ring nodes



<a name="container.ExtendedACLKey"></a>

### Message ExtendedACLKey



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [bytes](#bytes) |  | ID (container id) is a SHA256 hash of the container structure |


<a name="container.ExtendedACLValue"></a>

### Message ExtendedACLValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| EACL | [bytes](#bytes) |  | EACL carries binary representation of the table of extended ACL rules |
| Signature | [bytes](#bytes) |  | Signature carries EACL field signature |


<a name="container.GetExtendedACLRequest"></a>

### Message GetExtendedACLRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Key | [ExtendedACLKey](#container.ExtendedACLKey) |  | Key carries key to extended ACL information |
| Meta | [service.RequestMetaHeader](#service.RequestMetaHeader) |  | RequestMetaHeader contains information about request meta headers (should be embedded into message) |
| Verify | [service.RequestVerificationHeader](#service.RequestVerificationHeader) |  | RequestVerificationHeader is a set of signatures of every NeoFS Node that processed request (should be embedded into message) |


<a name="container.GetExtendedACLResponse"></a>

### Message GetExtendedACLResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ACL | [ExtendedACLValue](#container.ExtendedACLValue) |  | ACL carries extended ACL information |


<a name="container.GetRequest"></a>

### Message GetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CID | [bytes](#bytes) |  | CID (container id) is a SHA256 hash of the container structure |
| Meta | [service.RequestMetaHeader](#service.RequestMetaHeader) |  | RequestMetaHeader contains information about request meta headers (should be embedded into message) |
| Verify | [service.RequestVerificationHeader](#service.RequestVerificationHeader) |  | RequestVerificationHeader is a set of signatures of every NeoFS Node that processed request (should be embedded into message) |


<a name="container.GetResponse"></a>

### Message GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Container | [Container](#container.Container) |  | Container is a structure that contains placement rules and owner id |


<a name="container.ListRequest"></a>

### Message ListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| OwnerID | [bytes](#bytes) |  | OwnerID is a wallet address |
| Meta | [service.RequestMetaHeader](#service.RequestMetaHeader) |  | RequestMetaHeader contains information about request meta headers (should be embedded into message) |
| Verify | [service.RequestVerificationHeader](#service.RequestVerificationHeader) |  | RequestVerificationHeader is a set of signatures of every NeoFS Node that processed request (should be embedded into message) |


<a name="container.ListResponse"></a>

### Message ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CID | [bytes](#bytes) | repeated | CID (container id) is list of SHA256 hashes of the container structures |


<a name="container.PutRequest"></a>

### Message PutRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| MessageID | [bytes](#bytes) |  | MessageID is a nonce for uniq container id calculation |
| Capacity | [uint64](#uint64) |  | Capacity defines amount of data that can be stored in the container (doesn't used for now). |
| OwnerID | [bytes](#bytes) |  | OwnerID is a wallet address |
| rules | [netmap.PlacementRule](#netmap.PlacementRule) |  | Rules define storage policy for the object inside the container. |
| BasicACL | [uint32](#uint32) |  | BasicACL of the container. |
| Meta | [service.RequestMetaHeader](#service.RequestMetaHeader) |  | RequestMetaHeader contains information about request meta headers (should be embedded into message) |
| Verify | [service.RequestVerificationHeader](#service.RequestVerificationHeader) |  | RequestVerificationHeader is a set of signatures of every NeoFS Node that processed request (should be embedded into message) |


<a name="container.PutResponse"></a>

### Message PutResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CID | [bytes](#bytes) |  | CID (container id) is a SHA256 hash of the container structure |


<a name="container.SetExtendedACLRequest"></a>

### Message SetExtendedACLRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Key | [ExtendedACLKey](#container.ExtendedACLKey) |  | Key carries key to extended ACL information |
| Value | [ExtendedACLValue](#container.ExtendedACLValue) |  | Value carries extended ACL information |
| Meta | [service.RequestMetaHeader](#service.RequestMetaHeader) |  | RequestMetaHeader contains information about request meta headers (should be embedded into message) |
| Verify | [service.RequestVerificationHeader](#service.RequestVerificationHeader) |  | RequestVerificationHeader is a set of signatures of every NeoFS Node that processed request (should be embedded into message) |


<a name="container.SetExtendedACLResponse"></a>

### Message SetExtendedACLResponse



 <!-- end messages -->

 <!-- end enums -->



<a name="container/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## container/types.proto


 <!-- end services -->


<a name="container.Container"></a>

### Message Container
The Container service definition.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| OwnerID | [bytes](#bytes) |  | OwnerID is a wallet address. |
| Salt | [bytes](#bytes) |  | Salt is a nonce for unique container id calculation. |
| Capacity | [uint64](#uint64) |  | Capacity defines amount of data that can be stored in the container (doesn't used for now). |
| Rules | [netmap.PlacementRule](#netmap.PlacementRule) |  | Rules define storage policy for the object inside the container. |
| BasicACL | [uint32](#uint32) |  | BasicACL with access control rules for owner, system, others and permission bits for bearer token and extended ACL. |

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

