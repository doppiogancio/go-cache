package go_cache

import (
	"testing"
)

func TestFileSystemCache(t *testing.T) {
	cache := NewFileSystemCache("./cache")
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
	if err.Error() != "open ./cache/DAY: no such file or directory" {
		t.Error("unexpected error message")
	}
}
