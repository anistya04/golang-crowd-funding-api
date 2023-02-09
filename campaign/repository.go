package campaign

import "gorm.io/gorm"

type Repository interface {
	FindByUserId(Id int) ([]Campaign, error)
	FindAll() ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func (r repository) NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) FindByUserId(Id int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Find(&campaigns, "user_id = ?", Id).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r repository) FindAll() ([]Campaign, error) {

	var campaigns []Campaign
	err := r.db.Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
