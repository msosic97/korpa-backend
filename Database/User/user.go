//Logika vezana za usera i bazu

package db

import (
	"log"
	"time"

	"github.com/msosic97/korpa-backend/models"
	"github.com/msosic97/korpa-backend/routes/helper"
	"github.com/msosic97/korpa-backend/tools/db"
)

// Registration a new user in DB
func RegisterUser(user models.User) error {

	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
		return err
	}

	// Insert the user into the database

	// Prepare the call to the stored procedure
	stmt, err := db.DB.Prepare("CALL InsertUserDB(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("Failed to prepare statement for insert of user:", err)
		return err
	}
	defer stmt.Close()

	// Execute the stored procedure
	_ , err = stmt.Exec(user.Username, hashedPassword, user.Email, user.FirstName, user.LastName, user.Phone, time.Now().Unix())
	if err != nil {
		log.Printf("Failed to create new user: %v", err)
		return err
	}

	
	return nil
}
