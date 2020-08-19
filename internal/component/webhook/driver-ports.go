package webhook

type Service interface {
	Create(wh *Webhook) (id int, err error)
	Update(customerID, webhookID int, url string) (id int, err error)
	Test(wh *Webhook) (id int, err error)
}
