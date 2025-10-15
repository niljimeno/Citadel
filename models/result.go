package models

import (
	"strings"

	"github.com/niljimeno/citadel/utils"
)

type Result struct {
	Name       string
	Site       string
	Desc       string
	Tags       []string
	Searchable string
}

func NewResult(data []string) Result {
	r := Result{
		Name: utils.TryGet(data, 0),
		Site: utils.TryGet(data, 1),
		Desc: utils.TryGet(data, 2),
		Tags: strings.Split(utils.TryGet(data, 3), ";"),
	}

	r.Searchable = strings.ToLower(
		strings.Join(
			[]string{
				r.Name,
				r.Site,
				r.Desc,
				strings.Join(r.Tags, " ")},
			" ",
		),
	)

	return r
}
