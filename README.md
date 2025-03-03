# GUM - Git User Manager

GUM (Git User Manager) is a CLI tool that allows users to manage multiple Git profiles effortlessly. It enables users to switch between profiles, set project-specific Git configurations, and maintain a `.gumrc` file for easy profile activation.

## Features

- **Manage Multiple Git Profiles**: Add, remove, list, and switch between Git profiles.
- **Persistent Config Storage**: Profiles are stored in `~/.config/gum/profiles.json`.
- **Project-Specific Profiles**: Use `.gumrc` to associate a profile with a project directory.
- **Manual Profile Activation**: Easily apply the `.gumrc` profile using the `gum auto` command.

## Installation

You can install GUM by cloning the repository and building it with Go:

```sh
# Clone the repository
git clone https://github.com/rockingrohit9639/gum.git
cd gum

# Build the binary
go build -o gum

# Move the binary to a directory in your PATH
mv gum /usr/local/bin/
```

## Usage

### Add a New Profile

```sh
gum add
```

This will prompt you to enter a profile name and email.

### List All Profiles

```sh
gum list
```

Displays all stored profiles.

### Use a Profile

```sh
gum use
```

Applies the selected Git profile globally.

### Remove a Profile

```sh
gum remove <profile>
```

Deletes the specified profile.

### Show Current Profile

```sh
gum current
```

Displays the currently active Git profile.

### Set Up a Project-Specific Profile

```sh
gum rc <profile>
```

Creates a `.gumrc` file in the current directory with the specified profile.

### Activate a `.gumrc` Profile

```sh
gum auto
```

Reads the `.gumrc` file in the current directory and switches to the specified profile.

## Configuration

GUM stores profiles in `~/.config/gum/profiles.json`. Users can manually edit this file if needed.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License.
