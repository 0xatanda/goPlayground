package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"crypto/tls"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

const (
	DBHost  = "127.0.0.1"
	DBPort  = ":3306"
	DBUser  = "root"
	DBPass  = "ethereum"
	DBDbase = "cms"
	PORT    = ":8080"
	MQHost  = "127.0.0.1"
	MQPort  = ":5672"
)

var database *sql.DB
var sessionStore = sessions.NewCookieStore([]byte("our-social-network-application"))
var UserSession Session
var WelcomeTitle = "You've successfully registered!"
var WelcomeEmail = "Welcome to our CMS, {{Email}}! We're glad you could join us."
var channel *amqp.Channel

// var conn *amqp.Connection

type Page struct {
	Id         int
	Title      string
	Content    template.HTML
	RawContent string
	Date       string
	GUID       string
	Comments   []Comment
	Session    Session
}

type User struct {
	Id   int
	Name string
}

type Session struct {
	Id              string
	Authenticated   bool
	Unauthenticated bool
	User            User
}

type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	MaxAge     int
	RawExpires string
	Secure     bool
	HttpOnly   bool
	Raw        string
	Unparsed   []string
}

type Comment struct {
	Id          int
	Name        string
	Email       string
	CommentText string
}

type JSONResponse struct {
	Fields map[string]string
}

type RegistrationData struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

func APIPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	thisPage := Page{}

	// Correct SQL query with placeholder
	err := database.QueryRow("SELECT page_title, page_content, page_date FROM pages WHERE page_guid=?", pageGUID).Scan(&thisPage.Title, &thisPage.RawContent, &thisPage.Date)
	thisPage.Content = template.HTML(thisPage.RawContent)

	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("Error fetching page:", err)
		return
	}

	APIOutput, err := json.Marshal(thisPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Correct Content-Type
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(APIOutput))
}

func APICommentPost(w http.ResponseWriter, r *http.Request) {
	var commentAdded bool
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error())
	}
	name := r.FormValue("name")
	email := r.FormValue("email")
	comments := r.FormValue("comments")
	res, err := database.Exec("INSERT INTO comments SET comment_name=?, comment_email=?, comment_text=?", name, email, comments)
	if err != nil {
		log.Println(err.Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		commentAdded = false
	} else {
		commentAdded = true
	}
	commentAddedBool := strconv.FormatBool(commentAdded)
	var resp JSONResponse
	resp.Fields["id"] = string(id)
	resp.Fields["added"] = commentAddedBool
	jsonResp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, jsonResp)
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
	res, err := database.Exec("UPDATE comments SET comment_name=?, comment_email=?, comment_text=? WHERE comment_id=?", name, email, comments, id)
	fmt.Println(res)
	if err != nil {
		log.Println(err.Error())
	}

	var resp JSONResponse
	jsonResp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, jsonResp)
}

func servePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]

	thisPage := Page{}

	// Correct SQL query with placeholder
	err := database.QueryRow("SELECT page_title, page_content, page_date FROM pages WHERE page_guid=?", pageGUID).Scan(&thisPage.Title, &thisPage.RawContent, &thisPage.Date)
	thisPage.Content = template.HTML(thisPage.RawContent)
	if err != nil {
		// Return 404 if the page is not found
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("Couldn't get page:", err)
		return
	}

	comments, err := database.Query("SELECT id, comment_name as Name, comment_email, comment_text FROM comments WHERE page_id=?", thisPage.Id)
	if err != nil {
		log.Println(err)
	}

	for comments.Next() {
		var comment Comment
		comments.Scan(&comment.Id, &comment.Name, &comment.Email, &comment.CommentText)
		thisPage.Comments = append(thisPage.Comments, comment)
	}

	// Load template
	t, err := template.ParseFiles("templates/blog.html")
	if err != nil {
		log.Println("Couldn't load template: ", err)
		http.Error(w, "Couldn't load template", http.StatusInternalServerError)
		return
	}

	// Execute the template with page data
	err = t.Execute(w, thisPage)
	if err != nil {
		log.Println("Template execution error: ", err)
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
	for i := range p.Content {
		chars++
		if chars > 150 {
			return string(p.Content[:i]) + `...`
		}
	}
	return string(p.Content)
}

func MQPublish(message []byte) error {
	err := channel.Publish(
		"email", // exchange
		"",      // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	return err
}

func RegisterPOST(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("user_name")
	email := r.FormValue("user_email")
	pass := r.FormValue("user_password")
	pageGUID := r.FormValue("referrer")

	// Generate a user GUID based on the name
	gure := regexp.MustCompile("[^A-Za-z0-9]+")
	guid := gure.ReplaceAllString(name, "")

	// Hash the password (using a weak hash function here; replace with a strong one)
	password := weakPasswordHash(pass)

	// Insert user into the database
	_, err = database.Exec("INSERT INTERNATIONAL INTO users (user_name, user_guid, user_email, user_password) VALUES (?, ?, ?, ?)", name, guid, email, password)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Prepare welcome email content
	Email := RegistrationData{Email: email, Message: ""}
	messageTemplate, err := template.New("email").Parse(WelcomeEmail)
	if err != nil {
		http.Error(w, "Failed to prepare email template", http.StatusInternalServerError)
		return
	}
	var mbuf bytes.Buffer
	err = messageTemplate.Execute(&mbuf, Email)
	if err != nil {
		http.Error(w, "Failed to execute email template", http.StatusInternalServerError)
		return
	}

	// Convert the email message to JSON and publish to MQ
	emailMessage, err := json.Marshal(mbuf.String())
	if err != nil {
		http.Error(w, "Failed to encode email message", http.StatusInternalServerError)
		return
	}

	err = MQPublish(emailMessage)
	if err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	// Redirect the user after successful registration
	http.Redirect(w, r, "/page/"+pageGUID, http.StatusMovedPermanently)
}

// generate hash password
func weakPasswordHash(password string) []byte {
	hash := sha1.New()
	io.WriteString(hash, password)
	return hash.Sum(nil)
}

func RedirIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
}

func getSessionUID(sid string) int {
	user := User{}
	err := database.QueryRow("Select user_id FROM sessions WHERE session_id=?", sid).Scan(user.Id)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return user.Id
}

func updateSession(sid string, uid int) {
	const timeFmt = "2006-01-02T15:04:05.999999999"
	tstsmp := time.Now().Format(timeFmt)
	_, err := database.Exec("INSERT INTO sessions SET session_id=?, user_id=?, session_update=?, ON DUPLICATE KEY UPDATE user_id=?, session_update=?", sid, uid, tstsmp, uid, tstsmp)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func generateSessionId() string {
	sid := make([]byte, 24)
	_, err := io.ReadFull(rand.Reader, sid)
	if err != nil {
		log.Fatal("Could not generate session id")
	}
	return base64.URLEncoding.EncodeToString(sid)
}

func validateSession(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "app-session")
	if sid, valid := session.Values["sid"]; valid {
		currentUID := getSessionUID(sid.(string))
		updateSession(sid.(string), currentUID)
		UserSession.Id = string(currentUID)
	} else {
		newsID := generateSessionId()
		session.Values["sid"] = newsID
		session.Save(r, w)
		UserSession.Id = newsID
		updateSession(newsID, 0)
	}
	fmt.Println(session.ID)
}

func LoginPOST(w http.ResponseWriter, r *http.Request) {
	validateSession(w, r)
	u := User{}
	name := r.FormValue("user_name")
	pass := r.FormValue("user_password")
	password := weakPasswordHash(pass)
	err := database.QueryRow("SELECT user_id, user_name FROM users WHERE user_name=? and user_password=?", name, password).Scan(&u.Id, u.Name)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		u.Id = 0
		u.Name = " "
	} else {
		updateSession(UserSession.Id, u.Id)
		fmt.Fprintln(w, u.Name)
	}
}

func MQConnect() (*amqp.Connection, *amqp.Channel, error) {
	url := "amqp://" + MQHost + MQPort
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, nil, err
	}
	channel, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	if _, err := channel.QueueDeclare("", false, true, false, false, nil); err != nil {
		return nil, nil, err
	}
	return conn, channel, nil
}

func main() {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s", DBUser, DBPass, DBHost, DBDbase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Panicln("Couldn't connect to db:", err)
	}
	database = db
	defer db.Close()

	routes := mux.NewRouter()

	// Correct route and handler configuration
	routes.HandleFunc("/", RedirIndex)
	routes.HandleFunc("/home", ServeIndex)
	routes.HandleFunc("/page/{guid:[0-9a-zA\\-]+}", servePage)
	routes.HandleFunc("/api/pages", APIPage).Methods("GET").Schemes("https")
	routes.HandleFunc("/api/pages/{guid:[0-9a-zA\\-]+}", APIPage).Methods("GET").Schemes("https")
	routes.HandleFunc("/api/comments", APICommentPost).Methods("POST")
	routes.HandleFunc("/api/comments/{id:[\\w\\d\\-]+}", APICommentPut).Methods("PUT")
	routes.HandleFunc("/register", RegisterPOST).Methods("POST").Schemes("https")
	routes.HandleFunc("/login", LoginPOST).Methods("POST").Schemes("https")

	// Start the server
	http.Handle("/", routes)
	fmt.Println("Server is running on port", PORT)
	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}

	certificates, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatal("Error msg: ", err)
	}
	tlsConf := tls.Config{Certificates: []tls.Certificate{certificates}}
	tls.Listen("tcp", PORT, &tlsConf)
}
