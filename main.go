package main

import (
	"log"
	"net/http"
	"time"
)

// YourWuName : {json: with your given WuName}
type YourWuName struct {
	Message string `json:"message" `
}

func main() {
	url := "https://wunameaas.herokuapp.com/enterthewu/"

	spaceClient := http.Client{
		//max is 2 seconds
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url+"julio", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)

	if getErr != nil {
		log.Fatal(getErr)
	}

}
