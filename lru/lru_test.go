package lru

import (
	"testing"
)

func TestLRU_BasicPutAndGet(t *testing.T) {
	cache := New[string](2)

	Put(cache, "apple")
	Put(cache, "banana")

	val := Get(cache, "apple")
	if val == nil || *val != "apple" {
		t.Errorf("Expected 'apple', got %v", val)
	}

	val = Get(cache, "banana")
	if val == nil || *val != "banana" {
		t.Errorf("Expected 'banana', got %v", val)
	}
}

func TestLRU_CacheMiss(t *testing.T) {
	cache := New[int](2)
	Put(cache, 1)

	// Note: Your current implementation returns a pointer to the zero value (new(T)) on miss.
	// We are testing against that behavior here, though returning nil is usually preferred.
	val := Get(cache, 2)
	if val != nil && *val != 0 {
		t.Errorf("Expected zero value for cache miss, got %v", *val)
	}
}

func TestLRU_Eviction(t *testing.T) {
	cache := New[int](2)

	Put(cache, 1)
	Put(cache, 2)
	Put(cache, 3) // This should evict 1, as the capacity is 2

	// Check that 1 was removed from the hash map
	if _, exists := cache.hash[1]; exists {
		t.Errorf("Expected 1 to be evicted from the hash map")
	}

	// Check that 2 and 3 are still present
	if _, exists := cache.hash[2]; !exists {
		t.Errorf("Expected 2 to still be in the cache")
	}
	if _, exists := cache.hash[3]; !exists {
		t.Errorf("Expected 3 to still be in the cache")
	}
}

func TestLRU_PutUpdatesRecent(t *testing.T) {
	cache := New[int](2)

	Put(cache, 1)
	Put(cache, 2)
	Put(cache, 1) // 1 is accessed again, making 2 the least recently used
	Put(cache, 3) // This should evict 2, not 1

	if _, exists := cache.hash[2]; exists {
		t.Errorf("Expected 2 to be evicted")
	}

	val := Get(cache, 1)
	if val == nil || *val != 1 {
		t.Errorf("Expected 1 to still be in the cache")
	}
}

func TestLRU_GetUpdatesRecent(t *testing.T) {
	cache := New[int](2)

	Put(cache, 1)
	Put(cache, 2)
	_ = Get(cache, 1) // 1 is read, making 2 the least recently used
	Put(cache, 3)     // This should evict 2

	if _, exists := cache.hash[2]; exists {
		t.Errorf("Expected 2 to be evicted after Get(1) updated LRU order")
	}

	val := Get(cache, 1)
	if val == nil || *val != 1 {
		t.Errorf("Expected 1 to still be in the cache")
	}
}
