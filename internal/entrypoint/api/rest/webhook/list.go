package webhook

import (
	"net/http"
	"strconv"

	responses "github.com/darkcrux/mux-responses"
	"github.com/darkcrux/webhook-manager/internal/component/webhook"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func listWebhooks(service webhook.Service) http.HandlerFunc {

	log.Info("handler registered: webhooks::list")

	return func(res http.ResponseWriter, req *http.Request) {

		log.Info("list Webhook started")

		customerIDStr := mux.Vars(req)["id"]
		customerID, err := strconv.Atoi(customerIDStr)
		if err != nil {
			log.WithError(err).Error("customer ID should be a number")
			responses.WriteUnreadableRequestError(res)
			return
		}

		webhooks, err := service.List(customerID)
		if err != nil {
			log.WithError(err).Error("unable to list webhook")
			responses.WriteGatewayTimeout(res)
			return
		}

		responses.WriteOKWithEntity(res, webhooks)

		log.Info("create webhook success")
	}
}
