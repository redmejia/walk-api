package dbutils

// table model
type Product struct {
	ProID uint8   `json:"pro_id"`
	Name  string  `json:"name"`
	Color string  `json:"color"`
	Size  string  `json:"size"`
	Price float32 `json:"price"`
}

type Order struct {
	ProID uint8   `json:"pro_id"`
	Name  string  `json:"name"`
	Color string  `json:"color"`
	Size  string  `json:"size"`
	Total float32 `json:"total"`
}

// register and signin table
type registerForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"pwd"`
}

type SigninForm struct {
	Email    string `json:"email"`
	Password string `json:"pwd"`
}
