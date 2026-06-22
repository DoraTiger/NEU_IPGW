# Changelog

[简体中文](./CHANGELOG.md) | [English](./CHANGELOG.en.md)

## [v0.3.2] - 2026-06-23

### Added
- Add `power` command to query dormitory electricity bills
- Unified output format: remaining kWh + balance (converted at 0.51 CNY/kWh)
- Auto-translate dormitory numbers (based on official lookup table)
- Support manual room number query (`-r` flag)
- Support room info only query (`-i` flag)
- Support full output (`-a` flag)
- Unified SSO login flow (CAS + service login)

### Changed
- Refactor SSO login logic as reusable component
- Centralize all config constants, eliminate hardcoded values

## [v0.3.1] - 2026-03-31

### Added
- Surface account, traffic, duration, balance, and IP after a successful `login`
- Introduce `info` subcommand to query the current online status without re-login
- Add bilingual output with `--lang` flag or `NEU_IPGW_LANG` env (defaults to Chinese, English optional)

### Fixed
- Handle numeric responses from `rad_user_info` to avoid JSON parsing failures on certain gateways

### Changed
- Share the localized formatter between login/info outputs to keep zh/en strings consistent

## [v0.3.0] - 2026-03-30

### Added
- Add local encrypted credential storage with optional `--save`
- Add multi-account management by username and `accounts list`
- Add `--account` selector for login/logout credential operations
- Add runtime overrides for credential directory and master key via environment variables
- Add `install` CLI subcommand for user/system installs on Linux/macOS/FreeBSD
- Add `install-user`, `install-system`, and `uninstall-user` Makefile targets with XDG-compliant layout

### Changed
- Expand cross-platform build matrix in `Makefile`
- Update documentation with Linux user-level installation layout and symlink guidance
- Restructure README install instructions to cover automated and manual flows per platform

## [v0.2.2] - 2026-03-10

### Added
- Implement RSA encryption for account credentials storage

### Changed
- Update account data structure to support encrypted credentials
- Enhance config module for secure credential handling 

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
