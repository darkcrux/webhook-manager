package webhook

import (
	"net/http"

	"github.com/darkcrux/webhook-manager/internal/component/webhook"
	"github.com/gorilla/mux"
)

type Controller struct {
	service webhook.Service
}

func NewController(service webhook.Service) *Controller {
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
		Path("/customers/{id}/webhooks").
		HandlerFunc(createWebhook(service))

	tx.
		Methods(http.MethodGet, http.MethodOptions).
		Path("/customers/{id}/webhooks").
		HandlerFunc(listWebhooks(service))

	tx.
		Methods(http.MethodPost, http.MethodOptions).
		Path("/customers/{customer-id}/webhooks/{webhook-id}").
		HandlerFunc(updateWebhook(service))

}
