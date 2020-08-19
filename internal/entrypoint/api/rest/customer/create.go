package customer

import (
	"encoding/json"
	"net/http"

	responses "github.com/darkcrux/mux-responses"
	log "github.com/sirupsen/logrus"

	"github.com/darkcrux/webhook-manager/internal/component/customer"
)

func createCustomer(service customer.Service) http.HandlerFunc {

	log.Info("handler registered: customers::create")

	return func(res http.ResponseWriter, req *http.Request) {

		log.Info("create customer started")

		var request customer.Customer
		if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
			log.WithError(err).Error("unable to read create customer request")
			responses.WriteUnreadableRequestError(res)
			return
		}

		id, err := service.Register(&request)
		if err != nil {
			log.WithError(err).Error("unable to register account")
			responses.WriteGatewayTimeout(res)
			return
		}

		response := map[string]interface{}{
			"customer-id": id,
		}
		responses.WriteOKWithEntity(res, response)

		log.Info("create customer success")
	}
}
