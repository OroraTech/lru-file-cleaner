package adapters

import "time"

type File struct {
	FilePath     string
	Size         int64
	LastAccessed time.Time
	IsDir        bool
}

type Adapter interface {
	ListContents() ([]File, error)
	DeleteFiles([]File) error
}
