package main

import (
	"fmt"
	"github.com/hueyjj/fuse/pkg/sundermodule"
	"os"
)

func buildShellOptions() []sundermodule.ShellOption {
	downloadOpt, err := sundermodule.NewShellOptionBuilder().
		SetName("yt_download_video").
		SetAppName("youtube-dl").
		//AddCommand("").
		AddArgs("download", sundermodule.Arg{Flag: "d"}).
		Build()

	if err != nil {

	}

	return []sundermodule.ShellOption{
		*downloadOpt,
	}
}

func main() {
	var im sundermodule.IncomingMessage
	if im, err := sundermodule.CheckCmdArgs(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n%s\n", im, err)
		os.Exit(1)
	}

	fmt.Println(im)

	so, err := sundermodule.NewShellOptionBuilder().
		SetName(im.Name).
		SetAppName("Foo").
		AddCommand("smoething").
		AddArgs("download", sundermodule.Arg{Flag: "d"}).
		Build()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	if so.Args["download"].Value == "" {
		fmt.Println("Download value is nil")
	}
	if _, ok := so.Args["foo"]; !ok {
		fmt.Println("No foo option")
	}

	fmt.Printf("%+v\n", so)
	//sundermodule.Execute(?)
}
