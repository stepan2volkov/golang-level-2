package auth

//easyjson:json
type User struct {
	ID          int
	Username    string
	Email       string
	Permissions []string
}
