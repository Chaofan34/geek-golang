package week02

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func FindProductByCode(db *gorm.DB, code string) (*Product, error) {
	var product Product
	err := db.First(&product, "code = ?", code).Error
	if err != nil {
		return nil, errors.Wrap(err, "find product failed")
	}
	return &product, nil
}
