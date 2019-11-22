# Changelog
This is the changelog for NeoFS Proto

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
