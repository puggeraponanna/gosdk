package sqlite

import (
	"context"
	"gosdk/store"

	"gorm.io/gorm"
)

type SqliteStore[T store.Storable] struct {
	db *gorm.DB
}

func NewSqliteStore[T store.Storable](db *gorm.DB) (store.Store[T], error) {
	var t T
	err := db.Table(t.StoreName()).AutoMigrate(&t)
	if err != nil {
		return nil, err
	}
	return &SqliteStore[T]{
		db: db,
	}, nil
}

func (str *SqliteStore[T]) FindOne(ctx context.Context, filter map[string]any) (T, error) {
	var t T
	err := str.db.Where(filter).First(&t).Error
	if err != nil {
		return t, err
	}
	return t, nil

}
