# _murex_ Shell Docs

## Command Reference: `summary` 

> Defines a summary help text for a command

## Description

`summary` define help text for a command. This is effectively like a tooltip
message that appears, by default, in blue in the interactive shell.

Normally this text is populated from the `man` pages or `murex-docs`, however
if neither exist or if you wish to override their text, then you can use
`summary` to define that text.

## Usage

Define a commands summary

    summary command description
    
Undefine a summary

    !summary command

## Examples

Define a commands summary

    » summary: foobar "Hello, world!"
    » runtime: --summaries -> [ foobar ]
    Hello, world! 
    
Undefine a summary

    » !summary: foobar

## Synonyms

* `summary`
* `!summary`


## See Also

* [`bexists`](../commands/bexists.md):
  Check which builtins exist
* [`builtins`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [`config`](../commands/config.md):
  Query or define _murex_ runtime settings
* [`exec`](../commands/exec.md):
  Runs an executable
* [`fid-list`](../commands/fid-list.md):
  Lists all running functions within the current _murex_ session
* [`murex-docs`](../commands/murex-docs.md):
  Displays the man pages for _murex_ builtins
* [`murex-update-exe-list`](../commands/murex-update-exe-list.md):
  Forces _murex_ to rescan $PATH looking for exectables
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_