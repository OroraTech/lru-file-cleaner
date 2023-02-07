package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alexw9988/lru-file-cleaner/adapters"
	"github.com/alexw9988/lru-file-cleaner/cache"
)

func main() {
	dirpath := flag.String("dirpath", "", "target directory")
	quota := flag.Int64("quota", -1, "size quota in bytes")
	flag.Parse()

	if len(*dirpath) == 0 {
		fmt.Println("error: no dirpath was provided")
		os.Exit(1)
	}

	if *quota == -1 {
		fmt.Println("error: no directory size quota was specified!")
		os.Exit(1)
	}

	adapter := adapters.FilesystemBackend{Dirpath: *dirpath}
	files, _ := adapter.ListContents()
	cache := cache.FileCache{Files: files}
	cache.DeleteOldest(*quota)
}
