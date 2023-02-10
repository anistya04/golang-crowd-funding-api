package campaign

type Campaign struct {
	Id               uint64 `json:"id"`
	UserId           uint64 `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Goal             uint   `json:"goal"`
	CurrentAmount    uint   `json:"current_amount"`
	Description      string `json:"description"`
	Perks            string `json:"perks"`
	Slug             string `json:"slug"`
	BakerCount       uint   `json:"baker_count"`
	CreatedAt        uint64
	UpdatedAt        uint64
	CampaignImages   []CampaignImages `json:"images" gorm:"foreignKey:CampaignId;references:id"`
}

type CampaignImages struct {
	Id         uint64 `json:"id"`
	CampaignId uint64 `json:"campaign_id"`
	FileName   string `json:"file_name"`
	IsPriority bool   `json:"is_priority"`
	CreatedAt  uint64
	UpdatedAt  uint64
}
