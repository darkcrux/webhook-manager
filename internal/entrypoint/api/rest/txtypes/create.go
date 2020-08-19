package txtypes

import (
	"encoding/json"
	"net/http"

	responses "github.com/darkcrux/mux-responses"
	log "github.com/sirupsen/logrus"

	"github.com/darkcrux/webhook-manager/internal/component/txtypes"
)

func createTxType(service txtypes.Service) http.HandlerFunc {

	log.Info("handler registered: transaction-types::create")

	return func(res http.ResponseWriter, req *http.Request) {

		log.Info("create transaction type started")

		var request txtypes.TransactionType
		if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
			log.WithError(err).Error("unable to read create transaction type request")
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
			"transaction-type-id": id,
		}
		responses.WriteOKWithEntity(res, response)

		log.Info("create transaction type success")
	}
}
