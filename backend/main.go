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

func listEmployeeLeavesYears(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	yrs, _ := strconv.Atoi(vars["yr"])

	thisyear := time.Now().Year()
	results := []result{}
	var res result
	var leaves []models.Leave
	emp := models.Employee{}
	db.Find(&emp, id)
	if emp.StartDate.Year() > thisyear-yrs {
		yrs = thisyear - emp.StartDate.Year()
	}

	for yy := 0; yy <= yrs; yy++ {
		st := time.Date(thisyear-yy, 1, 1, 0, 0, 0, 0, time.UTC)
		ft := time.Date(thisyear-yy+1, 1, 1, 0, 0, 0, 0, time.UTC)
		res = result{}
		res.Days = make(map[string]LeaveDay)
		plds := []LeaveDay{}
		uplds := []LeaveDay{}
		q := db.
			Table("leaves").
			Select(`leave_names.name as leave_name,
				leave_category_id as cat_id, count(leaves.id) as count,
				leave_categories.name as cat_name`).
			Joins("inner join leave_categories on leave_category_id = leave_categories.id").
			Joins("inner join leave_names on leave_name_id = leave_names.id").
			Where("employee_id = ?", id).
			Where("leave_date >= ?", st).
			Where("leave_date < ?", ft).
			Where("EXTRACT(dow FROM leave_date) IN (1,2,3,4,5)").
			Group("leave_category_id, leave_names.name,leave_categories.name").
			Order("count DESC, leave_categories.name")

		q.Where("leave_names.id = 1").Scan(&plds)
		q.Where("leave_names.id = 2").Scan(&uplds)
		res.PlCounts = plds
		res.UplCounts = uplds

		dows := []DayCounts{}
		db.
			Table("leaves").
			Select("EXTRACT(dow FROM leave_date) as DOW, to_char(leave_date,'day') as Day, count(leave_date) as Count").
			Where("leave_name_id = 2").
			Where("employee_id = ?", id).
			Where("leave_date >= ?", st).
			Where("leave_date < ?", ft).
			Where("EXTRACT(dow FROM leave_date) IN (1,2,3,4,5)").
			Group("EXTRACT(dow FROM leave_date),to_char(leave_date,'day')").
			Order("Count DESC, DOW").
			Scan(&dows)
		res.Dows = dows
		leaves = []models.Leave{}
		db.
			// Debug().
			Where("employee_id = ?", id).
			Where("leave_date >= ?", st).
			Where("leave_date < ?", ft).
			Order("leave_date").
			Find(&leaves)
		res.Total = len(leaves)
		res.Title = fmt.Sprintf("%s %s absences in %d", emp.FirstName, emp.LastName, st.Year())
		res.FileTitle = fmt.Sprintf("Absences%s_%s_%d", emp.FirstName, emp.LastName, st.Year())
		res.Year = st.Year()
		for _, l := range leaves {
			res.Days[l.LeaveDate.Format("02-01")] = LeaveDay{
				Count:  -1,
				NameID: l.LeaveNameID,
				CatId:  l.LeaveCategoryID,
				TypeId: l.LeaveTypeID}
		}
		results = append(results, res)

	}

	// fmt.Println(leaves)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)

}

func CollectTeamLeavesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	yrs, _ := strconv.Atoi(vars["yr"])

	emp := models.Employee{}
	ok := !db.Find(&emp, id).RecordNotFound()
	if !ok {
		// log.Printf("Count not find user with id: %d \n", id)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	res := team_result{}

	count := 0
	db.
		Model(&models.Employee{}).
		Where("manager_id = ?", emp.ID).
		// Where("ManagerID = ?", emp.ID).
		Count(&count)

	if count == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return
	}

	res.PDays = make(map[string]uint)
	res.UDays = make(map[string]uint)
	res.Year = time.Now().Year()
	res.Title = fmt.Sprintf("%s %s's Direct Reports Absences for %d years as of %d",
		emp.FirstName, emp.LastName, yrs, res.Year)
	res.FileTitle = strings.Replace(res.Title, " ", "_", 0)
	type cc struct {
		Dom    string
		Pcount uint
		Ucount uint
	}
	tcs := []cc{}
	db.Raw("SELECT * FROM team_leaves(?,?,?)", id, yrs, false).Scan(&tcs)
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
	res.PMax = pmax
	res.PMin = pmin
	res.UMax = umax
	res.UMin = umin

	st := time.Date(res.Year-10, 1, 1, 0, 0, 0, 0, time.UTC)
	ft := time.Date(res.Year+1, 1, 1, 0, 0, 0, 0, time.UTC)
	dows := []DayCounts{}
	db.
		Table("leaves").
		Joins("inner join employees on leaves.employee_id = employees.id").
		Select("EXTRACT(dow FROM leave_date) as DOW, to_char(leave_date,'day') as Day, count(leave_date) as Count").
		Where("leave_name_id = 2").
		Where("employees.manager_id = ?", id).
		Where("leave_date >= ?", st).
		Where("leave_date < ?", ft).
		Where("EXTRACT(dow FROM leave_date) IN (1,2,3,4,5)").
		Group("EXTRACT(dow FROM leave_date),to_char(leave_date,'day')").
		Order("Count DESC, DOW").
		Scan(&dows)
	res.Dows = dows

	plds := []LeaveDay{}
	uplds := []LeaveDay{}
	q := db.
		Table("leaves").
		Select(`leave_names.name as leave_name,
				leave_category_id as cat_id, count(leaves.id) as count,
				leave_categories.name as cat_name`).
		Joins("inner join leave_categories on leave_category_id = leave_categories.id").
		Joins("inner join leave_names on leave_name_id = leave_names.id").
		Joins("inner join employees on leaves.employee_id = employees.id").
		Where("employees.manager_id = ?", id).
		Where("leave_date >= ?", st).
		Where("leave_date < ?", ft).
		Where("EXTRACT(dow FROM leave_date) IN (1,2,3,4,5)").
		Group("leave_category_id, leave_names.name,leave_categories.name").
		Order("count DESC, leave_categories.name")

	q.Where("leave_names.id = 1").Scan(&plds)
	q.Where("leave_names.id = 2").Scan(&uplds)
	res.PlCounts = plds
	res.UplCounts = uplds

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func isLeap(year int) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}
