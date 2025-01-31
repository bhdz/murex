package docs

func init() {

	Definition["bexists"] = "# _murex_ Shell Docs\n\n## Command Reference: `bexists`\n\n> Check which builtins exist\n\n## Description\n\n`bexists` takes an array of parameters and returns which commands have been\ncompiled into _murex_. The 'b' in `bexists` stands for 'builtins'\n\n## Usage\n\n    bexists command... -> <stdout>\n\n## Examples\n\n    » bexists: qr gzip runtime config\n    {\n        \"Installed\": [\n            \"runtime\",\n            \"config\"\n        ],\n        \"Missing\": [\n            \"qr\",\n            \"gzip\"\n        ]\n    }\n\n## Detail\n\nThis builtin dates back to the start of _murex_ when all of the builtins were\nconsidered optional. This was intended to be a way for scripts to determine\nwhich builtins were compiled. Since then `runtime` has absorbed and centralized\na number of similar commands which have since been deprecated. The same fate\nmight also happen to `bexists` however it is in use by a few modules and for\nthat reason alone it has been spared from the axe.\n\n## See Also\n\n* [Modules and Packages](../user-guide/modules.md):\n  An introduction to _murex_ modules and packages\n* [`fexec` ](../commands/fexec.md):\n  Execute a command or function, bypassing the usual order of precedence.\n* [`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of _murex_"

}
