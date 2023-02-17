package campaign

type Formatter struct {
	ID               uint64
	UserId           uint64
	Name             string
	ShortDescription string
	ImageUrl         string
	GoalAmount       uint
	CurrentAmount    uint
}

func FormatSingleCampaign(c Campaign) Formatter {
	formatted := Formatter{
		ID:               c.Id,
		UserId:           c.UserId,
		Name:             c.Name,
		ShortDescription: c.ShortDescription,
		GoalAmount:       c.Goal,
		CurrentAmount:    c.CurrentAmount,
	}

	if len(c.CampaignImages) > 0 {

		for _, val := range c.CampaignImages {
			if val.IsPriority == true {
				formatted.ImageUrl = val.FileName
			}
		}
	}
	return formatted
}

func FormatCampaignCollection(c []Campaign, formatSingleCampaigns func(Campaign) Formatter) []Formatter {
	var formatted []Formatter

	for _, campaign := range c {
		formatted = append(formatted, formatSingleCampaigns(campaign))
	}

	return formatted
}
