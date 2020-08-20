package customer

type Repository interface {
	Save(c *Customer) (id int, err error)
	GetByID(id int) (c *Customer, err error)
}
