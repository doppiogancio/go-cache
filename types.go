package go_cache

type (
	Cache interface {
		Get(key string) ([]byte, error)
		Set(key string, value []byte) error
	}
)
