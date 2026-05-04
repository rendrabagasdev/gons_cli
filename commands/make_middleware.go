package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

var MakeMiddlewareCmd = &cobra.Command{
	Use:   "make:middleware [name]",
	Short: "Create a new middleware class",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		createMiddleware(name)
	},
}

func createMiddleware(name string) {
	// Format name: Auth -> AuthMiddleware
	baseName := strings.Title(name)
	if !strings.HasSuffix(baseName, "Middleware") {
		baseName += "Middleware"
	}

	// Filename: auth_middleware.go
	fileName := ""
	for i, r := range baseName {
		if i > 0 && r >= 'A' && r <= 'Z' {
			fileName += "_"
		}
		fileName += string(r)
	}
	fileName = strings.ToLower(fileName) + ".go"

	targetDir := "app/http/middlewares"
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Printf("Error: %s directory not found.\n", targetDir)
		return
	}

	targetPath := filepath.Join(targetDir, fileName)

	if _, err := os.Stat(targetPath); err == nil {
		fmt.Printf("Error: Middleware %s already exists!\n", fileName)
		return
	}

	tmplContent, err := TemplatesFS.ReadFile("template/middleware.tmpl")
	if err != nil {
		fmt.Printf("Error reading template: %v\n", err)
		return
	}

	tmpl, err := template.New("middleware").Parse(string(tmplContent))
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

	// FuncName: AuthMiddleware
	data := struct {
		FuncName string
	}{
		FuncName: baseName,
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return
	}

	fmt.Printf("Middleware %s created successfully at %s\n", baseName, targetPath)
}
