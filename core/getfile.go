package core

import (
	"net/http"
	"io"
	"os"
	"crypto/sha256"
	"io/ioutil"
	"encoding/hex"
	"strings"
	"errors"
	"strconv"
)

/*
	Struct that contains methods to pull remote files
	from a web server to the local file system.
*/
type GetFile struct {
	LocalChecksum string
	RemoteChecksum string
}

/*
	Pull the file from the remote server and
	validate the checksum.
*/
func (getFile *GetFile) PullAndCheck(filePath, url string) (bool, error) {
	if err := getFile.Pull(filePath, url); err != nil {
		return false, err
	}

	checksumPath := filePath + ".sha256"
	if err := getFile.Pull(checksumPath, url + ".sha256"); err != nil {
		return false, err
	}

	if err := getFile.GetChecksum(filePath); err != nil {
		return false, err
	}

	remoteChecksum, err := ioutil.ReadFile(checksumPath)
	if err != nil {
		return false, err
	}
	getFile.RemoteChecksum = strings.Replace(string(remoteChecksum), "\n", "", -1)

	return getFile.LocalChecksum == getFile.RemoteChecksum, nil
}

/*
	Pull the remote file form the web server.
*/
func (getFile *GetFile) Pull(filePath, url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New("HTTP ERROR CODE: " + strconv.Itoa(res.StatusCode))
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = io.Copy(file, res.Body); err != nil {
		return err
	}

	return nil
}

/*
	Generate and return the SHA256 checksum.
*/
func (getFile *GetFile) GetChecksum(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err = io.Copy(hasher, file); err != nil {
		return err
	}

	getFile.LocalChecksum = hex.EncodeToString(hasher.Sum(nil))
	
	return nil
}