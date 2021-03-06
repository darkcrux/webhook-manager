package txtypes

type Repository interface {
	Save(tx *TransactionType) (id int, err error)
	List() (txTypes []TransactionType, err error)
	Get(id int) (t *TransactionType, err error)
}
