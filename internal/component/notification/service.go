package notification

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/darkcrux/webhook-manager/internal/component/customer"
	"github.com/darkcrux/webhook-manager/internal/component/transport"
	"github.com/darkcrux/webhook-manager/internal/component/txtypes"
	"github.com/darkcrux/webhook-manager/internal/component/webhook"
)

const (
	maxRetry  = 5
	retryWait = 5 * time.Second
)

type DefaultService struct {
	repo          Repository
	txtypeService txtypes.Service
	cService      customer.Service
	whService     webhook.Service
	tService      transport.Service
}

func NewDefaultService(repo Repository, whService webhook.Service, txtypeService txtypes.Service, cService customer.Service, tService transport.Service) Service {
	return &DefaultService{
		repo:          repo,
		whService:     whService,
		cService:      cService,
		tService:      tService,
		txtypeService: txtypeService,
	}
}

func (s *DefaultService) Test(webhookID int) (id int, err error) {
	log.Infof("Testing Webhook %d ...", webhookID)
	wh, err := s.whService.Get(webhookID)
	if err != nil {
		log.WithError(err).Error("Testing Webhook failed")
		return
	}
	t, err := s.txtypeService.Get(wh.TransactionTypeID)
	if err != nil {
		log.WithError(err).Error("Testing Webhook failed")
		return
	}
	notif := &Notification{
		WebhookID: webhookID,
		Payload:   t.SamplePayload,
	}
	id, err = s.Send(notif)
	if err != nil {
		log.WithError(err).Error("Testing Webhook failed")
		return
	}
	log.Info("Testing Webhook success")
	return
}

func (s *DefaultService) SendInternal(typeID, custID int, payload interface{}) (id int, err error) {
	log.Info("Send Internal...")
	wh, err := s.whService.GetByTxAndCust(typeID, custID)
	if err != nil {
		log.WithError(err).Error("unable to get tx and cust")
		return
	}
	id, err = s.Send(&Notification{
		WebhookID: *wh.ID,
		Payload:   payload,
	})
	if err != nil {
		log.WithError(err).Error("Sending failed")
		return
	}
	log.Info("Send Internal success")
	return
}

func (s *DefaultService) Send(notif *Notification) (id int, err error) {
	log.Infof("Sending Notification...")
	uuid, err := genUUID()
	if err != nil {
		log.WithError(err).Error("Sending Notification failed")
		return
	}
	notif.Status = "PENDING"
	notif.IdemKey = uuid
	notif.Payload, err = json.Marshal(notif.Payload)
	if err != nil {
		log.WithError(err).Error("Sending Notification failed")
		return
	}
	id, err = s.repo.Create(notif)
	if err != nil {
		log.WithError(err).Error("Sending Notification failed")
		return
	}
	if err = s.publishNotification(notif); err != nil {
		log.WithError(err).Error("Sending Notification failed")
		return
	}
	log.Infof("Sending Notification success")
	return
}

func (s *DefaultService) Retry(customerID, notificationID int) (id int, err error) {
	log.Info("Retrying Notification...")
	notif, err := s.repo.Get(notificationID)
	if err != nil {
		log.WithError(err).Error("Retrying Notification failed")
		return
	}
	wh, err := s.whService.Get(notif.WebhookID)
	if err != nil {
		log.WithError(err).Error("Retrying Notification failed")
		return
	}
	if wh.CustomerID != customerID {
		log.WithError(err).Error("Retrying Notification failed")
		err = errors.New("only the same customer can retry a notification")
		return
	}
	if err = s.publishNotification(notif); err != nil {
		log.WithError(err).Error("Retrying Notification failed")
		return
	}
	log.Info("Retrying Notification success")
	return
}

func (s *DefaultService) List(customerID int) (notifs []Notification, err error) {
	log.Info("Getting Notification List...")
	notifs, err = s.repo.List(customerID)
	if err != nil {
		log.WithError(err).Error("Getting Notification List failed")
		return
	}
	log.Info("Getting Notification List success")
	return
}

func (s *DefaultService) UpdateStatus(notifID int, status string) (err error) {
	log.Infof("Updating Notification Status for %d ...", notifID)
	if err = s.repo.UpdateStatus(notifID, status); err != nil {
		log.WithError(err).Error("Updating Notification Status failed")
		return
	}
	log.Info("Updating Notification Status success")
	return
}

func (s *DefaultService) StartLiseners() (err error) {
	log.Info("Starting Notification Listeners...")
	if err = s.tService.SubsribeNotification(s.notifHandler); err != nil {
		log.WithError(err).Error("Notif Listener failed to start")
		return
	}
	log.Info("notification listerner started")
	if err = s.tService.SubscribeNotificationStatus(s.notifStatusHandler); err != nil {
		log.WithError(err).Error("Notif Status Listener failed to start")
		return
	}
	log.Info("notification status listerner started")
	return
}

func (s *DefaultService) publishNotification(notif *Notification) (err error) {
	log.Info("Publishing Notification to Message Bus...")
	wh, err := s.whService.Get(notif.WebhookID)
	if err != nil {
		log.WithError(err).Error("Publishing Notification to Message Bus failed")
		return
	}

	customer, err := s.cService.Get(wh.CustomerID)
	if err != nil {
		log.WithError(err).Error("Publishing Notification to Message Bus failed")
		return
	}

	pubNotif := &transport.Notification{
		ID:        *notif.ID,
		UniqueKey: customer.UniqueKey,
		IdemKey:   notif.IdemKey,
		URL:       wh.WebhookURL,
		Payload:   notif.Payload,
		Retry:     0,
	}

	if err = json.Unmarshal(pubNotif.Payload.([]byte), &pubNotif.Payload); err != nil {
		log.WithError(err).Error("unable to read JSON payload from DB")
		return
	}

	if err = s.tService.PublishNotification(pubNotif); err != nil {
		log.WithError(err).Error("Publishing Notification to Message Bus failed")
		s.UpdateStatus(*notif.ID, "FAIL")
		return
	}

	log.Info("Publishing Notification to Message Bus success")
	return
}

func genUUID() (string, error) {
	log.Info("Generating UUID as idempotent ID")
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.WithError(err).Error("Generating UUID as idempotent ID failed")
		return "", err
	}
	return uuid.String(), nil
}

func (s *DefaultService) notifHandler(notif *transport.Notification) error {
	log.Info("Handling new Notification...")
	data, err := json.Marshal(notif.Payload)
	if err != nil {
		log.WithError(err).Error("Unable to marshal request to JSON")
		s.tService.PublishNotificationStatus(&transport.NotificationStatus{
			ID:                   notif.ID,
			Status:               "FAIL",
			FailReason:           err.Error(),
			PreviousNotification: *notif,
		})
		return err
	}
	dataBuf := bytes.NewBuffer(data)
	req, err := http.NewRequest(http.MethodPost, notif.URL, dataBuf)
	if err != nil {
		log.WithError(err).Error("Unable to create new HTTP Request")
		s.tService.PublishNotificationStatus(&transport.NotificationStatus{
			ID:                   notif.ID,
			Status:               "FAIL",
			FailReason:           err.Error(),
			PreviousNotification: *notif,
		})
		return err
	}
	req.Header.Add("WEBHOOK-UNIQUE-KEY", notif.UniqueKey)
	req.Header.Add("NOTIF-IDEM-KEY", notif.IdemKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.WithError(err).Error("Unable to call webhook")
		s.tService.PublishNotificationStatus(&transport.NotificationStatus{
			ID:                   notif.ID,
			Status:               "FAIL",
			FailReason:           err.Error(),
			PreviousNotification: *notif,
		})
		return err
	}
	if res.StatusCode != http.StatusOK {
		log.Errorf("Webhook did not return OK: %d", res.StatusCode)
		s.tService.PublishNotificationStatus(&transport.NotificationStatus{
			ID:                   notif.ID,
			Status:               "FAIL",
			FailReason:           "webhook did not return OK",
			PreviousNotification: *notif,
		})
		return err
	}
	err = s.tService.PublishNotificationStatus(&transport.NotificationStatus{
		ID:                   notif.ID,
		Status:               "SUCCESS",
		PreviousNotification: *notif,
	})
	if err != nil {
		log.WithError(err).Error("Unable to publish notification")
		return err
	}
	log.Info("Notif successful sent to webhook")
	return nil
}

func (s *DefaultService) notifStatusHandler(notifStatus *transport.NotificationStatus) error {
	log.Info("Handling new Notification Status Update Mesasge")

	// retry until max is reached
	if notifStatus.Status == "FAIL" {
		notif := notifStatus.PreviousNotification
		if notif.Retry < maxRetry {
			// retry threshold not met, retrying
			notif.Retry = notif.Retry + 1
			log.Infof("Retrying failed notification: notif-id: %d, retry: %d", notif.ID, notif.Retry)
			time.AfterFunc(retryWait, func() {
				s.tService.PublishNotification(&notif)
			})
			return nil
		}
		log.Info("Retry Limit reached, failing")
	}

	// else update status
	err := s.repo.UpdateStatus(notifStatus.ID, notifStatus.Status)
	if err != nil {
		log.WithError(err).Error("Unable to update notification status")
		return err
	}
	log.Info("Handling new Notification Status Update Mesasge success")
	return nil
}
