package proc

import (
	"errors"
	"sync"
)

// Aliases is a table of aliases
type Aliases struct {
	aliases map[string][]string
	mutex   sync.Mutex
}

// NewAliases creates a new table of aliases
func NewAliases() (a Aliases) {
	a.aliases = make(map[string][]string)
	return
}

// Add creates an alias
func (a *Aliases) Add(name string, alias []string) {
	a.mutex.Lock()
	a.aliases[name] = alias
	a.mutex.Unlock()
}

// Exists checks if alias exists in table
func (a *Aliases) Exists(name string) bool {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return len(a.aliases[name]) > 0
}

// Get the aliased code
func (a *Aliases) Get(name string) (alias []string) {
	a.mutex.Lock()
	alias = a.aliases[name]
	a.mutex.Unlock()
	return
}

// Delete an alias
func (a *Aliases) Delete(name string) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if len(a.aliases[name]) == 0 {
		return errors.New("Alias does not exist.")
	}
	delete(a.aliases, name)
	return nil
}

// Dump returns the complete alias table
func (a *Aliases) Dump() map[string][]string {
	a.mutex.Lock()
	dump := a.aliases
	a.mutex.Unlock()
	return dump
}
