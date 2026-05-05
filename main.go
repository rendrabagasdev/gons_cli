package main

import (
	"fmt"
	"os"

	"github.com/rendrabagasdev/gons/commands"
	"github.com/spf13/cobra"
)

var version = "v1.0.0"

var rootCmd = &cobra.Command{
	Use:     "gons",
	Short:   "Gons CLI Framework",
	Version: version,
}

func init() {
	rootCmd.AddCommand(commands.NewProjectCmd)
	rootCmd.AddCommand(commands.MakeModelCmd)
	rootCmd.AddCommand(commands.MakeControllerCmd)
	rootCmd.AddCommand(commands.MakeRequestCmd)
	rootCmd.AddCommand(commands.MakeServiceCmd)
	rootCmd.AddCommand(commands.MakeMiddlewareCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
