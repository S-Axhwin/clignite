# cldo CLI

**cldo** is a terminal-based developer todo list manager.  
It helps developers track tasks, manage states, assign priorities, and optionally use AI to suggest tags or priorities — all from the command line.

---

## Table of Contents

- [Features](#features)  
- [Folder Structure](#folder-structure)  
- [Functional Requirements](#functional-requirements)  
- [Sample Commands](#sample-commands)  
- [Installation](#installation)  
- [Usage](#usage)  
- [Future Enhancements](#future-enhancements)  

---

## Features

- Add, list, and remove tasks from CLI  
- Track task states: `todo`, `in-progress`, `blocked`, `done`  
- Assign task priorities: `low`, `mid`, `high`  
- Optional AI-assisted tag and priority suggestions  
- Context-aware task listing by project or directory  
- Persistent storage using SQLite  
- Cross-platform binary support (macOS, Linux)  
- Undo task state changes  

---

## Folder Structure

```

cldo/
├─ cmd/
│   └─ cldo/
│       └─ main.go        # CLI entrypoint, parses commands
├─ internal/
│   ├─ db/
│   │   ├─ db.go          # SQLite connection & setup
│   │   └─ task_table.go  # Table creation / migrations
│   ├─ task/
│   │   ├─ task.go        # Task struct & CRUD operations
│   │   └─ task_ai.go     # Optional AI integration
│   └─ cli/
│       └─ parser.go      # Command parsing & validation
├─ go.mod                 # Go module definition
├─ go.sum                 # Dependency lock file
├─ README.md              # Project documentation
└─ Makefile / build.sh    # Build & packaging scripts

````

---

## Functional Requirements

### Core Commands

| Command                  | Description |
|---------------------------|------------|
| `cldo add "task title"` | Add a new task |
| `cldo ls`               | List all tasks |
| `cldo rm <id>`          | Remove a task |
| `cldo done <id>`        | Mark task as done |
| `cldo start <id>`       | Mark task as in progress |
| `cldo block <id>`       | Mark task as blocked |
| `cldo undo <id>`        | Undo last state change |

### Task Attributes

| Field       | Type        | Default / Notes |
|-------------|------------|----------------|
| id          | integer    | Auto-increment |
| title       | string     | Required       |
| state       | string     | Default: `todo`; options: `todo`, `in-progress`, `blocked`, `done` |
| priority    | string     | Default: `mid`; options: `low`, `mid`, `high` |
| tags        | array/text | Optional; user-defined (`backend`, `bug`, etc.) |
| context     | string     | Optional; e.g., current directory or project |
| created_at  | datetime   | Auto timestamp |

---

## Sample Commands

```bash
# Add a task
cldo add "Fix login bug"

# List all tasks
cldo ls

# Mark a task as in progress
cldo start 1

# Mark a task as blocked
cldo block 2

# Mark a task as done
cldo done 1

# Remove a task
cldo rm 3

# Undo last state change
cldo undo 2

# Optional: AI-assisted add
cldo add "Optimize database queries" --ai
# Output: Suggests tags and priority automatically
````

---

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/cldo.git
cd cldo
```

2. Build the binary:

```bash
go build -o cldo ./cmd/cldo
```

3. Move binary to a folder in your PATH (optional):

```bash
sudo mv cldo /usr/local/bin/
```

Now you can run `cldo` from anywhere.

---

## Usage

* Database is stored in `~/.cldo/cldo.db`
* CLI commands are self-contained and cross-platform
* AI integration (optional) requires an API key if using an external service

---

## Future Enhancements

* Full AI-assisted task categorization
* Context-aware filtering (`ls --tag backend --state todo`)
* Batch import/export of tasks
* Interactive CLI mode (like `git add -p`)
* Windows binary support

---

**cldo** is designed for developers who prefer working in terminal environments, keeping task management fast, lightweight, and integrated with their workflow.

```
