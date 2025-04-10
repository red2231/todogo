package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)
func insert(t Task) {
db , err:= sql.Open("mysql", "root:erick@unix(/var/run/mysqld/mysqld.sock)/golang")
if err !=nil{
	log.Fatal("Erro na conexão!")
	
}
defer db.Close()
_, erro := db.Exec("insert into task(scheduled, name, description) values(?, ?, ?)", t.Scheduled, t.Name, t.Description)
if erro!=nil{
	log.Fatal(erro.Error())
}


}
func post(w http.ResponseWriter, r *http.Request ){

channel := make(chan Task)
go func(){
	channel<-normalizetask(*r)
}()
task := <-channel

go insert(task)
}


func normalizetask(i http.Request) Task{
ch := make(chan []byte)
er := make(chan error)

go func(){
	val, err:=io.ReadAll(i.Body)
	if err!=nil{
		er<-err
		log.Fatal("erro")
	}
	ch<-val
}()
values := <-ch
var task Task
err := json.Unmarshal(values, &task)
if err !=nil{
	log.Fatal("Erro")
	
}
timer, err := time.Parse("2006-01-02 15:04:05 -0700 MST", task.Scheduled.String())
if err !=nil{
	fmt.Print(err.Error())
}
task.Scheduled = timer
return task
}
func GET(w http.ResponseWriter, r *http.Request){
		vars:= mux.Vars(r)
		id, _:=strconv.Atoi(vars["id"])
		db , err:= sql.Open("mysql", "root:erick@unix(/var/run/mysqld/mysqld.sock)/golang?parseTime=true")
if err !=nil{
	log.Fatal("Erro na conexão!")
	
}
defer db.Close()
var task Task
sqll:="select scheduled, name, description from task where id =? limit 1"
values:=db.QueryRow(sqll, id)


err=values.Scan(&task.Scheduled, &task.Name, &task.Description)
if err !=nil{
	http.Error(w, "nao encontrado", http.StatusNotFound)
	log.Print("nao encontrado")
	return
}
value, err:= json.Marshal(task)
if err!=nil{
	log.Fatal("err")
}
 io.WriteString(w, string(value))
 
}
	
func main(){
	r := mux.NewRouter()
	 r.HandleFunc("/task", post).Methods("POST")
 r.HandleFunc("/task/{id}", GET ).Methods("GET")
	 http.ListenAndServe(":5000", r)
}
type Task struct{
	 Scheduled time.Time `json:"scheduled"`
	 Name string `json:"name"`
	 Description string `json:"description"`
}