package model

import "net/http"

type Page interface {
	GetTitle(r *http.Request) string
}
