package sundermodule

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

// IncomingMessage represents a command to execute on some data
type IncomingMessage struct {
	CommandName string                   `json:"command_name"`
	Args        []string                 `json:"args"`
	Options     map[FlagAlias]UserOption `json:"options"`
}

type UserOption map[string]string

// CheckIncMsg checks stdin if a JSON object was passed and is valid
func CheckIncMsg(cliCmds map[string]CliCmd) (IncomingMessage, error) {
	var msg string
	var im IncomingMessage
	if len(os.Args) == 2 {
		msg = os.Args[1]

		if msg == "get" {
			log.Printf("%s\n", "get command received")
			printAPI(cliCmds)
			os.Exit(0)
		} else if msg == "prettyget" {
			log.Printf("%s\n", "prettyget command received")
			prettyPrintAPI(cliCmds)
			os.Exit(0)
		} else if im, ok := parseCommand(msg, cliCmds); ok {
			//log.Printf("%v+\n", msg)
			return im, nil
		}
	} else if len(os.Args) > 2 {
		return im, fmt.Errorf("usage: %s [JSON string]", os.Args[0])
	}

	return im, errors.New("CheckIncMsg: required JSON object as a string as input")
}

// parseCommand checks if the string is of type IncomingMessage
func parseCommand(str string, cliCmds map[string]CliCmd) (IncomingMessage, bool) {
	var im IncomingMessage
	if json.Unmarshal([]byte(str), &im) != nil {
		return im, false
	}

	if im.CommandName == "" {
		return im, false
	}

	// Check if commands exist
	if _, ok := cliCmds[im.CommandName]; ok {
		// Check if the data matches

		//log.Printf("command exists")
	}

	return im, true
}

func printAPI(cliCmds map[string]CliCmd) {
	jsonCliCmds, err := json.Marshal(cliCmds)
	if err != nil {
		log.Panicf("Unable to convert api to json format")
		os.Exit(1)
	}
	os.Stdout.Write(jsonCliCmds)
}

func prettyPrintAPI(cliCmds map[string]CliCmd) {
	jsonCliCmds, err := json.MarshalIndent(cliCmds, "", "  ")
	if err != nil {
		log.Panicf("Unable to convert api to json format")
		os.Exit(1)
	}
	os.Stdout.Write(jsonCliCmds)
}
