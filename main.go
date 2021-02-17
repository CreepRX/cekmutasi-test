package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/test", printResponse)

	http.ListenAndServe(":80", r)
}

func printResponse(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method)

	bd, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("error converting body to string")

		return
	}

	bdStr := string(bd)
	fmt.Println("\n body : \n" + bdStr)

}
