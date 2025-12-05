package cmd

import (
	"fmt"

	"github.com/ccxdev/opssh/internal/agent"
	"github.com/ccxdev/opssh/internal/config"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch <profile>",
	Short: "Switch to a different SSH key profile",
	Long: `Switch to a different SSH key profile by updating the agent.toml file.
The specified profile must exist in your config file.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		profileName := args[0]

		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		profile, err := config.GetProfile(cfg, profileName)
		if err != nil {
			return fmt.Errorf("profile '%s' not found. Use 'opssh list' to see available profiles", profileName)
		}

		if err := agent.WriteAgentConfig(profile); err != nil {
			return fmt.Errorf("failed to update agent.toml: %w", err)
		}

		fmt.Printf("✓ Switched to profile '%s'\n", profileName)
		fmt.Printf("  Account: %s\n", profile.Account)
		fmt.Printf("  Vault: %s\n", profile.Vault)
		fmt.Printf("  Item: %s\n", profile.Item)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
}

