package walk

type PurchaseStatus struct {
	Status          string `json:"status"`
	TransactionCode uint8  `json:"transaction_code"` // 00 error card num or cv not valid, 02 ok,  05 not enough to compleate 0.0 balance or purchase is grather than amount
}

// ClientCardInfo ...
type ClientCardInfo struct {
	CardNumber     string  `json:"card_number"`
	CvNumber       uint8   `json:"cv_number"`
	PurchaseAmount float64 `json:"purchase_amount"`
}

// ClientOrder ... new order
type ClientOrder struct {
	Client ClientInfo     `json:"client"`
	Items  []OrderProduct `json:"items"`
	Total  float64        `json:"total"`
}

// ClientInfo ... information
type ClientInfo struct {
	PurchaseID int    `json:"purchase_id"`
	UserId     int    `json:"user_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	State      string `json:"state"`
	Zip        int    `json:"zip"`
	NameOnCard string `json:"name_on_card"`
	CardNumber string `json:"card_number"`
	CvNumber   uint8  `json:"cv_number"`
}

// OrderProduct ... product information
type OrderProduct struct {
	PurchaseID int     `json:"purchase_id"`
	ProductId  int     `json:"product_id"`
	ProName    string  `json:"pro_name"`
	Color      string  `json:"color"`
	Size       string  `json:"size"`
	Qty        int     `json:"qty"`
	Img        string  `json:"img"`
	Price      float64 `json:"price"`
	Total      float64 `json:"total"`
	StatusCode int     `json:"status_code"`
}

type Order struct {
	Client  ClientInfo   `json:"client"`
	Product OrderProduct `json:"product"`
}

type Purchase struct {
	Order []Order `json:"orders"`
}
