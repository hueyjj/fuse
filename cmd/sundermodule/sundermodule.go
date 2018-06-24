package main

import (
	"fmt"
	"github.com/hueyjj/fuse/pkg/sundermodule"
	"os"
)

func main() {
	err := sundermodule.CheckCmdArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
}
