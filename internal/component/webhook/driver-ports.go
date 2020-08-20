package webhook

type Service interface {
	Create(wh *Webhook) (id int, err error)
	Update(customerID, webhookID int, url string) (id int, err error)
	Get(id int) (wh *Webhook, err error)
	List(id int) (whs []Webhook, err error)
	GetByTxAndCust(txId, custID int) (wh *Webhook, err error)
}
