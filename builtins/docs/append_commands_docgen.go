package docs

func init() {

	Definition["append"] = "# _murex_ Shell Docs\n\n## Command Reference: `append`\n\n> Add data to the end of an array\n\n## Description\n\n`append` data to the end of an array.\n\n## Usage\n\n    <stdin> -> append: value -> <stdout>\n\n## Examples\n\n    » a: [Monday..Sunday] -> append: Funday\n    Monday\n    Tuesday\n    Wednesday\n    Thursday\n    Friday\n    Saturday\n    Sunday\n    Funday\n\n## Detail\n\n`prepend` and `append` are data type aware:\n\n    » tout json [1,2,3] -> append 4 5 6 bob\n    Error in `append` (1,22): cannot convert 'bob' to a floating point number: strconv.ParseFloat: parsing \"bob\": invalid syntax\n\n## Synonyms\n\n* `append`\n* `list.append`\n\n\n## See Also\n\n* [commands/`@[` (range) ](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [commands/`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [commands/`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [commands/`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [commands/`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [commands/`addheading` ](../commands/addheading.md):\n  Adds headings to a table\n* [commands/`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* [commands/`count`](../commands/count.md):\n  Count items in a map, list or array\n* [commands/`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [commands/`match`](../commands/match.md):\n  Match an exact value in an array\n* [commands/`msort` ](../commands/msort.md):\n  Sorts an array - data type agnostic\n* [commands/`mtac`](../commands/mtac.md):\n  Reverse the order of an array\n* [commands/`prepend` ](../commands/prepend.md):\n  Add data to the start of an array\n* [commands/`regexp`](../commands/regexp.md):\n  Regexp tools for arrays / lists of strings"

}
