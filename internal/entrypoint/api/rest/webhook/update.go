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

func updateWebhook(service webhook.Service) http.HandlerFunc {

	log.Info("handler registered: webhooks::update")

	return func(res http.ResponseWriter, req *http.Request) {

		log.Info("update Webhook started")

		log.Info("vars", mux.Vars(req))
		customerIDStr := mux.Vars(req)["customer-id"]
		customerID, err := strconv.Atoi(customerIDStr)
		if err != nil {
			log.WithError(err).Error("customer ID should be a number")
			responses.WriteUnreadableRequestError(res)
			return
		}

		webhookIDStr := mux.Vars(req)["webhook-id"]
		webhookID, err := strconv.Atoi(webhookIDStr)
		if err != nil {
			log.WithError(err).Error("webhook ID should be a number")
			responses.WriteUnreadableRequestError(res)
			return
		}

		var request webhook.Webhook
		if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
			log.WithError(err).Error("unable to read update Webhook request")
			responses.WriteUnreadableRequestError(res)
			return
		}

		id, err := service.Update(customerID, webhookID, request.WebhookURL)
		if err != nil {
			log.WithError(err).Error("unable to update webhook")
			responses.WriteGatewayTimeout(res)
			return
		}

		response := map[string]interface{}{
			"webhook-id": id,
		}
		responses.WriteOKWithEntity(res, response)

		log.Info("update webhook success")
	}
}
