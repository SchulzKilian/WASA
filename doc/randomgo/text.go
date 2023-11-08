package main

import (
	"fmt"
	"net/http"
)

func main() {
	i := 5
	fmt.Println(i)
	http.ListenAndServe(":8090",nil)


}

func evenRandomNumber(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","text")
}