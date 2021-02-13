package categories

import "database/sql"

type Categorie struct {
	ProID uint8   `json:"pro_id"`
	Name  string  `json:"name"`
	Color string  `json:"color"`
	Size  string  `json:"size"`
	Price float32 `json:"price"`
}

func retriveCategories(db *sql.DB, query string) ([]Categorie, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var categorie []Categorie
	for rows.Next() {
		var c Categorie
		rows.Scan(&c.ProID, &c.Name, &c.Color, &c.Size, &c.Price)
		categorie = append(categorie, c)
	}
	return categorie, nil
}
