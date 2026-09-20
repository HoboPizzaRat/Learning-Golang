package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"crypto/tls"

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
	Id         int
	Title      string
	RawContent string
	Content    template.HTML
	Date       string
	Comments   []Comment
	GUID       string
}
type Comment struct {
	Id          int
	Name        string
	Email       string
	CommentText string
}

func (p Page) TruncatedText() string {
	if len(p.RawContent) > 150 {
		return p.RawContent[:150] + " ..."
	}
	return p.RawContent
}

type JSONResponse struct {
	Fields map[string]string
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
		QueryRow("SELECT id,page_title,page_content,page_date,page_guid FROM pages WHERE page_guid=?", pageGUID).
		Scan(
			&thisPage.Id,
			&thisPage.Title,
			&thisPage.RawContent,
			&thisPage.Date,
			&thisPage.GUID,
		)

	thisPage.Content = template.HTML(thisPage.RawContent)

	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("Couldn't get page!")
		return
	}
	comments, err := database.
		Query("SELECT id, comment_name as Name, comment_email, comment_text FROM comments WHERE page_id=?", thisPage.Id)
	if err != nil {
		log.Println(err)
	}

	for comments.Next() {
		var comment Comment
		comments.Scan(&comment.Id, &comment.Name, &comment.Email, &comment.CommentText)
		thisPage.Comments = append(thisPage.Comments, comment)
	}

	t, err := template.ParseFiles("templates/blog.html")
	if err != nil {
		log.Println("Couldn't parse template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	t.Execute(w, thisPage)
	if err != nil {
		log.Println("Couldn't execute template:", err)
	}
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
func APICommentPut(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error())
	}

	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)

	name := r.FormValue("name")
	email := r.FormValue("email")
	comments := r.FormValue("comments")

	res, err := database.
		Exec("UPDATE comments SET comment_name=?, comment_email=?, comment_text=? WHERE comment_id=?", name, email, comments, id)
	fmt.Println(res)
	if err != nil {
		log.Println(err.Error())
	}
	var resp JSONResponse
	jsonResp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, jsonResp)
}
func APICommentPost(w http.ResponseWriter, r *http.Request) {
	var commentAdded bool

	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error())
	}

	guid := r.FormValue("guid")
	name := r.FormValue("name")
	email := r.FormValue("email")
	comments := r.FormValue("comments")

	var pageID int
	err = database.
		QueryRow("SELECT id FROM pages WHERE page_guid=?", guid).
		Scan(&pageID)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	res, err := database.
		Exec("INSERT INTO comments SET comment_name=?, comment_email=?, comment_text=?, page_id=?, comment_id=?, comment_date=?", name, email, comments, pageID)
	if err != nil {
		log.Println(err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		commentAdded = false
	} else {
		commentAdded = true
	}

	resp := JSONResponse{
		Fields: make(map[string]string),
	}

	resp.Fields["id"] = strconv.FormatInt(id, 10)
	resp.Fields["added"] = strconv.FormatBool(commentAdded)

	jsonResp, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, jsonResp)
}
func main() {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s%s)/%s", DBUser, DBPass, DBHost, DBPort, DBDbase)
	fmt.Println(dbConn)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("Couldn't connect to" + DBDbase)
		log.Println(err.Error())
	}
	database = db
	routes := mux.NewRouter()
	routes.HandleFunc("/api/pages", APIPage).
		Methods("GET").
		Schemes("https")
	routes.HandleFunc("/api/pages/{guid:[0-9a-zA-Z\\-]+}", APIPage).
		Methods("GET").
		Schemes("https")
	routes.HandleFunc("/api/comments", APICommentPost).
		Methods("POST").
		Schemes("https")
	routes.HandleFunc("/api/comments/{id:[\\w\\d\\-]+}", APICommentPut).
		Methods("PUT").
		Schemes("https")
	routes.HandleFunc("/page/{guid:[0-9a-zA-Z\\-]+}", ServePage).
		Methods("GET").
		Schemes("https")
	routes.HandleFunc("/", RedirIndex)
	routes.HandleFunc("/home", ServeIndex)
	http.Handle("/", routes)

	certificates, err := tls.LoadX509KeyPair("certificate.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}
	tlsConf := tls.Config{Certificates: []tls.Certificate{certificates}}
	listener, err := tls.Listen("tcp", PORT, &tlsConf)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.Serve(listener, routes))
}
