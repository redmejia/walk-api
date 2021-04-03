package clients

type registerForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"pwd"`
}

type signinForm struct {
	Email    string `json:"email"`
	Password string `json:"pwd"`
}
