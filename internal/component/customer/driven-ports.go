package customer

type Repository interface {
	Save(c *Customer) (id int, err error)
}
