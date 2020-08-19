package customer

type DefaultService struct {
	repo Repository
}

func NewDefaultService(repo Repository) Service {
	return &DefaultService{repo}
}

func (s *DefaultService) Register(c *Customer) (id int, err error) {
	return s.repo.Save(c)
}
