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
	AppName string          `json:"app_name"`
	Data    json.RawMessage `json:"data"`
}

// CheckIncMsg checks stdin if a JSON object was passed and is valid
func CheckIncMsg(cliCmds map[string]CliCmd) (IncomingMessage, error) {
	var msg string
	var im IncomingMessage
	if len(os.Args) == 2 {
		msg = os.Args[1]
		log.Printf("%v+\n", msg)

		if msg == "get" {
			printAPI(cliCmds)
			os.Exit(0)
		} else if msg == "prettyget" {
			prettyPrintAPI(cliCmds)
			os.Exit(0)
		} else if im, ok := commandExists(msg, cliCmds); ok {
			return im, nil
		}
	} else if len(os.Args) > 2 {
		return im, fmt.Errorf("usage: %s [JSON string]", os.Args[0])
	}

	return im, errors.New("CheckIncMsg: required JSON object as a string as input")
}

// commandExists checks if the string is of type IncomingMessage
func commandExists(str string, cliCmds map[string]CliCmd) (IncomingMessage, bool) {
	var im IncomingMessage
	isValid := json.Unmarshal([]byte(str), &im) == nil

	if im.AppName == "" || im.Data == nil {
		return im, false
	}

	// Check if commands exist
	if _, ok := cliCmds[im.AppName]; ok {
		// Check if the data matches

		log.Printf("command exists")
	}

	return im, isValid
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
