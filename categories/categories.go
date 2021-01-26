package categories

import (
	"encoding/json"
	"net/http"
)

const (
	// mens
	mensBoots = "mens-boots"
	mensSport = "mens-sport"
	// womens
	womensBoots  = "womens-boots"
	womensSports = "womens-sports"
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
