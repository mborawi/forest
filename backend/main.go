package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/mborawi/forest/backend/config"
	"github.com/mborawi/forest/backend/models"
)

var db *gorm.DB
var conf config.Config

func main() {
	config.ReadConfig(&conf)

	con := fmt.Sprintf("user=%s password=%s dbname=%s port=%s  sslmode=disable",
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.DbName,
		conf.Database.Port)

	var err error
	db, err = gorm.Open("postgres", con)
	if err != nil {
		log.Fatalln("failed to connect database", err)
	}
	defer db.Close()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/leaves/list", listleavesHandler)
	router.HandleFunc("/api/leaves/{year:[0-9]+}/", listleavesHandler)
	router.HandleFunc("/api/branch/list", listBranches)
	router.HandleFunc("/api/department/{id:[0-9]+}", listDepartments)
	router.HandleFunc("/api/availability", listAvailability)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../frontEnd/dist/static/"))))
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "../frontEnd/dist/index.html") })
	router.NotFoundHandler = http.HandlerFunc(notFound)
	log.Printf("Launching Server Now on %s...", conf.Server.Port)
	log.Fatal(http.ListenAndServe(conf.Server.Port, router))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	log.Println("Flux Reactor not ready doc!...")
}

func listDepartments(w http.ResponseWriter, r *http.Request) {
	deps := []models.Department{}
	vars := mux.Vars(r)
	br_id, _ := strconv.Atoi(vars["id"])
	db.
		Order("id").
		Where("branch_id = ?", br_id).
		Find(&deps)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deps)
}

func listBranches(w http.ResponseWriter, r *http.Request) {
	brs := []models.Branch{}
	db.Order("id").Find(&brs)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(brs)
}

func listleavesHandler(w http.ResponseWriter, r *http.Request) {
	lvs := []models.LeaveType{}
	db.
		Order("id").
		Find(&lvs)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lvs)
}

func listAvailability(w http.ResponseWriter, r *http.Request) {
	BSL := 0
	CostCentre := 0
	years := 10
	if len(r.FormValue("bsl")) > 0 {
		if tmp, err := strconv.Atoi(r.FormValue("bsl")); err == nil {
			BSL = tmp
		}
	}
	if len(r.FormValue("cost")) > 0 {
		if tmp, err := strconv.Atoi(r.FormValue("cost")); err == nil {
			CostCentre = tmp
		}
	}
	if len(r.FormValue("years")) > 0 {
		if tmp, err := strconv.Atoi(r.FormValue("years")); err == nil {
			years = tmp
		}
	}
	log.Println(years, BSL, CostCentre)
	res := team_result{}
	res.PDays = make(map[string]uint)
	res.UDays = make(map[string]uint)
	res.Year = time.Now().Year()
	res.Title = fmt.Sprintf("Leaves for %d years as of %d",
		years, res.Year)
	res.FileTitle = strings.Replace(res.Title, " ", "_", 0)

	// st := time.Date(res.Year-years, 1, 1, 0, 0, 0, 0, time.UTC)
	// ft := time.Date(res.Year+1, 1, 1, 0, 0, 0, 0, time.UTC)
	// dows := []DayCounts{}
	// db.
	// 	Table("leaves").
	// 	Select("EXTRACT(dow FROM leave_date) as DOW, to_char(leave_date,'day') as Day, count(leave_date) as Count").
	// 	Where("leave_date >= ?", st).
	// 	Where("leave_date < ?", ft).
	// 	Where("EXTRACT(dow FROM leave_date) IN (1,2,3,4,5)").
	// 	Group("EXTRACT(dow FROM leave_date),to_char(leave_date,'day')").
	// 	Order("Count DESC, DOW").
	// 	Scan(&dows)
	// res.Dows = dows

	type cc struct {
		Dom    string
		Pcount uint
		Ucount uint
	}
	tcs := []cc{}
	if CostCentre == 0 && BSL == 0 {
		db.Raw("SELECT * FROM team_leaves_years(?)", years).Scan(&tcs)
	} else if CostCentre == 0 {
		db.Raw("SELECT * FROM team_leaves_bsl(?,?)", years, BSL).Scan(&tcs)
	} else {
		log.Println("==>>", years, CostCentre)
		db.Raw("SELECT * FROM team_leaves_cost(?,?)", years, CostCentre).Scan(&tcs)
	}
	pmax := uint(0)
	umax := uint(0)
	pmin := uint(math.MaxUint32)
	umin := uint(math.MaxUint32)
	for _, t := range tcs {
		if t.Dom == "29-02" && !isLeap(res.Year) {
			continue
		}
		if t.Pcount != 0 {
			res.PDays[t.Dom] = t.Pcount
			res.PTotal += t.Pcount
			if pmax < t.Pcount {
				pmax = t.Pcount
			}
			if pmin > t.Pcount {
				pmin = t.Pcount
			}
		}
		if t.Ucount != 0 {
			res.UDays[t.Dom] = t.Ucount
			res.UTotal += t.Ucount

			if umax < t.Ucount {
				umax = t.Ucount
			}
			if umin > t.Ucount {
				umin = t.Ucount
			}
		}

	}
	res.UMax = umax
	res.UMin = umin
	res.PMax = pmax
	res.PMin = pmin
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func isLeap(year int) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}
