package notification

import (
	"net/http"
	"strconv"

	responses "github.com/darkcrux/mux-responses"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/darkcrux/webhook-manager/internal/component/notification"
)

func retryNotification(service notification.Service) http.HandlerFunc {

	log.Info("handler registered: notifications::retry")

	return func(res http.ResponseWriter, req *http.Request) {

		log.Info("retry notification started")

		customerIDStr := mux.Vars(req)["customer-id"]
		customerID, err := strconv.Atoi(customerIDStr)
		if err != nil {
			log.WithError(err).Error("unable to read customerID")
			responses.WriteUnreadableRequestError(res)
			return
		}

		notificationIDStr := mux.Vars(req)["notification-id"]
		notificationID, err := strconv.Atoi(notificationIDStr)
		if err != nil {
			log.WithError(err).Error("unable to read notification")
			responses.WriteUnreadableRequestError(res)
			return
		}

		id, err := service.Retry(customerID, notificationID)
		if err != nil {
			log.WithError(err).Error("unable to send notification")
			responses.WriteGatewayTimeout(res)
			return
		}

		response := map[string]interface{}{
			"notification-id": id,
		}

		responses.WriteOKWithEntity(res, response)

		log.Info("send notification success")
	}
}
