package database

import (
	"github.com/luizhenrique-dev/go-products-api/internal/entity"
	"gorm.io/gorm"
)

const ASC = "asc"
const DESC = "desc"

type ProductRepository struct {
	DB *gorm.DB	
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (p *ProductRepository) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *ProductRepository) FindById(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (p *ProductRepository) Update(product *entity.Product) error {
	_, err := p.FindById(product.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(product).Error
}

func (p *ProductRepository) Delete(id string) error {
	product, err := p.FindById(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}

func (p *ProductRepository) FindAll(page, limit int, sort string) ([]*entity.Product, error) {
	var products []*entity.Product
	
	if sort != "" && sort != ASC && sort != DESC {
		sort = ASC
	}
	if page <= 0 || limit <= 0 {
		return products, nil
	}

	offset := (page - 1) * limit
	err := p.DB.Offset(offset).Limit(limit).Order("created_at " + sort).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}