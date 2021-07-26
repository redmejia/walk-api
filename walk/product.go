package walk

// Store product information models
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
