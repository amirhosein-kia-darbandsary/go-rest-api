// package main

// import (
// 	"log"

// 	"github.com/jmoiron/sqlx"
// 	_ "github.com/mattn/go-sqlite3"
// )

// func main() {
// 	// Connect to SQLite database
// 	db, err := sqlx.Connect("sqlite3", "test.db")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer db.Close()

// 	// Create the customers table
// 	createTableQuery := `
//     CREATE TABLE IF NOT EXISTS customers (
//         id INTEGER PRIMARY KEY AUTOINCREMENT,
//         name TEXT NOT NULL,
//         age INTEGER NOT NULL
//     );`
// 	db.MustExec(createTableQuery)

// 	// Insert 10 fake data entries
// 	insertDataQuery := `
//     INSERT INTO customers (name, age) VALUES 
//     ('Alice', 25),
//     ('Bob', 30),
//     ('Charlie', 22),
//     ('Diana', 28),
//     ('Eve', 35),
//     ('Frank', 40),
//     ('Grace', 27),
//     ('Hank', 33),
//     ('Ivy', 29),
//     ('Jack', 31);`
// 	db.MustExec(insertDataQuery)
// }
