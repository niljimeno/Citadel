package repository

import (
	"encoding/csv"
	"os"

	"github.com/niljimeno/citadel/models"
)

// need a way to order alphabetically
// both tags and website names

var db = []models.Result{}
var tag = make(map[string]*models.Result)

func Connect() error {
	file, err := os.Open("db.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, r := range data {
		db = append(db, models.NewResult(r))
	}

	for _, r := range db {
		for _, t := range r.Tags {
			tag[t] = &r
		}
	}

	return nil
}
