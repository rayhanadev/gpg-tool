# GPG Tool

> Note: This tool is waiting in [this gopenpgp PR](https://github.com/ProtonMail/gopenpgp/pull/258) to go through for OpenPGP v3 support.

Minor exploration of CLI tools written in Golang, a simple tool to make encrypted PGP messages and decrypt them with a nicer UI.

## Install

```sh
git clone https://github.com/rayhanadev/gpg-tool.git
cd gpg-tool
go build -o gpgtool
```

## Usage

```sh
A simple GPG CLI tool to encrypt and decrypt messages using GPG keys.

Usage:
  gpgtool [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  decrypt     Decrypt a message
  encrypt     Encrypt a message
  help        Help about any command

Flags:
  -h, --help   help for gpgtool

Use "gpgtool [command] --help" for more information about a command.
```
