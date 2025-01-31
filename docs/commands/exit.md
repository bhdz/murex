# _murex_ Shell Docs

## Command Reference: `exit`

> Exit murex

## Description

Exit's _murex_ with either a exit number of 0 (by default if no parameters
supplied) or a custom value specified by the first parameter.

`exit` is not scope aware; if it is included in a function then the whole
shell will still exist and not just that function.

## Usage

    exit
    exit number

## Examples

    » exit
    
    » exit 42

## See Also

* [`break`](../commands/break.md):
  terminate execution of a block within your processes scope
* [`die`](../commands/die.md):
  Terminate murex with an exit number of 1
* [`null`](../commands/devnull.md):
  null function. Similar to /dev/null