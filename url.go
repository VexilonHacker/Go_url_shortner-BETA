package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	filename string = "data.csv"
	surl     string = "http://l.sh"
	test     int    = 0
)

func main() {
	check_file()
	banner()
	data := readCsv()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the link here : ")
	url, err := reader.ReadString('\n')
	url = strings.TrimSpace(url)
	handleError(err)
	ch := isValidUrl(url)
	if !ch {
		fmt.Println("Invalid Url")
		os.Exit(0)
	}
	rsh_url, check := check_url_repetation(data, url)
	if check {
		rnd := random_value(data, 3)
		no, err := strconv.Atoi(data[len(data)-1][0])
		handleError(err)
		last_no := strconv.Itoa(no + 1)
		shurl := fmt.Sprintf("%s/%s", surl, rnd)
		result := []string{
			last_no,
			shurl,
			rnd,
			url,
			fmt.Sprintf("%d", time.Now().Unix()),
		}
		writeCsv(result)
		fmt.Printf("Here it is your short url : %q\n", shurl)
	} else {
		fmt.Printf("Here it is your short url : %q\n", rsh_url)
	}
}

func check_file() {
	_, err := os.Stat(filename)
	if err != nil {
		file, err := os.Create(filename)
		handleError(err)
		defer file.Close()
		info := "no,shorten_url,id,long_url,date\n1,http://example.com/test,test,https://example.com,1728850133\n"
		_, err = file.WriteString(info)
		handleError(err)
		fmt.Println("database have been created")
		file.Close()
	}
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
	conntent, err := csv_file.ReadAll()
	if err != nil {
		handleError(err)
	}
	if len(conntent) == 0 {
		fmt.Println("The CSV file is empty.")
		os.Exit(1)
	}
	return conntent
}

func writeCsv(data []string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	handleError(err)
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write(data)
	// defer file.Close()
}

func id(complexity int) string {
	random_value := ""
	if complexity != 0 {
		for i := 0; i < complexity; i++ {
			x := rand.Intn(26)
			random_value += string(rune(x+'a')) + strconv.Itoa(x)
		}
		return random_value
	} else {
		return "0"
	}
}

func in(value string, ls []string) bool {
	for _, v := range ls {
		if v == value {
			return true
		}
	}
	return false
}

func random_value(data [][]string, complexity int) string {
	rg := id(complexity)
	ids := []string{}
	var rng string

	for _, i := range data[1:] {
		ids = append(ids, i[2])
	}
	reapted_check := in(rg, ids)
	if reapted_check {
		for {
			test += 1
			rg := id(3)
			reapted_check := in(rg, ids)
			if reapted_check == false {
				rng += rg
				break
			}
		}
		return rng
	} else {
		return rg
	}
}

func check_url_repetation(data [][]string, url string) (string, bool) {
	for _, i := range data[1:] {
		urls := string(i[3])
		if url == urls {
			return string(i[1]), false
		}
	}
	return "0", true
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func banner() {
	fmt.Println(
		"\n /$$   /$$/$$$$$$$ /$$              /$$$$$$ /$$                          /$$                                        ",
	)
	fmt.Println(
		"| $$  | $| $$__  $| $$             /$$__  $| $$                         | $$                                        ",
	)
	fmt.Println(
		"| $$  | $| $$  \\ $| $$            | $$  \\__| $$$$$$$  /$$$$$$  /$$$$$$ /$$$$$$   /$$$$$$ /$$$$$$$  /$$$$$$  /$$$$$$ ",
	)
	fmt.Println(
		"| $$  | $| $$$$$$$| $$            |  $$$$$$| $$__  $$/$$__  $$/$$__  $|_  $$_/  /$$__  $| $$__  $$/$$__  $$/$$__  $$",
	)
	fmt.Println(
		"| $$  | $| $$__  $| $$             \\____  $| $$  \\ $| $$  \\ $| $$  \\__/ | $$   | $$$$$$$| $$  \\ $| $$$$$$$| $$  \\__/",
	)
	fmt.Println(
		"| $$  | $| $$  \\ $| $$             /$$  \\ $| $$  | $| $$  | $| $$       | $$ /$| $$_____| $$  | $| $$_____| $$      ",
	)
	fmt.Println(
		"|  $$$$$$| $$  | $| $$$$$$$$      |  $$$$$$| $$  | $|  $$$$$$| $$       |  $$$$|  $$$$$$| $$  | $|  $$$$$$| $$      ",
	)
	fmt.Println(
		" \\______/|__/  |__|________/       \\______/|__/  |__/\\______/|__/        \\___/  \\_______|__/  |__/\\_______|__/      \n",
	)
}
