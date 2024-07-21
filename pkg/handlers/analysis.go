package handlers

import (
	"html/template"
	"net/http"
)

type AnalysisPage struct {
	Title    string
	IsResult bool
}

func (h AnalysisPage) GetTitle(_ *http.Request) string {
	return h.Title
}

func AnalysisHandler(w http.ResponseWriter, _ *http.Request) {
	p := AnalysisPage{
		Title: "Analysis",
	}
	t, err := template.ParseFiles("pkg/pages/analysis.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
