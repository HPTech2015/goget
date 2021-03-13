package core

import (
	"net/http"
	"io"
	"os"
)

/*
	Struct that contains methods to pull remote files
	from a web server to the local file system.
*/
type GetFile struct {
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