package webhook

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

type DefaultService struct {
	repo Repository
}

func NewDefaultService(repo Repository) Service {
	return &DefaultService{repo}
}

func (s *DefaultService) Create(wh *Webhook) (id int, err error) {
	log.Info("Creating new Webhook...")
	id, err = s.repo.Save(wh)
	if err != nil {
		log.WithError(err).Error("unable to create new webhook")
		return
	}
	log.Info("Creating new Webhook success")
	return
}

func (s *DefaultService) Update(customerID, webhookID int, url string) (id int, err error) {
	log.Info("Updating Webhook details...")
	wh, err := s.Get(webhookID)
	if err != nil {
		log.WithError(err).Error("unable to retrieve existing webhook")
		return
	}
	if wh.CustomerID != customerID {
		err = errors.New("wrong customer")
		log.WithError(err).Error("Customer ID does not match webhook's customer ID")
		return
	}
	wh.WebhookURL = url
	id, err = s.repo.Save(wh)
	if err != nil {
		log.WithError(err).Error("unable to save webhook changes")
		return
	}
	log.Info("Updating Webhook details success")
	return
}

func (s *DefaultService) Get(id int) (wh *Webhook, err error) {
	log.Info("Getting Webhook...")
	wh, err = s.repo.GetByID(id)
	if err != nil {
		log.WithError(err).Error("unable to get webhook")
		return
	}
	log.Info("Getting Webhook success")
	return
}

func (s *DefaultService) List(id int) (whs []Webhook, err error) {
	log.Info("Getting list of webhooks...")
	whs, err = s.repo.List(id)
	if err != nil {
		log.WithError(err).Error("unable to get list of webhooks")
		return
	}
	log.Info("Getting list of webhooks success")
	return
}

func (s *DefaultService) GetByTxAndCust(txId, custID int) (wh *Webhook, err error) {
	log.Info("Getting Webhook by tx and cust ID...")
	wh, err = s.repo.GetByTxIDAndCustomerID(txId, custID)
	if err != nil {
		log.WithError(err).Error("unable to get webhook by tx and customer")
		return
	}
	log.Info("Getting Webhook by tx and cust ID success")
	return
}
