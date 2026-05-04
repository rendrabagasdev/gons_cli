package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

var MakeControllerCmd = &cobra.Command{
	Use:   "make:controller [name]",
	Short: "Create a new controller class",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		createController(name)
	},
}

func createController(name string) {
	structName := strings.Title(name)
	if !strings.HasSuffix(structName, "Controller") {
		structName += "Controller"
	}
	fileName := structName + ".go"
	
	targetDir := "app/http/controllers"
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Printf("Error: %s directory not found.\n", targetDir)
		return
	}

	targetPath := filepath.Join(targetDir, fileName)

	if _, err := os.Stat(targetPath); err == nil {
		fmt.Printf("Error: Controller %s already exists!\n", fileName)
		return
	}

	tmplContent, err := TemplatesFS.ReadFile("template/controller.tmpl")
	if err != nil {
		fmt.Printf("Error reading template: %v\n", err)
		return
	}

	tmpl, err := template.New("controller").Parse(string(tmplContent))
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

	fmt.Printf("Controller %s created successfully at %s\n", structName, targetPath)
}
