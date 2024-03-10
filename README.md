# roadblock

roadblock is designed to analyze and either approve or reject interactive shell commands. In case of approval, roadblock exits with status code 0, otherwise, with status code 1, logging which rule forbids the analyzed command

By configuring your shell properly, you can let roadblock intercept the commands you enter and, if they are forbidden, prevent them from running

## Configuration

roadblock accepts the following CLI arguments

| Argument | Description                                                                                                                    | Optional | Default                                     |
| -------- | ------------------------------------------------------------------------------------------------------------------------------ | -------- | ------------------------------------------- |
| `-t`     | The command to be analyzed                                                                                                     | Yes      | Empty string                                |
| `-c`     | Root directory for the rule configuration files                                                                                | Yes      | `{USER_CONFIG_DIR}/roadblock`               |
| `-g`     | Path to the global configuration file                                                                                          | Yes      | `{USER_CONFIG_DIR}/roadblock/roadblock.yml` |
| `-s`     | If passed, when a command is forbidden, the path of the source configuration file containing the rule that forbids it is shown | Yes      | false                                       |

A self-documenting global configuration file can be found in `docs/configuration-examples/roadblock.yml`

## Install

To compile roadblock **the Go toolchain is required**

```
$ git clone https://github.com/deivshon/roadblock
$ cd roadblock
$ make release
# make install
```

## Fail-safes

You can set the environment variable `ROADBLOCK_SKIP` to any value to make roadblock completely skip all checks

This can also be needed in the unfortunate case where roadblock keeps crashing and prevents you from entering any command

## More documentation

- `docs/CONFIGS.md`: information about rule configuration files
- `docs/SHELLS.md`: information about how to configure your shell to let roadblock intercept commands
