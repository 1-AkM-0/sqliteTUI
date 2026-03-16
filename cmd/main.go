package main

import (
	"log"
	"os"

	"github.com/1-AkM-0/sqliteTUI/internal/db"
)

func main() {
	_, err := db.Open(os.Getenv("DB_PATH"))
	if err != nil {
		log.Fatal(err)
	}
	/*
			_, err = db.Tables()
			if err != nil {
				log.Fatal(err)
			}

				columns, err := db.Columns("applications")
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(columns)
		rr, err := db.Execute("SELECT title FROM jobs")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(rr.Row)

	*/
}
