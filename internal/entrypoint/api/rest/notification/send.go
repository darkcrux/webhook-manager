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

		var request notification.Notification
		json.NewDecoder(req.Body).Decode(&request)

		notifs, err := service.Send(&request)
		if err != nil {
			log.WithError(err).Error("unable to send notification")
			responses.WriteGatewayTimeout(res)
			return
		}

		responses.WriteOKWithEntity(res, notifs)

		log.Info("send notification success")
	}
}
