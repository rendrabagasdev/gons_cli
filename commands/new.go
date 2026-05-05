package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var NewProjectCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Create a new Gons project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		createNewProject(name)
	},
}

func createNewProject(name string) {
	repoURL := "https://github.com/rendrabagasdev/gons_framework"

	fmt.Printf("Creating new project: %s...\n", name)

	// 1. Clone the repository
	cloneCmd := exec.Command("git", "clone", repoURL, name)
	cloneCmd.Stdout = os.Stdout
	cloneCmd.Stderr = os.Stderr

	err := cloneCmd.Run()
	if err != nil {
		fmt.Printf("Error cloning repository: %v\n", err)
		return
	}

	// 2. Remove .git directory
	gitDir := filepath.Join(name, ".git")
	err = os.RemoveAll(gitDir)
	if err != nil {
		fmt.Printf("Warning: Could not remove .git directory: %v\n", err)
	}

	fmt.Printf("\nProject %s created successfully!\n", name)
	fmt.Printf("To get started:\n")
	fmt.Printf("  cd %s\n", name)
	fmt.Printf("  go mod tidy\n")
}
