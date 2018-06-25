package main

import (
	"fmt"
	"github.com/hueyjj/fuse/pkg/sundermodule"
	// "os"
	"bytes"
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
	//var im sundermodule.IncomingMessage
	//if im, err := sundermodule.CheckCmdArgs(); err != nil {
	//	fmt.Fprintf(os.Stderr, "%+v\n%s\n", im, err)
	//	os.Exit(1)
	//}

	// Validate incoming

	shellOpts := buildShellOptions()

	cmdOpt := shellOpts["yt_download_video"]
	cmdOpt.AddCommand("https://www.youtube.com/watch?v=64DtWBXjU2Y")
	cmdOpt.SetArgValue("format", "m4a")

	cmd := cmdOpt.BuildCmd()
	cmd.Dir = "C:\\Users\\JJ\\Downloads"
	fmt.Printf("%+v\n", cmd)

	var outbuf, errbuf bytes.Buffer

	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start cmd: %v", err)
		panic(err)
	}

	stdout := outbuf.String()
	stderr := errbuf.String()

	log.Println(stdout)
	log.Println(stderr)

	log.Println("DOing other stuff...")

	if err := cmd.Wait(); err != nil {
		stdout := outbuf.String()
		stderr := errbuf.String()

		log.Println(stdout)
		log.Println(stderr)
		log.Printf("Cmd returned error: %v", err)
		//panic(err)
	}
	log.Printf("Finished running")

	//sundermodule.Execute(?)
}
