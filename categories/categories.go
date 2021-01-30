package categories

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/redmejia/connection"
)

const (
	// mens
	mensBoots = "mens-boots"
	mensSport = "mens-sport"
	// womens
	womensBoots  = "womens-boots"
	womensSports = "womens-sports" // heels
)

type MensCat struct {
	Name string    `json:"name"`
	Size []float32 `json:"size"`
}

func Categories(w http.ResponseWriter, r *http.Request) {
	rQ := r.URL.Query().Get("cat")
	switch rQ {
	case mensBoots:
		men := []MensCat{
			MensCat{
				Name: "One",
				Size: []float32{1., 4., 7., 8.},
			},
			MensCat{
				Name: "two",
				Size: []float32{1., 4., 7., 8.},
			},
			MensCat{
				Name: "three",
				Size: []float32{1., 4., 7., 8.},
			},
		}
		db, err := connection.Dbconn()
		if err != nil {
			fmt.Println("error", err)
			return
		}
		defer db.Close()
		fmt.Println("runing")
		rows, _ := db.Query("SELECT name, color, size, price FROM boots_mens")
		var name string
		var color string
		var size string
		var price float32
		for rows.Next() {
			rows.Scan(&name, &color, &size, &price)
			fmt.Println(name, color, size, price)
		}
		json.NewEncoder(w).Encode(men)
	case mensSport:
		men := []MensCat{
			MensCat{
				Name: "One",
				Size: []float32{1., 4., 7., 8.},
			},
			MensCat{
				Name: "two",
				Size: []float32{1., 4., 7., 8.},
			},
		}
		json.NewEncoder(w).Encode(men)
	case womensBoots:
		women := []MensCat{
			MensCat{
				Name: "Wo-One",
				Size: []float32{1., 4., 7., 8.},
			},
			MensCat{
				Name: "Wo-two",
				Size: []float32{1., 4., 7., 8.},
			},
			MensCat{
				Name: "Wo-three",
				Size: []float32{1., 4., 7., 8.},
			},
		}
		json.NewEncoder(w).Encode(women)
	case womensSports:
		women := []MensCat{
			MensCat{
				Name: "Wo-two",
				Size: []float32{1., 4., 7., 8.},
			},
			MensCat{
				Name: "Wo-three",
				Size: []float32{1., 4., 7., 8.},
			},
		}
		json.NewEncoder(w).Encode(women)
	default:

		http.Error(w, "SOMETHIG GOES WRONG", http.StatusInternalServerError)
		return
	}
}
