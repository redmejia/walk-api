package dbutils

// table model
// Products ...
type Products struct {
	ProID     uint8   `json:"pro_id"`
	ProductID int     `json:"product_id"`
	ProName   string  `json:"pro_name"`
	Color     string  `json:"color"`
	Size      string  `json:"size"`
	Price     float32 `json:"price"`
}

// Product ...
type Product struct {
	ProductID int     `json:"product_id"`
	ProName   string  `json:"pro_name"`
	Color     string  `json:"color"`
	Size      string  `json:"size"`
	Price     float32 `json:"price"`
}

// Order ...
type Order struct {
	ProID uint8   `json:"pro_id"`
	Name  string  `json:"name"`
	Color string  `json:"color"`
	Size  string  `json:"size"`
	Total float32 `json:"total"`
}

// Register ...
type Register struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"pwd"`
}

// Signin ...
type Signin struct {
	UserId   int    `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"pwd"`
}
