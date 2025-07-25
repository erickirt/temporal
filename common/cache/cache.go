package cache

import (
	"time"

	"go.temporal.io/server/common/clock"
	"go.temporal.io/server/common/dynamicconfig"
)

// A Cache is a generalized interface to a cache.  See cache.LRU for a specific
// implementation (bounded cache with LRU eviction)
type Cache interface {
	// Get retrieves an element based on a key, returning nil if the element
	// does not exist
	Get(key interface{}) interface{}

	// Put adds an element to the cache, returning the previous element
	Put(key interface{}, value interface{}) interface{}

	// PutIfNotExist puts a value associated with a given key if it does not exist
	PutIfNotExist(key interface{}, value interface{}) (interface{}, error)

	// Delete deletes an element in the cache
	Delete(key interface{})

	// Release decrements the ref count of a pinned element. If the ref count
	// drops to 0, the element can be evicted from the cache.
	Release(key interface{})

	// Iterator returns the iterator of the cache
	Iterator() Iterator

	// Size returns current size of the Cache, the size definition is implementation of SizeGetter interface
	// for the entry size, if the entry does not implement SizeGetter interface, the size is 1
	Size() int
}

type StoppableCache interface {
	Cache

	// Stop halts any background processing, and should be called when the cache will no longer be used.
	Stop()
}

// Options control the behavior of the cache.
type Options struct {
	// TTL controls the time-to-live for a given cache entry.  Cache entries that
	// are older than the TTL will not be returned.
	TTL time.Duration

	// Pin prevents in-use objects from getting evicted.
	Pin bool

	// TimeSource is an optional clock to use for time-skipping and testing. If this is nil, a real clock will be used.
	TimeSource clock.TimeSource

	OnPut func(val any)

	OnEvict func(val any)

	// BackgroundEvict configures background scanning for expired entries.
	BackgroundEvict func() dynamicconfig.CacheBackgroundEvictSettings
}

// SimpleOptions provides options that can be used to configure SimpleCache.
type SimpleOptions struct {
	// RemovedFunc is an optional function called when an element
	// is scheduled for deletion
	RemovedFunc RemovedFunc
}

// RemovedFunc is a type for notifying applications when an item is
// scheduled for removal from the Cache. If f is a function with the
// appropriate signature and i is the interface{} scheduled for
// deletion, Cache calls go f(i)
type RemovedFunc func(interface{})

// Iterator represents the interface for cache iterators.
type Iterator interface {
	// Close closes the iterator
	// and releases any allocated resources
	Close()
	// HasNext return true if there is more items to be returned
	HasNext() bool
	// Next return the next item
	Next() Entry
}

// Entry represents a key-value entry within the map.
type Entry interface {
	// Key represents the key
	Key() interface{}
	// Value represents the value
	Value() interface{}
	// CreateTime represents the time when the entry is created
	CreateTime() time.Time
}
