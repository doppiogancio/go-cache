package go_cache

import "errors"

type (
	InMemoryCache struct {
		db map[string][]byte
	}
)

func NewInMemoryCache() Cache {
	return &InMemoryCache{
		db: make(map[string][]byte),
	}
}

func (c *InMemoryCache) Get(key string) ([]byte, error) {
	value, ok := c.db[key]
	if !ok {
		return nil, errors.New("key not found")
	}

	return value, nil
}

func (c *InMemoryCache) Set(key string, value []byte) error {
	c.db[key] = value
	return nil
}
