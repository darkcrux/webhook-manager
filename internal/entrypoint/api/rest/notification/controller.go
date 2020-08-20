package notification

import (
	"net/http"

	"github.com/darkcrux/webhook-manager/internal/component/notification"
	"github.com/gorilla/mux"
)

type Controller struct {
	service notification.Service
}

func NewController(service notification.Service) *Controller {
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
		Methods(http.MethodGet, http.MethodOptions).
		Path("/customers/{customer-id}/notifications").
		HandlerFunc(listNotification(service))

	tx.
		Methods(http.MethodPost, http.MethodOptions).
		Path("/send-notification").
		HandlerFunc(sendNotification(service))

	tx.
		Methods(http.MethodPost, http.MethodOptions).
		Path("/customers/{customer-id}/notifications/{notification-id}/retry").
		HandlerFunc(retryNotification(service))

	tx.
		Methods(http.MethodPost, http.MethodOptions).
		Path("/customers/{customer-id}/webhooks/{webhook-id}/test").
		HandlerFunc(testNotification(service))

}

func (c *Controller) StartListener() error {
	return c.service.StartLiseners()
}
