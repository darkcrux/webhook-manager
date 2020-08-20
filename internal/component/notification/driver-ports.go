package notification

type Service interface {
	Test(webhookID int) (id int, err error)
	SendInternal(typeID, custID int, payload interface{}) (id int, err error)
	Send(notif *Notification) (id int, err error)
	Retry(customerID, notificationID int) (id int, err error)
	List(customerID int) (notifs []Notification, err error)
	UpdateStatus(notifID int, status string) (err error)
	StartLiseners() (err error)
}
