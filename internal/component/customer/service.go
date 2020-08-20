package customer

import (
	log "github.com/sirupsen/logrus"
)

type DefaultService struct {
	repo Repository
}

func NewDefaultService(repo Repository) Service {
	return &DefaultService{repo}
}

func (s *DefaultService) Register(c *Customer) (id int, err error) {
	log.Info("Saving New Customer...")
	id, err = s.repo.Save(c)
	if err != nil {
		log.WithError(err).Error("Saving New Customer failed")
		return
	}
	log.Info("Saving New Customer success")
	return
}

func (s *DefaultService) Get(id int) (c *Customer, err error) {
	log.Infof("Getting Customer by ID: %d ...", id)
	c, err = s.repo.GetByID(id)
	if err != nil {
		log.WithError(err).Error("Getting New Customer failed")
		return
	}
	log.Info("Getting New Customer success...")
	return
}
