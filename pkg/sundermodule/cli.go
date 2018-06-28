package sundermodule

import (
	"os/exec"
	//"strings"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type CliCmd struct {
	// Alias for easier API readability
	CommandName string /* Client doesn't modify */

	// Actual application name
	AppName string /* Client doesn't modify */

	// Application command
	Command string /* Client doesn't modify */

	// FIXME Command line are order sensitive. Must give args position
	// when building the full command line.
	// List of args
	Args []string /* Client can modify */

	// Flags and values to be passed
	Options map[FlagAlias]*Option /* Client can modify */
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
	log.Printf("%s running...\n", cliCmd.CommandName)
	log.Printf("%s %s\n", cliCmd.AppName, strings.Join(args, " "))

	go logOutput(stdout)
	go logOutput(stderr)

	if err = cmd.Wait(); err != nil {
		log.Printf("%s returned error: %v", cliCmd.CommandName, err)
		// TODO Customize the error here, do not end the program here
	}

	return err
}

// logOutput logs a reader (stdout, stderr)
func logOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		// Write to Stdout even if io.Reader connects to stdout/stderr pipe
		os.Stdout.Write([]byte(scanner.Text()))
		os.Stdout.Write([]byte("\n"))
	}
}

func getCommand(cliCmds map[string]CliCmd, im IncomingMessage) (CliCmd, error) {
	if cmd, ok := cliCmds[im.CommandName]; ok {
		return cmd, nil
	}
	return CliCmd{}, fmt.Errorf("Unable to find %s in available command options", im.CommandName)
}

func FillCommand(cliCmd CliCmd, im IncomingMessage) CliCmd {
	cliCmd.Args = append(cliCmd.Args, im.Args...)
	for flagAlias, userOption := range im.Options {
		cliCmd.Options[flagAlias].Value = userOption["value"]
	}
	return cliCmd
}
