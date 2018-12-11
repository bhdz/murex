package shellruntime

import (
	"errors"
	"runtime"

	"github.com/lmorg/murex/builtins/events"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	proc.GoFunctions["runtime"] = cmdRuntime

	defaults.AppendProfile(`
        autocomplete set runtime { [{
            "Dynamic": ({ runtime --help }),
            "AllowMultiple": true
        }] }
    `)
}

func cmdRuntime(p *proc.Process) error {
	const (
		fVars          = "--vars"
		fAliases       = "--aliases"
		fConfig        = "--config"
		fNamedPipes    = "--named-pipes"
		fPipes         = "--pipes"
		fFuncs         = "--funcs"
		fFids          = "--fids"
		fReadArrays    = "--readarray"
		fReadMaps      = "--readmap"
		fWriteArrays   = "--writearray"
		fIndexes       = "--indexes"
		fMarshallers   = "--marshallers"
		fUnmarshallers = "--unmarshallers"
		fEvents        = "--events"
		fFlags         = "--flags"
		fMemstats      = "--memstats"
		fAstCache      = "--astcache"
		fTests         = "--tests"
		fHelp          = "--help"
	)

	flags := map[string]string{
		fVars:          types.Boolean,
		fAliases:       types.Boolean,
		fConfig:        types.Boolean,
		fPipes:         types.Boolean,
		fNamedPipes:    types.Boolean,
		fFuncs:         types.Boolean,
		fFids:          types.Boolean,
		fReadArrays:    types.Boolean,
		fReadMaps:      types.Boolean,
		fWriteArrays:   types.Boolean,
		fIndexes:       types.Boolean,
		fMarshallers:   types.Boolean,
		fUnmarshallers: types.Boolean,
		fEvents:        types.Boolean,
		fFlags:         types.Boolean,
		fMemstats:      types.Boolean,
		fAstCache:      types.Boolean,
		fTests:         types.Boolean,
		fHelp:          types.Boolean,
	}

	help := func() (s []string) {
		for f := range flags {
			s = append(s, f)
		}
		return
	}

	p.Stdout.SetDataType(types.Json)

	f, _, err := p.Parameters.ParseFlags(
		&parameters.Arguments{
			Flags:           flags,
			AllowAdditional: false,
		},
	)

	if err != nil {
		return err
	}

	if len(f) == 0 {
		return errors.New("Please include one or more parameters")
	}

	ret := make(map[string]interface{})
	for flag := range f {
		switch flag {
		case fVars:
			ret[fVars[2:]] = p.Variables.Dump()
		case fAliases:
			ret[fAliases[2:]] = proc.GlobalAliases.Dump()
		case fConfig:
			//ret[fConfig[2:]] = proc.ShellProcess.Config.Dump()
			ret[fConfig[2:]] = p.Config.Dump()
		case fNamedPipes:
			ret[fNamedPipes[2:]] = proc.GlobalPipes.Dump()
		case fPipes:
			ret[fPipes[2:]] = stdio.DumpPipes()
		case fFuncs:
			ret[fFuncs[2:]] = proc.MxFunctions.Dump()
		case fFids:
			ret[fFids[2:]] = proc.GlobalFIDs.Dump()
		case fReadArrays:
			ret[fReadArrays[2:]] = stdio.DumpArray()
		case fReadMaps:
			ret[fReadMaps[2:]] = stdio.DumpMap()
		case fWriteArrays:
			ret[fWriteArrays[2:]] = stdio.DumpArray()
		case fIndexes:
			ret[fIndexes[2:]] = define.DumpIndex()
		case fMarshallers:
			ret[fMarshallers[2:]] = define.DumpMarshaller()
		case fUnmarshallers:
			ret[fUnmarshallers[2:]] = define.DumpUnmarshaller()
		case fEvents:
			ret[fEvents[2:]] = events.DumpEvents()
		case fFlags:
			ret[fFlags[2:]] = autocomplete.ExesFlags
		case fMemstats:
			var mem runtime.MemStats
			runtime.ReadMemStats(&mem)
			ret[fMemstats[2:]] = mem
		case fAstCache:
			ret[fAstCache[2:]] = lang.AstCache.Dump()
		case fTests:
			ret[fTests[2:]] = p.Tests.Dump()
		case fHelp:
			ret[fHelp[2:]] = help()
		default:
			return errors.New("Unrecognised parameter: " + flag)
		}
	}

	var b []byte
	if len(ret) == 1 {
		var obj interface{}
		for _, obj = range ret {
		}
		b, err = json.Marshal(obj, p.Stdout.IsTTY())
		if err != nil {
			return err
		}

	} else {
		b, err = json.Marshal(ret, p.Stdout.IsTTY())
		if err != nil {
			return err
		}
	}

	_, err = p.Stdout.Write(b)
	return err
}
