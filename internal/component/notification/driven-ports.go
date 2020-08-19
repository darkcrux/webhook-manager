package notification

type Repository interface {
	Create(notif *Notification) (id int, err error)
	UpdateStatus(id int, status string) (err error)
}
