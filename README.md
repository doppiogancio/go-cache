# Go Cache Package

A lightweight and flexible caching package for Go applications that provides both in-memory and filesystem-based caching implementations.

## Features

- Simple and clean interface
- Multiple cache implementations:
    - In-memory cache for fast access
    - Filesystem cache for persistence
- Thread-safe operations
- Easy to extend with new implementations

## Installation

```bash
go get github.com/doppiogancio/go-cache
```

## Usage

### Basic Example

```go
package main

import (
	go_cache "github.com/doppiogancio/go-cache"
)

func main() {
	cache := go_cache.NewInMemoryCache("./cache")
	cache.Set("config:1", []byte("configuration data"))
	data, err := cache.Get("config:1")
}

```

### Using Filesystem Cache

```go
// Create a filesystem cache
cache := go_cache.NewFileSystemCache("/path/to/cache/directory")

// Usage is identical to in-memory cache
err := cache.Set("users.txt", []byte("a text"))
data, err := cache.Get("users.txt")
```

## Interface

The package provides a simple `Cache` interface that all implementations must satisfy:

```go
type Cache interface {
    Get(key string) ([]byte, error)
    Set(key string, value []byte) error
}
```

### Available Implementations

1. **InMemoryCache**: Stores data in memory using a thread-safe map
2. **FileSystemCache**: Persists data to the filesystem with automatic directory management


## License

This project is licensed under the MIT License - see the LICENSE file for details.