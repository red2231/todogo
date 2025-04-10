package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)
func insert(t Task) {
db , err:= sql.Open("mysql", "root:erick@unix(/var/run/mysqld/mysqld.sock)/golang")
if err !=nil{
	log.Fatal("Erro na conexão!")
	return
}
defer db.Close()
_, erro := db.Exec("insert into task(scheduled, name, description) values(?, ?, ?)", t.Scheduled, t.Name, t.Description)
if erro!=nil{
	erro.Error()
}


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
timer, err := time.Parse("2006-01-02 15:04:05 -0700 MST", task.Scheduled.String())
if err !=nil{
	fmt.Print(err.Error())
}
task.Scheduled = timer
fmt.Print(task.Scheduled.String())
go insert(task)
}
func main(){
	http.HandleFunc("POST /task", post)
	http.ListenAndServe(":5000", nil)

}
type Task struct{
	 Scheduled time.Time `json:"scheduled"`
	 Name string `json:"name"`
	 Description string `json:"description"`
}