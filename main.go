package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "omni-notes")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT title, content FROM notes")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var title string
		var content string
		err := rows.Scan(&title, &content)
		if err != nil {
			log.Fatal(err)
		}

		if title == "" {
			title = uuid.New().String()
		}

		filename := title + ".md"
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = file.WriteString(content)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Created file %s\n", filename)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
