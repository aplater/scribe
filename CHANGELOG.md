# Changelog
All notable changes to this project will be documented in this file.

## [1.1.1] - 2019-06-23
### Added
- Add ca-certificates to alpine image in Dockerfile

### Changed
- Fix change creation when there aren't Work Items associated to a deploy
- Add Azure Devops terminology and improve documentation

## [1.1.0] - 2018-12-08
### Added
- Add Work Items to Freshservice change when available
- Add more service info to health endpoint
- Improve log messages
- Add Work Items from all release artifacts and improve deploy items message  

### Changed
- Refactor TFS API client 

## 1.0.0 - 2018-12-07
### Added
- Base project, initial structure.
- Freshservice API client
- Main session workflow for VSTS deployment events
- API service health endpoint
- Basic auth for `/vss` subroutes
- Multi-stage Dockerfile

[Unreleased]: https://github.com/payvision-development/scribe/compare/v1.1.1...HEAD
[1.1.1]: https://github.com/payvision-development/scribe/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/payvision-development/scribe/compare/v1.0.0...v1.1.0
