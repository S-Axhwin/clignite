# Devtodo CLI

**Devtodo** is a terminal-based developer todo list manager.  
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

devtodo/
├─ cmd/
│   └─ devtodo/
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
| `devtodo add "task title"` | Add a new task |
| `devtodo ls`               | List all tasks |
| `devtodo rm <id>`          | Remove a task |
| `devtodo done <id>`        | Mark task as done |
| `devtodo start <id>`       | Mark task as in progress |
| `devtodo block <id>`       | Mark task as blocked |
| `devtodo undo <id>`        | Undo last state change |

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
devtodo add "Fix login bug"

# List all tasks
devtodo ls

# Mark a task as in progress
devtodo start 1

# Mark a task as blocked
devtodo block 2

# Mark a task as done
devtodo done 1

# Remove a task
devtodo rm 3

# Undo last state change
devtodo undo 2

# Optional: AI-assisted add
devtodo add "Optimize database queries" --ai
# Output: Suggests tags and priority automatically
````

---

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/devtodo.git
cd devtodo
```

2. Build the binary:

```bash
go build -o devtodo ./cmd/devtodo
```

3. Move binary to a folder in your PATH (optional):

```bash
sudo mv devtodo /usr/local/bin/
```

Now you can run `devtodo` from anywhere.

---

## Usage

* Database is stored in `~/.devtodo/devtodo.db`
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

**Devtodo** is designed for developers who prefer working in terminal environments, keeping task management fast, lightweight, and integrated with their workflow.

```
