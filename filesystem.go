package go_cache

import "log"

import (
	"fmt"
	"os"
)

type (
	FileSystemCache struct {
		path string
	}
)

func NewFileSystemCache(path string) CacheInterface {
	err := os.MkdirAll(path, os.ModePerm) // os.ModePerm sets permissions to 0777
	if err != nil {
		log.Fatalf("Error creating folder: %v\n", err)
	}

	fmt.Printf("Folder created or already exists: %s\n", path)

	return &FileSystemCache{path: path}
}

func (c *FileSystemCache) Get(key string) ([]byte, error) {
	return os.ReadFile(c.path + "/" + key)
}

func (c *FileSystemCache) Set(key string, value []byte) error {
	return os.WriteFile(c.path+"/"+key, value, 0644)
}
