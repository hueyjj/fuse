package sundermodule

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// IncomingMessage represents a command to execute on some data
type IncomingMessage struct {
	Command string `json:"command"`
	Data    json.RawMessage
}

// CheckCmdArgs Checks stdin if a JSON object was passed and is valid
func CheckCmdArgs() error {
	var msg string
	if len(os.Args) == 2 {
		msg = os.Args[1]
		fmt.Println(msg)
		if isValidInput(msg) {
			return nil
		}
	} else if len(os.Args) > 2 {
		return fmt.Errorf("usage: %s [JSON string]", os.Args[0])
	}

	return errors.New("CheckCmdArgs: required JSON object as a string as input")
}

func isValidInput(str string) bool {
	var im IncomingMessage
	return json.Unmarshal([]byte(str), &im) == nil
}

//func Parse() {
//}
