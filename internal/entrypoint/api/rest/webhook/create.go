package webhook

import (
	"encoding/json"
	"net/http"
	"strconv"

	responses "github.com/darkcrux/mux-responses"
	"github.com/darkcrux/webhook-manager/internal/component/webhook"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func createWebhook(service webhook.Service) http.HandlerFunc {

	log.Info("handler registered: webhooks::create")

	return func(res http.ResponseWriter, req *http.Request) {

		log.Info("create Webhook started")

		customerIDStr := mux.Vars(req)["id"]
		customerID, err := strconv.Atoi(customerIDStr)
		if err != nil {
			log.WithError(err).Error("customer ID should be a number")
			responses.WriteUnreadableRequestError(res)
			return
		}

		var request webhook.Webhook
		if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
			log.WithError(err).Error("unable to read create Webhook request")
			responses.WriteUnreadableRequestError(res)
			return
		}

		request.CustomerID = customerID

		id, err := service.Create(&request)
		if err != nil {
			log.WithError(err).Error("unable to create webhook")
			responses.WriteGatewayTimeout(res)
			return
		}

		response := map[string]interface{}{
			"webhook-id": id,
		}
		responses.WriteOKWithEntity(res, response)

		log.Info("create webhook success")
	}
}
