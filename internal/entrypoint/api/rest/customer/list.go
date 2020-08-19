package customer

import (
	"net/http"

	responses "github.com/darkcrux/mux-responses"
	log "github.com/sirupsen/logrus"

	"github.com/darkcrux/webhook-manager/internal/component/txtypes"
)

func listTxType(service txtypes.Service) http.HandlerFunc {

	log.Info("handler registered: transaction-types::list")

	return func(res http.ResponseWriter, req *http.Request) {

		log.Info("list transaction type started")

		types, err := service.List()
		if err != nil {
			log.WithError(err).Error("unable to get a list of transaction types")
			responses.WriteGatewayTimeout(res)
			return
		}

		responses.WriteOKWithEntity(res, types)

		log.Info("list transaction type success")
	}
}
