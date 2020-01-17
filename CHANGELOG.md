# Changelog
This is the changelog for NeoFS Proto

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

[0.2.0]: https://github.com/nspcc-dev/neofs-proto/compare/v0.1.0...v0.2.0
[0.2.1]: https://github.com/nspcc-dev/neofs-proto/compare/v0.2.0...v0.2.1
[0.2.2]: https://github.com/nspcc-dev/neofs-proto/compare/v0.2.1...v0.2.2
[0.2.3]: https://github.com/nspcc-dev/neofs-proto/compare/v0.2.2...v0.2.3
[0.2.4]: https://github.com/nspcc-dev/neofs-proto/compare/v0.2.3...v0.2.4
[0.2.5]: https://github.com/nspcc-dev/neofs-proto/compare/v0.2.4...v0.2.5
[0.2.6]: https://github.com/nspcc-dev/neofs-proto/compare/v0.2.5...v0.2.6
[0.2.7]: https://github.com/nspcc-dev/neofs-proto/compare/v0.2.6...v0.2.7
[0.2.8]: https://github.com/nspcc-dev/neofs-proto/compare/v0.2.7...v0.2.8
[0.2.9]: https://github.com/nspcc-dev/neofs-proto/compare/v0.2.8...v0.2.9
