package categories

import "database/sql"

type Product struct {
	ProID uint8   `json:"pro_id"`
	Name  string  `json:"name"`
	Color string  `json:"color"`
	Size  string  `json:"size"`
	Price float32 `json:"price"`
}

func retriveProducts(db *sql.DB, query string) ([]Product, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var products []Product
	for rows.Next() {
		var p Product
		rows.Scan(&p.ProID, &p.Name, &p.Color, &p.Size, &p.Price)
		products = append(products, p)
	}
	return products, nil
}
