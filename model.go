package main

import (
	"fmt"
)

type County struct {
	ID     int
	CityID int
	Name   string
}

type Town struct {
	ID         int
	CountyId   int
	Name       string
	Latitude   float64
	Longitude  float64
	PostalCode int
}

func GetCounty(city_id int, name string) (*County, error) {

	query := fmt.Sprintf("SELECT * FROM counties WHERE city_id=%d AND name='%s' limit 1;", city_id, name)
	//log.Println("QUERY: ", query)

	// prepare
	stmt, err := Conn.Prepare(query)
	if err != nil {
		return nil, err
	}

	// query
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = stmt.Close()
		_ = rows.Close()
	}()

	county := &County{}
	for rows.Next() {
		if err := rows.Scan(&county.ID, &county.CityID, &county.Name); err != nil {
			return nil, err
		}
	}

	return county, nil
}

func GetTown(county_id int, name string) (*Town, error) {

	query := fmt.Sprintf("SELECT * FROM towns WHERE county_id=%d AND name='%s' limit 1;", county_id, name)
	//log.Println("QUERY: ", query)

	// prepare
	stmt, err := Conn.Prepare(query)
	if err != nil {
		return nil, err
	}

	// query
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = stmt.Close()
		_ = rows.Close()
	}()

	town := &Town{}
	for rows.Next() {
		if err := rows.Scan(&town.ID, &town.CountyId, &town.Name, &town.Latitude, &town.Longitude, &town.PostalCode); err != nil {
			return nil, err
		}
	}

	return town, nil
}

func TownExists(county_id int, name string) (bool, error) {
	exists := 0

	query := fmt.Sprintf("SELECT count(*) FROM towns WHERE county_id=%d AND name='%s' limit 1;", county_id, name)
	//log.Println("QUERY: ", query)

	// prepare
	stmt, err := Conn.DB.Prepare(query)
	if err != nil {
		return false, err
	}

	// query
	rows, err := stmt.Query()
	if err != nil {
		return false, err
	}
	defer func() {
		_ = stmt.Close()
		_ = rows.Close()
	}()
	for rows.Next() {
		if err := rows.Scan(&exists); err != nil {
			return false, err
		}
	}
	return exists > 0, nil
}

func TownCreate(county_id int, postal_code, name string) error {

	query := "INSERT INTO towns (county_id, name, postal_code) VALUES ($1, $2, $3) RETURNING id;"

	stmt, err := Conn.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(county_id, name, postal_code)
	return err
}
