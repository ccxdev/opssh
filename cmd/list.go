package cmd

import (
	"fmt"

	"github.com/ccxdev/opssh/internal/agent"
	"github.com/ccxdev/opssh/internal/config"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available SSH key profiles",
	Long: `List all SSH key profiles stored in your config file.
The currently active profile will be marked with an asterisk (*).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		if len(cfg.Profiles) == 0 {
			fmt.Println("No profiles found. Use 'opssh add' to add a profile.")
			return nil
		}

		var currentProfile *config.Profile
		currentProfile, err = agent.GetCurrentActiveProfile()
		if err != nil {
			currentProfile = nil
		}

		fmt.Println("Available profiles:")
		fmt.Println()

		for name, profile := range cfg.Profiles {
			marker := " "
			if currentProfile != nil &&
				currentProfile.Account == profile.Account &&
				currentProfile.Vault == profile.Vault &&
				currentProfile.Item == profile.Item {
				marker = "*"
			}

			fmt.Printf("%s %s\n", marker, name)
			fmt.Printf("   Account: %s\n", profile.Account)
			fmt.Printf("   Vault: %s\n", profile.Vault)
			fmt.Printf("   Item: %s\n", profile.Item)
			fmt.Println()
		}

		if currentProfile != nil {
			fmt.Println("* = currently active")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

