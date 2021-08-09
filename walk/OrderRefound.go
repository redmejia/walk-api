package walk

// OrderRefound delete order and refoun.
type OrderRefound struct {
	CardNumber string  `json:"card_number"`
	CvNumber   uint8   `json:"cv_number"`
	Refound    float64 `json:"refound"`
}
