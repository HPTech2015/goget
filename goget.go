package main

import (
	"fmt"
	"goget/core"
)

func main() {
	// Import Settings
	settings := core.Settings{Version: "v1.0.0"}

	// Parse command line arguments.
	argParser := core.ArgParser{}
	err := argParser.ArgParse(&settings)
	if err != nil {
		panic(err)
	}

	// Get file and checksum from remote web server.
	getFile := core.GetFile{}
	validSig, err := getFile.PullAndCheck(settings.LocalTarget, settings.RemoteTarget)
	if err != nil {
		panic(err)
	}

	/*
		Compare the checksum from the remote server with
		the checksum from the file downloaded.
	*/
	if !validSig {
		fmt.Println("File Signature Validation Failed!")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("Expected Signature:")
		fmt.Println("    ", getFile.RemoteChecksum)
		fmt.Println("File Signature:")
		fmt.Println("    ", getFile.LocalChecksum)
	}
}