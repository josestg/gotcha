package gotcha_test

import (
	"github.com/bxcodec/gotcha"
	"github.com/bxcodec/gotcha/cache"
	"testing"
)

func TestGotcha(t *testing.T) {
	t.Run("lru", GotchaTestFactory(gotcha.New()))
	t.Run("lfu", GotchaTestFactory(gotcha.NewWithOption(gotcha.DefaultOption().SetAlgorithm(cache.LFUAlgorithm))))
	t.Run("default client", func(t *testing.T) {
		t.Run("set", func(t *testing.T) {
			err := gotcha.Set("name", "John Snow")
			if err != nil {
				t.Fatalf("expected: %v, got %v", nil, err)
			}
			err = gotcha.Set("kingdom", "North Kingdom")
			if err != nil {
				t.Fatalf("expected: %v, got %v", nil, err)
			}
		})

		t.Run("get", func(t *testing.T) {
			val, err := gotcha.Get("name")
			if err != nil {
				t.Fatalf("expected: %v, got %v", nil, err)
			}
			if val.(string) != "John Snow" {
				t.Fatalf("expected: %v, got %v", "John Snow", val)
			}
		})

		t.Run("get-keys", func(t *testing.T) {
			keys, err := gotcha.GetKeys()
			if err != nil {
				t.Fatalf("expected: %v, got %v", nil, err)
			}
			var contains = func(keys []string, k string) bool {
				for _, item := range keys {
					if item == k {
						return true
					}
				}
				return false
			}
			expectedKeys := []string{"name", "kingdom"}
			for _, k := range expectedKeys {
				if !contains(keys, k) {
					t.Fatalf("expected: %v, got: %v", true, false)
				}
			}
		})

		t.Run("delete", func(t *testing.T) {
			// Ensure the key is exists
			val, err := gotcha.Get("kingdom")
			if err != nil {
				t.Fatalf("expected: %v, got %v", nil, err)
			}
			if val.(string) != "North Kingdom" {
				t.Fatalf("expected: %v, got %v", "John Snow", val)
			}

			// Delete the Keys

			err = gotcha.Delete("kingdom")
			if err != nil {
				t.Fatalf("expected: %v, got %v", nil, err)
			}

			// Re-Ensure the keys is deleted
			val, err = gotcha.Get("kingdom")
			if err == nil {
				t.Fatalf("expected: %v, got %v", "error", err)
			}

			if val != nil {
				t.Fatalf("expected: %v, got %v", nil, val)
			}
		})

		t.Run("clear-cache", func(t *testing.T) {
			// Ensure the cache is still contains item
			keys, err := gotcha.GetKeys()
			if err != nil {
				t.Fatalf("expected: %v, got %v", nil, err)
			}
			if len(keys) == 0 {
				t.Fatalf("expected: %v, got %v", "not zero", len(keys))
			}

			err = gotcha.ClearCache()
			if err != nil {
				t.Fatalf("expected: %v, got %v", nil, err)
			}

			// Re-Ensure the cache already cleared
			keys, err = gotcha.GetKeys()
			if err != nil {
				t.Fatalf("expected: %v, got %v", nil, err)
			}
			if len(keys) != 0 {
				t.Fatalf("expected: %v, got %v", "zero", len(keys))
			}
		})
	})
}

func GotchaTestFactory(cacheImpl cache.Cache) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("set", func(t *testing.T) {
			err := cacheImpl.Set("name", "John Snow")
			if err != nil {
				t.Fatalf("expected: %v, got %v", nil, err)
			}
			err = cacheImpl.Set("kingdom", "North Kingdom")
			if err != nil {
				t.Fatalf("expected: %v, got %v", nil, err)
			}
		})

		t.Run("get", func(t *testing.T) {
			val, err := cacheImpl.Get("name")
			if err != nil {
				t.Fatalf("expected: %v, got %v", nil, err)
			}
			if val.(string) != "John Snow" {
				t.Fatalf("expected: %v, got %v", "John Snow", val)
			}
		})

		t.Run("get-keys", func(t *testing.T) {
			keys, err := cacheImpl.GetKeys()
			if err != nil {
				t.Fatalf("expected: %v, got %v", nil, err)
			}
			var contains = func(keys []string, k string) bool {
				for _, item := range keys {
					if item == k {
						return true
					}
				}
				return false
			}
			expectedKeys := []string{"name", "kingdom"}
			for _, k := range expectedKeys {
				if !contains(keys, k) {
					t.Fatalf("expected: %v, got: %v", true, false)
				}
			}
		})

		t.Run("delete", func(t *testing.T) {
			// Ensure the key is exists
			val, err := cacheImpl.Get("kingdom")
			if err != nil {
				t.Fatalf("expected: %v, got %v", nil, err)
			}
			if val.(string) != "North Kingdom" {
				t.Fatalf("expected: %v, got %v", "John Snow", val)
			}

			// Delete the Keys

			err = cacheImpl.Delete("kingdom")
			if err != nil {
				t.Fatalf("expected: %v, got %v", nil, err)
			}

			// Re-Ensure the keys is deleted
			val, err = cacheImpl.Get("kingdom")
			if err == nil {
				t.Fatalf("expected: %v, got %v", "error", err)
			}

			if val != nil {
				t.Fatalf("expected: %v, got %v", nil, val)
			}
		})

		t.Run("clear-cache", func(t *testing.T) {
			// Ensure the cache is still contains item
			keys, err := cacheImpl.GetKeys()
			if err != nil {
				t.Fatalf("expected: %v, got %v", nil, err)
			}
			if len(keys) == 0 {
				t.Fatalf("expected: %v, got %v", "not zero", len(keys))
			}

			err = cacheImpl.ClearCache()
			if err != nil {
				t.Fatalf("expected: %v, got %v", nil, err)
			}

			// Re-Ensure the cache already cleared
			keys, err = cacheImpl.GetKeys()
			if err != nil {
				t.Fatalf("expected: %v, got %v", nil, err)
			}
			if len(keys) != 0 {
				t.Fatalf("expected: %v, got %v", "zero", len(keys))
			}
		})
	}
}
