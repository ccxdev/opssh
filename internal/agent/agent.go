package agent

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ccxdev/opssh/internal/config"
)

const AgentConfigPath = ".config/1Password/ssh/agent.toml"

func GetAgentConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return filepath.Join(homeDir, AgentConfigPath), nil
}

func ReadAgentConfig() ([]string, error) {
	configPath, err := GetAgentConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read agent.toml: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	return lines, nil
}

func WriteAgentConfig(profile *config.Profile) error {
	lines, err := ReadAgentConfig()
	if err != nil {
		return err
	}

	firstSectionIdx := -1
	for i, line := range lines {
		if strings.Contains(line, "[[ssh-keys]]") {
			firstSectionIdx = i
			break
		}
	}

	var newLines []string

	if firstSectionIdx > 0 {
		newLines = append(newLines, lines[:firstSectionIdx]...)
	}

	newLines = append(newLines, "[[ssh-keys]]")
	newLines = append(newLines, fmt.Sprintf(`account = "%s"`, profile.Account))
	newLines = append(newLines, fmt.Sprintf(`vault = "%s"`, profile.Vault))
	newLines = append(newLines, fmt.Sprintf(`item = "%s"`, profile.Item))

	configPath, err := GetAgentConfigPath()
	if err != nil {
		return err
	}

	content := strings.Join(newLines, "\n")
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write agent.toml: %w", err)
	}

	return nil
}

func GetCurrentActiveProfile() (*config.Profile, error) {
	lines, err := ReadAgentConfig()
	if err != nil {
		return nil, err
	}

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "[[ssh-keys]]" {
			profile := &config.Profile{}

			for j := i + 1; j < len(lines) && j < i+10; j++ {
				nextLine := strings.TrimSpace(lines[j])

				if strings.Contains(nextLine, "[[ssh-keys]]") {
					break
				}

				if strings.HasPrefix(nextLine, "account =") {
					re := regexp.MustCompile(`account\s*=\s*"([^"]+)"`)
					matches := re.FindStringSubmatch(nextLine)
					if len(matches) > 1 {
						profile.Account = matches[1]
					}
				}

				if strings.HasPrefix(nextLine, "vault =") {
					re := regexp.MustCompile(`vault\s*=\s*"([^"]+)"`)
					matches := re.FindStringSubmatch(nextLine)
					if len(matches) > 1 {
						profile.Vault = matches[1]
					}
				}

				if strings.HasPrefix(nextLine, "item =") {
					re := regexp.MustCompile(`item\s*=\s*"([^"]+)"`)
					matches := re.FindStringSubmatch(nextLine)
					if len(matches) > 1 {
						profile.Item = matches[1]
					}
				}
			}

			if profile.Account != "" && profile.Vault != "" && profile.Item != "" {
				return profile, nil
			}
		}
	}

	return nil, fmt.Errorf("no active profile found")
}
