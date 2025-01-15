package go_cache

type (
	CacheInterface interface {
		Get(key string) ([]byte, error)
		Set(key string, value []byte) error
	}
)
