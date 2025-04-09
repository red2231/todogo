package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)
func insert(t Task) {
db := sql.OpenDB()
}
func post(w http.ResponseWriter, r *http.Request ){
values, err := io.ReadAll(r.Body)
if err!=nil{
	log.Fatal("Erro!")
	return
}
var task Task
err = json.Unmarshal(values, &task)
if err !=nil{
	log.Fatal("Erro")
	return
}
timer, err :=time.Parse("03/06/2006 15:04:05", task.scheduled)
data :=timer.String()
task.scheduled = data

}
func main(){
	http.HandleFunc("POST /task", post)
	http.ListenAndServe(":5000", nil)

}
type Task struct{
	 scheduled string `json:"scheduled"`
	 name string `json:"nome"`
	 description string `json:"description"`
}