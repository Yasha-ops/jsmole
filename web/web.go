package web

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type WebQuery struct {
	Url  string
	Body []byte
}

func isURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	return err == nil
}

func CreateNewWebQuery(url string) (*WebQuery, error) {
	if !isURL(url) {
		return nil, fmt.Errorf("not a valid url")
	}

	webQuery := WebQuery{Url: url}
	return &webQuery, nil
}

func (webQuery *WebQuery) getScriptUrls() []string {
	// Retrieving js url
	scriptUrlsTags := regexp.MustCompile(`[^'"]+\.js`).FindAllStringSubmatch(string(webQuery.Body), -1)

	// Check script urls
	var scriptUrls []string
	for _, tag := range scriptUrlsTags {
		src := tag[0]

		if src[0] == '/' {
			src = fmt.Sprintf("%s%s", webQuery.Url, src)
		}

		scriptUrls = append(scriptUrls, src)
	}

	return scriptUrls
}

func isMapAvailabe(url string) bool {

	if !strings.HasPrefix(url, ".map") {
		url = url + ".map"
	}

	_, err := http.Get(url)

	if err != nil {
		fmt.Println("Website not reachable", err)
		return false
	}

	return true
}

func (webQuery *WebQuery) GetMaps() ([]string, error) {
	resp, err := http.Get(webQuery.Url)
	if err != nil {
		return []string{}, fmt.Errorf("website not reachable %v", err)
	}

	webQuery.Body, err = io.ReadAll(resp.Body)
	if err != nil {
		return []string{}, fmt.Errorf("can't readall %v", err)
	}

	scriptUrls := webQuery.getScriptUrls()

	var mappedScripts []string
	for _, url := range scriptUrls {
		if isMapAvailabe((url)) {
			mappedScripts = append(mappedScripts, url+".map")
		}
	}

	defer resp.Body.Close()
	return mappedScripts, nil
}
