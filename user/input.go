package user

type RegisterInput struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Occupation string `json:"occupation" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UniqueEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
