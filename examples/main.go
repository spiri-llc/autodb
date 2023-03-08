package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spiri-llc/autodb/autoslice"
)

func main() {
	conn, err := sql.Open("mysql", "test:test@tcp(0.0.0.0:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	conn.SetConnMaxLifetime(time.Minute)
	conn.SetMaxOpenConns(100)
	conn.SetMaxIdleConns(100)

	db := autoslice.AutoDB{DB: conn}
	defer db.DB.Close()

	db.DB.Exec("INSERT INTO test (msg) VALUES ('test'),('test'),('test'),('test'),('test'),('test'),('test'),('test'),('test'),('test'),('test'),('test')")

	if err != nil {
		log.Fatal(err)
	}
	alltest := db.AutoQuery("SELECT * FROM test")
	for k := range alltest {
		log.Println(k)
		ids := db.AutoQuery("SELECT * FROM test WHERE id IN (5,7)")
		for id := range ids {
			log.Println(id)
		}
	}
}
