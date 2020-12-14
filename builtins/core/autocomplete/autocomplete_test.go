package cmdautocomplete_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
)

func TestAutocomplete(t *testing.T) {
	count.Tests(t, 1)

	block := []rune(
		`autocomplete set foobar { [{
			"AnyValue": true,
			"AllowMultiple": true,
			"ExecCmdline": true,
			"Flags": [ "hello", "world" ]
		}] }
		
		autocomplete get foobar -> pretty`)

	expected := `[
    {
        "IncFiles": false,
        "IncDirs": false,
        "IncExePath": false,
        "Flags": [
            "hello",
            "world"
        ],
        "FlagsDesc": null,
        "Dynamic": "",
        "DynamicDesc": "",
        "ListView": false,
        "FlagValues": null,
        "Optional": false,
        "AllowMultiple": true,
        "Alias": "",
        "NestedCommand": false,
        "AnyValue": true,
        "AutoBranch": false,
        "ExecCmdline": true
    }
]
`

	lang.InitEnv()

	fork := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_CREATE_STDERR)
	exitNum, err := fork.Execute(block)
	if exitNum != 0 {
		t.Errorf("None zero exit number: %d", exitNum)
	}
	if err != nil {
		t.Errorf("Error in fork.Execute(): %s", err.Error())
	}

	stdout, err := fork.Stdout.ReadAll()
	if err != nil {
		t.Errorf("Error in fork.Stdout.ReadAll(): %s", err.Error())
	}

	stderr, err := fork.Stderr.ReadAll()
	if err != nil {
		t.Errorf("Error in fork.Stderr.ReadAll(): %s", err.Error())
	}
	if len(stderr) > 0 {
		t.Errorf("Stderr should be empty: %s", string(stderr))
	}

	if string(stdout) != expected {
		t.Error("Stdout doesn't match expected:")
		t.Logf("  Expected: %s", expected)
		t.Logf("  Actual: %s", string(stdout))
	}
}
