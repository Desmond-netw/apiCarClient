package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Parse all templates once
var templates = template.Must(
	template.ParseGlob("templates/*.html"),
)

// Handling list of API
// Making sure
// JSON API data ->Go struct
// Go struct - HTML Templates

// -------------------------
// home page handler
// -------------------------
func homeHandler(wr http.ResponseWriter, req *http.Request) {

	var filteredCategories []Category
	var filteredManufacturers []Manufacturer

	cat, modes, manus := LoadData()

	//filter available cars by category
	for _, category := range cat {
		for _, carModel := range modes {
			if carModel.CategoryId == category.ID {
				filteredCategories = append(filteredCategories, category)
				break
			}
		}
	}

	// filter available cars by manufactures
	for _, manufacture := range manus {
		for _, carModel := range modes {
			if carModel.ManufacturerID == manufacture.ID {
				filteredManufacturers = append(filteredManufacturers, manufacture)
				break
			}
		}
	}

	// passing filtered
	data := homePageData{
		Title:                 "Home",
		Page:                  "home",
		FilteredCategories:    filteredCategories,
		FilteredManufacturers: filteredManufacturers,
	}
	wr.WriteHeader(http.StatusOK)
	renderTemplate(wr, "layout.html", data)
}

// Model handler to send model data to html templ
func modelsHandler(wr http.ResponseWriter, req *http.Request) {
	models, err := fetchModel()
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	manufacturers, _ := fetchManufacturers()
	categories, _ := fetchCategory()

	manMap := manufacturerMap(manufacturers)
	catMap := categoryMap(categories)

	// join data here
	for i := range models {
		models[i].ManufacturerName = manMap[models[i].ManufacturerID]
		models[i].CategoryName = catMap[models[i].CategoryId]
	}

	fmt.Println("Models fetchd", len(models))
	data := PageData{
		Title:  "Car Models",
		Page:   "models",
		Models: models,
	}

	renderTemplate(wr, "layout.html", data)
}

// --------------------------------------------------
// Manufacturer  Page data
// ---------------------------------------------
func manufacturerHandler(wr http.ResponseWriter, req *http.Request) {
	// geting manuf data
	manufacturer, err := fetchManufacturers()
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Manufacturers fetch", len(manufacturer)) // sending signal to cli
	data := PageData{
		Title:         "Car Manufacturers",
		Page:          "manufacturers",
		Manufacturers: manufacturer,
	}

	// render template to html output
	renderTemplate(wr, "layout.html", data)

}

// ------------------
// Categories Page
// -------------------
func categoryHandler(wr http.ResponseWriter, req *http.Request) {
	category, err := fetchCategory()
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("category fetch", len(category))
	data := PageData{
		Title:      "Car Category",
		Page:       "categories",
		Categories: category,
	}

	renderTemplate(wr, "layout.html", data)

}

// func to handle template rendering
func renderTemplate(wr http.ResponseWriter, tmpl string, data interface{}) {
	//passing specific html templates
	err := templates.ExecuteTemplate(wr, tmpl, data)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
}
