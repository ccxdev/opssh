# opssh

**SSH Key Profile Switcher** - A CLI tool for managing SSH key switching for 1Password's SSH agent.

## Description

`opssh` simplifies switching between different SSH key profiles (e.g., personal vs work) by managing the 1Password SSH agent configuration file (`agent.toml`). It allows you to:

- Store multiple SSH key profiles in a simple YAML configuration
- Quickly switch between profiles without manually editing configuration files
- View and manage your SSH key profiles

## Installation

### Prerequisites

- Go 1.25.5 or later

### Install from Source

```bash
go install github.com/ccxdev/opssh
```

### Build from Source

```bash
git clone https://github.com/ccxdev/opssh.git
cd opssh
go build -o opssh

./opssh <command>
```

## Usage

### Add a Profile

Add a new SSH key profile:

```bash
opssh add <profile-name>
```

You'll be prompted to enter:
- Account name
- Vault name
- Item name

### List Profiles

View all configured profiles:

```bash
opssh list
```

The currently active profile will be marked with an asterisk (`*`).

### Switch Profile

Switch to a different SSH key profile:

```bash
opssh switch <profile-name>
```

This updates the `~/.config/1Password/ssh/agent.toml` file with the selected profile.

### Show Current Profile

Display the currently active SSH key profile:

```bash
opssh current
```

### Remove a Profile

Remove a profile from your configuration:

```bash
opssh remove <profile-name>
```

**Note:** If the profile is currently active, you'll receive a warning.

## Configuration

Profiles are stored in `~/.opssh.yaml` with the following structure:

```yaml
profiles:
  personal:
    account: "myaccount"
    vault: "Personal"
    item: "My SSH Key"
  work:
    account: "myaccount"
    vault: "Work"
    item: "Work SSH Key"
```

## Global Flags

- `--verbose, -v`: Enable verbose output
- `--config`: Specify a custom config file path (default: `~/.opssh.yaml`)

## Examples

### Quick Setup

```bash
# Add a profile
opssh add work

# Switch to work profile
opssh switch work

# Verify current profile
opssh current
```

### Manual Profile Management

```bash
# Add a profile manually
opssh add personal

# List all profiles
opssh list

# Switch between profiles
opssh switch personal
opssh switch work

# Remove a profile
opssh remove old-profile
```

## How It Works

`opssh` manages SSH key profiles by:

1. **Storing profiles** in `~/.opssh.yaml` - a simple YAML configuration file
2. **Updating agent.toml** at `~/.config/1Password/ssh/agent.toml` when switching profiles
3. **Reading agent.toml** to determine the currently active profile

The tool works with a simple YAML configuration file to manage your SSH key profiles.

## Requirements

- **1Password SSH Agent**: This tool is designed to work with 1Password's SSH agent feature

### Agent.toml Not Found

If `agent.toml` doesn't exist, 1Password SSH agent may not be set up. Ensure:
- 1Password SSH agent is enabled
- The directory `~/.config/1Password/ssh/` exists

### Profile Not Found

If you get "profile not found" errors:
- Use `opssh list` to see available profiles
- Check your `~/.opssh.yaml` file for typos

## License

[MIT](./LICENSE)

## Contributing

Please feel free to open an issue or submit a pull request.

