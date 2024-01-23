package main

import (
	"net/http"
)

func main() {
    err := http.ListenAndServe(":8090", nil)
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

/*func evenRandomNumber(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","text")
}*/