package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/mborawi/heat/backend/models"
)

var db *gorm.DB

func main() {
	con := fmt.Sprintf("user=%s dbname=%s port=%s",
		"mido",
		"heat",
		"5432")
	var err error
	db, err = gorm.Open("postgres", con)
	if err != nil {
		log.Fatalln("failed to connect database", err)
	}
	defer db.Close()
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/api/list/{id:[0-9]+}", listEmployeeLeaves)
	router.HandleFunc("/api/search", SearchHandler)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../frontend/static/"))))
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "../frontend/static/index.html") })
	// router.NotFoundHandler = http.HandlerFunc(notFound)
	log.Println("Launching Server Now on 8088...")
	log.Fatal(http.ListenAndServe(":8088", router))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
}
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("query")
	if len(q) < 3 {
		http.Error(w, "Query was not found", http.StatusNotAcceptable)
		return
	}
	query := fmt.Sprintf("%%%s%%", q)
	emps := []models.Employee{}
	db.
		Where("first_name ILIKE ?", query).
		Or("last_name ILIKE ?", query).
		Limit(20).
		Find(&emps)
	type suggestion struct {
		Text  string `json:"text"`
		Value uint   `json:"value"`
	}
	suggestions := []suggestion{}
	for _, e := range emps {
		n := fmt.Sprintf("%s %s", e.FirstName, e.LastName)
		suggestions = append(suggestions, suggestion{Value: e.ID, Text: n})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestions)
}

func listEmployeeLeaves(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	now := time.Now()
	past := now.AddDate(-1, 0, 0)
	leaves := []models.Leave{}
	db.Debug().
		Where("employee_id = ?", id).
		Where("leave_date >= ?", past).
		Where("leave_date < ?", now).
		Order("leave_date").
		Find(&leaves)

	results := []Result{}
	for _, l := range leaves {
		results = append(results, Result{Count: 1, Date: l.LeaveDate})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
