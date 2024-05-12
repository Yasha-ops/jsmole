package web

import (
	"fmt"
	"io"
	"jsmole/sourcemap"
	"jsmole/utils"
	"net/http"
	"net/url"
  "crypto/tls"
	"os"
	"strings"
	"sync"

	"github.com/schollz/progressbar/v3"
)

func parseUrl(urlString string) (string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return "", fmt.Errorf("error parsing URL %v", err)
	}
	return fmt.Sprintf("%s://%s/", u.Scheme, u.Host), nil
}

func download(downloadUrl string) (*os.File, error) {
	tempFile, err := os.CreateTemp("", "main.js")
	if err != nil {
		return nil, err
	}

  http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, _ := http.NewRequest("GET", downloadUrl, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	f, _ := os.OpenFile(tempFile.Name(), os.O_CREATE|os.O_WRONLY, 0644)

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"Downloading js map",
	)

	io.Copy(io.MultiWriter(f, bar), resp.Body)

	return tempFile, nil
}

func ProcessMap(url string, output string) error {
	// Downloading the file
	mainJsFile, err := download(url)
	if err != nil {
		return err
	}

	// Retrieving the content
	mainJSContent, err := os.ReadFile(mainJsFile.Name())
	if err != nil {
		return err
	}

	// Parsing the map file
	smap, err := sourcemap.Parse(url, mainJSContent)
	if err != nil {
		return fmt.Errorf("can't parse file_data: %v", err)
	}

	sources := smap.GetSources()
	parsedUrl, err := parseUrl(url)
	if err != nil {
		return err
	}

	bar := progressbar.DefaultBytes(
		int64(len(sources)),
		"Processing sources",
	)

	// Looping through the sources
	var wg sync.WaitGroup
	for _, source := range sources {
		wg.Add(1)
		go func(source string) {
			defer wg.Done()
			sourcePath := strings.TrimPrefix(source, parsedUrl)

			sourceContent := smap.SourceContent(source)
			utils.CreateFileWithDirectories(fmt.Sprintf("%s/%s", output, sourcePath), sourceContent)
			bar.Add(1)
		}(source)
	}

	wg.Wait()
	fmt.Println(parsedUrl)

	return nil
}
