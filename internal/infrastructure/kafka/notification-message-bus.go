package kafka

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"

	"github.com/darkcrux/webhook-manager/internal/component/transport"
)

type KafkaNotifMessageBus struct {
	notifWriter       *kafka.Writer
	notifStatusWriter *kafka.Writer
	notifReader       *kafka.Reader
	notifStatusReader *kafka.Reader
}

func NewKafkaNotifMessageBus(notifWriter, notifStatusWriter *kafka.Writer, notifReader, notifStatusReader *kafka.Reader) transport.MessageBus {
	return &KafkaNotifMessageBus{
		notifWriter:       notifWriter,
		notifStatusWriter: notifStatusWriter,
		notifReader:       notifReader,
		notifStatusReader: notifStatusReader,
	}
}

func (mb *KafkaNotifMessageBus) PublishNotification(notif *transport.Notification) error {
	log.Info("Publishing Notification to Kafka...")
	data, err := json.Marshal(notif)
	if err != nil {
		log.WithError(err).Error("notification cannot be marhsaled to JSON")
		return err
	}
	if err := mb.notifWriter.WriteMessages(context.Background(), kafka.Message{Value: data}); err != nil {
		log.WithError(err).Error("error publishing notification to kafka")
		return err
	}
	log.Info("Publishing Notification to Kafka success")
	return nil
}

func (mb *KafkaNotifMessageBus) PublishNotificationStatus(notifStatus *transport.NotificationStatus) error {
	log.Info("Publishing Notification Status Update to Kafka...")
	data, err := json.Marshal(notifStatus)
	if err != nil {
		log.WithError(err).Error("notification status update cannot be marshaled to JSON")
		return err
	}
	if err := mb.notifStatusWriter.WriteMessages(context.Background(), kafka.Message{Value: data}); err != nil {
		log.WithError(err).Error("error publishing notification status update to kafka")
		return err
	}
	log.Info("Publishing Notification Status Update to Kafka success")
	return nil
}

func (mb *KafkaNotifMessageBus) SubsribeNotification(handler func(notif *transport.Notification) error) error {
	log.Info("Subscribing to Kafka for Notification Messages...")
	go func() {
		for {
			log.Info("Waiting for Notification Message...")
			msg, err := mb.notifReader.ReadMessage(context.Background())
			if err != nil {
				log.WithError(err).Error("Error reading new message, ignore message")
				// ignoring for now
				continue
			}
			var notif transport.Notification
			if err := json.Unmarshal(msg.Value, &notif); err != nil {
				log.WithError(err).Error("Error reading message as JSON, ignore message")
				// ignoring for now
				continue
			}
			if err := handler(&notif); err != nil {
				log.WithError(err).Error("Error received from message handler")
				// ignoring for now
				continue
			}
		}
	}()
	log.Info("Subscribing to Kafka for Notification Messages success")
	return nil
}

func (mb *KafkaNotifMessageBus) SubscribeNotificationStatus(handler func(notifStatus *transport.NotificationStatus) error) error {
	log.Info("Subscribing to Kafka for Notification Status Update Messages...")
	go func() {
		for {
			msg, err := mb.notifStatusReader.ReadMessage(context.Background())
			if err != nil {
				log.WithError(err).Error("Error reading new message, ignore message")
				// ignoring for now
				continue
			}
			var notifStatus transport.NotificationStatus
			if err := json.Unmarshal(msg.Value, &notifStatus); err != nil {
				log.WithError(err).Error("Error reading message as JSON, ignore message")
				// ignoring for now
				continue
			}
			if err := handler(&notifStatus); err != nil {
				log.WithError(err).Error("Error received from message handler")
				// ignoring for now
				continue
			}
		}
	}()
	log.Info("Subscribing to Kafka for Notification Status Update Messages success")
	return nil
}
