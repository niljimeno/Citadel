package repository

import (
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/niljimeno/citadel/models"
)

func Search(q string) []models.Result {
	results := []models.Result{}
	normalize := func(s string) string {
		return strings.ToLower(s)
	}

	highPriority := []models.Result{}

	// high priority search
	for _, v := range db {
		if strings.Contains(normalize(v.Name), normalize(q)) {
			results = append(results, v)
		} else {
			highPriority = append(highPriority, v)
		}
	}

	lowPriority := []models.Result{}

	for _, v := range highPriority {
		if strings.Contains(normalize(v.Desc), normalize(q)) {
			results = append(results, v)
		} else {
			lowPriority = append(lowPriority, v)
		}
	}

	// low priority search
	for _, v := range lowPriority {
		if fuzzy.Match(normalize(q), normalize(v.Name)) ||
			fuzzy.Match(normalize(q), normalize(v.Desc)) {
			results = append(results, v)
		}
	}

	return results
}
