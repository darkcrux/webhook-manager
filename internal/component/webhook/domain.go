package webhook

// Webhook represents a webhook that contains the transaction type, customer, and the webhook URL
type Webhook struct {
	ID                *int   `json:"id,omitempty"`
	CustomerID        int    `json:"customer-id,required"`
	TransactionTypeID int    `json:"transaction-type-id,required"`
	WebhookURL        string `json:"webhook-url,required"`
}
