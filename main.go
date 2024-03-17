package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("hollala")

	http.HandleFunc("GET /", func (w http.ResponseWriter, r *http.Request) {
		holla := "helloword"
		w.Header().Set("Content-Type", "text/plain")
		// w.Write([]byte(holla))	

		fmt.Fprint(w, holla)
	})	
	http.ListenAndServe(":9000",nil)
	fmt.Println("am i even executing")
}
