// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	request1()
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "welcome to go web API")
	fmt.Println("endpoint hit: homepage")
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	who := "sai duddukuri"
	fmt.Fprintf(response, "a little bit about sai")
	fmt.Println("end point hit: ", who)

}

func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
