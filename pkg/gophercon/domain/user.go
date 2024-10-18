package domain

// User is used to model a object
type User struct {
	ID        *string `json:"id"`
	Username  string  `json:"username,omitempty"`
	FirstName string  `json:"first_name,omitempty"`
	LastName  string  `json:"last_name,omitempty"`
	Email     string  `json:"email,omitempty"`
	UserType  string  `json:"user_type,omitempty"`
}
