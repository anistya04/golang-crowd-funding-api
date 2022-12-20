package user

type RegisterInput struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Password   string `json:"password" binding:"required"`
}
