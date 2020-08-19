package notification

type DefaultService struct {
	repo Repository
}

func NewDefaultService(repo Repository) Service {
	return &DefaultService{repo}
}

func (s *DefaultService) Send(notif *Notification) (id int, err error) {
	id, err = s.repo.Create(notif)
	if err != nil {
		// log
		return
	}
	// send notification here
	return
}

func (s *DefaultService) Retry(customerID, notificationID int) (id int, err error) {
	// get notif
	// check customer id
	// send notification here
	return
}

func (s *DefaultService) List() (notifs []Notification, err error) {

	return
}

func (s *DefaultService) UpdateStatus(notifID int, status string) (err error) {

	return
}
