package txtypes

import (
	"encoding/json"
)

type DefaultService struct {
	repo Repository
}

func NewDefaultService(repo Repository) Service {
	return &DefaultService{repo}
}

func (s *DefaultService) Register(tx *TransactionType) (id int, err error) {
	payload, err := payloadToString(tx.SamplePayload)
	if err != nil {
		// log
		return
	}
	tx.SamplePayload = payload
	return s.repo.Save(tx)
}

func (s *DefaultService) List() (txTypes []TransactionType, err error) {
	txTypes, err = s.repo.List()
	if err != nil {
		// log
		return
	}
	for i, txtype := range txTypes {
		payloadStr := string(txtype.SamplePayload.(string))
		txTypes[i].SamplePayload, err = stringToPayload(payloadStr)
		if err != nil {
			// log
			continue
		}
	}
	return
}

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
		// log
		return
	}
	data = d
	return
}
