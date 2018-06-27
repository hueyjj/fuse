package sundermodule

import (
	"os/exec"
	//"strings"
	//"fmt"
	"bufio"
	"io"
	"log"
	"strings"
)

type CliCmd struct {
	CommandName string

	AppName string
	Command string
	Args    []string
	Options map[FlagAlias]Option
}

type Option struct {
	Flag  string
	Value string
}

type FlagAlias string
type Flag string

func removeEmptyStrings(a []string) []string {
	b := a[:0]
	for _, x := range a {
		if x != "" {
			b = append(b, x)
		}
	}
	return b
}

func (so *CliCmd) buildOpts() []string {
	var opts []string
	for _, opt := range so.Options {
		opts = append(opts, opt.Flag)
		opts = append(opts, opt.Value)
	}

	return opts
}

// BuildCmd builds the command to be executed by exec
func (so *CliCmd) BuildCmd() *exec.Cmd {
	args := []string{so.Command}
	args = append(args, so.Args...)
	args = append(args, so.buildOpts()...)

	args = removeEmptyStrings(args)

	//log.Println(so.AppName)
	//log.Println(len(so.buildOpts()))
	//log.Println(so.buildOpts())
	log.Printf("Command built with %d args:", len(args))
	log.Printf("> %s %v ", so.AppName, args)
	cmd := exec.Command(
		so.AppName,
		args...,
	)
	return cmd
}

// RunCommand runs a command
func RunCommand(cliCmd CliCmd, dir string) error {
	log.Printf("%v+\n", cliCmd)
	cmd := cliCmd.BuildCmd()
	cmd.Dir = dir

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	args := append(cliCmd.Args, cliCmd.buildOpts()...)
	log.Printf("Starting command: %s\n%s %s\n", cliCmd.CommandName, cliCmd.AppName, strings.Join(args, " "))

	go LogOutput(stdout)
	go LogOutput(stderr)

	if err = cmd.Wait(); err != nil {
		log.Printf("%s returned error: %v", cliCmd.CommandName, err)
		// TODO Customize the error here, do not end the program here
	}

	return err
}

// LogOutput logs a reader (stdout, stderr)
func LogOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}
}
