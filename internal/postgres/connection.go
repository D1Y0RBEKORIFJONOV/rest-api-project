package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" 
	"github.com/spf13/viper"
)

type DB struct {
	DB *sql.DB
}

func (db *DB) ConnectDB() error {
	viper.SetConfigFile("internal/postgres/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err 
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.dbname"))

	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err := dbConn.Ping(); err != nil {
		return err
	}

	db.DB = dbConn
	return nil
}
