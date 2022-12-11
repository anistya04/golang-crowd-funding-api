package user

type User struct {
	Id           uint16
	Name         string
	Email        string
	Occupation   string
	PasswordHash string
	Avatar       string
	Role         string
	Token        string
	CreatedAt    int64
	UpdatedAt    int64
}
