package main

import (
	// "bytes"
	"fmt"

	"io"
	"log"

	// "mime/multipart"
	"net/http"
	// "os"
	"time"
)

// type messageData struct {
// 	Message string `json:"message"`
// }

// func getDataAndReturnResponse() messageData {
// 	r, err := http.Get("http://localhost:8080")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer r.Body.Close()

// 	data, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	message := messageData{}
// 	err = json.Unmarshal(data, &message)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return message
// }

// func postDataAndReturnResponse(msg messageData) messageData {
// 	jsonBytes, _ := json.Marshal(msg)
// 	buffer := bytes.NewBuffer(jsonBytes)

// 	r, err := http.Post("http://localhost:8080", "application/json", buffer)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer r.Body.Close()

// 	message := messageData{}
// 	err = json.NewDecoder(r.Body).Decode(&message)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return message
// }

// func postFileAndReturnResponse(filename string) string {
// 	fileDataBuffer := bytes.Buffer{}
// 	mpWritter := multipart.NewWriter(&fileDataBuffer)

// 	file, err := os.Open(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	formFile, err := mpWritter.CreateFormFile("myFile", file.Name())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if _, err := io.Copy(formFile, file); err != nil {
// 		log.Fatal(err)
// 	}
// 	mpWritter.Close()

// 	r, err := http.Post("http://localhost:8080", mpWritter.FormDataContentType(), &fileDataBuffer)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer r.Body.Close()

// 	data, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return string(data)
// }

func getDataWithCustomOptionsAndReturnResponse() string {
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	http.DefaultClient.Timeout = time.Second * 5

	req.Header.Set("Authorization", "superSecretToken")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func main() {
	// data := getDataAndReturnResponse()

	// msg := messageData{Message: "hi server!"}
	// data := postDataAndReturnResponse(msg)
	// fmt.Println(data.Message)

	// data := postFileAndReturnResponse("./test.txt")

	data := getDataWithCustomOptionsAndReturnResponse()
	fmt.Println(data)
}
