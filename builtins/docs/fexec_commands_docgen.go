package docs

func init() {

	Definition["fexec"] = "# _murex_ Shell Docs\n\n## Command Reference: `fexec` \n\n> Execute a command or function, bypassing the usual order of precedence.\n\n## Description\n\n`fexec` allows you to execute a command or function, bypassing the usual order\nof precedence.\n\n## Usage\n\n    fexec: flag command [ parameters... ] -> <stdout>\n    ``` \n\n## Examples\n\n    fexec: private /source/builtin/autocomplete.alias\n\n## Flags\n\n* `builtin`\n    Execute a _murex_ builtin\n* `function`\n    Execute a _murex_ public function\n* `help`\n    Display help message\n* `private`\n    Execute a _murex_ private function\n\n## Detail\n\n### Order of precedence\n\nThere is an order of precedence for which commands are looked up:\n1. `test` and `pipe` functions because they alter the behavior of the compiler\n2. Aliases - defined via `alias`. All aliases are global\n3. _murex_ functions - defined via `function`. All functions are global\n4. private functions - defined via `private`. Private's cannot be global and\n   are scoped only to the module or source that defined them. For example, You\n   cannot call a private function from the interactive command line\n5. variables (dollar prefixed) - declared via `set` or `let`\n6. auto-globbing prefix: `@g`\n7. murex builtins\n8. external executable files\n\n## See Also\n\n* [commands/`alias`](../commands/alias.md):\n  Create an alias for a command\n* [commands/`autocomplete`](../commands/autocomplete.md):\n  Set definitions for tab-completion in the command line\n* [commands/`bg`](../commands/bg.md):\n  Run processes in the background\n* [commands/`builtins`](../commands/runtime.md):\n  Returns runtime information on the internal state of _murex_\n* [commands/`event`](../commands/event.md):\n  Event driven programming for shell scripts\n* [commands/`exec`](../commands/exec.md):\n  Runs an executable\n* [commands/`fg`](../commands/fg.md):\n  Sends a background process into the foreground\n* [commands/`function`](../commands/function.md):\n  Define a function block\n* [commands/`jobs`](../commands/fid-list.md):\n  Lists all running functions within the current _murex_ session\n* [commands/`open`](../commands/open.md):\n  Open a file with a preferred handler\n* [commands/`private`](../commands/private.md):\n  Define a private function block\n* [commands/`source` ](../commands/source.md):\n  Import _murex_ code from another file of code block"

}
