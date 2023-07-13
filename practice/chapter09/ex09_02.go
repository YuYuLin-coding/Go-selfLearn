package main

import (
	"errors"
	"log"
)

// func main() {
// 	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
// 	log.Println("Start of our app")
// 	err := errors.New("application aborted!")

// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Println("End of our app")
// }

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile | log.Lmsgprefix)
	log.Println("Start of our app")
	err := errors.New("application aborted")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("End of our app")
}
