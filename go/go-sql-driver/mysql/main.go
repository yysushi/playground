package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password@/database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO squarenum VALUES( ?, ? )")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	stmtOut, err := db.Prepare("SELECT squareNumber FROM squarenum WHERE number = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	for i := 0; i < 25; i++ {
		_, err = stmtIns.Exec(i, (i * i))
		if err != nil {
			panic(err.Error())
		}
	}

	var squareNum int

	err = stmtOut.QueryRow(13).Scan(&squareNum)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("The square number of 13 is: %d", squareNum)

	err = stmtOut.QueryRow(1).Scan(&squareNum)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("The square number of 1 is: %d", squareNum)
}
