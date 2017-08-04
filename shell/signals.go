package shell

import (
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils"
	"os"
	"os/signal"
	"syscall"
)

func SigHandler() {
	defer func() {
		if r := recover(); r != nil {
			os.Stderr.WriteString(fmt.Sprintln("Exception caught: ", r))
			SigHandler()
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGQUIT)
	go func() {
		for {
			sig := <-c
			switch sig.String() {
			case syscall.SIGTERM.String():
				Instance.Terminal.ExitRawMode()
				os.Stderr.WriteString("Shell received SIGTERM!" + utils.NewLineString)
				os.Exit(1)
			case os.Interrupt.String():
				if Instance == nil {
					go proc.KillForeground()
					//os.Stderr.WriteString("^c")
				} else {
					p := proc.ForegroundProc
					for p.Id != 0 {
						parent := p.Parent
						if p.Kill != nil {
							p.Kill()
						}
						p = parent
					}
					//os.Stderr.WriteString("^c")
				}
			case syscall.SIGQUIT.String():
				os.Stderr.WriteString("Shell received SIGQUIT!" + utils.NewLineString)
				os.Exit(2)
			default:
				os.Stderr.WriteString("Unhandled signal: " + sig.String())
			}
		}
	}()
}
