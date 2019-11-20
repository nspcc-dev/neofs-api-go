# NeoFS-proto 

NeoFS-proto repository contains implementation of core NeoFS structures that
can be used for integration with NeoFS. 

## Description

Repository contains 13 packages that implement NeoFS core structures. These
packages mostly contain protobuf files with service and structure definitions 
or NeoFS core types with complemented functions. 

There is a auto-generated documentation for protobuf files in [docs](docs).

### Accounting

Accounting package defines services and structures for accounting operations: 
balance request and `cheque` operations for withdraw. `Cheque` is a structure
with inner ring signatures, which approve that user can withdraw requested
amount of assets. NeoFS smart contract takes binary formatted `cheque` as a 
parameter in withdraw call. 

#### API
Accounting package defines: 
- [Accounting service](docs/accounting.md#accounting.Accounting)
- [Withdraw service](docs/accounting.md#accounting.Withdraw)


### Bootstrap

Bootstrap package defines bootstrap service which is used by storage nodes to
connect to the storage network. 

#### API
Bootstrap package defines:
- [Bootstrap service](docs/bootstrap.md#bootstrap.Bootstrap)


### Chain

Chain package contains util functions for operations with NEO Blockchain types:
wallet addresses, script-hashes. 

### Container

Container package defines service and structures for operations with containers.
Objects in NeoFS are stored in containers. Container defines storage 
policy for the objects. 

#### API
Bootstrap package defines:
- [Container service](docs/container.md#container.Service)

### Decimal

Decimal defines custom decimal implementation which is used in accounting
operations.

### Hash

Hash package defines homomorphic hash type.

### Internal

Internal package defines constant error type and proto interface for custom
protobuf structures. 

### Object

Object package defines service and structures for object operations. Object is 
a core storage structure in NeoFS. Package contains detailed information
about object internal structure.

#### API
Object package defines:
- [Object service](docs/object.md#object.Service)

### Query

Query package defines structure for object search requests. 

### Refs

Refs package defines core identity types: Object ID, Container ID, etc.

### Service

Service package defines util structure and functions for all NeoFS services
operations: TTL and request signature management, node roles, epoch retriever. 

### Session

Session package defines service and structures for session obtain. Object
operations require an established session with pair of session keys signed by 
owner of the object. 

#### API
Session package defines:
- [Session service](docs/session.md#session.Session)

### State

State package defines service and structures for metrics gathering. 

#### API
State package defines:
- [Status service](docs/state.md#state.Status)

## How to use

NeoFS-proto packages contain godoc documentation. Examples of using most of 
these packages can be found in NeoFS-CLI repository. CLI implements and 
demonstrates all basic interactions with NeoFS: container, object, storage 
group, and accounting operations. 

Protobuf files are recompiled with the command:

```
$ make protoc
```

## Contributing

At this moment, we do not accept contributions. 

## License

This project is licensed under the GPLv3 License - 
see the [LICENSE.md](LICENSE.md) file for details
