# Rule configuration files

Rule configuration files are the way roadblock is configured to detect the commands you want to block

Rule configuration files are YAML files named `config.yml` and are recursively searched in the roadblock root configuration directory (by default in `{USER_CONFIG_DIR}/roadblock`, but this can be overwritten with the `-c` CLI argument)

Each configuration file has two sections: `conditions` and `rules`. The `conditions` section allows you to define conditions that either match or don't match a command, and the `rules` section allows you to combine these conditions into rules with the common boolean operators (and, or, not)

## Conditions

The `conditions` section is a list of conditions, each condition must have an `id` and is in turn divided into two sections: `select` and `evaluate`. The `select` field decides _what part_ of the command should be evaluated, and the `evaluate` field decides _how_ it will be evaluated

Below are tables listing all the available options for each of these sections

| Selector    | Value            | Behavior                                                                                                                 |
| ----------- | ---------------- | ------------------------------------------------------------------------------------------------------------------------ |
| `command`   | true             | Selects the whole command                                                                                                |
| `wordIndex` | positive integer | Selects a single _word_ of the command at the specified index (words are zero-indexed, i.e. the first one is at index 0) |
| `anyWord`   | true             | Selects all command words, if any evaluates to true, the condition is satisfied                                          |
| `everyWord` | true             | Selects all command words, if all evaluate to true, the condition is satisfied                                           |

| Evaluator  | Value  | Behavior                                                 |
| ---------- | ------ | -------------------------------------------------------- |
| `equals`   | string | Checks if the selected part is equal to the given string |
| `contains` | string | Checks if the selected part contains the given string    |
| `regex`    | regex  | Checks if the selected part matches the given regex      |

#

Note:

- In a condition, **exactly one selector and evaluator must be active**
- The scope of a condition is the file it's contained in

#

Correct example

```yml
conditions:
  - id: example-condition
    select:
      wordIndex: 0
    evaluate:
      regex: git$
rules: ...
```

Wrong example (multiple selectors active)

```yml
conditions:
  - id: example-condition
    select:
      command: true
      wordIndex: 0
    evaluate:
      regex: git$
rules: ...
```

## Rules

The rules section is a list of rules. Each rule must have a `name` and a `rule` field, where the actual conditions combination will be laid out

The available operators are

| Operator      | Value             | Behavior                                                                              |
| ------------- | ----------------- | ------------------------------------------------------------------------------------- |
| `allOf`       | List of operators | Evaluates to true if all the children operators evaluate to true                      |
| `oneOf`       | List of operators | Evaluates to true if at least one of the children operators evaluates to true         |
| `not`         | Single operator   | Evaluates to true if the children operator evaluates to false                         |
| `conditionId` | string            | Evaluates to true if the referenced condition is satisfied given the analyzed command |

#

Note:

- **Exactly one operator must be active per level**

#

Assume these conditions are contained in the example file for the following examples

```yml
conditions:
  - id: git-command
    select:
      wordIndex: 0
    evaluate:
      regex: git$
  - id: git-push-subcommand
    select:
      wordIndex: 1
    evaluate:
      equals: push
```

Correct example

```yml
rules:
  - name: prevent-git-push
    rule:
      allOf:
        - conditionId: git-command
        - conditionId: git-push-subcommand
```

Wrong example (first level has two active operators, `not` and `allOf`)

```yml
rules:
  - name: prevent-git-push
    rule:
      not:
        conditionId: git-command
      allOf:
        - conditionId: git-command
        - conditionId: git-push-subcommand
```

Wrong example (the referenced `conditionId` "`some-random-id`" does not exist in the conditions)

```yml
rules:
  - name: prevent-git-push
    rule:
      allOf:
        - conditionId: some-random-id
        - conditionId: git-push-subcommand
```
