package store

import "context"

type Storable interface {
	StoreName() string
}

type Cacheable interface {
	CacheKey() string
}

type StorableCacheable interface {
	Storable
	Cacheable
}

type Store[T Storable] interface {
	FindOne(ctx context.Context, filter map[string]any) (T, error)
}
