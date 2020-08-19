package webhook

import "errors"

type DefaultService struct {
	repo Repository
}

func NewDefaultService(repo Repository) Service {
	return &DefaultService{repo}
}

func (s *DefaultService) Create(wh *Webhook) (id int, err error) {
	return s.repo.Save(wh)
}

func (s *DefaultService) Update(customerID, webhookID int, url string) (id int, err error) {
	wh, err := s.repo.GetByID(webhookID)
	if err != nil {
		// log
		return
	}
	if wh.CustomerID != customerID {
		// wrong customer, log
		err = errors.New("wrong customer")
		return
	}
	wh.WebhookURL = url
	return s.repo.Save(wh)
}

func (s *DefaultService) Test(wh *Webhook) (id int, err error) {
	// get url from db (webhook)
	// get sample payload from db (tx)
	// oooooor... set up the repo to get all details. (so it's just one call for now.)
	// call notif service
	return
}
