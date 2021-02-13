package dbutils

// table model
type Categorie struct {
	ProID uint8   `json:"pro_id"`
	Name  string  `json:"name"`
	Color string  `json:"color"`
	Size  string  `json:"size"`
	Price float32 `json:"price"`
}
