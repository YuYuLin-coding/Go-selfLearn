package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://TBD"
	body := strings.NewReader("{\"userId\":1,\"articleId\":1,\"comment\":\"A comment\"}")
	response, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(body)))

	if err != nil {
		fmt.Println("err", err)
	}
	b, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(b))
}
