package main

import (
	"fmt"
	"net/http"

	"github.com/MISHRA7752/handler"
)

var PORT = ":8080"

func main() {
	http.HandleFunc("/get", handler.GetHandler)
	http.HandleFunc("/set", handler.SetHandler)
	http.HandleFunc("/delete", handler.DeleteHandler)

	fmt.Println("Server starting at ...", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println("Server failed ", err)
	}
}
