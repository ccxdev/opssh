package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ccxdev/opssh/internal/config"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <name>",
	Short: "Add a new SSH key profile",
	Long: `Add a new SSH key profile to your config file.
You'll be prompted to enter the account, vault, and item details.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		profileName := args[0]

		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		if _, err := config.GetProfile(cfg, profileName); err == nil {
			return fmt.Errorf("profile '%s' already exists", profileName)
		}

		var profile *config.Profile

		profile, err = promptManualProfile()

		if err != nil {
			return err
		}

		if err := config.AddProfile(cfg, profileName, *profile); err != nil {
			return fmt.Errorf("failed to add profile: %w", err)
		}

		if err := config.SaveConfig(cfg); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		fmt.Printf("✓ Added profile '%s'\n", profileName)
		fmt.Printf("  Account: %s\n", profile.Account)
		fmt.Printf("  Vault: %s\n", profile.Vault)
		fmt.Printf("  Item: %s\n", profile.Item)

		var switchNow bool
		switchPrompt := &survey.Confirm{
			Message: "Switch to this profile now?",
			Default: false,
		}
		if err := survey.AskOne(switchPrompt, &switchNow); err == nil && switchNow {
			fmt.Println("Use 'opssh switch " + profileName + "' to switch to this profile.")
		}

		return nil
	},
}

func promptManualProfile() (*config.Profile, error) {
	var account, vault, item string

	accountPrompt := &survey.Input{
		Message: "Account:",
	}
	if err := survey.AskOne(accountPrompt, &account); err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	vaultPrompt := &survey.Input{
		Message: "Vault:",
	}
	if err := survey.AskOne(vaultPrompt, &vault); err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	itemPrompt := &survey.Input{
		Message: "Item:",
	}
	if err := survey.AskOne(itemPrompt, &item); err != nil {
		return nil, fmt.Errorf("failed to get item: %w", err)
	}

	return &config.Profile{
		Account: account,
		Vault:   vault,
		Item:    item,
	}, nil
}

func init() {
	rootCmd.AddCommand(addCmd)
}
