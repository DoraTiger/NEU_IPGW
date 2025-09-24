# Changelog

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
