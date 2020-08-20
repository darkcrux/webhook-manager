package webhook

type Repository interface {
	Save(wh *Webhook) (id int, err error)
	GetByID(id int) (wh *Webhook, err error)
	GetByTxIDAndCustomerID(txID, customerID int) (wh *Webhook, err error)
	List(id int) (whs []Webhook, err error)
}
