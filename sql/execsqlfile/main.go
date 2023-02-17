package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	port     = flag.Int("port", 1433, "The sqlserver port")
	host     = flag.String("host", "", "The sqlserver host")
	user     = flag.String("user", "", "The sqlserver user")
	password = flag.String("password", "", "The sqlserver password")
	sqlFile  = flag.String("sqlfile", "", "The sqlserver password")
)

func main() {

	flag.Parse()

	if flag.NFlag() != 5 {
		log.Fatalln(flag.ErrHelp)
	}

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?encrypt=disable&timeout=5", *user, *password, *host, *port)

	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		log.Fatal("err1 >>>", err)
	}

	defer db.Close()

	b, err := os.ReadFile(*sqlFile)
	if err != nil {
		log.Fatal("err2 >>>", err)
	}

	r, err := db.Exec(string(b))
	if err != nil {
		log.Fatal("err3 >>>", err)
	}

	rs, err := r.RowsAffected()
	if err != nil {
		log.Fatal("err4 >>>", err)
	}

	log.Printf("RowsAffected %d", rs)
}
