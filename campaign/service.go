package campaign

import "fmt"

type Service interface {
	GetCampaigns(id int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(r *repository) *service {
	return &service{r}
}

func (s service) GetCampaigns(UserId int) ([]Campaign, error) {

	if UserId != 0 {
		fmt.Println("this is", UserId)
		campaigns, err := s.repository.FindByUserId(UserId)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
