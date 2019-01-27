package test

import (
	"strings"
	"testing"

	//_ "github.com/lmorg/murex/builtins/core/typemgmt" // import murex builtins
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
)

// MurexTest is a basic framework to test murex code.
// Please note this shouldn't be confused with the murex scripting language's inbuilt testing framework!
type MurexTest struct {
	Block   string
	Stdout  string
	Stderr  string
	ExitNum int
}

// RunMurexTests runs through all the test cases for MurexTest
func RunMurexTests(tests []MurexTest, t *testing.T) {
	defaults.Defaults(lang.InitConf, false)
	lang.InitEnv()

	for i := range tests {
		stdout := streams.NewStdin()
		stderr := streams.NewStdin()
		hasError := false

		exitNum, err := lang.RunBlockShellConfigSpace([]rune(tests[i].Block), nil, stdout, stderr)
		if err != nil {
			t.Errorf("Cannot execute script on test %d", i)
			t.Log(err)
			continue
		}

		bErr, err := stderr.ReadAll()
		if err != nil {
			t.Errorf("Cannot ReadAll() from Stderr on test %d", i)
			t.Log(err)
			continue
		}

		if string(bErr) != tests[i].Stderr {
			hasError = true
		}

		bOut, err := stdout.ReadAll()
		if err != nil {
			t.Errorf("Cannot ReadAll() from Stdout on test %d", i)
			t.Log(err)
			continue
		}

		if string(bOut) != tests[i].Stdout {
			hasError = true
		}

		if exitNum != tests[i].ExitNum {
			hasError = true
		}

		if hasError {
			t.Errorf("Code block doesn't return expected values in test %d", i)
			t.Log("  Code block:      ", tests[i].Block)

			t.Log("  Expected Stdout: ", strings.Replace(tests[i].Stdout, "\n", `\n`, -1))
			t.Log("  Actual Stdout:   ", strings.Replace(string(bOut), "\n", `\n`, -1))
			t.Log("  eo bytes:        ", []byte(tests[i].Stdout))
			t.Log("  ao bytes:        ", bOut)

			t.Log("  Expected Stderr: ", strings.Replace(tests[i].Stderr, "\n", `\n`, -1))
			t.Log("  Actual Stderr:   ", strings.Replace(string(bErr), "\n", `\n`, -1))
			t.Log("  eo bytes:        ", []byte(tests[i].Stderr))
			t.Log("  ao bytes:        ", bErr)

			t.Log("  Expected exitnum:", tests[i].ExitNum)
			t.Log("  Actual exitnum:  ", exitNum)
		}
	}
}
