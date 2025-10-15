package repository

import (
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/niljimeno/citadel/models"
)

func normalize(s string) string {
	return strings.ToLower(s)
}

func Search(query string) []models.Result {
	query = normalize(query)

	results, remainder := filter(query, db, filterByName)

	var r []models.Result
	r, remainder = filter(query, remainder, filterHighPrio)
	results = append(results, r...)

	r, remainder = filter(query, remainder, filterLowPrio)
	results = append(results, r...)

	return results
}

func filter(query string, source []models.Result, f func(string, models.Result) bool) ([]models.Result, []models.Result) {
	results := []models.Result{}
	remainder := []models.Result{}

	for _, v := range source {
		if f(query, v) {
			results = append(results, v)
		} else {
			remainder = append(remainder, v)
		}
	}

	return results, remainder
}

func filterByName(query string, r models.Result) bool {
	return strings.Contains(normalize(r.Name), query)
}

func filterHighPrio(query string, r models.Result) bool {
	return strings.Contains(r.Searchable, query)
}

func filterLowPrio(query string, r models.Result) bool {
	return fuzzy.Match(query, r.Searchable)
}
