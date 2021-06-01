package dbutils

// table model
// Products ...
type Products struct {
	ProID     uint8   `json:"pro_id"`
	ProductID int     `json:"product_id"`
	ProName   string  `json:"pro_name"`
	Price     float32 `json:"price"`
}

// Product ...
type Product struct {
	ProductID int     `json:"product_id"`
	ProName   string  `json:"pro_name"`
	Price     float32 `json:"price"`
}

// ProductInfo ...
type ProductInfo struct {
	Product
	Size   []string `json:"sizes"`
	Colors []string `json:"colors"`
}

// Sizes ...
type Sizes struct {
	ProductID int
	SizeOne   string
	SizeTwo   string
	SizeThree string
	SizeFour  string
}

// Colors ...
type Colors struct {
	ProductID  int
	ColorOne   string
	ColorTwo   string
	ColorThree string
	ColorFour  string
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
