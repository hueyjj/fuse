package main

import (
	"github.com/hueyjj/fuse/pkg/sundermodule"
	// "os"
	//"bytes"
	"log"
)

func buildShellOptions() map[string]sundermodule.ShellOption {
	downloadOpt, err := sundermodule.NewShellOptionBuilder().
		SetName("yt_download_video").
		SetAppName("youtube-dl").
		//AddCommand("").
		AddArg("embed_thumbnail", sundermodule.Arg{Flag: "--embed-thumbnail"}).
		AddArg("add_meta_data", sundermodule.Arg{Flag: "--add-metadata"}).
		AddArg("format", sundermodule.Arg{Flag: "--format"}).
		Build()

	if err != nil {

	}

	return map[string]sundermodule.ShellOption{
		downloadOpt.Name: *downloadOpt,
	}
}

func main() {
	shellOpts := buildShellOptions()

	//var im sundermodule.IncomingMessage
	//if im, err := sundermodule.CheckCmdArgs(); err != nil {
	//	fmt.Fprintf(os.Stderr, "%+v\n%s\n", im, err)
	//	os.Exit(1)
	//}

	// Validate incoming

	opt := shellOpts["yt_download_video"]
	opt.AddCommand("https://www.youtube.com/watch?v=64DtWBXjU2Y")
	opt.SetArgValue("format", "m4a")

	err := sundermodule.RunCommand(opt, "C:\\Users\\JJ\\Downloads")
	if err != nil {
		log.Fatal(err)
	}
}
