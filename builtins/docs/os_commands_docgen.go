package docs

func init() {

	Definition["os"] = "# _murex_ Shell Docs\n\n## Command Reference: `os`\n\n> Output the auto-detected OS name\n\n## Description\n\nOutput the auto-detected OS name.\n\n## Usage\n\n    os -> <stdout>\n    \n    os string -> <stdout>\n    ``` \n\n## Examples\n\n    » os\n    linux\n    \nOr if you want to check if the host is one of a number of platforms:\n\n    # When run on Linux or FreeBSD\n    » os linux freebsd\n    true\n    \n    # When run on another platform, eg Windows or Darwin (OSX)\n    # (exit number would also be `1`)\n    » os linux freebsd\n    false\n    \n`posix` is also supported:\n\n    # When run on Linux, FreeBSD or Darwin (for example)\n    » os posix\n    true\n    \n    # When run on Windows or Plan 9\n    # (exit number would also be `1`)\n    » os posix\n    false\n    \nPlease note that although Plan 9 shares similarities with POSIX, it is not\nPOSIX-compliant. For that reason, `os` returns false with the `posix`\nparameter when run on Plan 9. If you want to include Plan 9 in the check\nthen please write it as `os posix plan9`.\n\n## See Also\n\n* [`cpuarch`](../commands/cpuarch.md):\n  Output the hosts CPU architecture\n* [`cpucount`](../commands/cpucount.md):\n  Output the number of CPU cores available on your host"

}
