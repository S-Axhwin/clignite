# cldo CLI

**cldo** is a fast, terminal-based todo manager built for developers.
It helps you track tasks, manage states and priorities, and optionally use AI suggestions — entirely from the command line.

No servers. No accounts. One local binary per OS.

---

## Features

* Add, list, update, and remove tasks from the CLI
* Task states: `todo`, `in-progress`, `blocked`, `done`
* Task priorities: `low`, `mid`, `high`
* Context-aware tasks (per project / directory)
* Tree view to visualize tasks per directory
* Persistent local storage using SQLite
* Undo task state changes
* Single-binary CLI (per OS)

---

## Platform Support (Important)

> **cldo uses SQLite via CGO**

Because of this:

* ✅ macOS — supported (Intel + Apple Silicon)
* ❌ Linux — not yet available
* ❌ Windows — not yet available

Linux and Windows support will be added later via **GitHub Actions**.

---

## Installation (macOS only)

### Intel (x86_64)

```bash
curl -LO https://github.com/S-Axhwin/clignite/releases/download/v0.2.0/cldo_0.2.0_darwin_amd64.tar.gz
tar -xzf cldo_0.2.0_darwin_amd64.tar.gz
sudo mv cldo /usr/local/bin/
```

### Apple Silicon (ARM64)

```bash
curl -LO https://github.com/S-Axhwin/clignite/releases/download/v0.2.0/cldo_0.2.0_darwin_arm64.tar.gz
tar -xzf cldo_0.2.0_darwin_arm64.tar.gz
sudo mv cldo /usr/local/bin/
```

---

## Verify Installation

```bash
cldo --version
```

Expected output:

```
cldo v0.2.0
```

---

## Usage

Run commands like this:

```bash
go run cmd/cldo/main.go
```

or after installing the binary:

```bash
cldo <command>
```

---

### Commands

| Command                 | Description                              |
| ----------------------- | ---------------------------------------- |
| `cldo add "task title"` | Add a new task in the current directory  |
| `cldo ls`               | List tasks in the current directory      |
| `cldo ls -a`            | List **all tasks** grouped by directory  |
| `cldo tree`             | Show tasks in **tree view** by directory |
| `cldo done <id>`        | Mark task as done                        |
| `cldo rm <id>`          | Remove a task                            |

---

## Data Storage

* Database path: `~/.cldo/cldo.db`
* Fully local, no cloud dependency
* Safe to back up or delete anytime

---

## Folder Structure (Contributors)

```
cldo/
├─ cmd/cldo/main.go        # CLI entrypoint
├─ internal/
│  ├─ db/                 # SQLite setup & migrations
│  ├─ task/               # Task logic & AI integration
│  └─ cli/                # Command parsing
├─ go.mod
├─ go.sum
├─ Makefile
├─ .goreleaser.yml
└─ README.md
```

---

## Development Setup (Contributors Only)

```bash
git clone https://github.com/S-Axhwin/clignite.git
cd clignite
CGO_ENABLED=1 go build -o cldo ./cmd/cldo
```

> SQLite **requires CGO**. Builds with `CGO_ENABLED=0` will fail.

---

## Roadmap

* GitHub Actions–based releases (Linux + Windows)
* Homebrew install (`brew install cldo`)
* Interactive TUI mode
* Task filtering & export
* Plugin system
* Optional AI-assisted features

---

## License

MIT
