package core

import (
	"errors"
)

/*
	A struct to store settings accessible
	to the entire application.
*/
type Settings struct {
	RemoteTarget string
	LocalTarget string
	Version string
}

/*
	Set the version of this release.
*/
func (settings *Settings) SetVersion(version string) bool {
	settings.Version = version

	return true
}

/*
	Get the version of this release.
*/
func (settings *Settings) GetVersion() (string, error) {
	var localError error

	if settings.Version == "" {
		localError = errors.New("Version was not set.")
	}

	return settings.Version, localError
}

/*
	Set the remote target file download and checksum.
*/
func (settings *Settings) SetRemoteTarget(remoteTarget string) error {
	var localError error

	if remoteTarget == "" {
		localError = errors.New("Remote target cannot be empty!")
	}

	settings.RemoteTarget = remoteTarget

	return localError
}

/*
	Get the remote target file download and checksum.
*/
func (settings *Settings) GetRemoteTarget() (string, error) {
	var localError error

	if settings.RemoteTarget == "" {
		localError = errors.New("The remote target has not yet been set!")
	}

	return settings.RemoteTarget, localError
}

/*
	Set the local path, to download the remote file to.
*/
func (settings *Settings) SetLocalTarget(localTarget string) error {
	var localError error

	if localTarget == "" {
		localError = errors.New("Local target cannot be empty!")
	}

	settings.LocalTarget = localTarget

	return localError
}

/*
	Get the local path, to download the remote file to.
*/
func (settings *Settings) GetLocalTarget() (string, error) {
	var localError error

	if settings.LocalTarget == "" {
		localError = errors.New("The local target has not yet been set!")
	}

	return settings.LocalTarget, localError
}