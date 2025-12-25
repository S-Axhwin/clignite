package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbPath := getDBPath()

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable(db)

	addTask(db, "First task from CLI")
	listTasks(db)
}

// ---------- helpers ----------

func getDBPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	dir := filepath.Join(home, ".devtodo")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(dir, "devtodo.db")
}

func createTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		state TEXT NOT NULL,
		priority TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func addTask(db *sql.DB, title string) {
	query := `
	INSERT INTO tasks (title, state, priority)
	VALUES (?, ?, ?);
	`

	_, err := db.Exec(query, title, "todo", "mid")
	if err != nil {
		log.Fatal(err)
	}
}

func listTasks(db *sql.DB) {
	rows, err := db.Query(`SELECT id, title, state, priority FROM tasks`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Tasks:")
	for rows.Next() {
		var id int
		var title, state, priority string

		rows.Scan(&id, &title, &state, &priority)
		fmt.Printf("[%d] %s (%s, %s)\n", id, title, state, priority)
	}
}
