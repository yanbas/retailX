package main

import (
	"fmt"
	"log"

	"github.com/hapus_api/config"
)

func main() {
	config := config.Config{}
	config.New("./config.json")
	log.Print("Load Configuration Files")
	db, err := config.ConnectDB()
	defer db.Close()
	if err != nil {
		log.Fatal("errror db")
	}
	rows, err := db.Raw("SELECT grouping_id,group_name FROM grouping_material").Rows()

	var id string
	var name string

	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		rows.Scan(&id, &name)
	}

	fmt.Println(id)
	fmt.Println(name)

}
