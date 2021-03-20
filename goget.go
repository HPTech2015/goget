package main

import (
	"goget/core"
	"log"
)

func main() {
	// Import Settings
	settings := core.Settings{Version: "1.0.0"}

	// Parse command line arguments.
	argParser := core.ArgParser{}
	err := argParser.ArgParse(&settings)
	if err != nil {
		log.Fatal(err)
	}

	if settings.RemoteTarget == "" {
		return
	}

	// Get file and checksum from remote web server.
	getFile := core.GetFile{}
	validSig, err := getFile.PullAndCheck(settings.LocalTarget, settings.RemoteTarget)
	if err != nil {
		log.Fatal(err)
	}

	if settings.SkipValidation {
		return
	}

	/*
		Compare the checksum from the remote server with
		the checksum from the file downloaded.
	*/
	if !validSig {
		log.Fatalf("File Signature Validation Failed!\n\n" + 
		"Expected Signature:\n    %v\n" +
		"File Signature:\n    %v", 
		getFile.RemoteChecksum, getFile.LocalChecksum)
	}
}