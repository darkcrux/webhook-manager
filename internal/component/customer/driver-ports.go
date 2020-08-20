package customer

type Service interface {
	Register(tx *Customer) (id int, err error)
	Get(id int) (c *Customer, err error)
}
