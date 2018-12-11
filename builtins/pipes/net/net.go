package net

import (
	"context"
	"net"
	"os"
	"sync"

	"github.com/lmorg/murex/utils"
)

// Net Io interface
type Net struct {
	mutex      sync.Mutex
	ctx        context.Context
	forceClose func()
	bRead      uint64
	bWritten   uint64
	dependants int
	conn       net.Conn
	dataType   string
	protocol   string
}

// DefaultDataType is unavailable for net Io interfaces
func (n *Net) DefaultDataType(bool) {}

// IsTTY always returns false because net Io interfaces are not a pseudo-TTY
func (n *Net) IsTTY() bool { return false }

// MakePipe turns the stream.Io interface into a named pipe
func (n *Net) MakePipe() {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	//n.isParent = true
	n.dependants++
}

// SetDataType assigns a data type to the stream.Io interface
func (n *Net) SetDataType(dt string) {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	n.dataType = dt
}

// GetDataType read the stream.Io interface's data type
func (n *Net) GetDataType() string {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	return n.dataType
}

// Stats returns the bytes written and bytes read to the net Io interface
func (n *Net) Stats() (bytesWritten, bytesRead uint64) {
	n.mutex.Lock()
	bytesWritten = n.bWritten
	bytesRead = n.bRead
	n.mutex.Unlock()
	return
}

// Open net Io interface
func (n *Net) Open() {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	n.dependants++
}

// Close net Io interface
func (n *Net) Close() {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	n.dependants--

	if n.dependants == 0 {
		err := n.conn.Close()
		if err != nil {
			os.Stderr.WriteString(err.Error() + utils.NewLineString)
		}
	}

	if n.dependants < 0 {
		panic("more closed dependants than open")
	}
}

// ForceClose forces the net Io interface to close. This should only be called on reader interfaces
func (n *Net) ForceClose() {
	n.forceClose()
}
