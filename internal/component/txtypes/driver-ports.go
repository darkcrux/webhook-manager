package txtypes

type Service interface {
	Register(tx *TransactionType) (id int, err error)
	List() (txTypes []TransactionType, err error)
	Get(id int) (t *TransactionType, err error)
}
