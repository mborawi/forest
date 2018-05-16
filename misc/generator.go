package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/icrowley/fake"
)

func main() {
	N := 20000
	fmt.Println("creating fake data")
	f, err := os.Create("schema.sql")
	if err != nil {
		fmt.Println(err)
		return
	}
	rand.Seed(time.Now().Unix())
	w := bufio.NewWriter(f)
	TableName := "leaves"
	fmt.Fprintln(w, getTableHead2(TableName))
	TableName = "employees"
	fmt.Fprintln(w, getTableHead(TableName))
	fmt.Fprintln(w, CreateLeavesGroups())
	fmt.Fprintln(w, CreateLeavesCategories())
	fmt.Fprintln(w, CreateLeavesTypes())
	w.Flush()
	f.Close()

	f, err = os.Create(TableName + ".csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	w = bufio.NewWriter(f)
	for i := 1; i <= N; i++ {
		fname := fake.FirstName()
		lname := fake.LastName()
		email := fmt.Sprintf("%s.%s@fakemail.com", fname, lname)
		phone := fake.Phone()
		jobTitle := fake.JobTitle()
		branchId := rand.Intn(10) + 1
		managerID := rand.Intn(i) + 1
		if i == 1 {
			managerID = 0
		}
		if i != i && managerID == 0 {
			managerID = 1
		}
		fmt.Fprintf(w, "%d,%s,%s,%s,%s,%s,%d,%d\n",
			i, fname, lname, email, phone, jobTitle, branchId, managerID)
	}
	w.Flush()
	f.Close()
	TableName = "leaves"
	f, err = os.Create(TableName + ".csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	w = bufio.NewWriter(f)
	c := 1
	for i := 1; i <= N; i++ {
		for y := 2008; y < 2018; y++ {
			l := 20 + rand.Intn(20)
			pm := 0
			pd := 0
			for j := 0; j < l; j++ {
				m := random(pm, 13)
				d := random(pd, 30)
				if m == 2 {
					d = Max(d, 28)
				}
				LDate := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
				LeaveId := rand.Intn(20) + 1
				fmt.Fprintf(w, "%d, %d, %s, %d\n",
					c, i, LDate.Format("2006-01-02"), LeaveId)
				c += 1
				pm = m
				pd = d
				if pd > 27 {
					pd = 0
				}
				if pm > 11 {
					pm = 0
				}
			}
		}
	}
	w.Flush()

	//
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func random(min, max int) int {
	return rand.Intn(max-min) + min + 1
}

func getTableHead2(tn string) string {
	return "DROP TABLE IF EXISTS " + tn + " CASCADE;\n CREATE TABLE " + tn + ` (
	id SERIAL PRIMARY KEY,
	employee_id Integer,
	leave_date date,
	leave_id integer  
);` + "\n CREATE INDEX date_idx ON " + tn + " (leave_date);" +
		"\n CREATE INDEX employee_idx ON " + tn + " (employee_id);"
}

func getTableHead(tn string) string {
	return "DROP TABLE IF EXISTS " + tn + " CASCADE;\n create table " + tn + ` (
	id SERIAL PRIMARY KEY,
	first_name VARCHAR(50),
	last_name VARCHAR(50),
	email VARCHAR(80),
	phone VARCHAR(25),
	job_title VARCHAR(100),
	branch_id integer,
	manager_id integer 
);`
}

func CreateLeavesGroups() string {
	return "DROP TABLE IF EXISTS  lnames CASCADE;\n create table lnames(" +
		`id SERIAL PRIMARY KEY,
		name VARCHAR(50));`
}
func CreateLeavesCategories() string {
	return "DROP TABLE IF EXISTS lcategories CASCADE;\n create table lcategories(" +
		`id SERIAL PRIMARY KEY,
		name VARCHAR(50));`
}

func CreateLeavesTypes() string {
	return "DROP TABLE IF EXISTS ltypes CASCADE;\n create table ltypes(" +
		`id SERIAL PRIMARY KEY,
		name VARCHAR(50));`
}
