package crawler

import (
	"archive/zip"
	"fmt"
)

type ZipReader struct {
	path string
}

func NewZipReader(path string) *ZipReader {
	return &ZipReader{path}
}

func (z *ZipReader) Files(zipName string) []*zip.File {
	archive, err := zip.OpenReader(fmt.Sprintf("%s/%s", z.path, zipName))
	if err != nil {
		panic(err)
	}
	return archive.File
}
