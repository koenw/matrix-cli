# matrix-cli

Command line client for the [matrix](https://matrix.org) decentralized
communication network.

This is still under heavy development, so don't expect a lot of functionality,
command line arguments may change, etc.

## Usage

```
usage: matrix-cli --username=USERNAME --password=PASSWORD [<flags>] <command> [<args> ...]

Command line client for the [matrix] decentralized communication network.

Flags:
  --help               Show context-sensitive help (also try --help-long and
                       --help-man).
  --server-url="https://matrix.org"
                       URL to the matrix server
  --username=USERNAME  matrix username
  --password=PASSWORD  matrix password

Commands:
  help [<command>...]
    Show help.

  send <room> <message>...
    Send a message to a matrix room
```

## License

Available under either the MIT or Apache license.
