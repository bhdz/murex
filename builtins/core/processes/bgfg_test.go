package processes_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/test/count"
)

func TestBg(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `bg { out "bg" }; sleep 1; out "fg"`,
			Stdout: "bg\nfg\n",
		},
		{
			Block:  `bg { sleep 1; out "bg" }; out "fg"`,
			Stdout: "fg\nbg\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestBgFg(t *testing.T) {
	count.Tests(t, 2)
	sleep := 3
	block := fmt.Sprintf(`bg { sleep %d }`, sleep)

	lang.InitEnv()

	go func() {
		fork := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR)
		i, err := fork.Execute([]rune(block))
		if i != 0 || err != nil {
			t.Error("Error executing block:")
			t.Logf("  Block:    %s", block)
			t.Logf("  Exit num: %d", i)
			t.Logf("  Error:    %v", err.Error())
		}
	}()

	time.Sleep(1 * time.Second)

	var p *lang.Process
	fids := lang.GlobalFIDs.ListAll()

	for i := range fids {

		if fids[i].Name.String() == "exec" {
			name, err := fids[i].Parameters.String(0)
			if err != nil && name != "sleep" {
				continue
			}
			duration, err := fids[i].Parameters.Int(1)
			if err != nil && duration != sleep {
				continue
			}

			p = fids[i]
			goto next
		}
	}

	t.Fatalf("Cannot find FID attached to `sleep %d`\n", sleep)

next:

	if !p.Background.Get() {
		t.Fatalf("`sleep 5` isn't set to background: p.Background == %v", p.Background.Get())
	}

	if runtime.GOOS == "windows" || runtime.GOOS == "plan9" || runtime.GOOS == "js" {
		// skip `fg` tests on systems that don't support foregrounding
		return
	}

	count.Tests(t, 2)
	block = fmt.Sprintf(`fg %d`, p.Id)

	fork := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR)
	i, err := fork.Execute([]rune(block))
	if i != 0 || err != nil {
		t.Error("Error executing block:")
		t.Logf("  Block:    %s", block)
		t.Logf("  Exit num: %d", i)
		t.Logf("  Error:    %v", err.Error())
	}

	if p.Background.Get() {
		t.Fatalf("`sleep 5` hasn't been set to foreground: p.Background == %v", p.Background.Get())
	}
}
