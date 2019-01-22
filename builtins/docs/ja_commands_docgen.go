package docs

func init() {

	Definition["ja"] = "# _murex_ Language Guide\n\n## Command Reference: `ja`\n\n> A sophisticated yet simply way to build a JSON array\n\n### Description\n\n_murex_ has a pretty sophisticated builtin for generating JSON arrays.\nIt works a little bit like Bash's `{1..9}` syntax but includes a few\nadditional nifty features.\n\n### Usage\n\n    ja: [start..end] -> <stdout>\n    ja: [start..end.base] -> <stdout>\n    ja: [start..end,start..end] -> <stdout>\n    ja: [start..end][start..end] -> <stdout>\n\n### Examples\n\n    a: [1..5] \n    [\n        \"1\",\n        \"2\",\n        \"3\",\n        \"4\",\n        \"5\"\n    ]\n    \n    » ja: [Monday..Sunday]\n    [\n        \"Monday\",\n        \"Tuesday\",\n        \"Wednesday\",\n        \"Thursday\",\n        \"Friday\",\n        \"Saturday\",\n        \"Sunday\"\n    ]\n    \nPlease note that as per the first example, all arrays generated by `ja` are\narrays of strings - even if you're command is ranging over integers.\n\n### Detail\n\nPlease read the documentation on `a` for a more detailed breakdown on of\n`ja`'s supported features.\n\n### See Also\n\n* [`@[` (range) ](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [`a`](../commands/a.md):\n  A sophisticated yet simply way to build an array or list\n* [`len` ](../commands/len.md):\n  Outputs the length of an array\n* [mtac](../commands/mtac.md):\n  "

}
