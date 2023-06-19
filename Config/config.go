// Config stores all configuration of the application
// Here is place for loading env. variables and conf for DB


package config


import (
    _ "fmt"
	"log"
	"os"
    "github.com/joho/godotenv"
)



// Config holds configuration values of application for DB
type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
}



// LoadConfig loads the application configuration from environment variables or a .env file
func LoadConfig() (*Config, error) {
    err := godotenv.Load()
    if err != nil {
        log.Println("Error: no .env file found!")
    }

    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    return &Config{
        DBHost:     dbHost,
        DBPort:     dbPort,
        DBUser:     dbUser,
        DBPassword: dbPassword,
        DBName:     dbName,
    }, nil
}