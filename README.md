```regex
██╗   ██╗ ██╗   ██╗ ██╗      ███╗   ██╗ ███████╗  ██████╗  ██████╗   ██████╗ ███████╗
██║   ██║ ██║   ██║ ██║      ████╗  ██║ ██╔════╝ ██╔═══██╗ ██╔══██╗ ██╔════╝ ██╔════╝
██║   ██║ ██║   ██║ ██║      ██╔██╗ ██║ █████╗   ██║   ██║ ██████╔╝ ██║  ███╗ █████╗
╚██╗ ██╔╝ ██║   ██║ ██║      ██║╚██╗██║ ██╔══╝   ██║   ██║ ██╔══██╗ ██║   ██║ ██╔══╝
 ╚████╔╝  ╚██████╔╝ ███████╗ ██║ ╚████║ ██║      ╚██████╔╝ ██║  ██║ ╚██████╔╝ ███████╗
  ╚═══╝    ╚═════╝  ╚══════╝ ╚═╝  ╚═══╝ ╚═╝       ╚═════╝  ╚═╝  ╚═╝  ╚═════╝  ╚══════╝
```

[![Cybersecurity Projects](https://img.shields.io/badge/Cybersecurity--Projects-Project%20%2311-red?style=flat&logo=github)](https://github.com/Kiliankm19/vulnforge)
[![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go&logoColor=white)](https://go.dev)
[![License: AGPLv3](https://img.shields.io/badge/License-AGPL_v3-purple.svg)](https://www.gnu.org/licenses/agpl-3.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/Kiliankm19/vulnforge)](https://goreportcard.com/report/github.com/Kiliankm19/vulnforge)
[![OSV.dev](https://img.shields.io/badge/OSV.dev-integrated-4285F4?style=flat)](https://osv.dev)

> Fast Python dependency updater and vulnerability scanner written in Go.

*This is a quick overview — security theory, architecture, and full walkthroughs are in the [learn modules](#learn).*

## What It Does

- Scans pyproject.toml and requirements.txt for known CVEs via OSV.dev
- Updates all Python dependencies to latest stable versions in one command
- Parallel queries against PyPI with local ETag caching for speed
- Full PEP 440 version parsing with automatic pre-release filtering
- Comment-preserving file updates that keep your formatting intact
- Configurable via .vulnforge.toml or [tool.vulnforge] in pyproject.toml

## Quick Start

```bash
go install github.com/Kiliankm19/vulnforge/cmd/vulnforge@latest
vulnforge scan
```

> [!TIP]
> This project uses [`just`](https://github.com/casey/just) as a command runner. Type `just` to see all available commands.
>
> Install: `curl -sSf https://just.systems/install.sh | bash -s -- --to ~/.local/bin`

## Commands

| Command | Description |
|---------|-------------|
| `vulnforge init` | Initialize a new .vulnforge.toml configuration file |
| `vulnforge update` | Update all Python dependencies to latest stable versions |
| `vulnforge check` | Preview available updates without modifying files |
| `vulnforge scan` | Scan dependencies for known CVEs via OSV.dev |
| `vulnforge cache clear` | Clear the local ETag and version cache |

## Learn

This project includes step-by-step learning materials covering security theory, architecture, and implementation.

| Module | Topic |
|--------|-------|
| [00 - Overview](learn/00-OVERVIEW.md) | Prerequisites and quick start |
| [01 - Concepts](learn/01-CONCEPTS.md) | Security theory and real-world breaches |
| [02 - Architecture](learn/02-ARCHITECTURE.md) | System design and data flow |
| [03 - Implementation](learn/03-IMPLEMENTATION.md) | Code walkthrough |
| [04 - Challenges](learn/04-CHALLENGES.md) | Extension ideas and exercises |


## License

AGPL 3.0
