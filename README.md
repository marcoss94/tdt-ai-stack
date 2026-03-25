<div align="center">

<h1>TDT AI Stack</h1>

<p><strong>One command. Any agent. Any OS. The TDT AI stack -- configured and ready.</strong></p>

<p>
<a href="https://github.com/marcoss94/tdt-ai-stack/releases"><img src="https://img.shields.io/github/v/release/marcoss94/tdt-ai-stack" alt="Release"></a>
<a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License: MIT"></a>
<img src="https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go&logoColor=white" alt="Go 1.24+">
<img src="https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey" alt="Platform">
</p>

</div>

---

## What It Does

TDT AI Stack is the public distribution for **`tdt-ai`**. Some internal technical IDs still remain legacy (mainly Go module/import paths and a few compatibility assets), but the supported public install flow today is GitHub Releases plus the install scripts in this repository.

This is NOT an AI agent installer. Most agents are easy to install. This is an **ecosystem configurator** -- it takes whatever AI coding agent(s) you use and supercharges them with the TDT stack: persistent memory, Spec-Driven Development workflow, curated coding skills, MCP servers, an AI provider switcher, and a teaching-oriented persona with security-first permissions.

**Before**: "I installed Claude Code / OpenCode / Cursor, but it's just a chatbot that writes code."

**After**: Your agent now has memory, skills, workflow, MCP tools, and a persona that actually teaches you.

---

## Quick Start

### macOS / Linux

```bash
curl -fsSL https://raw.githubusercontent.com/marcoss94/tdt-ai-stack/main/scripts/install.sh | bash
```

### Windows (PowerShell)

```powershell
irm https://raw.githubusercontent.com/marcoss94/tdt-ai-stack/main/scripts/install.ps1 | iex
```

This downloads the latest published release for your platform, installs the public `tdt-ai` binary, and leaves you ready to run `tdt-ai` to launch the interactive TUI. No Go toolchain required.

---

## Install

### macOS / Linux

```bash
curl -fsSL https://raw.githubusercontent.com/marcoss94/tdt-ai-stack/main/scripts/install.sh | bash
```

### Windows (PowerShell)

```powershell
irm https://raw.githubusercontent.com/marcoss94/tdt-ai-stack/main/scripts/install.ps1 | iex
```

### From releases

Download the binary for your platform from [GitHub Releases](https://github.com/marcoss94/tdt-ai-stack/releases), place it somewhere in your `PATH`, and run `tdt-ai`.

Supported public install paths today:

- install scripts from `marcoss94/tdt-ai-stack`
- direct binary download from GitHub Releases

Publicly supported command:

- `tdt-ai`

Not currently documented as public install paths:

- `brew`
- `go install`

Those legacy or internal routes stay undocumented until they are actually supported again on the public surface.

---

## Documentation

| Topic | Description |
|-------|-------------|
| [Agents](docs/agents.md) | Supported agents, feature matrix, config paths, and per-agent notes |
| [Components, Skills & Presets](docs/components.md) | All components, GGA behavior, skill catalog, and preset definitions |
| [Usage](docs/usage.md) | Persona modes, interactive TUI, CLI flags, and dependency management |
| [Platforms](docs/platforms.md) | Supported platforms, Windows notes, security verification, config paths |
| [Architecture & Development](docs/architecture.md) | Codebase layout, testing, and relationship to Gentleman.Dots / the upstream base |

---

<div align="center">
<a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License: MIT"></a>
</div>
