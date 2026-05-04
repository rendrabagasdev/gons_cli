package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

var MakeModelCmd = &cobra.Command{
	Use:   "make:model [name]",
	Short: "Create a new model class",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		createModel(name)
	},
}

func createModel(name string) {
	// Capitalize name for struct
	structName := strings.Title(name)
	fileName := structName + ".go"
	
	// Determine target path
	targetDir := "app/models"
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Printf("Error: %s directory not found. Are you in the root of gons framework?\n", targetDir)
		return
	}

	targetPath := filepath.Join(targetDir, fileName)

	// Check if file already exists
	if _, err := os.Stat(targetPath); err == nil {
		fmt.Printf("Error: Model %s already exists!\n", fileName)
		return
	}

	// Read template (embedded or from a known path)
	// For now, let's assume it's in a 'template' folder relative to the executable or a fixed path
	// Or we can just hardcode the string for simplicity in this task
	tmplContent, err := TemplatesFS.ReadFile("template/model.tmpl")
	if err != nil {
		fmt.Printf("Error reading template: %v\n", err)
		return
	}

	tmpl, err := template.New("model").Parse(string(tmplContent))
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	file, err := os.Create(targetPath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	data := struct {
		Name string
	}{
		Name: structName,
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return
	}

	fmt.Printf("Model %s created successfully at %s\n", structName, targetPath)
}
