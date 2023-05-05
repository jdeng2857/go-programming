package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Course struct {
	ID      string
	Details string
}

func GetRecords(db *sql.DB) {
	results, err := db.Query("Select * FROM Course")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var course Course
		err = results.Scan(&course.ID, &course.Details)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(course.ID, "-", course.Details)
	}
}

func InsertRecord(db *sql.DB, ID string, Details string) {
	result, err := db.Exec("INSERT INTO Course VALUES (?, ?)", ID, Details)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if count, err := result.RowsAffected(); err == nil {
			fmt.Println(count, "row(s) affected")
		}
	}
}

func EditRecord(db *sql.DB, ID string, Details string) {
	result, err := db.Exec("UPDATE Course SET Details=? WHERE ID=?", Details, ID)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if count, err := result.RowsAffected(); err == nil {
			fmt.Println(count, "row(s) affected")
		}
	}
}

func DeleteRecord(db *sql.DB, ID string) {
	result, err := db.Exec("DELETE FROM Course WHERE ID=?", ID)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if count, err := result.RowsAffected(); err == nil {
			fmt.Println(count, "row(s) affected")
		}
	}
}

func main() {
	db, err := sql.Open("mysql", "gouser:password@tcp(127.0.0.1:3306)/CoursesDB")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database object created")
		// InsertRecord(db, "BU2005", "Income Tax")
		// EditRecord(db, "BU2005", "Tax")
		DeleteRecord(db, "BU2005")
		GetRecords(db)
	}
	defer db.Close()
}
