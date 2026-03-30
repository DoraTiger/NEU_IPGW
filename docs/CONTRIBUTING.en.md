# Contribution Guide

[简体中文](./CONTRIBUTING.md) | [English](./CONTRIBUTING.en.md)

Thank you for your interest in and contribution to this project! To ensure code quality and collaboration efficiency, please read this guide before submitting Issues or Pull Requests.

## How to Contribute

### 🐞 Report Issues

- Search in [Issues](https://github.com/doratiger/neu_ipgw/issues) to check if a similar issue already exists.
- Provide clear reproduction steps, environment information (such as Go version, operating system), and expected/actual behavior.

### 💡 Suggest Features

- Discuss the necessity and design of new features through Issues first to avoid ineffective development.
- Clearly describe the usage scenarios and expected behavior.

### 🧑‍💻 Contribute Code

1. Fork this repository
2. Create a feature branch (e.g. `feat/qr-login` or `fix/timeout-issue`)
3. Write code and ensure it passes tests
4. Follow the **Code Standards** and **Commit Message Format** below
5. Submit a Pull Request and link to relevant Issues (if any)

## Code Standards

### Go Code Style

This project follows the [Go Official Code Review Comments](https://go.dev/wiki/CodeReviewComments) and [Uber Go Style Guide](https://github.com/uber-go/guide/). Please ensure:
- Format code with `go fmt`
- Avoid unnecessary complex logic
- Use clear and meaningful function and variable names

### Chinese Copywriting Layout

All Chinese documents (including comments, README, CHANGELOG, etc.) should follow the [Chinese Copywriting Guidelines](https://github.com/sparanoid/chinese-copywriting-guidelines):
- Add spaces between Chinese and English/digits (e.g. `support iOS 15`)
- Use full-width Chinese punctuation
- Do not use non-standard abbreviations (e.g. "前端" should not be written as "FED")

## Commit Message Format

We adopt the [Conventional Commits](https://www.conventionalcommits.org/) specification with the following format:

```text
<type>: <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

### `type` Description

- feat: New feature
- fix: Bug fix
- docs: Documentation update (including README, CHANGELOG, etc.)
- style: Code style adjustment (such as formatting, spacing, etc., without logic changes)
- refactor: Code refactoring (without functional changes)
- perf: Performance optimization
- test: Add or modify test cases
- chore: Modify build process, or change dependencies and tools
- ci: Continuous integration configuration changes
- revert: Version rollback

### `subject` Requirements

- Concise title describing the summary of this commit
- Use **imperative** and **present tense** (e.g. "fix bug" instead of "Fixed bug")
- Start with **lowercase**
- **No period**
- Recommended length <= 50 characters

### `body` Requirements (optional but recommended)

- Explain **why** these changes were made, not just what was done
- <= 72 characters per line
- `feat`/`fix`/`refactor` commits are recommended to include body

### `footer` Requirements (optional)

- Related Issues or PRs (e.g. `Close #123`)

## Examples

1. `feat`/`refactor` commit example

```text
feat(auth): add QR code login

- implement ScanQR in internal/client
- add --qr flag to CLI
- update README usage example

Closes #45
```

2. `docs` commit example

```text
docs: update README with logout command
```

3. `fix` commit example

```text
fix(client): handle timeout in weak network

Previous implementation would hang indefinitely.
Now use context.WithTimeout to enforce 10s limit.

Closes #38
```

## Documentation and Tests

- **New features must update `README.md`**
- **User-visible changes should update `CHANGELOG.md`**
- **Breaking changes should be documented in `UPDATING.md`**
- It is encouraged to add unit tests (*_test.go) for key logic

## License

All contributions are considered to be released under the project's [MIT License](../LICENSE)

## Getting Help

If you have any questions, please raise them in [Discussions](https://github.com/DoraTiger/NEU_IPGW/discussions) or relevant Issues.
