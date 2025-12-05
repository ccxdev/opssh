package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	verbose    bool
	configPath string
	version    = "dev"
)

var rootCmd = &cobra.Command{
	Use:   "opssh",
	Short: "SSH Key Profile Switcher",
	Long: `opssh is a CLI tool for managing SSH key switching for 1Password's SSH agent.
It allows you to easily switch between different SSH key profiles (e.g., personal vs work)
by managing the agent.toml configuration file.`,
	Version: version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "", "config file path (default is ~/.opssh.yaml)")
}
