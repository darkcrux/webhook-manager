package txtypes

import (
	"net/http"

	"github.com/darkcrux/webhook-manager/internal/component/txtypes"
	"github.com/gorilla/mux"
)

type Controller struct {
	service txtypes.Service
}

func NewController(service txtypes.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) Register(router *mux.Router) {
	tx := router.
		PathPrefix("").
		Subrouter()

	service := c.service

	tx.
		Methods(http.MethodPost, http.MethodOptions).
		Path("/transaction-types").
		HandlerFunc(createTxType(service))

	tx.
		Methods(http.MethodGet, http.MethodOptions).
		Path("/transaction-types").
		HandlerFunc(listTxType(service))

}
