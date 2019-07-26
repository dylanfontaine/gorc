package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := getTargetURL()
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var responseData map[string]interface{}
	err = json.Unmarshal([]byte(responseBody), &responseData)

	if err != nil {
		panic(err)
	}

	fmt.Println("Response:")
	fmt.Println(responseData)

}

func getTargetURL() string {
	return "https://xkcd.com/936/info.0.json"
}
