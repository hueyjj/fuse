package main

import (
	sm "github.com/hueyjj/fuse/pkg/sundermodule"
	"os"
	//"bytes"
	"log"
)

func buildShellOptions() map[string]sm.CliCmd {
	return map[string]sm.CliCmd{
		"yt_download_music": sm.CliCmd{
			CommandName: "yt_download_music",
			AppName:     "youtube-dl",
			Options: map[sm.FlagAlias]sm.Option{
				"embed_thumbnail": sm.Option{
					Flag: "--embed-thumbnail",
				},
				"add_meta_data": sm.Option{
					Flag: "--add-metadata",
				},
				"format": sm.Option{
					Flag: "--format", Value: "m4a",
				},
			},
		},
	}
}

func main() {
	shellOpts := buildShellOptions()

	im, err := sm.CheckIncMsg(shellOpts)
	if err != nil {
		log.Printf("%+v\n", im)
		log.Fatal(err)
		os.Exit(1)
	}
	log.Printf("%+v\n", im)

	//opt := shellOpts["yt_download_video"]
	//opt.AddCommand("https://www.youtube.com/watch?v=64DtWBXjU2Y")
	//opt.SetArgValue("format", "m4a")

	//err := sm.RunCommand(opt, "C:\\Users\\JJ\\Downloads")
	//if err != nil {
	//	log.Fatal(err)
	//}
}
