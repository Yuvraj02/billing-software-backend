package handlers

import (
	"fmt"
	"net/http"
)

func GetCategories(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "All Categories will be returned from here")
	
}