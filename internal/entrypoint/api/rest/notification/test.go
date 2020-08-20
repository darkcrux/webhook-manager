package notification

import (
	"net/http"
	"strconv"

	responses "github.com/darkcrux/mux-responses"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/darkcrux/webhook-manager/internal/component/notification"
)

func testNotification(service notification.Service) http.HandlerFunc {

	log.Info("handler registered: notifications::test")

	return func(res http.ResponseWriter, req *http.Request) {

		log.Info("test notification started")

		webhookIDStr := mux.Vars(req)["webhook-id"]
		webhookID, err := strconv.Atoi(webhookIDStr)
		if err != nil {
			log.WithError(err).Error("unable to read webhook ID")
			responses.WriteUnreadableRequestError(res)
			return
		}

		id, err := service.Test(webhookID)
		if err != nil {
			log.WithError(err).Error("unable to test notification")
			responses.WriteGatewayTimeout(res)
			return
		}

		response := map[string]interface{}{
			"notification-id": id,
		}
		responses.WriteOKWithEntity(res, response)

		log.Info("test notification success")
	}
}
