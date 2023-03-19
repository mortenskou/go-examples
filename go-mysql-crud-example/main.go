package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	Id       int
	Headline string
	Text     string
}

/*
Requirements for this to work:

Create a database called test, a user called test and a table called data in MySQL.

Create the database:
CREATE DATABASE test;

Create a user called test with the password 'some-password' with limited privileges:
GRANT SELECT, INSERT, UPDATE, DELETE on test.* to test@'127.0.0.1' identified by 'some-password';


The folloing SQL statement can be used to create the table in MySQL:
CREATE TABLE `data` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `headline` varchar(255) DEFAULT NULL,
  `text` text DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci

*/

func main() {
	// username:password@protocol(address:port)/dbname
	db, err := sql.Open("mysql", "test:some-password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Ping the database to check if it is alive
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Insert row into the database
	stmt, err := db.Prepare("INSERT INTO data (headline, text) VALUES (?, ?)")
	if err != nil {
		panic(err)
	}

	// Execute the statement
	res, err := stmt.Exec("Hello World", "This is a test")
	if err != nil {
		panic(err)
	}

	// Get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	// Print the last inserted id
	fmt.Println("Insert id", id)

	// Select all rows in the database
	rows, err := db.Query("SELECT * FROM data")
	if err != nil {
		panic(err)
	}

	// Close the rows when we are done
	defer rows.Close()

	// Loop through the rows and add them to a slice of Data structs
	var data []Data
	for rows.Next() {
		var d Data
		err := rows.Scan(&d.Id, &d.Headline, &d.Text)
		if err != nil {
			panic(err)
		}

		data = append(data, d)
	}

	// Loop through the data and print it out one row at a time
	for _, d := range data {
		fmt.Println(d.Id, d.Headline, d.Text)
	}

	// Update a row in the database
	stmt, err = db.Prepare("UPDATE data SET headline = ? WHERE id = ?")
	if err != nil {
		panic(err)
	}

	// Execute the statement to update the row with the id 1
	res, err = stmt.Exec("Hello World - row 1", 1)
	if err != nil {
		panic(err)
	}

	// Get the number of rows affected
	affectedrows, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	// Print the number of rows affected
	fmt.Println("Rows affected", affectedrows)

	// Delete a row in the database
	stmt, err = db.Prepare("DELETE FROM data WHERE id = ?")
	if err != nil {
		panic(err)
	}

	// Execute the statement to delete the row with the id 1
	res, err = stmt.Exec(1)
	if err != nil {
		panic(err)
	}

	// Get the number of rows affected
	affectedrows, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}

	// Print the number of rows affected
	fmt.Println("Rows affected", affectedrows)

}
