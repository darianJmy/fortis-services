package db

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type Curd[T any] struct {
	EmptyInstance T
	db            *gorm.DB
}

func newCurd[T any](emptyInstance T, db *gorm.DB) *Curd[T] {
	return &Curd[T]{EmptyInstance: emptyInstance, db: db}
}

func (c *Curd[T]) Create(ctx context.Context, object T) (T, error) {
	return object, c.db.WithContext(ctx).Create(object).Error
}

func (c *Curd[T]) Update(ctx context.Context, uid int64, attr map[string]interface{}) (T, error) {
	var (
		ret T
		err error
	)

	if err = c.db.WithContext(ctx).Model(c.EmptyInstance).Where("id = ?", uid).Updates(attr).Error; err != nil {
		return ret, err
	}

	if ret, err = c.Get(ctx, uid); err != nil {
		return ret, err
	}

	return ret, nil
}

func (c *Curd[T]) Delete(ctx context.Context, uid int64) (T, error) {
	var (
		ret T
		err error
	)

	if ret, err = c.Get(ctx, uid); err != nil {
		return ret, err
	}

	return ret, c.db.WithContext(ctx).Where("id = ?", uid).Delete(&ret).Error
}

func (c *Curd[T]) Get(ctx context.Context, uid int64) (T, error) {
	var ret T
	if err := c.db.WithContext(ctx).Where("id = ?", uid).First(&ret).Error; err != nil {
		return ret, err
	}

	return ret, nil
}

func (c *Curd[T]) List(ctx context.Context) ([]T, error) {
	var ret []T
	if err := c.db.WithContext(ctx).Find(&ret).Error; err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *Curd[T]) FindByWhere(ctx context.Context, query interface{}, args ...interface{}) ([]T, error) {
	var ret []T
	if err := c.db.WithContext(ctx).Where(query, args...).Find(&ret).Error; err != nil {
		return nil, err
	}

	if len(ret) == 0 {
		return nil, errors.New("is not found")
	}

	return ret, nil
}

func (c *Curd[T]) FirstByWhere(ctx context.Context, query interface{}, args ...interface{}) (T, error) {
	var ret T
	if err := c.db.WithContext(ctx).Where(query, args...).First(&ret).Error; err != nil {
		return ret, err
	}

	return ret, nil
}
