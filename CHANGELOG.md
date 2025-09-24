# Changelog

## [v0.2.1] - 2025-09-25

### Changed
- Optimize build process and improve Makefile configuration
- Add SHA256 checksum generation mechanism for releases

### Added
- Integrate automatic SHA256 checksum calculation and generation for binary files in release script
- Provide unified checksums.txt file for each release version

## [v0.2.0] - 2023-04-22

### Changed
- **BREAKING**: Switch to subcommand-based CLI (`login`, `logout`, `version`)

### Added
- Support `logout` command
- Flags for `login` : `-u` for username, `-p` for password
- FLags for `version` : `--verbose` for detailed version info

### Removed
- Flat-style flags (`--username`, `--password` at root level)

## [v0.1.0] - 2023-04-18

### Added
- Initial release
- Login via `--username` and `--password`
- Version info via `--version`
