package customer

import (
	"net/http"

	"github.com/darkcrux/webhook-manager/internal/component/customer"
	"github.com/gorilla/mux"
)

type Controller struct {
	service customer.Service
}

func NewController(service customer.Service) *Controller {
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
		Path("/customers").
		HandlerFunc(createCustomer(service))

}
