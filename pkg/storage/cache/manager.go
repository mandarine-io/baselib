package cache

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrCacheEntryNotFound = errors.New("cache entry not found")
)

type Manager interface {
	Get(ctx context.Context, key string, value interface{}) error
	Set(ctx context.Context, key string, value interface{}) error
	SetWithExpiration(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Delete(ctx context.Context, keys ...string) error
	Invalidate(ctx context.Context, keyRegex string) error
}
