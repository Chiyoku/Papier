package user

type User struct {
	ID           int
	Username     string
	Email        string
	PasswordHash string
	Permissions  uint
}

func ValidateJWT(jwt string) bool {
	return false
}
