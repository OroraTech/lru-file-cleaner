package cache

import (
	"fmt"
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
	startIndex := -1

	for index, file := range c.Files {
		quotaUsed = quotaUsed + file.Size
		if quotaUsed > quota {
			startIndex = index
			break
		}
	}

	if startIndex == -1 {
		return
	}

	deleteableFiles := c.Files[startIndex:]

	for _, file := range deleteableFiles {
		err = os.RemoveAll(file.FilePath)
		if err != nil {
			return
		}
		fmt.Println(file.FilePath)
	}

	return
}

func (c *FileCache) sortDescending() {
	sort.SliceStable(c.Files, func(i int, j int) bool { return c.Files[i].LastAccessed.After(c.Files[j].LastAccessed) })
}
