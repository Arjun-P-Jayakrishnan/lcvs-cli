package api

import (
	"fmt"
	"net/http"
)

func handlerErr(w http.ResponseWriter,r *http.Request){
	fmt.Println("server up and running")
	respondWithError(w,400,"Something from machine went wrong")
}