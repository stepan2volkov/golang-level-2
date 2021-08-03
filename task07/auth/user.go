package auth

//easyjson:json
type User struct {
	ID          int
	Login       string
	Email       string
	Permissions []string
}
