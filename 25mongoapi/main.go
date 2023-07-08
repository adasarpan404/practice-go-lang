package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adasarpan404/mongo/router"
)

func main() {
	fmt.Println("Mongo Api")

	r := router.Router()

	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ...")
}
