# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [state/service.proto](#state/service.proto)
 - Services
    - [Status](#state.Status)
    
  - Messages
    - [HealthRequest](#state.HealthRequest)
    - [HealthResponse](#state.HealthResponse)
    - [MetricsRequest](#state.MetricsRequest)
    - [MetricsResponse](#state.MetricsResponse)
    - [NetmapRequest](#state.NetmapRequest)
    

- [Scalar Value Types](#scalar-value-types)



<a name="state/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## state/service.proto




<a name="state.Status"></a>

### Service "state.Status"
Status service provides node's healthcheck and status info

```
rpc Netmap(NetmapRequest) returns (.bootstrap.SpreadMap);
rpc Metrics(MetricsRequest) returns (MetricsResponse);
rpc HealthCheck(HealthRequest) returns (HealthResponse);

```

#### Method Netmap

Netmap request allows to receive current [bootstrap.SpreadMap](bootstrap.md#bootstrap.SpreadMap)

| Name | Input | Output |
| ---- | ----- | ------ |
| Netmap | [NetmapRequest](#state.NetmapRequest) | [.bootstrap.SpreadMap](#bootstrap.SpreadMap) |
#### Method Metrics

Metrics request allows to receive metrics in prometheus format

| Name | Input | Output |
| ---- | ----- | ------ |
| Metrics | [MetricsRequest](#state.MetricsRequest) | [MetricsResponse](#state.MetricsResponse) |
#### Method HealthCheck

HealthCheck request allows to check health status of the node.
If node unhealthy field Status would contains detailed info.

| Name | Input | Output |
| ---- | ----- | ------ |
| HealthCheck | [HealthRequest](#state.HealthRequest) | [HealthResponse](#state.HealthResponse) |
 <!-- end services -->


<a name="state.HealthRequest"></a>

### Message HealthRequest
HealthRequest message to check current state



<a name="state.HealthResponse"></a>

### Message HealthResponse
HealthResponse message with current state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Healthy | [bool](#bool) |  | Healthy is true when node alive and healthy |
| Status | [string](#string) |  | Status contains detailed information about health status |


<a name="state.MetricsRequest"></a>

### Message MetricsRequest
MetricsRequest message to request node metrics



<a name="state.MetricsResponse"></a>

### Message MetricsResponse
MetricsResponse contains [][]byte,
every []byte is marshaled MetricFamily proto message
from github.com/prometheus/client_model/metrics.proto


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Metrics | [bytes](#bytes) | repeated |  |


<a name="state.NetmapRequest"></a>

### Message NetmapRequest
NetmapRequest message to request current node netmap


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

