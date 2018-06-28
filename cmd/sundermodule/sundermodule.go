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

	//cliCmd := cliCmds["yt_download_music"]
	//cliCmd.Args = append(cliCmd.Args, "https://www.youtube.com/watch?v=64DtWBXjU2Y")

	cliCmd := sm.FillCommand(cliCmds[im.CommandName], im)
	log.Printf("%+v\n", cliCmd)

	err = sm.RunCommand(cliCmd, `C:\Users\JJ\Downloads`)
	if err != nil {
		log.Fatal(err)
	}
}
