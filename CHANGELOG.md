# Changelog

## [Unreleased]

### Added
### Fixed
### Changed
### Updated
- Minimum required version of Go to 1.20
- `github.com/nspcc-dev/rfc6979` [v0.2.0 => v0.2.1](https://github.com/nspcc-dev/rfc6979/compare/v0.2.0...v0.2.1)
- `google.golang.org/grpc` [v1.59.0 => v1.62.0](https://github.com/grpc/grpc-go/compare/v1.59.0...v1.62.0)
- `google.golang.org/protobuf` [v1.31.0 => v1.32.0](https://github.com/protocolbuffers/protobuf-go/compare/v1.31.0...v1.32.0)

## [2.14.0] - 2022-10-17 - Anmado (안마도, 鞍馬島)

### Added
- Maintenance status and node state (#415)
- `NetmapSnapshot` RPC (#418)

### Fixed
- Marshal repeated `[]byte` fields correctly (#413)
- Block until client connection is up (#419)

### Changed
- Support new object session context format (#416)
- Mark all expiration methods as deprecated (#417)

### Updated
- Minimal go version to 1.17 (#412) 
- `neofs-crypto` to `v0.4.0` (#412)
- `google.golang.org/grpc` to `v1.48.0` (#415)
- `google.golang.org/protobuf` to `v1.28.0` (#415)

## [2.13.1] - 2022-08-01

### Added
- `EACL_NOT_FOUND` status code support (#409)

## [2.13.0] - 2022-07-04 - Yeonpyeongdo (연평도, 延坪島)

NeoFS API v2.13 support

### Added
- WalletConnect signature scheme (#386)
- `object.StatusOutOfRange` enum value (#407)
- Functionality to disable homomorphic hashing in container (#399)
- `status.SignatureVerificationFail` enum value (#405)

## [2.12.2] - 2022-06-03

### Added
- `util/protogen` tool to compile Protocol Buffers with stable marshaling

### Fixed
- `refs.ObjectIDNestedListMarshal` incorrect encoding (#401)
- incorrect conversion of `oneof` field in `GetRangeResponseBody.FromGRPCMessage` (#401)

### Changed
- Setters don't check call on `nil` (#388)
- `StableMarshal` methods don't return `error`
- `oneof` fields are now randomly initialized in test generators (#401)

## [2.12.1] - 2022-03-15

### Fixed
- Incompatible changes in signature scheme (#380)
 
### Added
- Public URI-parsing function `client.ParseURI()` (#383)

### Changed
- Slices with pointers `[]*T` replaced with regular structure slices `[]T` in
  v2 structures (#376)

## [2.12.0] - 2022-02-25 - Heuksando (흑산도, 黑山島)

NeoFS API v2.12 support

### Added
- Network magic field (`uint64`) to `session.RequestMetaHeader` (#367)
- Constants for attributes related to object notifications (#372)
- `refs.SignatureScheme` enum and corresponding field to `session.Signature` message (#55)
- New status codes in `container`, `object` and `session` packages (#371)
- `object.TypeLock` type enum value and `object.Lock` type of its content (#371)

### Changed
- API of `signature` package (#55)

### Deprecated
- `SessionToken`, `SessionTokenBody` and `SessionTokenContext` types from `session` package

## [2.11.1] - 2021-12-28

### Added
- Timeout for single message transmission in `rpc/client` (#366)

## [2.11.0] - 2021-12-02 - Sinjido (신지도, 薪智島)

NeoFS API v2.11 support. High level packages are moved to 
[neofs-sdk-go](https://github.com/nspcc-dev/neofs-sdk-go) repository. 
Repository restructured as Go module version 2 and synced with 
[neofs-api](https://github.com/nspcc-dev/neofs-api) release version.

### Fixed
- Protobuf recompile script handles path with special symbols correctly (#354)
- `make dep` command (#364)

### Added
- Status codes in response messages (#350)
- Storage subnetwork types (#355, #356, #360, #362)

### Changed
- neofs-api-go is now Go module version 2 (#201)

### Removed
- All packages from `pkg` moved to 
  [neofs-sdk-go](https://github.com/nspcc-dev/neofs-sdk-go) (#201)

## [1.30.0] - 2021-10-19 - Udo (우도, 牛島)

NeoFS API v2.10 support.

### Fixed

- Potential NPE in NEO3Wallet related functions (#340)
- Incorrect return of reputation intermediate announcement (#342)

### Added

- pkg/client callback to parse internal response information (#337)
- Service filter type in extended ACL from API v2.10 (#338)
- Enhanced network info structures from API v2.10 (#339)
- Well-known public-append basic ACL constant (#341) 
- Native contract names support (#351)

### Changed

- Updated grpc to v1.41.0 (#176)

## [1.29.0] - 2021-08-27 - Anmyeondo (안면도, 安眠島)

### Fixed

- Well-known BasicACL constants to fit the specification (#330, #334).
- Linter warnings (#335).

### Added

- Support for `COMMON_PREFIX` filter operation for object attributes (#331).
- Missing well-known eACL filter setters (#333).

### Removed

- `FilterObjectParent` well-known object related filter key (#333).

### Updated

- Go version to `1.16` (#335).
- Tests in CI to run on `1.16` and `1.17` Go versions (#335).

## [1.28.3] - 2021-07-07

### Fixed

- Regression of the session token corruption fix from v1.28.2 ([#326](https://github.com/nspcc-dev/neofs-api-go/issues/326)).

## [1.28.2] - 2021-07-06

### Fixed

- Data corruption of parameterized session token in `pkg/client.Client` ([#323](https://github.com/nspcc-dev/neofs-api-go/issues/323)).

## [1.28.1] - 2021-07-01

### Fixed

- Incorrect unsupported version error in `Client.GetContainer` of containers of newer versions.

### Removed

- No longer used `pkg.IsSupportedVersion` func.
- No longer used `container.NewVerifiedFromV2` func.

## [1.28.0] - 2021-06-28 - Muuido (무의도, 舞衣島)

### Added

- `String` / `FromString` methods to work with text format of enums from `pkg`.
- `Marshal(JSON)` / `Unmarshal(JSON)` methods to `container.ContainerContext` type.
- Ability to handle the `io.Reader` of the object payload in `Client.GetObject`.
- `NumberOfAddresses` / `IterateAddresses` methods to node info types for support of multiple addresses.

### Fixed

- Added leading slash to format of gRPC method names.

### Updated

- Neo Go library to v0.95.3.

## [1.27.1] - 2021-06-10

### Fixed

- SDK version was updated (actualized) to `2.7`.

### Changed

- `pkg` wrappers' `ToV2` methods return `nil` if called on `nil`.
- `pkg` wrappers' `NewFromV2` functions constructs `nil` if called with `nil` argument.

### Added

- Getters and setters for lifetime fields of `pkg/session.Token`.
- `MarshalHeaderJSON` method to `pkg/object.Object`.
- Generators for types from `pkg` (for testing).
- Descriptions of default fields for `pkg` wrappers and unit tests for its constructors.
- Unit tests for `ToV2` methods and `NewFromV2` functions.

## [1.27.0] - 2021-06-03 - Seongmodo (석모도, 席毛島)

### Added

- Message structures related to Container service sessions in `v2` and `pkg`.
- `session.Token` and `Signature` to `pkg/container.Container` and `pkg/acl/eacl.Table`.
- `Conn` method of clients to get the underlying connection.
- `WithTLSConfig` client option to specify TLS configuration.
- `WithNetworkURIAddress` client option to specify URI of the remote server.
- Generators of random container IDs, owner IDs and session tokens (for testing).

### Replaced

- `pkg/token.SessionToken` type to `pkg/session` package as `Token`. Old type is deprecated.
- `pkg/container.ID` type to `pkg/container/id` package. Old type is deprecated.

### Updated

- NEO Go library to v0.95.1.

## [1.26.1] - 2021-05-19

### Changed

- Updated neo-go to v0.95.0 release.

### Removed

- `pkg/errors` dependency (stdlib errors used instead).

## [1.26.0] - 2021-05-07 - Daecheongdo (대청도, 大靑島)

### Added

- Implementation of `v2/reputation` package.
- Implementation of reputation methods in `pkg/client`.
- Float64 stable marshaling wrappers in `util/proto`.

## [1.25.0] - 2021-03-22 - Jebudo (제부도, 濟扶島)

Raw client and support of NeoFS API v2.5.0 "Jebudo" release.

### Added

- Raw client for peer to peer communication.
- `client.WithKey` option to sign messages with different keys within single 
  client.
- `Content-Type` well-known object attribute constant.

### Changed

- Refactored `v2` sub-packages to support single raw client in all RPC methods.
- Client constructor returns `Client` interface instead of structure.

## [1.24.0] - 2021-02-26 - Ganghwado (강화도, 江華島)

Support changes from NeoFS API v2.4.0 "Ganghwado" release.

### Added 

- `netmap.NetworkInfo` definitions in `v2` and `pkg/netmap`.
- `netmap.NetworkInfo` RPC support in `pkg/client`.

### Changed

- Updated in-line docs from NeoFS API "Ganghwado" release.

## [1.23.0] - 2021-02-11 - Seonyudo (선유도, 仙遊島)

Support changes from NeoFS API v2.3.0 "Seonyudo" release.

### Added

- Fulfill backup factor for default attribute in placement.
- Support of `Container.AnnounceUsedSpace` RPC from NeoFS API.
- New `pkg/client.Client.AnnounceContainerUsedSpace` method.
- Support of `STRING_NOT_EQUAL` and `NOT_PRESENT` object search filters.
- Implementation of `json.Marshaler`/`json.Unmarshaler` on `v2/object/SearchFilter`.
- Implementation of `json.Marshaler`/`json.Unmarshaler` on `pkg/object/SearchFilters`.
- Named constants of well-known node attributes in `pkg/netmap`.

### Renamed

- `pkg/netmap/PriceAttr` to `pkg/netmap/AttrPrice`.
- `pkg/netmap/CapacityAttr` to `pkg/netmap/AttrCapacity`.

## [1.22.2] - 2021-01-27

### Fixed

- Fix size limit for grpc messages in object.Put operation.
- Fix `GetContainerNode()` function, so that it does not modify placement policy.

## [1.22.1] - 2021-01-15

Support changes from NeoFS API v2.2.1 release.

### Added

- Constant prefix of the reserved keys to X-headers (`__NEOFS__`).
- Constant string key to netmap epoch X-header (`__NEOFS__NETMAP_EPOCH`).
- Constant string key to netmap lookup depth X-header (`__NEOFS__NETMAP_LOOKUP_DEPTH`).

### Changed

- Linter's configuration in `.golangci.yml`.

### Fixed

- Remarks of the updated linter. 

## [1.22.0] - 2020-12-30 - Yeouido (여의도, 汝矣島)

Support changes from NeoFS API v2.2.0 "Yeouido" release.

### Added

- Payload hash field to `ShortHeader` message.
- Payload homomorphic hash field to `ShortHeader` message.
- Support of `StorageGroup` message.
- Support of `DataAuditResult` message.
- Stringer and string parser for `Checksum` type of client library.
- Stringer and string parser for `Type` message. 
- Stringer and string parser for `Type` type of client library.
- `AddTypeFilter` method on `SearchFilters` type of client library
  that adds filter by object type.
- Utility functions for working with `fixed64` protobuf type to `proto` library.
- Converters for `repeated` object ID messages in `v2` library.

## [1.21.2] - 2020-12-24

### Added

- `Container.NonceUUID` getter of container nonce in UUID format.
- `Container.SetNonceUUID` setter of container nonce in UUID format.
- `NewVerifiedContainerFromV2` container constructor that preliminary
  checks if container message argument meets NeoFS API V2 specification.
  
### Changed

- `Container.Nonce`/`Container.SetNonce` marked as deprecated.
- `Client.GetContainer` method returns an error if received
  container does not meet NeoFS API specification.

### Fixed

- `pkg.SDKVersion` to return version with minor `1`.
- `pkg.IsSupportedVersion` to consider `2.1` as supported.

## [1.21.1] - 2020-12-18

Support neofs-api v2.1.1.

### Added

- `client.GetVerifiedContainerStructure` function to check 
  that the container structure matches the requested identifier.

## [1.21.0] - 2020-12-11 - Modo (모도, 茅島)

### Added

- `SplitID` message support
- Search filter by `SplitID` field
- `SplitInfo` message support and related error
- `Raw` flag support in `Client.GetObject(Header)`
- Getters for parameter structures in `pkg/client` package
- `Tombstone` message support
- Tombstone address target parameter of `Client.DeleteObject` method
- `client.DeleteObject` helpful function
- Usage of default value for backup factor in placement builder

### Removed

- Object search filter by `CHILDFREE` property

### Renamed

- `AddLeafFilter` to `AddPhyFilter` 

### Fixed

- NPE in `eacl.NewTargetFromV2` function
- Processing `REP X` policies in placement builder


## [1.20.3] - 2020-11-25

### Added

- `AddObjectIDFilter` method of `SearchFilters` type
- `WithDialTimeout` option of v2 and SDK `Client`'s
- `GetEACLWithSignature` method of SDK `Client` type

### Fixed

- incorrect signature verification algorithm in `GetEACL` method of SDK `Client`

## [1.20.2] - 2020-11-17

### Fixed

- Readme badges

## [1.20.1] - 2020-11-17

### Fixed

- Signature check of head response in `pkg/client` (#202)

## [1.20.0] - 2020-11-16 - Jindo (진도, 珍島)

Major API refactoring and simplification. From now on this library will have 
backward compatibility and support of major versions of NeoFS-API by having 
**version specific** files in `vN` dirs and **version independent** SDK 
structures and client in `pkg`. This version supports NeoFS-API v2.0.X


### Added
- cross-protocol ```v2``` message types
- utility functions for message signing/verification
- ```v2```/ ```gRPC``` back and forth conversions
- primary SDK

### Removed
- v0 and v1 NeoFS API is not supported anymore

## [1.3.0] - 2020-07-23

### Changed

- Format of ```refs.OwnerID``` based on NEO3.
- Binary format of extended ACL table.
- ```acl``` package structure.

## [1.2.0] - 2020-07-08

### Added

- Extended ACL types.
- Getters and setters of ```EACLTable``` and its internal messages.
- Wrappers over ```EACLTable``` and its internal messages.
- Getters, setters and marshaling methods of wrappers.

### Changed

- Mechanism for signing requests on the principle of Matryoshka.

### Updated

- NeoFS API v1.1.0 => 1.2.0

## [1.1.0] - 2020-06-18

### Added

- `container.SetExtendedACL` rpc.
- `container.GetExtendedACL` rpc.
- Bearer token to all request messages.
- X-headers to all request messages.

### Changed

- Implementation and signatures of Sign/Verify request functions.

### Updated

- NeoFS API v1.0.0 => 1.1.0

## [1.0.0] - 2020-05-26

- Bump major release

### Updated

- NeoFS API v0.7.5 => v1.0.0
- github.com/golang/protobuf v1.4.0 => v1.4.2
- github.com/prometheus/client_golang v1.5.1 => v1.6.0
- github.com/spf13/viper v1.6.2 => v1.7.0
- google.golang.org/grpc v1.28.1 => v1.29.1

## [0.7.6] - 2020-05-19

### Added

- `session.PublicSessionToken` function for session public key bytes receiving.
- The implementation of `service.DataWithSignKeyAccumulator` methods on `object.IntegrityHeader`.

### Changed

- The implementation of `AddSignKey` method on `service.signedSessionToken` structure.
- `session.PrivateTOken` interface methods group.

### Removed

- `OwnerKey` from `service.SessionToken` signed payload.

### Fixed

- Incorrect `object.HeadRequest.ReadSignedData` method implementation.

## [0.7.5] - 2020-05-16

### Added

- Owner key to the `SessionToken` signed payload.

### Changed

- `OwnerKeyContainer` interface embedded to `SessionTokenInfo` interface.

### Updated

- NeoFS API v0.7.5

## [0.7.4] - 2020-05-12

### Added

- Stringify for `object.Object`.

### Changed

- Mechanism for creating and verifying request message signatures.
- Implementation and interface of private token storage.
- File structure of packages.

### Updated

- NeoFS API v0.7.4

## [0.7.1] - 2020-04-20

### Added

- Method to change current node state. (`state.ChangeState`)

### Updated

- NeoFS API v0.7.1

## [0.7.0] - 2020-04-16

### Updated
- NeoFS API v0.7.0

## [0.6.2] - 2020-04-16

### Updated
- NeoFS API v0.6.1
- Protobuf v1.4.0
- Netmap v1.7.0
- Prometheus Client v1.5.1
- Testify v1.5.1
- gRPC v1.28.1

### Fixed
- formatting
- test coverage for Object.PutRequest.CID method

## [0.6.1] - 2020-04-10

### Changed

- License changed to Apache 2.0

### Fixed

-  NPE in PutRequest.CID()


## [0.6.0] - 2020-04-03

### Added

- `RequestType` for object service requests
- `Type()` function in `Request` interface

### Changed

- Synced proto files with `neofs-api v0.6.0`

## [0.5.0] - 2020-03-31

### Changed
- Rename repo to `neofs-api-go`
- Used public proto files

## [0.4.2] - 2020-03-16

### Fixed
- NPE bug with CID method of object.PutRequest

## [0.4.1] - 2020-03-02

### Changed
- Updated neofs-crypto library to v0.3.0

## [0.4.0] - 2020-02-18

### Added
- Meta header for all gRPC responses. It contains epoch stamp and version number.
### Changed
- Endianness in accounting cheque. Now it uses little endian for cheaper
decoding in neofs smart-contract.

## [0.3.2] - 2020-02-10

### Added
- gRPC method DumpVars to State service
- add method `EncodeVariables` to encode debug variables to JSON (slice of bytes)
- increase test coverage for state package

### Updated
- state proto file
- documentation for state service and messages

## [0.3.1] - 2020-02-07
### Fixed
- bug with `tz.Concat`

### Updated
- dependencies:
    - github.com/nspcc-dev/tzhash `v1.3.0 => v1.4.0`
    - github.com/prometheus/client_golang `v1.4.0 => v1.4.1`
    - google.golang.org/grpc `v1.27.0 => v1.27.1`

## [0.3.0] - 2020-02-05

### Updated
- proto files
- dependencies
    - github.com/golang/protobuf `v1.3.2 => v1.3.3`
    - github.com/pkg/errors `v0.8.1 => v0.9.1`
    - github.com/prometheus/client_golang `v1.2.1 => v1.4.0`
    - github.com/prometheus/client_model `v0.0.0-20190812154241-14fe0d1b01d4 => v0.2.0`
    - github.com/spf13/viper `v1.6.1 => v1.6.2`
    - google.golang.org/grpc `v1.24.0 => v1.27.0`

### Changed
- make object.GetRange to be server-side streaming RPC
- GetRange response struct

### Added
- badges to readme

## [0.2.14] - 2020-02-04

### Fixed
- Readme

### Added
- Filename header

### Updated
- Object.Search now uses streams

## [0.2.13] - 2020-02-03

### Fixed
- Code format

### Changed
- Use separated proto repository
- Rename neofs-proto to neofs-api

## [0.2.12] - 2020-01-27

### Fixed
- Bug with ByteSize (0 bytes returns NaN)

## [0.2.11] - 2020-01-21

### Added
- Raw flag in object head and get queries with docs

## [0.2.10] - 2020-01-17

### Changed
- Private token contructor now takes public keys as an argument

## [0.2.9] - 2020-01-17

### Added
- Docs for container ACL field
- Public key header in the object with docs
- Public key field in the session token with docs

### Changed
- Routine to verify correct object checks if integrity header is last and
may use public key header if verification header is not present
- Routine to verify correct session token checks if keys in the token
associated with owner id
- Updated neofs-crypto to v0.2.3

### Removed
- Timestamp in object tombstone header

## [0.2.8] - 2019-12-21

### Added
- Container access control type definitions

### Changed
- Used sync.Pool for Sign/VerifyRequestHeader
- VerifiableRequest.Marshal method replace with MarshalTo and Size

## [0.2.7] - 2019-12-17

### Fixed
- Bug with DecodeMetrics (empty metrics returns)

## [0.2.6] - 2019-12-17

### Added
- Request to dump node config

## [0.2.5] - 2019-12-05

### Removed
- proto.Message in Maintainable/Verifiable requests

## [0.2.4] - 2019-12-03

### Added
- StorageGroup library

### Changed
- Storage group part of object library moved into separate package
- Updated proto documentation

## [0.2.3] - 2019-11-28

### Removed
- service: SignRequest / VerifyRequest and accompanying code
- proto: Signature field from requests
- object: bytefmt package not used anymore

### Changed
- service: rename EpochRequest to EpochHeader and merge with MetaHeader
- service: get status error even if it is wrapped

### Added
- service: RequestVerificationHeader's method to validate owner
- service: test coverage for CheckOwner
- service: test coverage for wrapped status errors

## [0.2.2] - 2019-11-22

### Changed
- ProcessRequestTTL don't changes status errors from TTLCondition

## [0.2.1] - 2019-11-22

### Changed
- Removed SendPutRequest
- MakePutRequestHeader sets only object and token

## [0.2.0] - 2019-11-21

### Added
- Container not found error
- GitHub Actions as CI and Codecov
- Auto-generated proto documentation
- RequestMetaHeader to all RPC requests
- RequestVerificationHeader to all RPC requests

### Changed
- Moved TTL and Epoch fields to RequestMetaHeader
- Renamed Version in object.SearchRequest to QueryVersion
- Removed SetTTL, GetTTL, SetEpoch, GetEpoch from all RPC requests

## 0.1.0 - 2019-11-18

Initial public release

[0.2.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.1.0...v0.2.0
[0.2.1]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.2.0...v0.2.1
[0.2.2]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.2.1...v0.2.2
[0.2.3]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.2.2...v0.2.3
[0.2.4]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.2.3...v0.2.4
[0.2.5]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.2.4...v0.2.5
[0.2.6]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.2.5...v0.2.6
[0.2.7]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.2.6...v0.2.7
[0.2.8]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.2.7...v0.2.8
[0.2.9]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.2.8...v0.2.9
[0.2.10]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.2.9...v0.2.10
[0.2.11]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.2.10...v0.2.11
[0.2.12]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.2.11...v0.2.12
[0.2.13]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.2.12...v0.2.13
[0.2.14]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.2.13...v0.2.14
[0.3.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.2.14...v0.3.0
[0.3.1]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.3.0...v0.3.1
[0.3.2]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.3.1...v0.3.2
[0.4.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.3.2...v0.4.0
[0.4.1]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.4.0...v0.4.1
[0.4.2]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.4.1...v0.4.2
[0.5.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.4.2...v0.5.0
[0.6.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.5.0...v0.6.0
[0.6.1]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.6.0...v0.6.1
[0.6.2]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.6.1...v0.6.2
[0.7.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.6.2...v0.7.0
[0.7.1]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.7.0...v0.7.1
[0.7.4]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.7.1...v0.7.4
[0.7.5]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.7.4...v0.7.5
[0.7.6]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.7.5...v0.7.6
[1.0.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v0.7.6...v1.0.0
[1.1.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.0.0...v1.1.0
[1.2.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.1.0...v1.2.0
[1.3.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.2.0...v1.3.0
[1.20.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.3.0...v1.20.0
[1.20.1]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.20.0...v1.20.1
[1.20.2]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.20.1...v1.20.2
[1.20.3]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.20.2...v1.20.3
[1.21.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.20.3...v1.21.0
[1.21.1]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.21.0...v1.21.1
[1.21.2]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.21.1...v1.21.2
[1.22.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.21.2...v1.22.0
[1.22.1]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.22.0...v1.22.1
[1.22.2]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.22.1...v1.22.2
[1.23.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.22.2...v1.23.0
[1.24.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.23.0...v1.24.0
[1.25.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.24.0...v1.25.0
[1.26.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.25.0...v1.26.0
[1.26.1]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.26.0...v1.26.1
[1.27.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.26.1...v1.27.0
[1.27.1]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.27.0...v1.27.1
[1.28.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.27.1...v1.28.0
[1.28.1]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.28.0...v1.28.1
[1.28.2]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.28.1...v1.28.2
[1.28.3]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.28.2...v1.28.3
[1.29.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.28.3...v1.29.0
[1.30.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.29.0...v1.30.0
[2.11.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v1.30.0...v2.11.0
[2.11.1]: https://github.com/nspcc-dev/neofs-api-go/compare/v2.11.0...v2.11.1
[2.12.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v2.11.1...v2.12.0
[2.12.1]: https://github.com/nspcc-dev/neofs-api-go/compare/v2.12.0...v2.12.1
[2.12.2]: https://github.com/nspcc-dev/neofs-api-go/compare/v2.12.1...v2.12.2
[2.13.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v2.12.2...v2.13.0
[2.13.1]: https://github.com/nspcc-dev/neofs-api-go/compare/v2.13.0...v2.13.1
[2.14.0]: https://github.com/nspcc-dev/neofs-api-go/compare/v2.13.1...v2.14.0
