package main

import (
	"net/http"
)

//	func hello(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "Hello World!")
//	}
type Refer struct {
	handler http.Handler
	refer   string
}

func (this *Refer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Referer() == this.refer {
		this.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is handler"))
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	// server := &http.Server{
	// 	Addr: "0.0.0.0:80",
	// }
	// http.HandleFunc("/", hello)
	// server.ListenAndServe()
	// referer := &Refer{
	// 	handler: http.HandlerFunc(myHandler),
	// 	refer:   "www.shirdon.com",
	// }
	http.HandleFunc("/hello", SayHello)
	// http.ListenAndServe(":8080", referer)
	http.ListenAndServe(":8080", nil)
}
