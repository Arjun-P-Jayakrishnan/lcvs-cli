package api

import (
	"fmt"
	"net/http"
)

func handlerReadiness(w http.ResponseWriter,r *http.Request){
	fmt.Println("server up and running")
	respondWithJSON(w,200,struct{}{})
}