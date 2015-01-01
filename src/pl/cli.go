package main

import (
	"fmt"
	"io"
	flag "mflag"
	"os"
	"reflect"
	"strings"
)

type PlCli struct {
	in      io.ReadCloser
	out     io.Writer
	err     io.Writer
	verbose bool
}

func (cli *PlCli) getMethod(args ...string) (func(...string) error, bool) {
	camelArgs := make([]string, len(args))
	for i, s := range args {
		if len(s) == 0 {
			return nil, false
		}
		camelArgs[i] = strings.ToUpper(s[:1])+strings.ToLower(s[1:])
	}
	methodName := "Cmd" + strings.Join(camelArgs, "")
	method := reflect.ValueOf(cli).MethodByName(methodName)
	if !method.IsValid() {
		return nil, false
	}
	return method.Interface().(func(...string) error), true
}

func (cli *PlCli) Cmd(args ...string) error {
	if len(args) > 0 {
		method, exists := cli.getMethod(args[:1]...)
		if !exists {
			fmt.Printf("Error: Command not found: %s\n", args[0])
			return cli.CmdHelp()
		}
		return method(args[1:]...)
	}
	return cli.CmdHelp()
}

func (cli *PlCli) SubCmd(name, signature, description string) *flag.FlagSet {
	flags := flag.NewFlagSet(name, flag.ContinueOnError)
	flags.Usage = func() {
		options := ""
		if flags.FlagCountUndeprecated() > 0 {
			options = "[OPTIONS] "
		}
		fmt.Fprintf(cli.err, "\nUsage: pl %s %s%s\n\n%s\n\n", name, options, signature, description)
		flags.PrintDefaults()
		os.Exit(2)
	}
	return flags
}

func CreatePlCli(in io.ReadCloser, out, err io.Writer, verbose bool) *PlCli {
	return &PlCli{
		in:      in,
		out:     out,
		err:     err,
		verbose: verbose,
	}
}
