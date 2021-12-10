package main

import (
	"os"

	"github.com/NicoNex/echotron/v3"
)

// DownloadFileToDisk it download a given echotron.Document at the given path + file name and returns the content
func DownloadFileToDisk(a echotron.API, document echotron.Document, path string) (content []byte, err error) {
	content, err = DownloadFile(a, document.FileID)
	if err != nil {
		return
	}

	err = os.WriteFile(path, content, os.ModePerm)

	return
}

// DownloadFile get the content of a given echotron.Document and returns the content
func DownloadFile(a echotron.API, fileID string) (content []byte, err error) {
	var res echotron.APIResponseFile

	res, err = a.GetFile(fileID)
	if err != nil {
		return
	}

	return a.DownloadFile(res.Result.FilePath)
}
