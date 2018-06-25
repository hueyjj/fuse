package sundermodule

import (
	"os/exec"
	//"strings"
	"fmt"
)

type ShellOption struct {
	Name string

	AppName  string
	Commands []string
	Args     map[FlagAlias]Arg
}

type Arg struct {
	Flag  string
	Value string
}

type FlagAlias string
type Flag string

type ShellOptionBuilder interface {
	SetName(name string) ShellOptionBuilder

	SetAppName(appName string) ShellOptionBuilder

	AddCommand(command string) ShellOptionBuilder

	AddArg(flagAlias FlagAlias, arg Arg) ShellOptionBuilder

	Build() (*ShellOption, error)
}

// NewShellOptionBuilder creates new builder
func NewShellOptionBuilder() ShellOptionBuilder {
	return &ShellOption{
		Args: make(map[FlagAlias]Arg),
	}
}

// SetName sets the name
func (so *ShellOption) SetName(name string) ShellOptionBuilder {
	so.Name = name
	return so
}

// SetAppName sets the app name
func (so *ShellOption) SetAppName(appName string) ShellOptionBuilder {
	so.AppName = appName
	return so
}

// AddCommand add a command
func (so *ShellOption) AddCommand(command string) ShellOptionBuilder {
	so.Commands = append(so.Commands, command)
	return so
}

// AddArg add argument
func (so *ShellOption) AddArg(flagAlias FlagAlias, arg Arg) ShellOptionBuilder {
	so.Args[flagAlias] = arg
	return so
}

// Build builds the option to be executed
func (so *ShellOption) Build() (*ShellOption, error) {
	return &ShellOption{
		Name:     so.Name,
		AppName:  so.AppName,
		Commands: so.Commands,
		Args:     so.Args,
	}, nil
}

// SetArgValue sets the argument value
func (so *ShellOption) SetArgValue(flagAlias FlagAlias, value string) {
	so.Args[flagAlias] = Arg{
		Flag:  so.Args[flagAlias].Flag,
		Value: value,
	}
}

func (so *ShellOption) buildArgs() []string {
	var args []string
	for _, arg := range so.Args {
		if arg.Flag != "" {
			args = append(args, arg.Flag)
		}
		if arg.Value != "" {
			args = append(args, arg.Value)
		}
	}

	return args
}

// BuildCmd builds the command to be executed by exec
func (so *ShellOption) BuildCmd() *exec.Cmd {
	args := append(so.Commands, so.buildArgs()...)
	cmd := exec.Command(
		so.AppName,
		args...,
	)
	return cmd
}
