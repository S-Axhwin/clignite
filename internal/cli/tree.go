package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"cldo/internal/task"
)

type treeNode struct {
	children map[string]*treeNode
	tasks    []task.Task
}

func printTree(tasks []task.Task, ctxMap map[int]string) {
	cwd, _ := os.Getwd()
	root := &treeNode{children: make(map[string]*treeNode)}

	for _, t := range tasks {
		relPath, err := filepath.Rel(cwd, ctxMap[t.ID])
		if err != nil {
			relPath = ctxMap[t.ID]
		}
		relPath = filepath.Clean(relPath)

		node := root

		if relPath != "." {
			parts := strings.Split(relPath, string(filepath.Separator))
			for _, p := range parts {
				if p == "" {
					continue
				}
				if node.children[p] == nil {
					node.children[p] = &treeNode{children: make(map[string]*treeNode)}
				}
				node = node.children[p]
			}
		}

		// always append task (including root folder tasks)
		node.tasks = append(node.tasks, t)
	}

	// Print tree with current folder as root
	rootName := filepath.Base(cwd)
	fmt.Println(rootName)
	printNode(root, "", false)
}

func printNode(node *treeNode, prefix string, isRoot bool) {
	for name, child := range node.children {
		if !isRoot {
			fmt.Println(prefix + "├── " + name)
		} else {
			fmt.Println(name)
		}

		newPrefix := prefix
		if !isRoot {
			newPrefix += "│   "
		}

		for _, t := range child.tasks {
			fmt.Printf("%s│   [%d][%s] %s\n", newPrefix, t.ID, t.State, t.Title)
		}

		printNode(child, newPrefix, false)
	}
}
