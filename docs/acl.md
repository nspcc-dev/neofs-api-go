# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [acl/types.proto](#acl/types.proto)

  - Messages
    - [EACLRecord](#acl.EACLRecord)
    - [EACLRecord.FilterInfo](#acl.EACLRecord.FilterInfo)
    - [EACLRecord.TargetInfo](#acl.EACLRecord.TargetInfo)
    - [EACLTable](#acl.EACLTable)
    

- [Scalar Value Types](#scalar-value-types)



<a name="acl/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## acl/types.proto


 <!-- end services -->


<a name="acl.EACLRecord"></a>

### Message EACLRecord
EACLRecord groups information about extended ACL rule.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| operation | [EACLRecord.Operation](#acl.EACLRecord.Operation) |  | Operation carries type of operation. |
| action | [EACLRecord.Action](#acl.EACLRecord.Action) |  | Action carries ACL target action. |
| Filters | [EACLRecord.FilterInfo](#acl.EACLRecord.FilterInfo) | repeated | Filters carries set of filters. |
| Targets | [EACLRecord.TargetInfo](#acl.EACLRecord.TargetInfo) | repeated | Targets carries information about extended ACL target list. |


<a name="acl.EACLRecord.FilterInfo"></a>

### Message EACLRecord.FilterInfo
FilterInfo groups information about filter.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [EACLRecord.FilterInfo.Header](#acl.EACLRecord.FilterInfo.Header) |  | Header carries type of header. |
| matchType | [EACLRecord.FilterInfo.MatchType](#acl.EACLRecord.FilterInfo.MatchType) |  | MatchType carries type of match. |
| HeaderName | [string](#string) |  | HeaderName carries name of filtering header. |
| HeaderVal | [string](#string) |  | HeaderVal carries value of filtering header. |


<a name="acl.EACLRecord.TargetInfo"></a>

### Message EACLRecord.TargetInfo
TargetInfo groups information about extended ACL target.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Target | [Target](#acl.Target) |  | Target carries target of ACL rule. |
| KeyList | [bytes](#bytes) | repeated | KeyList carries public keys of ACL target. |


<a name="acl.EACLTable"></a>

### Message EACLTable
EACLRecord carries the information about extended ACL rules.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Records | [EACLRecord](#acl.EACLRecord) | repeated | Records carries list of extended ACL rule records. |

 <!-- end messages -->


<a name="acl.EACLRecord.Action"></a>

### EACLRecord.Action
Action is an enumeration of EACL actions.

| Name | Number | Description |
| ---- | ------ | ----------- |
| ActionUnknown | 0 |  |
| Allow | 1 |  |
| Deny | 2 |  |



<a name="acl.EACLRecord.FilterInfo.Header"></a>

### EACLRecord.FilterInfo.Header
Header is an enumeration of filtering header types.

| Name | Number | Description |
| ---- | ------ | ----------- |
| HeaderUnknown | 0 |  |
| Request | 1 |  |
| ObjectSystem | 2 |  |
| ObjectUser | 3 |  |



<a name="acl.EACLRecord.FilterInfo.MatchType"></a>

### EACLRecord.FilterInfo.MatchType
MatchType is an enumeration of match types.

| Name | Number | Description |
| ---- | ------ | ----------- |
| MatchUnknown | 0 |  |
| StringEqual | 1 |  |
| StringNotEqual | 2 |  |



<a name="acl.EACLRecord.Operation"></a>

### EACLRecord.Operation
Operation is an enumeration of operation types.

| Name | Number | Description |
| ---- | ------ | ----------- |
| OPERATION_UNKNOWN | 0 |  |
| GET | 1 |  |
| HEAD | 2 |  |
| PUT | 3 |  |
| DELETE | 4 |  |
| SEARCH | 5 |  |
| GETRANGE | 6 |  |
| GETRANGEHASH | 7 |  |



<a name="acl.Target"></a>

### Target
Target of the access control rule in access control list.

| Name | Number | Description |
| ---- | ------ | ----------- |
| Unknown | 0 | Unknown target, default value. |
| User | 1 | User target rule is applied if sender is the owner of the container. |
| System | 2 | System target rule is applied if sender is the storage node within the container or inner ring node. |
| Others | 3 | Others target rule is applied if sender is not user or system target. |
| PubKey | 4 | PubKey target rule is applied if sender has public key provided in extended ACL. |


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

