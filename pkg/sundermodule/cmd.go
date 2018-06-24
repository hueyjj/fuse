package sundermodule

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// IncomingMessage represents a command to execute on some data
type IncomingMessage struct {
	Name string `json:"name"`
	Data json.RawMessage
}

// CheckCmdArgs checks stdin if a JSON object was passed and is valid
func CheckCmdArgs() (IncomingMessage, error) {
	var msg string
	var im IncomingMessage
	if len(os.Args) == 2 {
		msg = os.Args[1]
		fmt.Println(msg)
		if im, ok := isValidInput(msg); ok {
			return im, nil
		}
	} else if len(os.Args) > 2 {
		return im, fmt.Errorf("usage: %s [JSON string]", os.Args[0])
	}

	return im, errors.New("CheckCmdArgs: required JSON object as a string as input")
}

// isValidInput checks if the string is of type IncomingMessage
func isValidInput(str string) (IncomingMessage, bool) {
	var im IncomingMessage
	isValid := json.Unmarshal([]byte(str), &im) == nil

	// TODO: Check if command exists in dev defined map of commands

	return im, isValid
}

// Execute the command on the data
func Execute(im IncomingMessage) {
}
