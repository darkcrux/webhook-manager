package repository

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/darkcrux/webhook-manager/internal/component/notification"
)

type GormNotificationRepository struct {
	db *gorm.DB
}

func NewGormNotificationRepository(db *gorm.DB) notification.Repository {
	return &GormNotificationRepository{db}
}

func (repo *GormNotificationRepository) Create(notif *notification.Notification) (id int, err error) {
	err = repo.db.Save(notif).Error
	if err != nil {
		log.WithError(err).Error("unable to save new notification info")
		return
	}
	id = *notif.ID
	return
}

func (repo *GormNotificationRepository) Get(id int) (notif *notification.Notification, err error) {
	notif = &notification.Notification{}
	if err = repo.db.Find(notif, "id = ?", id).Error; err != nil {
		log.WithError(err).Error("unable to get notification info")
		return
	}
	return
}

func (repo *GormNotificationRepository) List(customerID int) (notifs []notification.Notification, err error) {
	notifs = []notification.Notification{}
	if err = repo.db.Find(&notifs).Error; err != nil {
		log.WithError(err).Error("unable to list notification infos for customer")
		return
	}
	return
}

func (repo *GormNotificationRepository) UpdateStatus(id int, status string) (err error) {
	notif, err := repo.Get(id)
	if err != nil {
		log.WithError(err).Error("unable to get notif info")
		return
	}
	notif.Status = status
	if err = repo.db.Save(notif).Error; err != nil {
		log.WithError(err).Error("unable to update notif info status")
	}
	return
}
