package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/tidwall/gjson"
)

func main() {
	subReddit := parseFlags()

	fmt.Printf("Fetching data from r/%s\n", subReddit)

	getPosts(subReddit)
}

func parseFlags() (ret string) {
	subReddit := flag.String("sub", "", "Sub Reddit")
	flag.Parse()

	if len(*subReddit) == 0 {
		fmt.Fprintf(os.Stderr, "You must specify a Sub Reddit")
	}

	return *subReddit
}

func getPosts(subReddit string) {
	url := fmt.Sprintf("https://www.reddit.com/r/%s.json", subReddit)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	value := gjson.Get(sb, "data.children.#.data.title")

	for i, title := range value.Array() {
		fmt.Printf("%d - %s\n", i, title)
	}
}
