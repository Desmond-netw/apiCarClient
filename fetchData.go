package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	BASE API URL : http://localhost:3000/api

	EndPoints:
	GET /models
	GET /manufacturers
	GET /categories
*/

const APIURL = "http://localhost:3000/api"

// Fetching API Data

/*
	Fetch data in Paraller & launch three goroutines,

each loading data from it's API endpoint
*/
func LoadData() ([]Category, []CarModel, []Manufacturer) {
	catChanel := make(chan []Category)
	carModelChanel := make(chan []CarModel)
	manuChanel := make(chan []Manufacturer)

	// load each data set
	go func() {
		categories, _ := fetchCategory()
		catChanel <- categories
	}()

	go func() {
		carModels, _ := fetchModel()
		carModelChanel <- carModels
	}()

	go func() {
		manufacturers, _ := fetchManufacturers()
		manuChanel <- manufacturers
	}()

	// waiting to catch responses
	cat := <-catChanel
	modes := <-carModelChanel
	manus := <-manuChanel

	return cat, modes, manus
}

// fetch car Models from json data and decode to Go struct
func fetchModel() ([]CarModel, error) {
	var models []CarModel
	resp, err := http.Get(APIURL + "/models")
	if err != nil {
		fmt.Println("Could not access API/car models")
		models = []CarModel{
			{Name: "SEVER ERROR: couldn't load cars right now ->API server down"},
		}
		return models, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&models)
	return models, err
}

/* ---------- Fetch Manufacturers data from API ------*/

func fetchManufacturers() ([]Manufacturer, error) {
	resp, err := http.Get(APIURL + "/manufacturers")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // close response body

	var manufacturers []Manufacturer                        // var to store address manuf
	err = json.NewDecoder(resp.Body).Decode(&manufacturers) // decoding json data from resp.body to Go struct
	return manufacturers, err
}

/*----------- Fetching categories data*/
func fetchCategory() ([]Category, error) {
	resp, err := http.Get(APIURL + "/categories")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // close response body

	var categories []Category
	err = json.NewDecoder(resp.Body).Decode(&categories)
	return categories, err
}

// func fetchItems(endpoint string) ([]Item, error) {
// 	baseURL := APIURL + endpoint // base url will be /api/models
// 	println("fetching url:", baseURL)
// 	// get response
// 	resp, err := http.Get(baseURL)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	// decode and read json data
// 	var items []Item
// 	err = json.NewDecoder(resp.Body).Decode(&items)

// 	println("Items fetched:", len(items)) // debug

// 	return items, err
// }
