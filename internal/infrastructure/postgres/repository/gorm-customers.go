package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/darkcrux/webhook-manager/internal/component/customer"
)

type GormCustomerRepository struct {
	db *gorm.DB
}

func NewGormCustomerRepository(db *gorm.DB) customer.Repository {
	return &GormCustomerRepository{db}
}

func (repo *GormCustomerRepository) Save(c *customer.Customer) (id int, err error) {
	if err = repo.db.Save(c).Error; err != nil {
		// what went wrong?
		return
	}
	id = *c.ID
	return
}
