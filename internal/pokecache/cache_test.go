package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		preAdd bool
		key    string
		val    []byte
	}{
		{
			preAdd: true,
			key:    "https://example.com",
			val:    []byte("testdata"),
		},
		{
			preAdd: true,
			key:    "https://example.com/path",
			val:    []byte("moretestdata"),
		},
		{
			preAdd: false,
			key:    "https://foobar.com",
			val:    []byte("anything"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			if c.preAdd {
				cache.Add(c.key, c.val)
			}
			val, ok := cache.Get(c.key)
			if c.preAdd {
				if !ok {
					t.Errorf("expected to find key")
					return
				}
				if string(val) != string(c.val) {
					t.Errorf("expected to find value")
					return
				}
			} else {
				if ok {
					t.Errorf("expected not to find key")
				}
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
