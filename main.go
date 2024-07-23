package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)



func main() {
	 dbHost := os.Getenv("MYSQL_DBHOST")
		
		 dbUser := os.Getenv("MYSQL_USER")
		 dbPassword := os.Getenv("MYSQL_PASSWORD")
		 dbName := os.Getenv("MYSQL_DATABASE")
	
	// Verbinde mit der MySQL-Datenbank
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Erstelle den HTTP-Handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		
		fmt.Fprintf(w, "Application version: %s\n", dbHost)
	
		fmt.Fprintf(w, "Pod IP: %s\n", dbUser)
		fmt.Fprintf(w, "Node Name: %s\n", dbPassword)
		fmt.Fprintf(w, "Pod Namespace: %s\n", dbName)
		
		
		rows, err := db.Query("SELECT * FROM sdvdemotabelle")
		if err != nil {
			fmt.Fprintf(w, "SDVDemo mit Image markusfelsner/config:1.15: ")
			
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var id int
		var name string
		var age int

		var result string

		for rows.Next() {
			err := rows.Scan(&id, &name, &age)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			result += fmt.Sprintf("ID: %d, Name: %s, Age: %d\n", id, name, age)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, result)
	})

	// Starte den HTTP-Server
	log.Println("Server gestartet, höre auf Port 8888...")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
