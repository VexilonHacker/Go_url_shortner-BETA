package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
)

var filename string = "data.csv"

func main() {
	mustBeRoot()
	ds := readCsv()[2:]
	for _, i := range ds {
		path := fmt.Sprintf("/%s", i[2])

		// Create a closure to capture the path variable
		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			// Redirect to the URL specified in i[3]
			http.Redirect(w, r, i[3], http.StatusFound)
		})
	}
	fmt.Println("http://l.sh Shorten urls are  working now")
	http.ListenAndServe(":80", nil)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func readCsv() [][]string {
	file, err := os.Open(filename)
	if err != nil {
		handleError(err)
	}
	defer file.Close()

	csv_file := csv.NewReader(file)
	content, err := csv_file.ReadAll()
	if err != nil {
		handleError(err)
	}
	if len(content) == 0 {
		fmt.Println("The CSV file is empty.")
		os.Exit(1)
	}
	return content
}

func mustBeRoot() {
	if os.Getuid() != 0 {
		fmt.Println("Error: must be run as root")
		os.Exit(1)
	}
}
