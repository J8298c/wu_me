package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

}
