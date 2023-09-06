package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type LandingpageData struct {
	PageTitle  string
	PageHeader string
}

func main() {
	http.HandleFunc("/", loadLandingPage)
	fmt.Println("MTG is listening on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadLandingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Reloading Page")
	tmpl, err := template.ParseFiles("./gohtml/landingpage.gohtml")
	if err != nil {
		fmt.Println("-------------------------------ERROR-------------------------------")
		panic(err.Error())
	}
	data := LandingpageData{
		PageTitle:  "MTG Abfalleimer",
		PageHeader: "The awesome MTG Card-Website",
	}
	tmpl.Execute(w, data)
}
