package notification

import (
	"net/http"

	responses "github.com/darkcrux/mux-responses"
	log "github.com/sirupsen/logrus"

	"github.com/darkcrux/webhook-manager/internal/component/notification"
)

func listNotification(service notification.Service) http.HandlerFunc {

	log.Info("handler registered: notifications::list")

	return func(res http.ResponseWriter, req *http.Request) {

		log.Info("create notification started")

		notifs, err := service.List()
		if err != nil {
			log.WithError(err).Error("unable to list notification")
			responses.WriteGatewayTimeout(res)
			return
		}

		responses.WriteOKWithEntity(res, notifs)

		log.Info("create notification success")
	}
}
