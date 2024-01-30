package adapters

import (
	"HexAndClean/core"

	"gorm.io/gorm"
)

type GormOderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) *GormOderRepository {
	return &GormOderRepository{db: db}
}

func (r *GormOderRepository) Save(order core.Order) error {
	if result := r.db.Create(&order); result.Error != nil {
		return result.Error
	}
	return nil
}
