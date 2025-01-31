package docs

func init() {

	Definition["gz"] = "# _murex_ Shell Docs\n\n## Optional Command Reference: `gz` \n\n> Compress or decompress a gzip file\n\n## Description\n\nAn optional builtin for compressing or decompressing a gzip stream from STDIN.\n\n## Usage\n\n    <stdin> -> gz -> <stdout>\n    \n    <stdin> -> !gz -> <stdout>\n\n## Synonyms\n\n* `gz`\n* `!gz`\n\n\n## See Also\n\n* [`!bz2` ](../optional/bz2.md):\n  Decompress a bz2 file\n* [`base64` ](../optional/base64.md):\n  Encode or decode a base64 string\n* [`escape`](../commands/escape.md):\n  Escape or unescape input \n* [`esccli`](../commands/esccli.md):\n  Escapes an array so output is valid shell code\n* [`eschtml`](../commands/eschtml.md):\n  Encode or decodes text for HTML\n* [`escurl`](../commands/escurl.md):\n  Encode or decodes text for the URL"

}
