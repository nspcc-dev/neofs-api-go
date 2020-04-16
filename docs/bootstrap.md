# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [bootstrap/service.proto](#bootstrap/service.proto)
 - Services
    - [Bootstrap](#bootstrap.Bootstrap)
    
  - Messages
    - [Request](#bootstrap.Request)
    

- [bootstrap/types.proto](#bootstrap/types.proto)

  - Messages
    - [NodeInfo](#bootstrap.NodeInfo)
    - [SpreadMap](#bootstrap.SpreadMap)
    

- [Scalar Value Types](#scalar-value-types)



<a name="bootstrap/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## bootstrap/service.proto




<a name="bootstrap.Bootstrap"></a>

### Service "bootstrap.Bootstrap"
Bootstrap service allows neofs-node to connect to the network. Node should
perform at least one bootstrap request in the epoch to stay in the network
for the next epoch.

```
rpc Process(Request) returns (SpreadMap);

```

#### Method Process

Process is method that allows to register node in the network and receive actual netmap

| Name | Input | Output |
| ---- | ----- | ------ |
| Process | [Request](#bootstrap.Request) | [SpreadMap](#bootstrap.SpreadMap) |
 <!-- end services -->


<a name="bootstrap.Request"></a>

### Message Request



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [int32](#int32) |  | Type is NodeType, can be InnerRingNode (type=1) or StorageNode (type=2) |
| info | [NodeInfo](#bootstrap.NodeInfo) |  | Info contains information about node |
| state | [Request.State](#bootstrap.Request.State) |  | State contains node status |
| Meta | [service.RequestMetaHeader](#service.RequestMetaHeader) |  | RequestMetaHeader contains information about request meta headers (should be embedded into message) |
| Verify | [service.RequestVerificationHeader](#service.RequestVerificationHeader) |  | RequestVerificationHeader is a set of signatures of every NeoFS Node that processed request (should be embedded into message) |

 <!-- end messages -->


<a name="bootstrap.Request.State"></a>

### Request.State
Node state

| Name | Number | Description |
| ---- | ------ | ----------- |
| Unknown | 0 | used by default |
| Online | 1 | used to inform that node online |
| Offline | 2 | used to inform that node offline |


 <!-- end enums -->



<a name="bootstrap/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## bootstrap/types.proto


 <!-- end services -->


<a name="bootstrap.NodeInfo"></a>

### Message NodeInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Address | [string](#string) |  | Address is a node [multi-address](https://github.com/multiformats/multiaddr) |
| PubKey | [bytes](#bytes) |  | PubKey is a compressed public key representation in bytes |
| Options | [string](#string) | repeated | Options is set of node optional information, such as storage capacity, node location, price and etc |
| Status | [uint64](#uint64) |  | Status is bitmap status of the node |


<a name="bootstrap.SpreadMap"></a>

### Message SpreadMap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Epoch | [uint64](#uint64) |  | Epoch is current epoch for netmap |
| NetMap | [NodeInfo](#bootstrap.NodeInfo) | repeated | NetMap is a set of NodeInfos |

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

