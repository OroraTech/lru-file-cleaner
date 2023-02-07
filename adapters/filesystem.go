package adapters

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type FilesystemBackend struct {
	Dirpath string
}

func (b *FilesystemBackend) ListContents() (files []File, err error) {
	fileInfos, err := ioutil.ReadDir(b.Dirpath)
	if err != nil {
		return
	}

	for _, fileInfo := range fileInfos {
		filepath := filepath.Join(b.Dirpath, fileInfo.Name())

		if fileInfo.IsDir() {
			var size int64
			size, err = dirSize(filepath)
			if err != nil {
				files = []File{}
				return
			}

			files = append(files, File{filepath, size, fileInfo.ModTime(), true})

		} else {
			files = append(files, File{filepath, fileInfo.Size(), fileInfo.ModTime(), false})
		}
	}

	return
}

func dirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
