package main

import (
	"fmt"
	"net/http"
	"io"
	"html"
	"encoding/json"
	"os"
)

import Calendar "calendar"

func hello(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w,"hello, %q\n", html.EscapeString(r.URL.Path))
}

func post (w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	fmt.Printf("Received POST request with body: %s\n", string(body))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("POST request received sucesssfully"))
}

type Response struct {
	Message string `json:"message"`
	Status int `json:"status"`
}

func getJson(w http.ResponseWriter, r *http.Request){
	response := Response{
		Message: "This is a json response",
		Status: 200,
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed ot encode JSON response", http.StatusInternalServerError)
		return
	}
}

type RequestData struct {
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
	Age int `json:"age"`
}

func postJson(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost{
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
}

	body, err := io.ReadAll(r.Body)
	if err != nil{
		http.Error(w, "Failed to read requestbody", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var requestData RequestData
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
return
	}

	response := Response{}
	password := os.Getenv("PASSWORD")
	if requestData.Name == "Bob" && requestData.Password == password{
		fmt.Printf("Success\n")
		response.Message = fmt.Sprintf("Secret")
		response.Status = http.StatusOK
	} else {
		response.Message = "Invalid credentials"
		response.Status = http.StatusUnauthorized
		
	}

	fmt.Printf("Received JSON: %+v\n", requestData)
	fmt.Printf("%s\n",requestData.Name)
	fmt.Printf("%d\n",requestData.Age)
	fmt.Printf("%s\n", requestData.Email)



	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}

func main() {

/*	http.HandleFunc("/hello", hello)
	http.HandleFunc("/post", post)
	http.HandleFunc("/getJson", getJson)
	http.HandleFunc("/postJson", postJson)
	http.ListenAndServe(":8080", nil)*/


	var cal calendar.Calendar
	cal.Test = "hi"

	fmt.Printf(cal.test)

} 
