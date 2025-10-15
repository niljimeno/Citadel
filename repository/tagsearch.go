package repository

import (
	"slices"

	"github.com/niljimeno/citadel/models"
)

func FilterByTag(q string) []models.Result {
	results := []models.Result{}
	for _, r := range db {
		if slices.Contains(r.Tags, q) {
			results = append(results, r)
			break
		}
	}

	return results
}
