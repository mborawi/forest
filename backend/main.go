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
	con := fmt.Sprintf("user=%s dbname=%s port=%s  sslmode=disable",
		"mido",
		"heat",
		"5432")
	var err error
	db, err = gorm.Open("postgres", con)
	if err != nil {
		log.Fatalln("failed to connect database", err)
	}
	defer db.Close()

	////
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/api/individual/{id:[0-9]+}", listEmployeeLeaves)

	log.Println("Launching Server Now on 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
	///
	// m := models.Employee{}
	// db.First(&m)

	// fmt.Println(m)
}

type Result struct {
	Count uint      `json:"count"`
	Date  time.Time `json:"date"`
}

func listEmployeeLeaves(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	now := time.Now()
	past := now.AddDate(-1, 0, 0)
	leaves := []models.Leave{}
	db.Debug().
		Where("employee_id = ?", id).
		Where("leave_date >= ? AND leave_date < ?", past, now).
		Find(&leaves)

	results := []Result{}
	for _, l := range leaves {
		results = append(results, Result{Count: 1, Date: *l.LeaveDate})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
