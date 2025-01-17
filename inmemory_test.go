package go_cache

import (
	"testing"
)

func TestInMemoryCache(t *testing.T) {
	cache := NewInMemoryCache()
	err := cache.Set("YEAR", []byte("2024"))
	if err != nil {
		t.Error(err)
	}

	val, err := cache.Get("YEAR")
	if err != nil {
		t.Error(err)
	}

	if string(val) != "2024" {
		t.Error("value is not equal to 2024")
	}

	val, err = cache.Get("DAY")
	if err.Error() != "key not found" {
		t.Error("unexpected error message")
	}
}
