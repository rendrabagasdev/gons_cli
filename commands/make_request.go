package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

var MakeRequestCmd = &cobra.Command{
	Use:   "make:request [name]",
	Short: "Create a new request class",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		createRequest(name)
	},
}

func createRequest(name string) {
	structName := strings.Title(name)
	if !strings.HasSuffix(structName, "Request") {
		structName += "Request"
	}
	
	// Convert to snake_case for filename
	fileName := ""
	for i, r := range structName {
		if i > 0 && r >= 'A' && r <= 'Z' {
			fileName += "_"
		}
		fileName += string(r)
	}
	fileName = strings.ToLower(fileName) + ".go"
	
	targetDir := "app/http/requests"
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Printf("Error: %s directory not found.\n", targetDir)
		return
	}

	targetPath := filepath.Join(targetDir, fileName)

	if _, err := os.Stat(targetPath); err == nil {
		fmt.Printf("Error: Request %s already exists!\n", fileName)
		return
	}

	tmplContent, err := TemplatesFS.ReadFile("template/request.tmpl")
	if err != nil {
		fmt.Printf("Error reading template: %v\n", err)
		return
	}

	tmpl, err := template.New("request").Parse(string(tmplContent))
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

	fmt.Printf("Request %s created successfully at %s\n", structName, targetPath)
}
