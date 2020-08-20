package transport

type Notification struct {
	ID        int         `json:"id"`
	UniqueKey string      `json:"unique-key"`
	IdemKey   string      `json:"idem-key"`
	URL       string      `json:"url"`
	Payload   interface{} `json:"payload"`
	Retry     int         `json:"retry"`
}

type NotificationStatus struct {
	ID                   int          `json:"id"`
	Status               string       `json:"string"`
	FailReason           string       `json:"fail-reason"`
	PreviousNotification Notification `json:"previous-notification"`
}
