package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rafaeltiribas/cardapio-uff/internal/config"
	"github.com/rafaeltiribas/cardapio-uff/internal/models"
	"log"
)

func OpenConnection() (*sql.DB, error) {
	connectionString := GetConnectionString()
	db, err := sql.Open(config.DbDriver, connectionString)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	return db, err
}

func GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName)
}

func InsertUser(user models.User) (id int64, err error) {
	exists, existingID, err := UserExists(user.UserID)
	if err != nil {
		log.Printf("Check if user exists error: %s\n", err)
		return -1, err
	}

	if exists {
		return existingID, nil
	}

	connection, err := OpenConnection()
	if err != nil {
		log.Printf("Open connection error: %s\n", err)
		return -1, err
	}
	defer connection.Close()

	sqlStatement := "INSERT INTO chatuser (user_id, register_date, lastuse_date) VALUES ($1, $2, $3) RETURNING id"

	err = connection.QueryRow(sqlStatement, user.UserID, user.RegisterDate, user.LastUseDate).Scan(&id)
	if err != nil {
		log.Printf("Insert error: %s\n", err)
		return -1, err
	}

	return id, nil
}

func UserExists(userID int64) (bool, int64, error) {
	connection, err := OpenConnection()
	if err != nil {
		log.Printf("Open connection error: %s\n", err)
		return false, -1, err
	}
	defer connection.Close()

	var existingID int64
	query := "SELECT id FROM chatuser WHERE user_id = $1"
	err = connection.QueryRow(query, userID).Scan(&existingID)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("User not exists: %s\n", err)
			return false, -1, nil
		}
		log.Printf("Check query error: %s\n", err)
		return false, -1, err
	}

	log.Printf("User already exists: %s\n", err)
	return true, existingID, nil
}
