package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/hello",helloHandler)
	http.HandleFunc("/mainpage",mainpageHandler)


	fmt.Printf("Starting server at port 8080\n")
	if err:=http.ListenAndServe(":8080",nil); err!=nil{
		log.Fatal(err)
	}
}

func mainpageHandler (w http.ResponseWriter , r *http.Request) {
	if r.URL.Path != "/mainpage"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
}

fmt.Fprintf(w, "hello")
}

func helloHandler( w http.ResponseWriter , r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w,"404 not found", http.StatusNotFound)
		return
}
fmt.Fprintf(w,"Hello")
}







