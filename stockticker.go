package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type stocksArray []map[string]string

var (
	err      error
	stocks   stocksArray
	response *http.Response
	body     []byte
)

func main() {
	// Use http://finance.google.com/finance/info?client=ig&q=NASDAQ:GOOG to get a JSON response
	response, err = http.Get("http://finance.google.com/finance/info?client=ig&q=NASDAQ:GOOG,NASDAQ:AAPL,NASDAQ:MSFT")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	// Read the data into a byte slice
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	// Remove whitespace from response
	data := bytes.TrimSpace(body)

	// Remove leading slashes and blank space to get byte slice that can be unmarshaled from JSON
	data = bytes.TrimPrefix(data, []byte("// "))

	// Unmarshal the JSON byte slice to a predefined struct
	err = json.Unmarshal(data, &stocks)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(stocks)

}
