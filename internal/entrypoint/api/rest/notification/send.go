package notification

import (
	"encoding/json"
	"net/http"

	responses "github.com/darkcrux/mux-responses"
	log "github.com/sirupsen/logrus"

	"github.com/darkcrux/webhook-manager/internal/component/notification"
)

func sendNotification(service notification.Service) http.HandlerFunc {

	log.Info("handler registered: notifications::send")

	return func(res http.ResponseWriter, req *http.Request) {

		log.Info("send notification started")

		var request NotificationRequest
		if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
			log.WithError(err).Error("unable to read request body")
			responses.WriteUnreadableRequestError(res)
			return
		}

		id, err := service.SendInternal(
			request.TransactionTypeID,
			request.CustomerID,
			request.Payload,
		)
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
