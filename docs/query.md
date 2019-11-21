# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [query/types.proto](#query/types.proto)

  - Messages
    - [Filter](#query.Filter)
    - [Query](#query.Query)
    

- [Scalar Value Types](#scalar-value-types)



<a name="query/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## query/types.proto


 <!-- end services -->


<a name="query.Filter"></a>

### Message Filter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [Filter.Type](#query.Filter.Type) |  | Type of filter |
| Name | [string](#string) |  | Name of field that should be filtered |
| Value | [string](#string) |  | Value that should be used for filter |


<a name="query.Query"></a>

### Message Query



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Filters | [Filter](#query.Filter) | repeated | Filters is set of filters, should not be empty |

 <!-- end messages -->


<a name="query.Filter.Type"></a>

### Filter.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| Exact | 0 | Exact sets when value of filter should be equal to the header value |
| Regex | 1 | Regex sets when value of filter should match the header value by the regular expression |


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

