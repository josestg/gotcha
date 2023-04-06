package cache

import (
	"errors"
	"time"
)

var (
	// ErrMissed ...
	ErrMissed = errors.New("Cache item's missing")
)

const (
	// Byte ...
	Byte uint64 = 1
	// KB ...
	KB = Byte * 1024
	// MB ...
	MB = KB * 1024
	// LRUAlgorithm ...
	LRUAlgorithm = "lru"
	// LFUAlgorithm ...
	LFUAlgorithm = "lfu"
	// DefaultSize ..
	DefaultSize = 100
	// DefaultExpiryTime ...
	DefaultExpiryTime = time.Second * 10
	// DefaultAlgorithm ...
	DefaultAlgorithm = LRUAlgorithm
	// DefaultMaxMemory ...
	DefaultMaxMemory = 10 * MB
)

// Document represent the Document structure stored in the cache
type Document struct {
	Key        string
	Value      interface{}
	StoredTime int64 // timestamp
}

// Option used for Cache configuration
type Option struct {
	AlgorithmType string        // represent the algorithm type
	ExpiryTime    time.Duration // represent the expiry time of each stored item
	MaxSizeItem   uint64        // Max size of item for eviction
	MaxMemory     uint64        // Max Memory of item stored for eviction
}

// OptionBuilder is a function type for building an Option by composing smaller option.
type OptionBuilder func(opt Option) Option

// NewOption creates a new OptionBuilder.
func NewOption() OptionBuilder {
	// returning an identity function.
	return func(opt Option) Option { return opt }
}

// Build builds the Option.
func (f OptionBuilder) Build() Option { return f(Option{}) }

// SetAlgorithm will set the algorithm value.
func (f OptionBuilder) SetAlgorithm(algorithm string) OptionBuilder {
	return func(opt Option) Option {
		opt = f(opt)
		opt.AlgorithmType = algorithm
		return opt
	}
}

// SetExpiryTime will set the expiry time.
func (f OptionBuilder) SetExpiryTime(expiry time.Duration) OptionBuilder {
	return func(opt Option) Option {
		opt = f(opt)
		opt.ExpiryTime = expiry
		return opt
	}
}

// SetMaxSizeItem will set the maximum size of item in cache.
func (f OptionBuilder) SetMaxSizeItem(size uint64) OptionBuilder {
	return func(opt Option) Option {
		opt = f(opt)
		opt.MaxSizeItem = size
		return opt
	}
}

// SetMaxMemory will set the maximum memory will be used for cache.
func (f OptionBuilder) SetMaxMemory(memory uint64) OptionBuilder {
	return func(opt Option) Option {
		opt = f(opt)
		opt.MaxMemory = memory
		return opt
	}
}

// Cache represent the public API that will available used by user
type Cache interface {
	Set(key string, value interface{}) error
	Get(key string) (val interface{}, err error)
	Delete(key string) (err error)
	GetKeys() (keys []string, err error)
	ClearCache() (err error)
}
