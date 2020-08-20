package repository

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/darkcrux/webhook-manager/internal/component/webhook"
)

type GormWebhookRepository struct {
	db *gorm.DB
}

func NewGormWebhookRepository(db *gorm.DB) webhook.Repository {
	return &GormWebhookRepository{db}
}

func (repo *GormWebhookRepository) Save(tx *webhook.Webhook) (id int, err error) {
	if err = repo.db.Save(tx).Error; err != nil {
		log.WithError(err).Error("unable to save new webhook")
		return
	}
	id = *tx.ID
	return
}

func (repo *GormWebhookRepository) GetByID(id int) (wh *webhook.Webhook, err error) {
	wh = &webhook.Webhook{}
	err = repo.db.Find(wh, "id = ?", id).Error
	if err != nil {
		log.WithError(err).Error("unable to get webhook")
	}
	return
}

func (repo *GormWebhookRepository) List(id int) (whs []webhook.Webhook, err error) {
	whs = []webhook.Webhook{}
	if err = repo.db.Find(&whs, "customer_id = ?", id).Error; err != nil {
		log.WithError(err).Error("unable to get a list of webhooks for customer")
	}
	return
}

func (repo *GormWebhookRepository) GetByTxIDAndCustomerID(txID, customerID int) (wh *webhook.Webhook, err error) {
	wh = &webhook.Webhook{}
	err = repo.db.Find(wh, "transaction_type_id = ? AND customer_id = ?", txID, customerID).Error
	if err != nil {
		log.WithError(err).Error("unable to get webhook for customer with transaction type")
	}
	return
}
