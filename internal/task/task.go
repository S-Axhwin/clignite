package task

import (
	"cldo/internal/db"
	"os"
	"path/filepath" // for filepath.Rel
	"strings"       // for strings.HasPrefix
)

type Task struct {
	ID    int
	Title string
	State string
}

// Add a new task
func Add(title string) error {
	cwd, _ := os.Getwd()
	context := filepath.Clean(cwd) // normalize context path

	_, err := db.DB.Exec(
		"INSERT INTO tasks(title, state, context) VALUES(?, ?, ?)",
		title,
		"todo",
		context,
	)
	return err
}

func ListAllGrouped() (map[string][]Task, error) {
	rows, err := db.DB.Query(
		`SELECT id, title, state, context
		 FROM tasks
		 ORDER BY context, id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string][]Task)

	for rows.Next() {
		var t Task
		var ctx string
		if err := rows.Scan(&t.ID, &t.Title, &t.State, &ctx); err != nil {
			return nil, err
		}
		result[ctx] = append(result[ctx], t)
	}

	return result, nil
}

func ListAllUnderRoot(root string) ([]Task, map[int]string, error) {
	root = filepath.Clean(root) // normalize root
	rows, err := db.DB.Query(
		`SELECT id, title, state, context FROM tasks ORDER BY context, id`,
	)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var tasks []Task
	ctxMap := make(map[int]string)

	for rows.Next() {
		var t Task
		var ctx string
		if err := rows.Scan(&t.ID, &t.Title, &t.State, &ctx); err != nil {
			return nil, nil, err
		}

		ctxClean := filepath.Clean(ctx) // normalize DB path
		rel, err := filepath.Rel(root, ctxClean)
		if err != nil {
			continue
		}

		// Skip tasks outside the root directory tree
		// If rel starts with "..", it's outside; otherwise it's inside or at root
		if strings.HasPrefix(rel, ".."+string(filepath.Separator)) || rel == ".." {
			continue
		}

		tasks = append(tasks, t)
		ctxMap[t.ID] = ctxClean
	}

	return tasks, ctxMap, nil
}

func ListAllWithContext() ([]Task, map[int]string, error) {
	rows, err := db.DB.Query(
		`SELECT id, title, state, context FROM tasks ORDER BY context, id`,
	)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var tasks []Task
	contexts := make(map[int]string)

	for rows.Next() {
		var t Task
		var ctx string
		if err := rows.Scan(&t.ID, &t.Title, &t.State, &ctx); err != nil {
			return nil, nil, err
		}
		tasks = append(tasks, t)
		contexts[t.ID] = ctx
	}

	return tasks, contexts, nil
}

func ListByContext(context string) ([]Task, error) {
	rows, err := db.DB.Query(
		"SELECT id, title, state FROM tasks WHERE context = ?",
		context,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		rows.Scan(&t.ID, &t.Title, &t.State)
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func List() ([]Task, error) {
	rows, err := db.DB.Query("SELECT id, title, state FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		rows.Scan(&t.ID, &t.Title, &t.State)
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func UpdateState(id int, state string) error {
	_, err := db.DB.Exec(
		"UPDATE tasks SET state=? WHERE id=?",
		state, id,
	)
	return err
}

func Remove(id int) error {
	_, err := db.DB.Exec("DELETE FROM tasks WHERE id=?", id)
	return err
}
