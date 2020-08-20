package notification

import (
	"net/http"
	"strconv"

	responses "github.com/darkcrux/mux-responses"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/darkcrux/webhook-manager/internal/component/notification"
)

func listNotification(service notification.Service) http.HandlerFunc {

	log.Info("handler registered: notifications::list")

	return func(res http.ResponseWriter, req *http.Request) {

		log.Info("create notification started")

		customerIDStr := mux.Vars(req)["customer-id"]
		customerID, err := strconv.Atoi(customerIDStr)
		if err != nil {
			log.WithError(err).Error("unable to read customer ID")
			responses.WriteUnreadableRequestError(res)
			return
		}

		notifs, err := service.List(customerID)
		if err != nil {
			log.WithError(err).Error("unable to list notification")
			responses.WriteGatewayTimeout(res)
			return
		}

		responses.WriteOKWithEntity(res, notifs)

		log.Info("create notification success")
	}
}
