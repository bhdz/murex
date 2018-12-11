package streams

import (
	"context"
	"sync"

	"github.com/lmorg/murex/lang/proc/stdio"
)

func init() {
	stdio.RegesterPipe("std", newStream)
}

// Stdin is the default stdio.Io interface.
// Despite it's name, this interface can and is used for Stdout and Stderr streams too.
type Stdin struct {
	mutex      sync.Mutex
	ctx        context.Context
	forceClose func()
	buffer     []byte
	bRead      uint64
	bWritten   uint64
	dependants int
	dataType   string
	dtLock     sync.Mutex
	max        int
}

// DefaultMaxBufferSize is the maximum size of buffer for stdin
var DefaultMaxBufferSize = 1024 * 1024 * 10 // 10 meg

func newStream(_ string) (io stdio.Io, err error) {
	io = NewStdin()
	return
}

// NewStdin creates a new stream.Io interface for piping data between processes.
// Despite it's name, this interface can and is used for Stdout and Stderr streams too.
func NewStdin() (stdin *Stdin) {
	stdin = new(Stdin)
	stdin.max = DefaultMaxBufferSize
	stdin.ctx, stdin.forceClose = context.WithCancel(context.Background())
	return
}

// NewStdinWithContext creates a new stream.Io interface for piping data between processes.
// Despite it's name, this interface can and is used for Stdout and Stderr streams too.
// This function is also useful as a context aware version of ioutil.ReadAll
func NewStdinWithContext(ctx context.Context, forceClose context.CancelFunc) (stdin *Stdin) {
	stdin = new(Stdin)
	stdin.max = DefaultMaxBufferSize
	stdin.ctx = ctx
	stdin.forceClose = forceClose
	return
}

// MakePipe is used for named pipes. Basically just used to relax the exception handling since we can make fewer
// guarantees about the state of named pipes.
func (stdin *Stdin) MakePipe() {
	stdin.mutex.Lock()
	stdin.dependants++
	stdin.mutex.Unlock()
}

// Open the stream.Io interface for another dependant
func (stdin *Stdin) Open() {
	stdin.mutex.Lock()
	defer stdin.mutex.Unlock()

	stdin.dependants++
}

// Close the stream.Io interface
func (stdin *Stdin) Close() {
	stdin.mutex.Lock()
	defer stdin.mutex.Unlock()

	stdin.dependants--

	if stdin.dependants < 0 {
		panic("More closed dependants than open")
	}
}

// ForceClose forces the stream.Io interface to close. This should only be called by a STDIN reader
func (stdin *Stdin) ForceClose() {
	if stdin.forceClose != nil {
		stdin.forceClose()
	}
}
