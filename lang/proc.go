package lang

import (
	"context"
	"errors"
	"os/exec"
	"sync"
	"time"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/runmode"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/proc/stdio"
)

// Process - Each process running inside the murex shell will be one of these objects.
// It is equivalent to the /proc directory on Linux, albeit queried through murex as JSON.
// External processes will also appear in the host OS's process list.
type Process struct {
	Context            context.Context
	Stdin              stdio.Io
	Stdout             stdio.Io
	Stderr             stdio.Io
	Parameters         parameters.Parameters
	ExitNum            int
	Name               string
	Id                 int
	Exec               shellExec
	PromptId           int
	Path               string
	IsMethod           bool
	Module             string
	Scope              *Process  `json:"-"`
	Parent             *Process  `json:"-"`
	Previous           *Process  `json:"-"`
	Next               *Process  `json:"-"`
	WaitForTermination chan bool `json:"-"`
	Done               func()    `json:"-"`
	Kill               func()    `json:"-"`
	IsNot              bool
	NamedPipeOut       string
	NamedPipeErr       string
	NamedPipeTest      string
	hasTerminatedM     sync.Mutex
	hasTerminatedV     bool
	State              state.FunctionState
	IsBackground       bool
	LineNumber         int
	ColNumber          int
	RunMode            runmode.RunMode
	Config             *config.Config
	Tests              *Tests
	Variables          *Variables
	FidTree            []int
	CreationTime       time.Time
	StartTime          time.Time
}

type shellExec struct {
	Pid int
	Cmd *exec.Cmd
}

// HasTerminated checks if process has terminated.
// This is a function because terminated state can be subject to race conditions
// so we need a mutex to make the state thread safe.
func (p *Process) HasTerminated() (state bool) {
	p.hasTerminatedM.Lock()
	state = p.hasTerminatedV
	p.hasTerminatedM.Unlock()
	return
}

// HasCancelled is a wrapper function around context because it's a pretty ugly API
func (p *Process) HasCancelled() (state bool) {
	if p.Context == nil {
		return false
	}

	select {
	case <-p.Context.Done():
		return true
	default:
		return false
	}
}

// SetTerminatedState sets the process terminated state.
// This is a function because terminated state can be subject to race conditions
// so we need a mutex to make the state thread safe.
func (p *Process) SetTerminatedState(state bool) {
	p.hasTerminatedM.Lock()
	p.hasTerminatedV = state
	p.hasTerminatedM.Unlock()
	return
}

// ErrIfNotAMethod returns a standard error message for builtins not run as methods
func (p *Process) ErrIfNotAMethod() (err error) {
	if !p.IsMethod {
		err = errors.New("`" + p.Name + "` expects to be pipelined")
	}
	return
}

// DeregisterProcess deregisters a murex process
func DeregisterProcess(p *Process) {
	p.State = state.Terminating

	p.Stdout.Close()
	p.Stderr.Close()

	p.SetTerminatedState(true)
	if !p.IsBackground {
		ForegroundProc = p.Next
	}

	go deregister(p)
}

// deregister FID and mark variables for garbage collection.
func deregister(p *Process) {
	p.State = state.AwaitingGC
	CloseScopedVariables(p)
	GlobalFIDs.Deregister(p.Id)
}
