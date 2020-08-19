package customer

type Service interface {
	Register(tx *Customer) (id int, err error)
}
