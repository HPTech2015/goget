package main

import (
	"fmt"
	"goget/core"
)

func main() {
	// Import Settings
	settings := core.Settings{Version: "v1.0.0"}

	argParser := core.ArgParser{}
	err := argParser.ArgParse(&settings)
	if err != nil {
		fmt.Println(err)
	}

	getFile := core.GetFile{}
	getFile.Pull(settings.LocalTarget, settings.RemoteTarget)
}