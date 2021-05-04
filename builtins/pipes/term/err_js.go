// +build js

package term

import (
	"html"
	"syscall/js"

	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/utils"
)

// Terminal: Standard Error

// Err is the Stderr interface for term
type Err struct {
	term
}

// Write is the io.Writer() interface for term
func (t *Err) Write(b []byte) (int, error) {
	t.mutex.Lock()
	t.bWritten += uint64(len(b))
	t.mutex.Unlock()

	jsDoc := js.Global().Get("document")
	outElement := jsDoc.Call("getElementById", "term")

	term := outElement.Get("innerHTML").String()
	new := html.EscapeString(string(b))
	outElement.Set("innerHTML", term+new)

	return len(b), nil
}

// Writeln writes an OS-specific terminated line to the stderr
func (t *Err) Writeln(b []byte) (int, error) {
	return t.Write(appendBytes(b, utils.NewLineByte...))
}

// WriteArray performs data type specific buffered writes to an stdio.Io interface
func (t *Err) WriteArray(dataType string) (stdio.ArrayWriter, error) {
	return stdio.WriteArray(t, dataType)
}
