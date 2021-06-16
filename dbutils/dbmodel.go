package dbutils

// table model
// Products ...
type Products struct {
	ProductID  int     `json:"product_id"`
	ProName    string  `json:"pro_name"`
	Price      float32 `json:"price"`
	ProductImg string  `json:"product_img"`
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
	Image  []string `json:"image"`
}

// ProductSizes ... four size can add more
type ProductSize struct {
	SizeOne   string
	SizeTwo   string
	SizeThree string
	SizeFour  string
}

// ProductColor ... four color can add more
type ProductColor struct {
	ColorOne   string
	ColorTwo   string
	ColorThree string
	ColorFour  string
}

// ProductImage ... two images per product you can add more
type ProductImage struct {
	ImgOne string
	ImgTwo string
}

// Order ...
type Order struct {
	ProID      uint8   `json:"pro_id"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Email      string  `json:"email"`
	Address    string  `json:"address"` // I will store on one columna but can be separate
	CardName   string  `json:"card_name"`
	CardNumber string  `json:"card_number"`
	CvNumber   uint8   `json:"cv_number"`
	Name       string  `json:"name"`
	Color      string  `json:"color"`
	Size       string  `json:"size"`
	Total      float32 `json:"total"`
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
