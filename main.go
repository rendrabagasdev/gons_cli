package main

import (
	"fmt"
	"os"
	"runtime/debug" // Import ini

	"github.com/rendrabagasdev/gons/commands"
	"github.com/spf13/cobra"
)

var version = "dev" // Biarkan sebagai default untuk development lokal

var rootCmd = &cobra.Command{
	Use:   "gons",
	Short: "Gons CLI Framework",
	// Kita tidak set Version di sini secara hardcoded lagi
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
	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Version != "" && info.Main.Version != "(devel)" {
			rootCmd.Version = info.Main.Version
		} else {
			rootCmd.Version = version
		}
	}

	rootCmd.Flags().BoolP("version", "v", false, "show version")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
