package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/darkcrux/webhook-manager/internal/entrypoint/api/rest/customer"
	"github.com/darkcrux/webhook-manager/internal/entrypoint/api/rest/notification"
	"github.com/darkcrux/webhook-manager/internal/entrypoint/api/rest/txtypes"
	"github.com/darkcrux/webhook-manager/internal/entrypoint/api/rest/webhook"
)

type (
	Config struct {
		Host    string
		Port    int
		Spec    string
		Version string
		Cors    CORSConfig
	}
	CORSConfig struct {
		AllowedOrigins []string
		AllowedHeaders []string
		AllowedMethods []string
	}
	API struct {
		config                 *Config
		router                 *mux.Router
		txtypesController      *txtypes.Controller
		customerController     *customer.Controller
		webhookController      *webhook.Controller
		notificationController *notification.Controller
	}
)

func NewRestAPI(config *Config, router *mux.Router, txtypesController *txtypes.Controller, customerController *customer.Controller, webhookController *webhook.Controller, notificationController *notification.Controller) *API {

	return &API{
		config:                 config,
		router:                 router,
		txtypesController:      txtypesController,
		customerController:     customerController,
		webhookController:      webhookController,
		notificationController: notificationController,
	}
}

func (api *API) Run() error {
	api.exposeSwagger()
	api.enableCORS()
	api.addMiddlewares()
	api.registerHandlers()
	api.startListeners()
	return http.ListenAndServe(api.address(), api.router)
}

func (api *API) address() string {
	return fmt.Sprintf("%s:%d", api.config.Host, api.config.Port)
}

func (api *API) exposeSwagger() {
	api.router.Path("/spec").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, api.config.Spec)
	})
	log.Infof("OpenAPI Spec accessible at http://%s/spec", api.address())
}

func (api *API) enableCORS() {
	cors := handlers.CORS(
		handlers.AllowedOrigins(api.config.Cors.AllowedOrigins),
		handlers.AllowedHeaders(api.config.Cors.AllowedHeaders),
		handlers.AllowedMethods(api.config.Cors.AllowedMethods),
	)
	api.router.Use(cors)
	log.Info("CORS filter enabled")
}

func (api *API) addMiddlewares() {
	// no middlewares yet
}

func (api *API) registerHandlers() {
	log.Info("Register Handlers")
	api.txtypesController.Register(api.router)
	api.customerController.Register(api.router)
	api.webhookController.Register(api.router)
	api.notificationController.Register(api.router)
}

func (api *API) startListeners() {
	log.Info("Starting Listeners")
	if err := api.notificationController.StartListener(); err != nil {
		log.WithError(err).Error("error starting listeners")
	}
}
