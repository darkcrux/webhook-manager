package notification

type Repository interface {
	Create(notif *Notification) (id int, err error)
	Get(id int) (notif *Notification, err error)
	List(customerID int) (notifs []Notification, err error)
	UpdateStatus(id int, status string) (err error)
}
