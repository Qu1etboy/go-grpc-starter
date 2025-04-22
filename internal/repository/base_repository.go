package repository

import "gorm.io/gorm"

type Repository[T any] interface {
	Save(entity *T) error
	FindById(id uint64) (*T, error)
}

type BaseRepository[T any] struct {
	DB *gorm.DB
}

func (r BaseRepository[T]) Save(entity *T) error {
	return r.DB.Save(entity).Error
}

func (r BaseRepository[T]) FindById(id uint64) (*T, error) {
	var entity T
	err := r.DB.First(&entity, id).Error
	return &entity, err
}
