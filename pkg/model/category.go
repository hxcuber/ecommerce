package model

import "github.com/hxcuber/ecommerce/internal/repository/orm"

type Category string

const (
	CategoryDefault     Category = "default"
	CategoryInteresting Category = "interesting"
	CategoryTesting     Category = "testing"
)

func FromOrmCategory(category orm.Category) Category {
	switch category {
	case orm.CategoryDefault:
		return CategoryDefault
	case orm.CategoryInteresting:
		return CategoryInteresting
	case orm.CategoryTesting:
		return CategoryTesting
	}
	return CategoryDefault
}

func (category Category) ToOrmCategory() orm.Category {
	switch category {
	case CategoryDefault:
		return orm.CategoryDefault
	case CategoryInteresting:
		return orm.CategoryInteresting
	case CategoryTesting:
		return orm.CategoryTesting
	}
	return orm.CategoryDefault
}
