package db

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() error {
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".cldo")
	_ = os.MkdirAll(dir, 0755)

	path := filepath.Join(dir, "cldo.db")

	var err error
	DB, err = sql.Open("sqlite3", path)
	if err != nil {
		return err
	}

	return DB.Ping()
}
