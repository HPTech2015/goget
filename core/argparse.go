package core

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"errors"
)

/*
	A struct with methods, for parsing
	the command line arguments.
*/
type ArgParser struct {
}

/*
	Parse char and strings arguments.
*/
func (argParser *ArgParser) ArgParse(settings *Settings) error {
	var localError error

	for i, arg := range os.Args {
		if i == 0 {
			continue
		}

		if matched, _ := regexp.MatchString("^-[a-zA-Z0-9]", arg); matched {
			argStrs, err := argParser.ParseArgChar(arg)
			if err != nil {
				localError = err
				break
			}

			for _, argStr := range argStrs {
				err := argParser.ArgInvoke(i, argStr, settings)
				if err != nil {
					localError = err
					break
				}
			}
		} else if matched, _ := regexp.MatchString("^--[a-zA-Z0-9]", arg); matched {
			err := argParser.ArgInvoke(i, arg, settings)
			if err != nil {
				localError = err
				break
			}
		}
	}

	return localError
}

/*
	Parse char arguments.

	Split each char from string and convert
	to matching string argument.
*/
func (argParser *ArgParser) ParseArgChar(arg string) ([]string, error) {
	var localError error
	argStrs := make([]string, 0)

	for _, argChar := range arg[1:] {
		switch string(argChar) {
		case "i":
			argStrs = append(argStrs, "--input-file")
		case "o":
			argStrs = append(argStrs, "--output-file")
		case "V":
			argStrs = append(argStrs, "--version")
		case "s":
			argStrs = append(argStrs, "--skip-validation")
		default:
			localError = errors.New("Invalid argument " + string(argChar))
		}
	}

	return argStrs, localError
}

/*
	Perform an action for the provided string argument.
*/
func (argParser *ArgParser) ArgInvoke(i int, arg string, settings *Settings) error {
	var localError error

	if matched, _ := regexp.MatchString("^--input-file", arg); matched {
		argVal, _ := argParser.ExtractArgVal(i, "--input-file", arg)
		if settings.LocalTarget == "" || settings.LocalTarget == "./" {
			remotePath := strings.Split(argVal, "/")
			pwd, _ := os.Getwd()
			settings.SetLocalTarget(pwd + "/" + remotePath[len(remotePath) - 1])
		}
		settings.SetRemoteTarget(argVal)
	} else if matched, _ := regexp.MatchString("^--skip-validation", arg); matched {
		settings.SetSkipValidation(true)
	} else if matched, _ := regexp.MatchString("^--output-file", arg); matched {
		argVal, _ := argParser.ExtractArgVal(i, "--output-file", arg)
		if argVal != "./" {
			settings.SetLocalTarget(argVal)
		}
	} else if matched, _ := regexp.MatchString("^--version$", arg); matched {
		v, err := settings.GetVersion()
		localError = err
		fmt.Println("LGPL GoGet", v, "built on GNU Linux.")
		fmt.Println("")
		fmt.Println("Original Author: Brendon Dobbs")
		fmt.Println("")
	} else {
		localError = fmt.Errorf("Argument %v does not exist.", arg)
	}

	return localError
}

/*
	Get the value for the command line argument.
*/
func (argParser *ArgParser) ExtractArgVal(i int, a, argAndVal string) (string, error) {
	var localError error
	var argVal string

	if matched, _ := regexp.MatchString("^--[a-zA-Z0-9_-]+=", argAndVal); matched {
		argVal = strings.Split(argAndVal, a + "=")[1]
	} else if len(os.Args) > i {
		argVal = os.Args[i+1]
	} else {
		localError = fmt.Errorf("%v requires a value.", a)
	}

	return argVal, localError
}
