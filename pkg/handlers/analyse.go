package handlers

import (
	"html/template"
	"net/http"
	"web-page-analyser/pkg/analyse"
	"web-page-analyser/pkg/model"
)

type AnalysisResultPage struct {
	Title    string
	Url      string
	Result   *model.Analysis
	IsResult bool
}

func AnalyseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pageURL := r.Form.Get("url")
	if pageURL == "" {
		http.Error(w, "Page URL not provided", http.StatusBadRequest)
		return
	}

	result, statusCode, err := analyse.AnalyseUrl(r.Form.Get("url"))
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	p := AnalysisResultPage{
		Title:    "Analysis",
		Url:      pageURL,
		Result:   result,
		IsResult: true,
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
	return
}
