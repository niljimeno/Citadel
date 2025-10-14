package repository

import (
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/niljimeno/citadel/models"
)

func normalize(s string) string {
	return strings.ToLower(s)
}

func Search(r string) []models.Result {
	contains := func(s, sub string) bool {
		return strings.Contains(normalize(s), normalize(sub))
	}

	containsTag := func(tags []string, query string) bool {
		for _, tag := range tags {
			if contains(query, tag) {
				return true
			}
		}

		return false
	}

	resembles := func(sub, s string) bool {
		return fuzzy.Match(normalize(sub), normalize(s))
	}

	tagResembles := func(tags []string, query string) bool {
		for _, tag := range tags {
			if resembles(tag, query) || resembles(query, tag) {
				return true
			}
		}

		return false
	}

	results := []models.Result{}
	highPriority := []models.Result{}

	// high priority search
	for _, v := range db {
		if contains(v.Name, r) {
			results = append(results, v)
		} else {
			highPriority = append(highPriority, v)
		}
	}

	lowPriority := []models.Result{}

	for _, v := range highPriority {
		if contains(v.Desc, r) || containsTag(v.Tags, r) {
			results = append(results, v)
		} else {
			lowPriority = append(lowPriority, v)
		}
	}

	// low priority search
	for _, v := range lowPriority {
		if resembles(r, v.Name) ||
			resembles(r, v.Desc) ||
			tagResembles(v.Tags, r) {
			results = append(results, v)
		}
	}

	return results
}

func FilterByTag(q string) []models.Result {
	results := []models.Result{}

	for _, r := range db {
		for _, tag := range r.Tags {
			if tag == q {
				results = append(results, r)
				break
			}
		}
	}

	return results
}
