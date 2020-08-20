package transport

type Service interface {
	PublishNotification(notif *Notification) error
	PublishNotificationStatus(notifUpdate *NotificationStatus) error
	SubsribeNotification(func(notif *Notification) error) error
	SubscribeNotificationStatus(func(notifStatus *NotificationStatus) error) error
}
