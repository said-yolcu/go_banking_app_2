package models

// This struct is not used in database, it is just
// used as a middle struct
type Signup struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	StateID  string `json:"state_id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	// Initial balance
	Balance int `json:"balance"`
}
