package users

type User struct {
	Id           int64
	Username     string
	Email        string
	PasswordHash string
	Permissions  uint
}

func ValidateJWT(jwt string) bool {
	return false
}
