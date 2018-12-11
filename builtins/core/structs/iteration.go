package structs

import (
	"errors"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["for"] = cmdFor
	proc.GoFunctions["foreach"] = cmdForEach
	proc.GoFunctions["formap"] = cmdForMap
	proc.GoFunctions["while"] = cmdWhile
	proc.GoFunctions["!while"] = cmdWhile
}

// Example usage:
// for ( i=1; i<6; i++ ) { echo $i }
func cmdFor(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Generic)

	cblock, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	block, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	parameters := strings.Split(string(cblock), ";")
	if len(parameters) != 3 {
		return errors.New("Invalid syntax. Must be ( variable; conditional; incremental )")
	}

	variable := "let " + parameters[0]
	conditional := "eval " + parameters[1]
	incremental := "let " + parameters[2]

	_, err = lang.RunBlockExistingConfigSpace([]rune(variable), nil, nil, p.Stderr, p)
	if err != nil {
		return err
	}

	for {
		if p.HasCancelled() {
			return errors.New(errCancelled)
		}

		stdout := streams.NewStdin()
		i, err := lang.RunBlockExistingConfigSpace([]rune(conditional), nil, stdout, p.Stderr, p)
		if err != nil {
			return err
		}

		b, err := stdout.ReadAll()
		if err != nil {
			return err
		}
		if !types.IsTrue(b, i) {
			return nil
		}

		// Execute block.
		lang.RunBlockExistingConfigSpace(block, nil, p.Stdout, p.Stderr, p)

		// Increment counter.
		_, err = lang.RunBlockExistingConfigSpace([]rune(incremental), nil, nil, p.Stderr, p)
		if err != nil {
			return err
		}
	}

	//return nil
}

func cmdForEach(p *proc.Process) (err error) {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	var (
		block   []rune
		varName string
	)

	switch p.Parameters.Len() {
	case 1:
		block, err = p.Parameters.Block(0)
		if err != nil {
			return err
		}

	case 2:
		block, err = p.Parameters.Block(1)
		if err != nil {
			return err
		}

		varName, err = p.Parameters.String(0)
		if err != nil {
			return err
		}

	default:
		return errors.New("Invalid number of parameters.")
	}

	err = p.Stdin.ReadArray(func(b []byte) {
		if len(b) == 0 || p.HasCancelled() {
			return
		}

		if varName != "" {
			p.Variables.Set(varName, string(b), dt)
		}

		stdin := streams.NewStdin()
		stdin.SetDataType(dt)
		stdin.Writeln(b)

		lang.RunBlockExistingConfigSpace(block, stdin, p.Stdout, p.Stderr, p)
	})

	return err
}

func cmdForMap(p *proc.Process) error {
	p.Stdout.SetDataType(types.Generic)
	dt := p.Stdin.GetDataType()
	//p.Stdout.SetDataType(dt)

	block, err := p.Parameters.Block(2)
	if err != nil {
		return err
	}

	varKey, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	varVal, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	err = p.Stdin.ReadMap(p.Config, func(key, value string, last bool) {
		if p.HasCancelled() {
			return
		}

		p.Variables.Set(varKey, key, types.String)
		p.Variables.Set(varVal, value, dt)

		lang.RunBlockExistingConfigSpace(block, nil, p.Stdout, p.Stderr, p)
	})

	return err
}

func cmdWhile(p *proc.Process) error {
	p.Stdout.SetDataType(types.Generic)

	switch p.Parameters.Len() {
	case 1:
		// Condition is taken from the while loop.
		block, err := p.Parameters.Block(0)
		if err != nil {
			return err
		}

		for {
			if p.HasCancelled() {
				return errors.New(errCancelled)
			}

			stdout := streams.NewStdin()
			i, err := lang.RunBlockExistingConfigSpace(block, nil, stdout, p.Stderr, p)
			if err != nil {
				return err
			}
			b, err := stdout.ReadAll()
			if err != nil {
				return err
			}

			_, err = p.Stdout.Write(b)
			if err != nil {
				return err
			}

			conditional := types.IsTrue(b, i)

			if (!p.IsNot && !conditional) ||
				(p.IsNot && conditional) {
				return nil
			}

		}

	case 2:
		// Condition is first parameter, while loop is second.
		ifBlock, err := p.Parameters.Block(0)
		if err != nil {
			return err
		}

		whileBlock, err := p.Parameters.Block(1)
		if err != nil {
			return err
		}

		for {
			if p.HasTerminated() {
				return nil
			}

			stdout := streams.NewStdin()
			i, err := lang.RunBlockExistingConfigSpace(ifBlock, nil, stdout, nil, p)
			if err != nil {
				return err
			}
			b, err := stdout.ReadAll()
			if err != nil {
				return err
			}
			conditional := types.IsTrue(b, i)

			if (!p.IsNot && !conditional) ||
				(p.IsNot && conditional) {
				return nil
			}

			lang.RunBlockExistingConfigSpace(whileBlock, nil, p.Stdout, p.Stderr, p)
		}

	default:
		// Error
		return errors.New("Invalid number of parameters. Please read usage notes.")
	}
}
