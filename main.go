package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Устанавливаем соединение с MySQL
	db, _ = sql.Open("mysql", "root:AvitoTech123@tcp(localhost:3306)/db_avito")
	defer db.Close()

	http.HandleFunc("/createSegment", createSegment)
	http.HandleFunc("/deleteSegment", deleteSegment)
	http.HandleFunc("/addUserToSegment", addUserToSegment)
	http.HandleFunc("/getUserSegments", getUserSegments)

	http.ListenAndServe(":8080", nil)
}

func createSegment(w http.ResponseWriter, r *http.Request) {
	// Получаем название сегмента из параметра запроса
	r.ParseForm()
	slug := r.FormValue("slug")

	_, err := db.Exec("INSERT INTO segments (slug) VALUES (?)", slug)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating segment: %s", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Segment created: %s", slug)
}

func deleteSegment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	slug := r.FormValue("slug")

	_, err := db.Exec("DELETE FROM segments WHERE slug = ?", slug)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting segment: %s", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Segment deleted: %s", slug)
}

func addUserToSegment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userID := r.FormValue("user_id")
	addSegments := r.Form["add_segments[]"]
	removeSegments := r.Form["remove_segments[]"]

	// Добавляем сегменты
	for _, segment := range addSegments {
		_, err := db.Exec("INSERT INTO user_segments (user_id, segment_id) VALUES (?, ?)", userID, segment)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error adding user to segment: %s", err), http.StatusInternalServerError)
			return
		}
	}

	// Удаляем сегменты
	for _, segment := range removeSegments {
		_, err := db.Exec("DELETE FROM user_segments WHERE user_id = ? AND segment_id = ?", userID, segment)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error removing user from segment: %s", err), http.StatusInternalServerError)
			return
		}
	}

	fmt.Fprintf(w, "User segments updated")
}

func getUserSegments(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userID := r.FormValue("user_id")

	rows, err := db.Query("SELECT segment_id FROM user_segments WHERE user_id = ?", userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting user segments: %s", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	segments := []string{}
	for rows.Next() {
		var segmentID string
		err := rows.Scan(&segmentID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error scanning user segments: %s", err), http.StatusInternalServerError)
			return
		}
		segments = append(segments, segmentID)
	}

	fmt.Fprintf(w, "User segments: %v", segments)
}
