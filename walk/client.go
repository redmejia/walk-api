package walk

// Client Register and Signin model forms
// ClientRegister ... client register with name and email
type ClientRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"pwd"`
}

// ClientSignin ... client signing user email and password
type ClientSignin struct {
	UserId   int    `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"pwd"`
}

// Message ... response message
type Message struct {
	Signin   bool   `json:"signin"`
	UserName string `json:"user_name"`
	UserId   int    `json:"user_id"`
}
