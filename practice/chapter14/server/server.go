package main

import (
	// "encoding/json"
	// "fmt"
	// "io"

	"log"
	"net/http"
	"time"
	// "strings"
)

type server struct{}

// type messageData struct {
// 	Message string `json:"message"`
// }

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// msg := `{"message": "hello world"}`
	// w.Write([]byte(msg))

	// message := messageData{}
	// err := json.NewDecoder(r.Body).Decode(&message)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(message)

	// message.Message = strings.ToUpper(message.Message)
	// jsonBytes, _ := json.Marshal(message)
	// w.Write(jsonBytes)

	// file, fileHeader, err := r.FormFile("myFile")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// fileContent, err := io.ReadAll(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = os.WriteFile(fmt.Sprintf("./%s", fileHeader.Filename), fileContent, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("%s uploaded!", fileHeader.Filename)
	// w.Write([]byte(fmt.Sprintf("%s uploaded!", fileHeader.Filename)))

	auth := r.Header.Get("Authorization")
	if auth != "superSecretToken" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Authorization token not recognized"))
		return
	}
	time.Sleep(time.Second * 2)
	msg := "Hello client!"
	w.Write([]byte(msg))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
