package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/darkcrux/webhook-manager/internal/component/txtypes"
)

type GormTxTypeRepository struct {
	db *gorm.DB
}

func NewGormTxTypeRepository(db *gorm.DB) txtypes.Repository {
	return &GormTxTypeRepository{db}
}

func (repo *GormTxTypeRepository) Save(tx *txtypes.TransactionType) (id int, err error) {
	if err = repo.db.Save(tx).Error; err != nil {
		// what went wrong?
		return
	}
	id = *tx.ID
	return
}

func (repo *GormTxTypeRepository) List() (txTypes []txtypes.TransactionType, err error) {
	txTypes = []txtypes.TransactionType{}
	err = repo.db.Find(&txTypes).Error
	if err != nil {
		// what went wrong?
		return
	}
	return
}

func (repo *GormTxTypeRepository) Get(id int) (t *txtypes.TransactionType, err error) {
	t = &txtypes.TransactionType{}
	if err = repo.db.Find(&t, "id = ?", id).Error; err != nil {
		// log
		return
	}
	return
}
