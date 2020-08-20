package transport

type MessageBus interface {
	PublishNotification(notif *Notification) error
	PublishNotificationStatus(notifStatus *NotificationStatus) error
	SubsribeNotification(func(notif *Notification) error) error
	SubscribeNotificationStatus(func(notifStatus *NotificationStatus) error) error
}
