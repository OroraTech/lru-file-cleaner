package cache

import (
	"os"
	"sort"

	"github.com/alexw9988/lru-file-cleaner/adapters"
)

type FileCache struct {
	Files []adapters.File
}

func (c *FileCache) DeleteOldest(quota int64) (err error) {
	c.sortDescending()

	var quotaUsed int64 = 0
	deleteableFiles := []adapters.File{}

	for _, file := range c.Files {
		quotaUsed = quotaUsed + file.Size
		if quotaUsed > quota {
			deleteableFiles = append(deleteableFiles, file)
		}
	}

	for _, file := range deleteableFiles {
		err = os.RemoveAll(file.FilePath)
		if err != nil {
			return
		}
	}

	return
}

func (c *FileCache) sortDescending() {
	sort.SliceStable(c.Files, func(i int, j int) bool { return c.Files[i].LastAccessed.After(c.Files[j].LastAccessed) })
}
