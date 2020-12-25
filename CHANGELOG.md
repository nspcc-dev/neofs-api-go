# Changelog

## [1.22.0] - 2020-12-XX - Yeouido (여의도, 汝矣島)

Support changes from NeoFS API Yeouido.

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
