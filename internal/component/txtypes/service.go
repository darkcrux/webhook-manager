package txtypes

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type DefaultService struct {
	repo Repository
}

func NewDefaultService(repo Repository) Service {
	return &DefaultService{repo}
}

func (s *DefaultService) Register(tx *TransactionType) (id int, err error) {
	log.Info("Registering new Transaction Type...")
	payload, err := payloadToString(tx.SamplePayload)
	if err != nil {
		log.WithError(err).Error("Unable to marshal Sample Payload")
		return
	}
	tx.SamplePayload = payload
	id, err = s.repo.Save(tx)
	if err != nil {
		log.WithError(err).Error("Unable to save new transaction type")
		return
	}
	log.Info("Registering new Transaction Type success")
	return
}

func (s *DefaultService) List() (txTypes []TransactionType, err error) {
	log.Info("Getting List of Transaction Types...")
	txTypes, err = s.repo.List()
	if err != nil {
		log.WithError(err).Error("Unable to get a list of transaction types")
		return
	}
	for i, txtype := range txTypes {
		payloadStr := txtype.SamplePayload.(string)
		txTypes[i].SamplePayload, err = stringToPayload(payloadStr)
		if err != nil {
			log.WithError(err).Error("Unable to Unmarshal sample payload, skipping")
			continue
		}
	}
	return
}

func (s *DefaultService) Get(id int) (t *TransactionType, err error) {
	log.Info("Getting Transaction Type...")
	t, err = s.repo.Get(id)
	if err != nil {
		log.WithError(err).Error("Unable to get transaction type")
		return
	}
	t.SamplePayload, err = stringToPayload(t.SamplePayload.(string))
	if err != nil {
		log.WithError(err).Error("unable to unmarshal sample payload")
	}
	return
}

// TODO: should probably move this to a UTIL

func payloadToString(p interface{}) (data string, err error) {
	out, err := json.Marshal(p)
	if err != nil {
		return
	}
	data = string(out)
	return
}

func stringToPayload(str string) (data interface{}, err error) {
	var d interface{}
	if err = json.Unmarshal([]byte(str), &d); err != nil {
		log.WithError(err).Error("")
		return
	}
	data = d
	return
}
