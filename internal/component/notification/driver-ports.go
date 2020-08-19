package notification

type Service interface {
	Send(notif *Notification) (id int, err error)
	Retry(customerID, notificationID int) (id int, err error)
	List() (notifs []Notification, err error)
	UpdateStatus(notifID int, status string) (err error)
}
