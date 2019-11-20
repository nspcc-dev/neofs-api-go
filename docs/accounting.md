# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [accounting/service.proto](#accounting/service.proto)
 - Services
    - [Accounting](#accounting.Accounting)
    
  - Messages
    - [BalanceRequest](#accounting.BalanceRequest)
    - [BalanceResponse](#accounting.BalanceResponse)
    

- [accounting/types.proto](#accounting/types.proto)

  - Messages
    - [Account](#accounting.Account)
    - [Balances](#accounting.Balances)
    - [ContainerCreateTarget](#accounting.ContainerCreateTarget)
    - [Lifetime](#accounting.Lifetime)
    - [LockTarget](#accounting.LockTarget)
    - [PayIO](#accounting.PayIO)
    - [Settlement](#accounting.Settlement)
    - [Settlement.Container](#accounting.Settlement.Container)
    - [Settlement.Receiver](#accounting.Settlement.Receiver)
    - [Settlement.Tx](#accounting.Settlement.Tx)
    - [Tx](#accounting.Tx)
    - [WithdrawTarget](#accounting.WithdrawTarget)
    

- [accounting/withdraw.proto](#accounting/withdraw.proto)
 - Services
    - [Withdraw](#accounting.Withdraw)
    
  - Messages
    - [DeleteRequest](#accounting.DeleteRequest)
    - [DeleteResponse](#accounting.DeleteResponse)
    - [GetRequest](#accounting.GetRequest)
    - [GetResponse](#accounting.GetResponse)
    - [Item](#accounting.Item)
    - [ListRequest](#accounting.ListRequest)
    - [ListResponse](#accounting.ListResponse)
    - [PutRequest](#accounting.PutRequest)
    - [PutResponse](#accounting.PutResponse)
    

- [Scalar Value Types](#scalar-value-types)



<a name="accounting/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## accounting/service.proto




<a name="accounting.Accounting"></a>

### Service "accounting.Accounting"
Accounting is a service that provides access for accounting balance
information

```
rpc Balance(BalanceRequest) returns (BalanceResponse);

```

#### Method Balance

Balance returns current balance status of the NeoFS user

| Name | Input | Output |
| ---- | ----- | ------ |
| Balance | [BalanceRequest](#accounting.BalanceRequest) | [BalanceResponse](#accounting.BalanceResponse) |
 <!-- end services -->


<a name="accounting.BalanceRequest"></a>

### Message BalanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| OwnerID | [bytes](#bytes) |  | OwnerID is a wallet address |
| TTL | [uint32](#uint32) |  | TTL must be larger than zero, it decreased in every neofs-node |


<a name="accounting.BalanceResponse"></a>

### Message BalanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Balance | [decimal.Decimal](#decimal.Decimal) |  | Balance contains current account balance state |
| LockAccounts | [Account](#accounting.Account) | repeated | LockAccounts contains information about locked funds. Locked funds appear when user pays for storage or withdraw assets. |

 <!-- end messages -->

 <!-- end enums -->



<a name="accounting/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## accounting/types.proto


 <!-- end services -->


<a name="accounting.Account"></a>

### Message Account



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| OwnerID | [bytes](#bytes) |  | OwnerID is a wallet address |
| Address | [string](#string) |  | Address is identifier of accounting record |
| ParentAddress | [string](#string) |  | ParentAddress is identifier of parent accounting record |
| ActiveFunds | [decimal.Decimal](#decimal.Decimal) |  | ActiveFunds is amount of active (non locked) funds for account |
| Lifetime | [Lifetime](#accounting.Lifetime) |  | Lifetime is time until account is valid (used for lock accounts) |
| LockTarget | [LockTarget](#accounting.LockTarget) |  | LockTarget is the purpose of lock funds (it might be withdraw or payment for storage) |
| LockAccounts | [Account](#accounting.Account) | repeated | LockAccounts contains child accounts with locked funds |


<a name="accounting.Balances"></a>

### Message Balances



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Accounts | [Account](#accounting.Account) | repeated | Accounts contains multiple account snapshots |


<a name="accounting.ContainerCreateTarget"></a>

### Message ContainerCreateTarget



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CID | [bytes](#bytes) |  | CID is container identifier |


<a name="accounting.Lifetime"></a>

### Message Lifetime



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| unit | [Lifetime.Unit](#accounting.Lifetime.Unit) |  | Unit describes how lifetime is measured in account |
| Value | [int64](#int64) |  | Value describes how long lifetime will be valid |


<a name="accounting.LockTarget"></a>

### Message LockTarget
LockTarget must be one of two options


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| WithdrawTarget | [WithdrawTarget](#accounting.WithdrawTarget) |  | WithdrawTarget used when user requested withdraw |
| ContainerCreateTarget | [ContainerCreateTarget](#accounting.ContainerCreateTarget) |  | ContainerCreateTarget used when user requested creation of container |


<a name="accounting.PayIO"></a>

### Message PayIO



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| BlockID | [uint64](#uint64) |  | BlockID contains id of the NEO block where withdraw or deposit call was invoked |
| Transactions | [Tx](#accounting.Tx) | repeated | Transactions contains all transactions that founded in block and used for PayIO |


<a name="accounting.Settlement"></a>

### Message Settlement



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Epoch | [uint64](#uint64) |  | Epoch contains an epoch when settlement was accepted |
| Transactions | [Settlement.Tx](#accounting.Settlement.Tx) | repeated | Transactions is a set of transactions |


<a name="accounting.Settlement.Container"></a>

### Message Settlement.Container



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CID | [bytes](#bytes) |  | CID is container identifier |
| SGIDs | [bytes](#bytes) | repeated | SGIDs is a set of storage groups that successfully passed the audit |


<a name="accounting.Settlement.Receiver"></a>

### Message Settlement.Receiver



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| To | [string](#string) |  | To is the address of funds recipient |
| Amount | [decimal.Decimal](#decimal.Decimal) |  | Amount is the amount of funds that will be sent |


<a name="accounting.Settlement.Tx"></a>

### Message Settlement.Tx



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| From | [string](#string) |  | From is the address of the sender of funds |
| Container | [Settlement.Container](#accounting.Settlement.Container) |  | Container that successfully had passed the audit |
| Receivers | [Settlement.Receiver](#accounting.Settlement.Receiver) | repeated | Receivers is a set of addresses of funds recipients |


<a name="accounting.Tx"></a>

### Message Tx



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [Tx.Type](#accounting.Tx.Type) |  | Type describes target of transaction |
| From | [string](#string) |  | From describes sender of funds |
| To | [string](#string) |  | To describes receiver of funds |
| Amount | [decimal.Decimal](#decimal.Decimal) |  | Amount describes amount of funds |
| PublicKeys | [bytes](#bytes) |  | PublicKeys contains public key of sender |


<a name="accounting.WithdrawTarget"></a>

### Message WithdrawTarget



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Cheque | [string](#string) |  | Cheque is a string representation of cheque id |

 <!-- end messages -->


<a name="accounting.Lifetime.Unit"></a>

### Lifetime.Unit
Unit can be Unlimited, based on NeoFS epoch or Neo block

| Name | Number | Description |
| ---- | ------ | ----------- |
| Unlimited | 0 |  |
| NeoFSEpoch | 1 |  |
| NeoBlock | 2 |  |



<a name="accounting.Tx.Type"></a>

### Tx.Type
Type can be withdrawal, payIO or inner

| Name | Number | Description |
| ---- | ------ | ----------- |
| Unknown | 0 |  |
| Withdraw | 1 |  |
| PayIO | 2 |  |
| Inner | 3 |  |


 <!-- end enums -->



<a name="accounting/withdraw.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## accounting/withdraw.proto




<a name="accounting.Withdraw"></a>

### Service "accounting.Withdraw"
Withdraw is a service that provides withdraw assets operations from the NeoFS

```
rpc Get(GetRequest) returns (GetResponse);
rpc Put(PutRequest) returns (PutResponse);
rpc List(ListRequest) returns (ListResponse);
rpc Delete(DeleteRequest) returns (DeleteResponse);

```

#### Method Get

Get returns cheque if it was signed by inner ring nodes

| Name | Input | Output |
| ---- | ----- | ------ |
| Get | [GetRequest](#accounting.GetRequest) | [GetResponse](#accounting.GetResponse) |
#### Method Put

Put ask inner ring nodes to sign a cheque for withdraw invoke

| Name | Input | Output |
| ---- | ----- | ------ |
| Put | [PutRequest](#accounting.PutRequest) | [PutResponse](#accounting.PutResponse) |
#### Method List

List shows all user's checks

| Name | Input | Output |
| ---- | ----- | ------ |
| List | [ListRequest](#accounting.ListRequest) | [ListResponse](#accounting.ListResponse) |
#### Method Delete

Delete allows user to remove unused cheque

| Name | Input | Output |
| ---- | ----- | ------ |
| Delete | [DeleteRequest](#accounting.DeleteRequest) | [DeleteResponse](#accounting.DeleteResponse) |
 <!-- end services -->


<a name="accounting.DeleteRequest"></a>

### Message DeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [bytes](#bytes) |  | ID is cheque identifier |
| OwnerID | [bytes](#bytes) |  | OwnerID is a wallet address |
| MessageID | [bytes](#bytes) |  | MessageID is a nonce for uniq request (UUIDv4) |
| Signature | [bytes](#bytes) |  | Signature is a signature of the sent request |
| TTL | [uint32](#uint32) |  | TTL must be larger than zero, it decreased in every neofs-node |


<a name="accounting.DeleteResponse"></a>

### Message DeleteResponse
DeleteResponse is empty



<a name="accounting.GetRequest"></a>

### Message GetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [bytes](#bytes) |  | ID is cheque identifier |
| OwnerID | [bytes](#bytes) |  | OwnerID is a wallet address |
| TTL | [uint32](#uint32) |  | TTL must be larger than zero, it decreased in every neofs-node |


<a name="accounting.GetResponse"></a>

### Message GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Withdraw | [Item](#accounting.Item) |  | Item is cheque with meta information |


<a name="accounting.Item"></a>

### Message Item



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [bytes](#bytes) |  | ID is a cheque identifier |
| OwnerID | [bytes](#bytes) |  | OwnerID is a wallet address |
| Amount | [decimal.Decimal](#decimal.Decimal) |  | Amount of funds |
| Height | [uint64](#uint64) |  | Height is the neo blockchain height until the cheque is valid |
| Payload | [bytes](#bytes) |  | Payload contains cheque representation in bytes |


<a name="accounting.ListRequest"></a>

### Message ListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| OwnerID | [bytes](#bytes) |  | OwnerID is a wallet address |
| TTL | [uint32](#uint32) |  | TTL must be larger than zero, it decreased in every neofs-node |


<a name="accounting.ListResponse"></a>

### Message ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Items | [Item](#accounting.Item) | repeated | Item is a set of cheques with meta information |


<a name="accounting.PutRequest"></a>

### Message PutRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| OwnerID | [bytes](#bytes) |  | OwnerID is a wallet address |
| Amount | [decimal.Decimal](#decimal.Decimal) |  | Amount of funds |
| Height | [uint64](#uint64) |  | Height is the neo blockchain height until the cheque is valid |
| MessageID | [bytes](#bytes) |  | MessageID is a nonce for uniq request (UUIDv4) |
| Signature | [bytes](#bytes) |  | Signature is a signature of the sent request |
| TTL | [uint32](#uint32) |  | TTL must be larger than zero, it decreased in every neofs-node |


<a name="accounting.PutResponse"></a>

### Message PutResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [bytes](#bytes) |  | ID is cheque identifier |

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

