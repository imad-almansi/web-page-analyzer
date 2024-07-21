package model

import "net/http"

type Page interface {
	GetTitle(r *http.Request) string
}

type Links struct {
	Internal     int
	External     int
	Inaccessible int
}
type Analysis struct {
	Version  string
	Title    string
	Headings map[string]int
	Links    Links
	HasLogin bool
}
