package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	db "github.com/techvisionus/golang-source/api/db/sqlc"
)

type connection struct {
	Host string
	Port string
	User string
	Password string
	DBName string
}
// Queries database
var Queries *db.Queries
// Init database
func Init() {
	err := godotenv.Load(".env")
	
	if err != nil {
		fmt.Printf("Error loading env error: %s\n", err.Error())
		return 
	}

	connInfo := connection {
		Host: os.Getenv("POSTGRES_URL"),
		Port: os.Getenv("POSTGRES_PORT"),
		User: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName: os.Getenv("POSTGRES_DB"),
	}

	if goEnv := os.Getenv("GO_ENV"); goEnv == "development" {
		connInfo.Host = "localhost"
	}

	dbSource := connectToString(connInfo)
	fmt.Printf("dbSource: %s\n", dbSource)
	conn, err := sql.Open("postgres", dbSource)
	err = conn.Ping()

	if err != nil {
		fmt.Printf("Cannot connect to db: %v", err)
	}

	fmt.Printf("Connect database success.\n")
	Queries = db.New(conn)
}

func connectToString(info connection) string{
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
	info.User, info.Password, info.Host, info.Port, info.DBName)
}