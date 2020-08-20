package notification

type Notification struct {
	ID         *int        `json:"id,omitempty"`
	IdemKey    string      `json:"idem-key"`
	WebhookID  int         `json:"webhook-id,required"`
	Payload    interface{} `json:"payload,required"`
	Status     string      `json:"status"`
	FailReason string      `json:"fail-reason"`
}
