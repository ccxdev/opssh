package cmd

import (
	"fmt"

	"github.com/ccxdev/opssh/internal/agent"
	"github.com/ccxdev/opssh/internal/config"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <profile>",
	Short: "Remove an SSH key profile",
	Long: `Remove an SSH key profile from your config file.
If the profile is currently active, you will be warned.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		profileName := args[0]

		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		profile, err := config.GetProfile(cfg, profileName)
		if err != nil {
			return fmt.Errorf("profile '%s' not found", profileName)
		}

		currentProfile, err := agent.GetCurrentActiveProfile()
		if err == nil {
			if currentProfile.Account == profile.Account &&
				currentProfile.Vault == profile.Vault &&
				currentProfile.Item == profile.Item {
				fmt.Printf("Warning: Profile '%s' is currently active. Removing it may cause issues.\n", profileName)
			}
		}

		if err := config.RemoveProfile(cfg, profileName); err != nil {
			return fmt.Errorf("failed to remove profile: %w", err)
		}

		if err := config.SaveConfig(cfg); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		fmt.Printf("✓ Removed profile '%s'\n", profileName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

