package notification

type NotificationRequest struct {
	TransactionTypeID int         `json:"transaction-type-id"`
	CustomerID        int         `json:"customer-id"`
	Payload           interface{} `json:"payload"`
}
