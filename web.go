package main

import (
	"fmt"
	"net/http"
)

func mainX(){
	http.HandleFunc("/", greet)
	http.ListenAndServe(PORT,nil)
}
func greet(w http.ResponseWriter, r *http.Request ){
	msg := "hello"
	fmt.Fprint(w, msg)
}