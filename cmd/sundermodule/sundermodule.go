package main

import (
	sm "github.com/hueyjj/fuse/pkg/sundermodule"
	"os"
	//"bytes"
	"log"
)

func buildCliCmds() map[string]sm.CliCmd {
	return map[string]sm.CliCmd{
		"yt_download_music": sm.CliCmd{
			CommandName: "yt_download_music",
			AppName:     "youtube-dl",
			Options: map[sm.FlagAlias]*sm.Option{
				"embed_thumbnail": &sm.Option{
					Flag: "--embed-thumbnail",
				},
				"add_meta_data": &sm.Option{
					Flag: "--add-metadata",
				},
				"format": &sm.Option{
					Flag: "--format", Value: "m4a",
				},
			},
		},
	}
}

func main() {
	cliCmds := buildCliCmds()

	im, err := sm.CheckIncMsg(cliCmds)
	if err != nil {
		log.Printf("%+v\n", im)
		log.Fatal(err)
		os.Exit(1)
	}

	cliCmd := sm.FillCommand(cliCmds[im.CommandName], im)

	err = sm.RunCommand(cliCmd, `C:\Users\JJ\Downloads`)
	if err != nil {
		//log.Fatal(err)
		os.Stdout.WriteString(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
