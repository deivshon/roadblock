# Shells configuration

To configure roadblock to block forbidden commands in your shell of choice, refer to the paragraphs below

Before pasting any of these snippets, make sure to thoroughly understand what they do

## Bash

Add the following to your `.bashrc`

```bash
shopt -s extdebug

command_check() {
    roadblock_output=$(roadblock -t "$BASH_COMMAND" 2>&1)
    roadblock_status=$?

    if [[ $roadblock_status -ne 0 ]]; then
        printf 1>&2 "%s\n" "$roadblock_output"
        false
    fi
}

trap 'command_check' DEBUG
```

Make sure you understand the implications of the `extdebug` option before proceeding

## Fish

Add the following to your `config.fish`

```sh
function command_check
    set -l roadblock_output (roadblock -t $(commandline) 2>&1)
    set -l roadblock_status $status

    if test $roadblock_status -ne 0
        printf 1>&2 "\n%s\n" "$roadblock_output"
        commandline -f repaint
        return
    end

    commandline -f execute
end

bind \r command_check
```

Note

- This will prevent you from using fish's multiline command editing
