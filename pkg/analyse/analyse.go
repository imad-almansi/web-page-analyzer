package analyse

import (
	"errors"
	"golang.org/x/net/html"
	"log/slog"
	"net/http"
	"net/url"
	"regexp"
	"web-page-analyser/pkg/model"
)

func AnalyseUrl(pageURL string) (*model.Analysis, int, error) {
	response, err := http.Get(pageURL)
	if err != nil {
		return nil, http.StatusServiceUnavailable, errors.New("failed to reach the requested page")
	}

	defer response.Body.Close()
	if response.StatusCode > 299 {
		return nil, response.StatusCode, errors.New("response failed")
	}

	doc, err := html.Parse(response.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("failed to parse html")
	}
	analysisResult = model.Analysis{}
	analysisError = nil
	parsedPageURL, err = url.Parse(pageURL)
	if err != nil {
		slog.Error("failed to parse url", "value", pageURL)
		return nil, http.StatusBadRequest, errors.New("failed to parse url")
	}
	analyseBody(pageURL, doc)
	if analysisError != nil {
		return nil, http.StatusInternalServerError, analysisError
	}

	return &analysisResult, http.StatusOK, nil
}

var (
	parsedPageURL  *url.URL
	analysisResult model.Analysis
	analysisError  error
)

func analyseBody(pageURL string, n *html.Node) {
	switch n.Type {
	case html.DoctypeNode:
		analysisResult.Version = "5"
		analysisResult.Version, analysisError = getHtmlVersion(n)
		if analysisError != nil {
			return
		}
		break
	case html.ElementNode:
		if n.Data == "title" {
			analysisResult.Title = n.FirstChild.Data
			break
		}

		reg := regexp.MustCompile("^h\\d$")
		if reg.MatchString(n.Data) {
			if analysisResult.Headings == nil {
				analysisResult.Headings = map[string]int{
					n.Data: 1,
				}
			} else {
				analysisResult.Headings[n.Data]++
			}
			break
		}

		if n.Data == "a" {
			links, err := getLinks(n)
			if err != nil {
				analysisError = err
				return
			}
			analysisResult.Links.Internal += links.Internal
			analysisResult.Links.External += links.External
			analysisResult.Links.Inaccessible += links.Inaccessible
			break
		}

		if n.Data == "input" {
			for _, attribute := range n.Attr {
				if attribute.Key == "type" {
					if attribute.Val == "password" {
						analysisResult.HasLogin = true
					}
					break
				}
			}
			break
		}
		break
	default:
		break
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		analyseBody(pageURL, c)
	}
}

func getHtmlVersion(n *html.Node) (string, error) {
	for _, attribute := range n.Attr {
		if attribute.Key == "public" {
			reg := regexp.MustCompile("html (\\d\\.\\d+)")
			matches := reg.FindStringSubmatch(attribute.Val)
			if len(matches) != 2 {
				slog.Error("failed to find html version", "value", attribute.Val)
				return "", errors.New("failed to find html version")
			}
			return matches[1], nil
		}
	}
	return "5", nil
}

func getLinks(n *html.Node) (links model.Links, analysisErr error) {
	for _, attribute := range n.Attr {
		if attribute.Key == "href" {
			parsedURL, err := url.Parse(attribute.Val)
			if err != nil {
				slog.Error("failed to parse url", "value", attribute.Val)
				analysisErr = errors.New("failed to parse url")
				return
			}
			if parsedURL.IsAbs() {
				links.External++
				_, err := http.Head(attribute.Val)
				if err != nil {
					links.Inaccessible++
					slog.Error("inaccessible link", "link", attribute.Val)
				}
			} else {
				links.Internal++
				parsedURL.Scheme = parsedPageURL.Scheme
				parsedURL.Host = parsedPageURL.Host
				_, err := http.Head(parsedURL.String())
				if err != nil {
					links.Inaccessible++
					slog.Error("inaccessible link", "link", parsedURL.String())
				}
			}
			return
		}
	}
	return
}
