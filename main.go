package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/xuri/excelize/v2"
)

var (
	Conn *DB
)

func init() {
	// Load Env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Load Env Error: %v", err)
	}

	Conn = &DB{}
	Conn.ConnectDatabase()
}

func main() {

	defer Conn.CloseDatabase()

	f, err := excelize.OpenFile("country.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rows, err := f.GetRows("all")
	if err != nil {
		log.Fatal(err)
	}

	var errList = []string{}
	// ADANA	ALADAĞ	AKPINAR	01720	1
	for i, row := range rows {
		if i == 0 {
			continue
		}
		// check county
		city_id, _ := strconv.Atoi(row[4])
		county, err := GetCounty(city_id, row[1])
		if county == nil || err != nil {
			errList = append(errList, fmt.Sprintf("county not found: %s", row[1]))
			continue
		}
		// check town
		townExists, err := TownExists(county.ID, row[2])
		if townExists || err != nil {
			errList = append(errList, fmt.Sprintf("town exists: %s", row[2]))
			continue
		}
		// town create
		townCreate := TownCreate(county.ID, row[3], row[2])
		if townCreate != nil {
			errList = append(errList, fmt.Sprintf("create err: %s", row[2]))
		} else {
			fmt.Printf("row: %d CREATED: %s\n", i, row[2])
		}
	}
	log.Println("ERR LIST", errList)
}
