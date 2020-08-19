package notification

type Notification struct {
	ID                *int        `json:"id,omitempty"`
	WebhookID         int         `json:"webhook-id,required"`
	TransactionTypeID int         `json:"transaction-type-id,required"`
	Payload           interface{} `json:"payload,required"`
	Status            string      `json:"status"`
	FailReason        string      `json:"fail-reason"`
}
