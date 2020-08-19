package common

import (
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/darkcrux/webhook-manager/internal/infrastructure/postgres/entity"
)

type GormRepository struct {
	*gorm.DB
}

func (cr *GormRepository) SaveEntity(e entity.Entity, data interface{}) (id interface{}, err error) {
	if err = copier.Copy(e, data); err != nil {
		log.WithError(err).Error("unable to copy data to entity")
		return
	}

	err = cr.Save(e).Error
	id = e.GetID()
	return
}

func (cr *GormRepository) FindEntity(e interface{}, o interface{}, where ...interface{}) error {
	if err := cr.Find(e, where...).Error; err != nil {
		log.WithError(err).Error("error accessing database")
		return err
	}
	if err := copier.Copy(o, e); err != nil {
		log.WithError(err).Error("unable to copy entity to data")
	}
	return nil
}

func (cr *GormRepository) DeleteEntity(e interface{}, where ...interface{}) error {
	return cr.Delete(e, where...).Error
}
