# Changelog
This is the changelog for NeoFS API

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
