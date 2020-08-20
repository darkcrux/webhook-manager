package transport

import log "github.com/sirupsen/logrus"

type DefaultService struct {
	mb MessageBus
}

func NewDefaultService(mb MessageBus) Service {
	return &DefaultService{
		mb: mb,
	}
}

func (s *DefaultService) PublishNotification(notif *Notification) error {
	log.Info("Publishing Notification...")
	if err := s.mb.PublishNotification(notif); err != nil {
		log.WithError(err).Error("Unable to publish notification to message bus")
		return err
	}
	log.Info("Publishing Notification success")
	return nil
}

func (s *DefaultService) PublishNotificationStatus(notifUpdate *NotificationStatus) error {
	log.Info("Publishing Notification Status Update...")
	if err := s.mb.PublishNotificationStatus(notifUpdate); err != nil {
		log.WithError(err).Error("Unable to publish notification status update to message bus")
		return err
	}
	log.Info("Publishing Notification Status Update success")
	return nil
}

func (s *DefaultService) SubsribeNotification(handler func(notif *Notification) error) error {
	log.Info("Subscribing to Notification Messages...")
	if err := s.mb.SubsribeNotification(handler); err != nil {
		log.WithError(err).Error("Unable to subscribe to notification messages")
		return err
	}
	log.Info("Subscribing to Notification Messages success")
	return nil
}

func (s *DefaultService) SubscribeNotificationStatus(handler func(notifStatus *NotificationStatus) error) error {
	log.Info("Subscribing to Notification Status Update Messages...")
	if err := s.mb.SubscribeNotificationStatus(handler); err != nil {
		log.WithError(err).Error("Unable to up subscribe to notification status update messages")
		return err
	}
	log.Info("Subscribing to Notification Status Update Messages success")
	return nil
}
