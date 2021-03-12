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

	// Print Settings
	remoteTarget, err := settings.GetRemoteTarget()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Remote Target:", remoteTarget)

	localTarget, err := settings.GetLocalTarget()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Local Target:", localTarget)
}