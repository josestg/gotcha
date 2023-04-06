package cache_test

import (
	"github.com/bxcodec/gotcha/cache"
	"testing"
)

func TestNewOption_WithoutOverride(t *testing.T) {
	opt := cache.NewOption().Build()
	if opt.MaxMemory != 0 {
		t.Fatalf("expected %v, actual %v", 0, opt.MaxMemory)
	}

	if opt.AlgorithmType != "" {
		t.Fatalf("expected %v, actual %v", 0, opt.AlgorithmType)
	}

	if opt.MaxSizeItem != 0 {
		t.Fatalf("expected %v, actual %v", 0, opt.MaxSizeItem)
	}

	if opt.ExpiryTime != 0 {
		t.Fatalf("expected %v, actual %v", 0, opt.ExpiryTime)
	}
}

func TestNewOption_WithOverride(t *testing.T) {
	opt := cache.NewOption().
		SetAlgorithm(cache.DefaultAlgorithm).
		SetMaxMemory(cache.DefaultMaxMemory).
		SetMaxSizeItem(cache.DefaultSize).
		SetExpiryTime(cache.DefaultExpiryTime).
		Build()

	if opt.AlgorithmType != cache.DefaultAlgorithm {
		t.Fatalf("expected %v, actual %v", 0, opt.AlgorithmType)
	}

	if opt.ExpiryTime != cache.DefaultExpiryTime {
		t.Fatalf("expected %v, actual %v", 0, opt.ExpiryTime)
	}

	if opt.MaxSizeItem != cache.DefaultSize {
		t.Fatalf("expected %v, actual %v", 0, opt.MaxSizeItem)
	}

	if opt.MaxMemory != cache.DefaultMaxMemory {
		t.Fatalf("expected %v, actual %v", 0, opt.MaxMemory)
	}
}

func TestOptionBuilder_SetAlgorithm(t *testing.T) {
	opt := cache.NewOption().SetAlgorithm(cache.DefaultAlgorithm).Build()
	if opt.AlgorithmType != cache.DefaultAlgorithm {
		t.Fatalf("expected %v, actual %v", 0, opt.AlgorithmType)
	}
}

func TestOptionBuilder_SetExpiryTime(t *testing.T) {
	opt := cache.NewOption().SetExpiryTime(cache.DefaultExpiryTime).Build()
	if opt.ExpiryTime != cache.DefaultExpiryTime {
		t.Fatalf("expected %v, actual %v", 0, opt.ExpiryTime)
	}
}

func TestOptionBuilder_SetMaxSizeItem(t *testing.T) {
	opt := cache.NewOption().SetMaxSizeItem(cache.DefaultSize).Build()
	if opt.MaxSizeItem != cache.DefaultSize {
		t.Fatalf("expected %v, actual %v", 0, opt.MaxSizeItem)
	}
}

func TestOptionBuilder_SetMaxMemory(t *testing.T) {
	opt := cache.NewOption().SetMaxMemory(cache.DefaultMaxMemory).Build()
	if opt.MaxMemory != cache.DefaultMaxMemory {
		t.Fatalf("expected %v, actual %v", 0, opt.MaxMemory)
	}
}
