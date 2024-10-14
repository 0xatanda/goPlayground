package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Database credentials
const (
	DBHost  = "127.0.0.1"
	DBPort  = ":3306"
	DBUser  = "root"
	DBPass  = "ethereum"
	DBDbase = "cms"
	PORT    = ":8080"
)

var database *sql.DB

// Page represents the structure of a page retrieved from the DB.
type Page struct {
	Title      string
	RawContent string
	Content    template.HTML
	Date       string
	GUID       string
}

// ServePage handles requests to serve a page based on its GUID.
func ServePage(w http.ResponseWriter, r *http.Request) {
	// Extract the GUID from the request's URL
	vars := mux.Vars(r)
	pageGUID := vars["id"]

	thisPage := Page{}

	// Query the database to get the page information
	err := database.QueryRow("SELECT page_title, page_content, page_date FROM pages WHERE page_guid=?", pageGUID).Scan(&thisPage.Title, &thisPage.RawContent, &thisPage.Date)
	thisPage.Content = template.HTML(thisPage.RawContent)
	if err != nil {
		// Return 404 if the page is not found
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("Couldn't get page:", err)
		return
	}

	// Parse and execute the HTML template
	t, err := template.ParseFiles("templates/blog.html")
	if err != nil {
		log.Println("Couldn't load template:", err)
		http.Error(w, "Couldn't load template", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, thisPage)
	if err != nil {
		log.Println("Template execution error:", err)
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
	}
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	var Pages = []Page{}
	pages, err := database.Query("SELECT page_title,page_content,page_date, page_guid FROM pages ORDER BY ? DESC", "page_date")
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
	defer pages.Close()
	for pages.Next() {
		thisPage := Page{}
		pages.Scan(&thisPage.Title, &thisPage.RawContent, &thisPage.Date, &thisPage.GUID)
		thisPage.Content = template.HTML(thisPage.RawContent)
		Pages = append(Pages, thisPage)
	}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, Pages)
}

func (p Page) TruncatedText() string {
	chars := 0
	for i, _ := range p.Content {
		chars++
		if chars > 150 {
			return string(p.Content[:i]) + `...`
		}
	}
	return string(p.Content)
}

func RedirIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
}

func main() {
	// Create the database connection string
	dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s", DBUser, DBPass, DBHost, DBDbase)

	// Open the MySQL connection
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("Couldn't connect to database:", err)
		panic(err.Error())
	}
	database = db
	defer db.Close()

	// Set up routes using gorilla/mux
	routes := mux.NewRouter()
	routes.HandleFunc("/page/{id:[0-9a-zA-Z\\-]+}", ServePage)
	routes.HandleFunc("/", RedirIndex)
	routes.HandleFunc("/home", ServeIndex)

	// Start the server
	http.Handle("/", routes)
	fmt.Println("Server is running on port", PORT)
	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
