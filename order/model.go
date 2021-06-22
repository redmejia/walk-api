package order

type PurchaseStatus struct {
	Status          string `json:"status"`
	TransactionCode uint8  `json:"transaction_code"` // 00 error card num or cv not valid, 02 ok,  05 not enough to compleate 0.0 balance or purchase is grather than amount
}

type ClientInfo struct {
	CardNumber     string  `json:"card_number"`
	CvNumber       uint8   `json:"cv_number"`
	PurchaseAmount float32 `json:"purchase_amount"`
}
