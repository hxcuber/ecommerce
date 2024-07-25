package model

import (
	"github.com/google/uuid"
	"github.com/hxcuber/ecommerce/internal/repository/orm"
	"time"
)

type Product struct {
	ProductID   uuid.UUID `boil:"product_id" json:"product_id" toml:"product_id" yaml:"product_id"`
	Name        string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Description string    `boil:"description" json:"description" toml:"description" yaml:"description"`
	Price       float32   `boil:"price" json:"price" toml:"price" yaml:"price"`
	Stock       int64     `boil:"stock" json:"stock" toml:"stock" yaml:"stock"`
	CategoryID  Category  `boil:"category_id" json:"category_id" toml:"category_id" yaml:"category_id"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
}

func FromOrmProduct(product *orm.Product) (Product, error) {
	id, err := uuid.Parse(product.ProductID)
	if err != nil {
		return Product{}, err
	}
	return Product{
		ProductID:   id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CategoryID:  FromOrmCategory(product.Category),
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}

func (p Product) ToOrmProduct() *orm.Product {
	return &orm.Product{
		ProductID:   p.ProductID.String(),
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		Category:    p.CategoryID.ToOrmCategory(),
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
