package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func main() {
	db, err := sql.Open("mysql", "root:123321@tcp(43.138.184.202:34306)/cmdb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("err = ", err)
		return
	}

	if err := db.Ping(); err != nil {
		log.Println("err = ", err)
		return
	}

	defer db.Close()

	db.SetMaxOpenConns(10)

	s := "select name,path from cmdb.web_cachename"

	r, _ := db.Query(s)

	for r.Next() {
		var name, path string
		r.Scan(&name, &path)
		log.Println(name, path)
	}

}
