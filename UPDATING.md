# Upgrading NEU_IPGW

This guide helps you migrate from older versions of NEU_IPGW to newer ones.

## Upgrading from v0.1.x to v0.2.0

### ⚠️ Breaking Changes

Version 0.2.0 introduces **breaking changes** to the command-line interface and project structure.

#### 1. Command Structure Changed

The CLI now uses subcommands. You must update your scripts and usage.

| Old (v0.1.x) | New (v0.2.0) |
|--------------|--------------|
| `neu-ipgw --username <user> --password <pass>` | `neu-ipgw login -u <user> -p <pass>` |
| `neu-ipgw --version` | `neu-ipgw version` |

> 💡 Note: The `logout` command is newly added in v0.2.0:
> ```bash
> neu-ipgw logout
> ```

### ✅ Migration Steps

1. Replace all usages of the old flat flags with the new `login` subcommand.
2. Add `logout` to your workflow if needed.
3. Update any automation scripts (shell, CI, etc.) accordingly.
