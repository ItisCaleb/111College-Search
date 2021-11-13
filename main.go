package main

import (
	"111college/routes"
	"fmt"
	"net/http"
	"strconv"
)

var port = 2500

func main() {
	loadRoutes()
	portStr := strconv.Itoa(port)
	fmt.Println("Server Start at port " + portStr)
	http.ListenAndServe(":"+portStr, nil)
}

func loadRoutes() {
	http.HandleFunc("/", routes.Root)
	http.HandleFunc("/data", routes.Data)
	http.HandleFunc("/search", routes.SearchSchool)
}
