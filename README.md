# swenv

Tired of manually rewriting .env file just to change the enviroment? This little CLI app does the job for you!

Prepare a couple of `.env.enviroment` files like `.env.dev` or `.env.prod` and switch quickly between them with `swenv sw <enviroment>`.

Selected enviroment will be copied into the `.env` file.

## Install

### Go install

```
go install github.com/BasileosFelices/swenv@latest
```

### GitHub Releases

Download the appropriate binary from the Releases page for your platform.

## Usage

You can allways run `swenv --help` for all available commands. Currently however there are just two commands:

- `swenv list/ls ` - Lists currently available enviroment files
- `swenv switch/sw <enviroment>` - Switches to desired enviroment

The current `.env` file is always backed up to `.swenv.envbackup`

## Recommended `.gitignore`

I recommend adding the following into your project `.gitignore` file. It prevents all enviroment files from being commited while allowing the examples.

Be careful as the backup may leak credentials if you do not add it into your gitignore as well.

```
# env files
.env*
!.env*.example
.swenv.*
```