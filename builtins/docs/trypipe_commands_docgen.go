package docs

func init() {

	Definition["trypipe"] = "# _murex_ Shell Docs\n\n## Command Reference: `trypipe`\n\n> Checks state of each function in a pipeline and exits block on error\n\n## Description\n\n`trypipe` checks the state of each function and exits the block if any of them\nfail. Where `trypipe` differs from regular `try` blocks is `trypipe` will check\nevery process along the pipeline as well as the terminating function (which\n`try` only validates against). The downside to this is that piped functions can\nno longer run in parallel.\n\n## Usage\n\n    trypipe { code-block } -> <stdout>\n    \n    <stdin> -> trypipe { -> code-block } -> <stdout>\n\n## Examples\n\n    trypipe {\n        out: \"Hello, World!\" -> grep: \"non-existent string\" -> cat\n        out: \"This command will be ignored\"\n    }\n    \nFormated pager (`less`) where the pager isn't called if the formatter (`pretty`) fails (eg input isn't valid JSON):\n\n    func pless {\n        -> trypipe { -> pretty -> less }\n    }\n\n## Detail\n\nA failure is determined by:\n\n* Any process that returns a non-zero exit number\n* Any process that returns more output via STDERR than it does via STDOUT\n\nYou can see which run mode your functions are executing under via the `fid-list`\ncommand.\n\n## See Also\n\n* [user-guide/Schedulers](../user-guide/schedulers.md):\n  Overview of the different schedulers (or 'run modes') in _murex_\n* [commands/`catch`](../commands/catch.md):\n  Handles the exception code raised by `try` or `trypipe` \n* [commands/`fid-list`](../commands/fid-list.md):\n  Lists all running functions within the current _murex_ session\n* [commands/`if`](../commands/if.md):\n  Conditional statement to execute different blocks of code depending on the result of the condition\n* [commands/`runmode`](../commands/runmode.md):\n  Alter the scheduler's behaviour at higher scoping level\n* [commands/`switch`](../commands/switch.md):\n  Blocks of cascading conditionals\n* [commands/`try`](../commands/try.md):\n  Handles errors inside a block of code"

}
