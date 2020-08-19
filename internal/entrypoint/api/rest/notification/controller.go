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
		Path("/notifications").
		HandlerFunc(listNotification(service))

	tx.
		Methods(http.MethodPost, http.MethodOptions).
		Path("/send-notification").
		HandlerFunc(listNotification(service))

	tx.
		Methods(http.MethodPost, http.MethodOptions).
		Path("/customers/{customer-id}/notifications/{notification-id}/retry").
		HandlerFunc(retryNotification(service))

}
