package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // driver mysql
)

// Config ...
// local config
type Config struct {
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
	Database struct {
		Host     string `json:"host"`
		Password string `json:"password"`
		User     string `json:"user"`
		Port     string `json:"port"`
		Dbname   string `json:"dbname"`
	} `json:"database"`
	JWT struct {
		SigningKey       string `json:"signing_key"`
		SigningAlgorithm string `json:"signing_algorithm"`
	} `json:"jwt"`
}

// New ...
func (c *Config) New(file string) {
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(c)
}

// ConnectDB ...
func (c *Config) ConnectDB() (db *gorm.DB, err error) {
	connectionString := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", c.Database.User, c.Database.Password, c.Database.Dbname)
	db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.DB().SetMaxOpenConns(5)

	return db, err
}
