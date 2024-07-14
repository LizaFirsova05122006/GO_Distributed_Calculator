package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Knetic/govaluate"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbPath = "users.db"
	jwtKey = "secret_key"
)

var (
	tmpl      = template.Must(template.ParseFiles("login.html"))
	dbPaths   = "rezults.db"
	userDB    *gorm.DB
	jwtSecret = []byte("super-secret")
	resultDB  *gorm.DB
)

type Request struct {
	Email string `json:"Email"`
}

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Email    string `not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Token    string `gorm:"not null" json:"token"`
}

type Result struct {
	ID         uint `gorm:"primaryKey"`
	Email      string
	Expression string
	Result     string
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

func main() {
	r := mux.NewRouter()
	initDatabase()
	http.HandleFunc("/api/v1/register", register)
	http.HandleFunc("/api/v1/login", login)
	http.HandleFunc("/api/v1/calculate", calculate)
	http.HandleFunc("/api/v1/results", results)
	r.HandleFunc("/user/{email}", userHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/authorized", authHandler)
	http.HandleFunc("/login", tokenHandler)
	http.HandleFunc("/user/", userHandler)
	http.HandleFunc("/calculate/", calculateHandler)
	http.HandleFunc("/postform", postformHandler)
	http.HandleFunc("/rezults/", rezultsHandler)
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func isExpressionValid(expression string) bool {
	// Проверяем, чтобы строка не была пустой
	if expression == "" {
		return false
	}

	// Проверяем, чтобы выражение не было числовым значением
	if _, err := strconv.ParseFloat(expression, 64); err == nil {
		return false
	}

	return true
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Println("Method not allowed")
		return
	}

	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Println("Bad request:", err)
		return
	}

	if credentials.Email == "" || credentials.Password == "" {
		http.Error(w, "Login and password cannot be empty", http.StatusBadRequest)
		log.Println("Login and password cannot be empty")
		return
	}

	var existingUser User
	if !userDB.Where("email = ?", credentials.Email).First(&existingUser).RecordNotFound() {
		http.Error(w, "Login already exists", http.StatusConflict)
		log.Println("Login already exists")
		return
	}

	token, err := generateJWT(credentials.Email)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error generating token:", err)
		return
	}

	newUser := User{
		Email:    credentials.Email,
		Password: credentials.Password,
		Token:    token,
	}

	if err := userDB.Create(&newUser).Error; err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		log.Println("Could not create user:", err)
		return
	}

	message := fmt.Sprintf("OK. Registration was successful.")
	response := RegisterResponse{
		Message: message,
		Token:   token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Email == "" || req.Password == "" {
		http.Error(w, "Invalid request: both login and password are required", http.StatusBadRequest)
		return
	}

	var user User
	result := userDB.Where("email = ? AND password = ?", req.Email, req.Password).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Invalid login or password", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	message := fmt.Sprintf("OK. Authorization was successful.")
	response := LoginResponse{
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func calculate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Println("Метод не разрешен", r.Method)
		return
	}

	var request struct {
		Email      string `json:"email"`
		Expression string `json:"expression"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Println("Некорректный запрос:", err)
		return
	}

	var user User
	if err := userDB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		log.Println("Пользователь не найден:", request.Email)
		return
	}

	if !isExpressionValid(request.Expression) {
		http.Error(w, "Invalid input data, not a valid expression", http.StatusBadRequest)
		log.Println("Invalid input data, not a valid expression")
		return
	}

	result, err := calculateExpression(request.Expression)
	if err != nil {
		http.Error(w, "Invalid expression", http.StatusBadRequest)
		log.Println("Некорректное выражение:", err)
		return
	}

	newResult := Result{
		Email:      request.Email,
		Expression: request.Expression,
		Result:     result, // Конвертация в строку
	}

	if err := resultDB.Create(&newResult).Error; err != nil {
		http.Error(w, "Could not save result", http.StatusInternalServerError)
		log.Println("Could not save result", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Expression accepted successfully"}`))
}

func results(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	results, err := queryDatabase(req.Email)
	if err != nil {
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}

	if len(results) == 0 {
		http.Error(w, "No results found for the given email", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func queryDatabase(email string) ([]Result, error) {
	db, err := sql.Open("sqlite3", "./results.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT expression, result FROM results WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []Result
	for rows.Next() {
		var result Result
		if err := rows.Scan(&result.Expression, &result.Result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func initDatabase() {
	var err error
	userDB, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		panic("failed to connect to database")
	}
	userDB.LogMode(true)
	userDB.AutoMigrate(&User{})
	resultDB, err = gorm.Open("sqlite3", "rezults.db")
	if err != nil {
		log.Fatalf("failed to connect to results database: %v", err)
	}
	if err := resultDB.AutoMigrate(&Result{}).Error; err != nil {
		log.Fatalf("failed to migrate results database: %v", err)
	}
}

func rezultsHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	log.Println("Received request for email:", email)

	db, err := sql.Open("sqlite3", dbPaths)
	if err != nil {
		log.Println("Error opening database:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	expressions, err := getExpressionsByEmail(db, email)
	if err != nil {
		log.Println("Error getting expressions:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Expressions found:", expressions)

	tmpl := template.Must(template.ParseFiles("rezult.html"))

	data := struct {
		Email       string
		Expressions []Expression
	}{
		Email:       email,
		Expressions: expressions,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getExpressionsByEmail(db *sql.DB, email string) ([]Expression, error) {
	query := "SELECT Email, Expression, Result FROM results WHERE Email = ?"
	rows, err := db.Query(query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expressions []Expression
	for rows.Next() {
		var exp Expression
		err := rows.Scan(&exp.Email, &exp.Expression, &exp.Result)
		if err != nil {
			return nil, err
		}
		expressions = append(expressions, exp)
	}

	return expressions, nil
}

type Expression struct {
	Email      string
	Expression string
	Result     string
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "calculate.html")
}

func postformHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	expression := r.FormValue("expression")

	if email == "" || expression == "" {
		http.Error(w, "Both email and expression are required", http.StatusBadRequest)
		return
	}

	result, err := calculateExpression(expression)
	if err != nil {
		http.Error(w, "Invalid expression", http.StatusBadRequest)
		return
	}

	db, err := sql.Open("sqlite3", dbPaths)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	err = createTables(db)
	if err != nil {
		log.Printf("Error creating tables: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Сохранить исходное выражение и результат в базу данных
	err = insertExper(db, email, expression, result)
	if err != nil {
		log.Printf("Error inserting expression: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.ServeFile(w, r, "postform.html")
}

func calculateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}

	result, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result), nil
}

func insertExper(db *sql.DB, email, expression, result string) error {
	query := "INSERT INTO results (Email, Expression, Result) VALUES (?, ?, ?)"
	_, err := db.Exec(query, email, expression, result)
	return err
}

// Создание таблиц, если они не существуют
func createTables(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS results (
        ID INTEGER PRIMARY KEY AUTOINCREMENT,
        Email TEXT NOT NULL,
        Expression TEXT NOT NULL,
        Result TEXT NOT NULL
    );`
	_, err := db.Exec(query)
	return err
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "register.html")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	createTable(db)

	exists, err := userExists(db, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	token, err := generateJWT(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = insertUser(db, email, password, token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User %s registered successfully!\nYour JWT Token: %s\nПерейдите на главную для входа: http://localhost:8080/login", email, token)
}

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderTokenTemplate(w, "", false)
		return
	}

	if r.Method == http.MethodPost {
		token := r.FormValue("token")

		if token == "" {
			http.Redirect(w, r, "http://127.0.0.1:5000/forgot", http.StatusSeeOther)
			return
		}

		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()
		var email string
		err = db.QueryRow("SELECT Email FROM users WHERE Token=?", token).Scan(&email)
		if err == sql.ErrNoRows {
			http.Redirect(w, r, "http://127.0.0.1:5000/forgot", http.StatusSeeOther)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/user/%s?email=%s", email, email), http.StatusSeeOther)
		return
	}

	http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
}

// Функция для рендера HTML-шаблона с формой
func renderTokenTemplate(w http.ResponseWriter, errorMessage string, showError bool) {
	tmpl.Execute(w, struct {
		ShowError    bool
		ErrorMessage string
	}{
		ShowError:    showError,
		ErrorMessage: errorMessage,
	})
}

// Пример создаст типичный маршрут с email в пути
func userHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	data := struct {
		Email string
	}{
		Email: email,
	}

	tmpll := template.Must(template.ParseFiles("user.html"))

	tmpll.Execute(w, data)
}

func createTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        ID INTEGER PRIMARY KEY AUTOINCREMENT,
        Email TEXT NOT NULL UNIQUE,
        Password TEXT NOT NULL,
        Token TEXT NOT NULL
    );`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

func insertUser(db *sql.DB, email, password, token string) error {
	query := "INSERT INTO users (Email, Password, Token) VALUES (?, ?, ?)"
	_, err := db.Exec(query, email, password, token)
	return err
}

func userExists(db *sql.DB, email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE Email=?)"
	err := db.QueryRow(query, email).Scan(&exists)
	return exists, err
}

func generateJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   jwt.TimeFunc().Add(time.Hour * 87600).Unix(), // Token будет действовать 10 лет
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
