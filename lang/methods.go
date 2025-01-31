package lang

import (
	"sync"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

type methods struct {
	mutex sync.Mutex
	dt    map[string][]string
}

func newMethods() *methods {
	m := new(methods)
	m.dt = make(map[string][]string)
	return m
}

func (m *methods) Exists(cmd, dataType string) bool {
	m.mutex.Lock()
	i := m.exists(cmd, dataType)
	m.mutex.Unlock()
	return i != -1
}

func (m *methods) exists(cmd, dataType string) int {
	cmds := m.dt[dataType]

	for i := range cmds {
		if cmds[i] == cmd {
			return i
		}
	}

	return -1
}

// Define creates a record of a new method
func (m *methods) Define(cmd, dataType string) {
	m.mutex.Lock()

	cmds := m.dt[dataType]

	if m.exists(cmd, dataType) != -1 {
		m.mutex.Unlock()
		debug.Log("method define", cmd, dataType, "exists")
		return
	}

	m.dt[dataType] = append(cmds, cmd)
	m.mutex.Unlock()
}

// Degroup takes the commands assigned to group types and sorts them back into individual types
func (m *methods) Degroup() {
	for group := range m.dt {
		if group[0] == '@' && group != types.Any {
			m.degroup(group, groups(group))
		}
	}
}

func (m *methods) degroup(group string, dataTypes []string) {
	cmds := m.Get(group)
	for i := range dataTypes {
		for j := range cmds {
			m.Define(cmds[j], dataTypes[i])
		}
	}
	m.mutex.Lock()
	delete(m.dt, group)
	m.mutex.Unlock()
}

func groups(group string) []string {
	switch group {
	case types.Text:
		return types.GroupText

	case types.Math:
		return types.GroupMath

	case types.Marshal:
		return DumpUnmarshaller()

	case types.Unmarshal:
		return DumpMarshaller()

	case types.ReadArray:
		return stdio.DumpReadArray()

	case types.ReadArrayWithType:
		return stdio.DumpReadArrayWithType()

	case types.WriteArray:
		return stdio.DumpWriteArray()

	case types.ReadMap:
		return stdio.DumpMap()

	case types.ReadIndex:
		return DumpIndex()

	case types.ReadNotIndex:
		return DumpNotIndex()

	default:
		panic("Group name doesn't have a programmed list of data types: " + group)
	}
}

// Get returns all the methods for a murex data type
func (m *methods) Get(dataType string) (cmds []string) {
	m.mutex.Lock()
	cmds = m.get(dataType)
	m.mutex.Unlock()
	return
}

func (m *methods) get(dataType string) []string {
	cmds := m.dt[dataType]

	if cmds == nil {
		return []string{}
	}

	s := make([]string, len(cmds))
	copy(s, cmds)

	return s
}

// Types returns all the data types supported by a command
func (m *methods) Types(cmd string) (dataTypes []string) {
	dump := m.Dump()

	for dt := range dump {
		if dt == types.Any {
			continue
		}

		for i := range dump[dt] {
			if dump[dt][i] == cmd {
				dataTypes = append(dataTypes, dt)
			}
		}
	}

	return
}

// Dump returns all methods for `runtime`
func (m *methods) Dump() map[string][]string {
	m.mutex.Lock()

	dump := make(map[string][]string)

	for dt := range m.dt {
		dump[dt] = m.get(dt)
	}

	m.mutex.Unlock()
	return dump
}
