package entities

type User struct {
	Id        UuidV7Str `json:"id"`
	UserName  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	// Email     string    `json:"email"`
	// Phone     string    `json:"phone"`
}
