package user

type User struct {
	Id           uint64
	Name         string
	Email        string
	Occupation   string
	PasswordHash string
	Avatar       string
	Role         string
	CreatedAt    int64
	UpdatedAt    int64
}
