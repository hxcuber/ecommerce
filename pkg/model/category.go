package model

import "github.com/hxcuber/ecommerce/internal/repository/orm"

type Category struct {
	CategoryID  int64  `boil:"category_id" json:"category_id" toml:"category_id" yaml:"category_id"`
	Description string `boil:"description" json:"description" toml:"description" yaml:"description"`
}

func FromOrmCategory(category *orm.Category) Category {
	return Category{
		CategoryID:  category.CategoryID,
		Description: category.Description,
	}
}

func (category Category) ToOrmCategory() *orm.Category {
	return &orm.Category{
		CategoryID:  category.CategoryID,
		Description: category.Description,
	}
}
