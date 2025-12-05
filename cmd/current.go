package cmd

import (
	"fmt"

	"github.com/ccxdev/opssh/internal/agent"
	"github.com/ccxdev/opssh/internal/config"
	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show the currently active SSH key profile",
	Long: `Display the currently active SSH key profile by reading the agent.toml file
and matching it against your configured profiles.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		activeProfile, err := agent.GetCurrentActiveProfile()
		if err != nil {
			return fmt.Errorf("failed to get current profile: %w", err)
		}

		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		var profileName string
		for name, profile := range cfg.Profiles {
			if profile.Account == activeProfile.Account &&
				profile.Vault == activeProfile.Vault &&
				profile.Item == activeProfile.Item {
				profileName = name
				break
			}
		}

		if profileName != "" {
			fmt.Printf("Current profile: %s\n", profileName)
		} else {
			fmt.Println("Current profile (not in config):")
		}

		fmt.Printf("  Account: %s\n", activeProfile.Account)
		fmt.Printf("  Vault: %s\n", activeProfile.Vault)
		fmt.Printf("  Item: %s\n", activeProfile.Item)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}

