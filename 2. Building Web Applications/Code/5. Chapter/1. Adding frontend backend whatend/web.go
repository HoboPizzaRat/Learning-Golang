package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	DBHost  = "127.0.0.1"
	DBPort  = ":3306"
	DBUser  = "root"
	DBPass  = "KukkaKekkonen1234"
	DBDbase = "testdb"
	PORT    = ":8080"
)

var database *sql.DB

type Page struct {
	Title      string
	RawContent string
	Content    template.HTML
	Date       string
	GUID       string
}

func (p Page) TruncatedText() string {
	if len(p.RawContent) > 150 {
		return p.RawContent[:150] + " ..."
	}
	return p.RawContent
}

func RedirIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", 301)
}
func ServePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	thisPage := Page{}
	fmt.Println(pageGUID)

	err := database.
		QueryRow("SELECT page_title,page_content,page_date FROM pages WHERE page_guid=?", pageGUID).
		Scan(&thisPage.Title, &thisPage.RawContent, &thisPage.Date)

	thisPage.Content = template.HTML(thisPage.RawContent)

	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("Couldn't get page!")
		return
	}

	t, _ := template.ParseFiles("templates/blog.html")
	t.Execute(w, thisPage)
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	var Pages = []Page{}

	pages, err := database.Query("SELECT page_title, page_content, page_guid, page_date FROM pages ORDER BY page_date DESC")

	if err != nil {
		fmt.Fprintln(w, err.Error)
	}
	defer pages.Close()
	for pages.Next() {
		thisPage := Page{}
		err := pages.Scan(
			&thisPage.Title,
			&thisPage.Content,
			&thisPage.Date,
			&thisPage.GUID,
		)
		if err != nil {
			fmt.Println("SCAN ERROR", err)
		}
		Pages = append(Pages, thisPage)
	}
	fmt.Println("Pages found:", len(Pages))

	for _, page := range Pages {
		fmt.Println(page.Title)
		fmt.Println(page.RawContent)
	}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, Pages)
}
func APIPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	thisPage := Page{}
	fmt.Println(pageGUID)
	err := database.
		QueryRow("SELECT page_title,page_content,page_date FROM pages WHERE page_guid=?", pageGUID).
		Scan(&thisPage.Title, &thisPage.RawContent, &thisPage.Date)

	thisPage.Content = template.HTML(thisPage.RawContent)
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println(err)
		return
	}
	APIOutput, err := json.Marshal(thisPage)
	fmt.Println(APIOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, thisPage)
}

func main() {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s%s)/%s", DBUser, DBPass, DBHost, DBPort, DBDbase)
	fmt.Println(dbConn)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("Couldn't connect to" + DBDbase)
		log.Println(err.Error)
	}
	database = db
	routes := mux.NewRouter()
	routes.HandleFunc("/api/pages", APIPage).
		Methods("GET").
		Schemes("https")
	routes.HandleFunc("/api/pages/{guid:[0-9a-zA\\-]+}", APIPage).
		Methods("GET").
		Schemes("https")
	routes.HandleFunc("/page/{guid:[0-9a-zA-Z\\-]+}", ServePage)
	routes.HandleFunc("/", RedirIndex)
	routes.HandleFunc("/home", ServeIndex)
	http.Handle("/", routes)
	http.ListenAndServe(PORT, nil)
}
