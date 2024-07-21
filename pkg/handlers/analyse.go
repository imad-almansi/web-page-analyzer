package handlers

import (
	"net/http"
)

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

	err = analyseUrl(r.Form.Get("url"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("Analyse: " + r.Form.Get("url")))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func analyseUrl(pageURL string) error {
	return nil
}
