# swenv

Tired of manually rewriting .env file just to change the enviroment? This little CLI app does the job for you!

Prepare a couple of `.env.enviroment` files like `.env.dev` or `.env.prod` and switch quickly between them with `swenv sw <enviroment>`.

Selected enviroment will be copied into the `.env` file.

## Install

### Homebrew (macOS and Linux)

```bash
brew install BasileosFelices/tap/swenv
```

To upgrade:
```bash
brew upgrade swenv
```

### Go install

```bash
go install github.com/BasileosFelices/swenv@latest
```

### GitHub Releases

Download the appropriate binary from the [Releases page](https://github.com/BasileosFelices/swenv/releases) for your platform.

## Usage

You can allways run `swenv --help` for all available commands. Currently however there are just two commands:

- `swenv list/ls ` - Lists currently available enviroment files
- `swenv switch/sw <enviroment>` - Switches to desired enviroment

For simple use you can use `swenv <enviroment>` without the `switch` command. No argument acts like the `list` command.

The current `.env` file is always backed up to `.env.swenv.envbackup`

## Troubleshooting

### macOS: "swenv cannot be opened because the developer cannot be verified"

This happens because the binary isn't code-signed. To fix:

1. Open **System Settings → Privacy & Security**
2. Scroll down and click **"Open Anyway"** next to the swenv warning

Or run:
```bash
xattr -d com.apple.quarantine $(which swenv)
```

## Recommended `.gitignore`

I recommend adding the following into your project `.gitignore` file. It prevents all enviroment files from being commited while allowing the examples.
```gitignore
# env files
.env*
!.env*.example
```
