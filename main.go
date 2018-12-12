package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"
)

// YourWuName : {json: with your given WuName}
type YourWuName struct {
	Message string `json:"message" `
}

func main() {
	url := "https://wunameaas.herokuapp.com/enterthewu/"

	resp, _ := http.Get(url + "julio")

	bytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("HTML:\n\n", string(bytes))

	resp.Body.Close()

	z := html.NewTokenizer(resp.Body)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := z.Token()

			isPara := t.Data == "p"

			if isPara {
				fmt.Println("FOund a p tag")
			}
		}
	}
}
