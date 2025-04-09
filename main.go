package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)
func handleGET(w http.ResponseWriter, r *http.Request){
	body, err :=io.ReadAll(r.Body)

	if err!=nil{
		log.Fatal("Erro!")
	http.Error(w, "erro ao ler", http.StatusInternalServerError)
	return
	}
	
	defer r.Body.Close()
	fmt.Print(string(body))

}
func main(){
http.HandleFunc("POST /task", handleGET)
fmt.Print("comecou")
http.ListenAndServe(":8000", nil)

}