package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

const (
	HTMLStarterRepo    = "https://github.com/your-username/gons-html-starter.git"
	InertiaStarterRepo = "https://github.com/your-username/gons-inertia-starter.git"
	APIStarterRepo     = "https://github.com/your-username/gons-api-starter.git"
)

var NewProjectCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new Gons project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		var selectedStack string
		prompt := &survey.Select{
			Message: "Choose the frontend/templating stack for your project:",
			Options: []string{
				"Pure HTML (Go Templates + Alpine.js + Tailwind)",
				"Inertia.js + React",
				"Pure API (No Frontend)",
			},
		}

		err := survey.AskOne(prompt, &selectedStack)
		if err != nil {
			fmt.Printf("Error selecting stack: %v\n", err)
			return
		}

		var repoURL string
		switch selectedStack {
		case "Pure HTML (Go Templates + Alpine.js + Tailwind)":
			repoURL = HTMLStarterRepo
		case "Inertia.js + React":
			repoURL = InertiaStarterRepo
		case "Pure API (No Frontend)":
			repoURL = APIStarterRepo
		}

		cloneRepository(repoURL, projectName)
	},
}

func cloneRepository(repoURL string, projectName string) {
	fmt.Printf("Creating new project: %s...\n", projectName)

	// Clone the repository
	cloneCmd := exec.Command("git", "clone", repoURL, projectName)
	cloneCmd.Stdout = nil
	cloneCmd.Stderr = nil

	err := cloneCmd.Run()
	if err != nil {
		fmt.Printf("Error cloning repository: %v\n", err)
		return
	}

	// Remove .git directory
	gitDir := filepath.Join(projectName, ".git")
	err = os.RemoveAll(gitDir)
	if err != nil {
		fmt.Printf("Warning: Could not remove .git directory: %v\n", err)
	}

	fmt.Printf("\nProject %s created successfully!\n", projectName)
	fmt.Printf("To get started:\n")
	fmt.Printf("  cd %s\n", projectName)
	fmt.Printf("  go mod tidy\n")
}
