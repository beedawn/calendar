package main

import (
	"fmt"
	"net/http"
	"html"
)

func hello(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w,"hello, %q\n", html.EscapeString(r.URL.Path))
}


func post (w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	fmt.Println("test")
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/post", post)

	http.ListenAndServe(":8080", nil)

} 
