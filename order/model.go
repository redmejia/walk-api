package order

type PurchaseStatus struct {
	Status          string `json:"status"`
	TransactionCode uint8  `json:"transaction_code"` // 00 error card num or cv not valid, 02 ok,  05 not enough to compleate 0.0 balance or purchase is grather than amount
}

type ClientInfo struct {
	CardNumber     string  `json:"card_number"`
	CvNumber       uint8   `json:"cv_number"`
	PurchaseAmount float64 `json:"purchase_amount"`
}

// Order ...
type Order struct {
	Client Client    `json:"client"`
	Items  []Product `json:"items"`
	Total  float64   `json:"total"`
}

type Client struct {
	PurchaseID int    `json:"purchase_id"`
	UserId     int    `json:"user_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Address    string `json:"address"` // I will store on one columna but can be separate
	State      string `json:"state"`
	Zip        int    `json:"zip"`
	NameOnCard string `json:"name_on_card"`
	CardNumber string `json:"card_number"`
	CvNumber   uint8  `json:"cv_number"`
}

type Product struct {
	PurchaseID int     `json:"purchase_id"`
	ProductId  int     `json:"product_id"`
	ProName    string  `json:"pro_name"`
	Color      string  `json:"color"`
	Size       string  `json:"size"`
	Qty        int     `json:"qty"`
	Img        string  `json:"img"`
	Price      float64 `json:"price"`
	StatusCode int     `json:"status_code"`
}

type Totals struct {
	PurchaseID int     `json:"purchase_id"`
	Total      float64 `json:"total"`
}
type MyOrder struct {
	Client  Client  `json:"client"`
	Product Product `json:"product"`
}
type Purchase struct {
	// Client []Client  `json:"client"`
	Orders []MyOrder `json:"orders"`
	Totals []Totals  `json:"totals"`
}
