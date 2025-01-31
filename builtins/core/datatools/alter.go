package datatools

import (
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/alter"
)

func init() {
	lang.DefineMethod("alter", cmdAlter, types.Unmarshal, types.Marshal)

	defaults.AppendProfile(`
		autocomplete: set alter { [{
			"AnyValue": true,
			"ExecCmdline": true,
			"AutoBranch": true,
			"Dynamic": ({ -> struct-keys }),
			"FlagsDesc": {
				"--merge": "Merge data structures rather than overwrite",
				"--sum": "Sum values in a map, merge items in an array"
			},
			"FlagValues": {
				"-m": [{
					"AnyValue": true,
					"ExecCmdline": true,
					"AutoBranch": true,
					"Dynamic": ({ -> struct-keys })
				}],
				"--merge": [{
					"AnyValue": true,
					"ExecCmdline": true,
					"AutoBranch": true,
					"Dynamic": ({ -> struct-keys })
				}],

				"-s": [{
					"AnyValue": true,
					"ExecCmdline": true,
					"AutoBranch": true,
					"Dynamic": ({ -> struct-keys })
				}],
				"--sum": [{
					"AnyValue": true,
					"ExecCmdline": true,
					"AutoBranch": true,
					"Dynamic": ({ -> struct-keys })
				}]
			}
		} ]}
	`)
}

func cmdAlter(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	const (
		alterAlter int = 0
		alterMerge int = iota + 1
		alterSum
	)

	var (
		action int
		offset int
	)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	v, err := lang.UnmarshalData(p, dt)
	if err != nil {
		return err
	}

	s, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	if s == "-m" || s == "--merge" {
		action = alterMerge
		offset++

		s, err = p.Parameters.String(1)
		if err != nil {
			return err
		}
	}

	if s == "-s" || s == "--sum" {
		action = alterSum
		offset++

		s, err = p.Parameters.String(1)
		if err != nil {
			return err
		}
	}

	new, err := p.Parameters.String(1 + offset)
	if err != nil {
		return err
	}

	path, err := alter.SplitPath(s)
	if err != nil {
		return err
	}

	switch action {
	default:
		v, err = alter.Alter(p.Context, v, path, new)
		if err != nil {
			return err
		}

	case alterMerge:
		v, err = alter.Merge(p.Context, v, path, new)
		if err != nil {
			return err
		}

	case alterSum:
		v, err = alter.Sum(p.Context, v, path, new)
		if err != nil {
			return err
		}
	}

	b, err := lang.MarshalData(p, dt, v)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
