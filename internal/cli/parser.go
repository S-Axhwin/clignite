package cli

import (
	"cldo/internal/db"
	"cldo/internal/task"
	"fmt"
	"os"
	"strconv"
)

func Run() error {
	if err := db.Init(); err != nil {
		return err
	}
	if err := db.Migrate(); err != nil {
		return err
	}

	if len(os.Args) < 2 {
		printHelp()
		return nil
	}

	ctx, err := currentContext()
	if err != nil {
		return err
	}

	switch os.Args[1] {
	case "add":
		return task.Add(os.Args[2])

	case "tree":
		cwd, _ := os.Getwd() // <-- current folder
		tasks, ctxMap, err := task.ListAllUnderRoot(cwd)
		if err != nil {
			return err
		}
		printTree(tasks, ctxMap)
		return nil
	case "ls":
		if len(os.Args) > 2 && os.Args[2] == "-a" {
			grouped, err := task.ListAllGrouped()
			if err != nil {
				return err
			}
			for ctx, tasks := range grouped {
				fmt.Println(ctx)
				for _, t := range tasks {
					fmt.Printf("  %d [%s] %s\n", t.ID, t.State, t.Title)
				}
				fmt.Println()
			}
			return nil
		}

		tasks, err := task.ListByContext(ctx)
		if err != nil {
			return err
		}
		for _, t := range tasks {
			fmt.Printf("%d [%s] %s\n", t.ID, t.State, t.Title)
		}

	case "done":
		id, _ := strconv.Atoi(os.Args[2])
		return task.UpdateState(id, "done")
	case "rm":
		id, _ := strconv.Atoi(os.Args[2])
		return task.Remove(id)
	default:
		printHelp()
	}

	return nil
}

func printHelp() {
	fmt.Println(`
cldo add "task title"
cldo ls           # list tasks in current directory
cldo ls -a        # list ALL tasks grouped by directory
cldo tree         # show tasks in tree view by directory
cldo done <id>
cldo rm <id>
`)
}
