package campaign

type Service interface {
	GetAll() ([]Campaign, error)
	GetByUserId(id int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(r *repository) *service {
	return &service{r}
}

func (s service) GetAll() ([]Campaign, error) {
	campaigns, err := s.repository.FindAll()

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s service) GetByUserId(id int) ([]Campaign, error) {
	campaigns, err := s.repository.FindByUserId(id)

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
