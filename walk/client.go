package walk

// Client for register and signin
type Client struct {
	UserId   int    `json:"user_id"`
	Name     string `json:"name"` // for register
	Email    string `json:"email"`
	Password string `json:"pwd"`
}

// Message ... response message
type Message struct {
	Signin   bool   `json:"signin"`
	UserName string `json:"user_name"`
	UserId   int    `json:"user_id"`
}
