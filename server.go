package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"
)

var filename string = "data.csv"

type URLData struct {
	shortID string
	longURL string
}

func main() {
	mustBeRoot()
	urlMappings := loadURLMappings()

	for _, data := range urlMappings {
		path := fmt.Sprintf("/%s", data.shortID)
		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			handleRedirect(w, r, data.longURL)
		})
	}

	fmt.Println("Server running at http://l.sh. Short URLs are now working.")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		handleError(err)
	}
}

func loadURLMappings() []URLData {
	data := readCsv()[1:]

	var urlMappings []URLData
	for _, row := range data {
		urlMappings = append(urlMappings, URLData{
			shortID: row[2],
			longURL: row[3],
		})
	}

	return urlMappings
}

func handleRedirect(w http.ResponseWriter, r *http.Request, longURL string) {
	ipAddr := r.RemoteAddr
	timestamp := time.Now().Format(time.RFC3339)

	fmt.Printf("[%s] Source IP: %s | Request URL: %s | Redirecting to: %s\n", timestamp, ipAddr, r.URL.Path, longURL)

	http.Redirect(w, r, longURL, http.StatusFound)
}

func readCsv() [][]string {
	file, err := os.Open(filename)
	if err != nil {
		handleError(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	content, err := csvReader.ReadAll()
	if err != nil {
		handleError(err)
	}

	if len(content) == 0 {
		fmt.Println("Error: The CSV file is empty.")
		os.Exit(1)
	}
	return content
}

func mustBeRoot() {
	if os.Getuid() != 0 {
		fmt.Println("Error: Server must be run as root.")
		os.Exit(1)
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
