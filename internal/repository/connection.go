package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rafaeltiribas/cardapio-uff/internal/config"
	"github.com/rafaeltiribas/cardapio-uff/internal/models"
	"log"
	"time"
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

func UpdateLastUseDate(chatID int64) error {
	connection, err := OpenConnection()
	if err != nil {
		log.Printf("Open connection error: %s\n", err)
	}
	defer connection.Close()

	var lastUseDate time.Time
	sqlStatement := "SELECT lastuse_date FROM chatuser WHERE user_id = $1"
	err = connection.QueryRow(sqlStatement, chatID).Scan(&lastUseDate)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("User not found: %s\n", err)
			return err
		}
		log.Printf("Select error: %s\n", err)
		return err
	}

	now := time.Now()
	sqlStatement = "UPDATE chatuser SET lastuse_date = $1 WHERE user_id = $2"
	_, err = connection.Exec(sqlStatement, now, chatID)
	if err != nil {
		log.Printf("Update last use date error: %s\n", err)
		return err
	}

	log.Printf("Updated lastuse_date for user_id: %d\n", chatID)
	return nil
}

func NewUser(chatID int64) {
	now := time.Now()
	user := models.User{UserID: chatID, RegisterDate: now, LastUseDate: now}

	id, err := InsertUser(user)
	if err != nil {
		log.Printf("Erro ao inserir usuário: %v", err)
	}

	log.Printf("ID: %d\n", id)
}

func GetAllChatIDs() ([]int64, error) {
	connection, err := OpenConnection()
	if err != nil {
		log.Printf("Open connection error: %s\n", err)
	}

	defer connection.Close()

	var chatIDs []int64

	sqlStatement := "SELECT user_id FROM chatuser"
	rows, err := connection.Query(sqlStatement)
	if err != nil {
		log.Printf("Query error: %s\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var chatID int64
		if err := rows.Scan(&chatID); err != nil {
			log.Printf("Scan error: %s\n", err)
			return nil, err
		}
		chatIDs = append(chatIDs, chatID)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Rows error: %s\n", err)
		return nil, err
	}

	return chatIDs, nil
}
