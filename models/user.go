package models

// user model comes here

type User struct {

	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	City     string `json:"city"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
	Role     string `json:"role"`

}
