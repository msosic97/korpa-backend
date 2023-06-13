package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//--------------------------------------------------------

	db, err := sql.Open("mysql", "root:<b16b>@tcp(127.0.0.1:3306)/test")
	if err != nil {
		// Handle the error
		panic(err.Error())
	}
	defer db.Close()

	//--------------------------------------------------------

	// err = db.Ping()
	// if err != nil {
	//     // Handle the error
	// 	panic(err.Error())
	// }

	// //--------------------------------------------------------

	// rows, err := db.Query("SELECT * FROM table_name")
	// if err != nil {
	//     // Handle the error
	// }
	// defer rows.Close()

	// for rows.Next() {
	//     // Process each row of data
	// }

	// if err = rows.Err(); err != nil {
	//     // Handle the error
	// 	panic(err.Error())
	// }

	// //--------------------------------------------------------

	// Close the connection at the end
	defer db.Close()
	fmt.Println("Success!")

}


// konekcija na bazu