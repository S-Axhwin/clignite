# cldo CLI

**cldo** is a fast, terminal-based todo manager built for developers.
It lets you track tasks, manage states and priorities, and optionally use AI suggestions — entirely from the command line.

No servers. No accounts. One binary.

---

## Features

* Add, list, update, and remove tasks from the CLI
* Task states: `todo`, `in-progress`, `blocked`, `done`
* Task priorities: `low`, `mid`, `high`
* Optional AI-assisted tag & priority suggestions
* Context-aware tasks (per project / directory)
* Persistent local storage using SQLite
* Cross-platform binaries (macOS, Linux, Windows)
* Undo task state changes

---

## Installation (Recommended)

### macOS (Intel & Apple Silicon)

```bash
curl -LO https://github.com/S-Axhwin/clignite/releases/download/v0.1.0/cldo_0.1.0_darwin_amd64.tar.gz
tar -xzf cldo_0.1.0_darwin_amd64.tar.gz

sudo mv cldo /usr/local/bin/
```

---

### Linux

```bash
curl -LO https://github.com/S-Axhwin/clignite/releases/download/v0.1.0/cldo_0.1.0_linux_amd64.tar.gz
tar -xzf cldo_0.1.0_linux_amd64.tar.gz
sudo mv cldo /usr/local/bin/
```

ARM64 (servers / Raspberry Pi):

```bash
curl -LO https://github.com/S-Axhwin/clignite/releases/download/v0.1.0/cldo_0.1.0_linux_arm64.tar.gz
tar -xzf cldo_0.1.0_linux_arm64.tar.gz
sudo mv cldo /usr/local/bin/
```

---

### Windows (PowerShell)

```powershell
Invoke-WebRequest `
  -Uri https://github.com/S-Axhwin/clignite/releases/download/v0.1.0/cldo_0.1.0_windows_amd64.tar.gz `
  -OutFile cldo.tar.gz

tar -xzf cldo.tar.gz
Move-Item cldo.exe C:\Windows\System32\
tar -xf cldo.tar.gz
Move-Item cldo.exe C:\Windows\System32\
```

Restart terminal, then:

```powershell
cldo --help
```

---

## Verify Installation

```bash
cldo --version
```

You should see something like:

```
cldo v0.1.0
```

If this fails, installation is broken.

---

## Usage

### Add a task

```bash
cldo add "Fix login bug"
```

### List tasks

```bash
cldo ls
```

### Start a task

```bash
cldo start 1
```

### Block a task

```bash
cldo block 2
```

### Complete a task

```bash
cldo done 1
```

### Undo last change

```bash
cldo undo 1
```

### AI-assisted task creation

```bash
cldo add "Optimize database queries" --ai
```

---

## Data Storage

* Database path: `~/.cldo/cldo.db`
* Fully local, no cloud dependency
* Safe to back up or delete anytime

---

## Folder Structure (For Contributors)

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
└─ README.md
```

---

## Development Install (Contributors Only)

```bash
git clone https://github.com/S-Axhwin/clignite.git
cd clignite
go build -o cldo ./cmd/cldo
```

---

## Roadmap

* Homebrew install (`brew install cldo`)
* Interactive TUI mode
* Task sync & export
* Plugin system

---