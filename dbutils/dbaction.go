package dbutils

import "database/sql"

func Retrive(db *sql.DB, query string) ([]Categorie, error) {
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
