package main

import (
	"fmt"
	"gotomongoimpl/routers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello Implementing MongoDB in GO") // go.mongodb.org
	r := routers.Router()
	fmt.Println("Server is getting Started....")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port : 4000")
}
