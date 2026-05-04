package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

var MakeServiceCmd = &cobra.Command{
	Use:   "make:service [name]",
	Short: "Create a new service class",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		createService(name)
	},
}

func createService(name string) {
	structName := strings.Title(name)
	if !strings.HasSuffix(structName, "Service") {
		structName += "Service"
	}
	fileName := structName + ".go"
	
	targetDir := "app/http/services"
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Printf("Error: %s directory not found.\n", targetDir)
		return
	}

	targetPath := filepath.Join(targetDir, fileName)

	if _, err := os.Stat(targetPath); err == nil {
		fmt.Printf("Error: Service %s already exists!\n", fileName)
		return
	}

	tmplContent, err := TemplatesFS.ReadFile("template/service.tmpl")
	if err != nil {
		fmt.Printf("Error reading template: %v\n", err)
		return
	}

	tmpl, err := template.New("service").Parse(string(tmplContent))
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

	fmt.Printf("Service %s created successfully at %s\n", structName, targetPath)
}
