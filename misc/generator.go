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
	fmt.Println("creating fake data")
	f, err := os.Create("FakeData.sql")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	TableName := "employee"

	fmt.Fprintln(w, getTableHead(TableName))
	N := 1000
	for i := 1; i <= N; i++ {
		fname := fake.FirstName()
		lname := fake.LastName()
		email := fake.EmailAddress()
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
		fmt.Fprintf(w, "insert into %s (id, first_name,last_name,email,phone,job_title, branchId,  managerId) values", TableName)
		fmt.Fprintf(w, "('%d','%s', '%s','%s', '%s', '%s','%d', '%d');\n",
			i, fname, lname, email, phone, jobTitle, branchId, managerID)
	}
	w.Flush()

	TableName = "Leaves"
	fmt.Fprintln(w, getTableHead2(TableName))
	for i := 1; i <= N; i++ {
		l := 20 + rand.Intn(20)
		for j := 0; j < l; j++ {
			LDate := time.Date(rand.Intn(10)+2008, time.Month(rand.Intn(12)+1), rand.Intn(28)+1, 0, 0, 0, 0, time.UTC)
			LeaveId := rand.Intn(20) + 1
			fmt.Fprintf(w, "insert into %s (employeeId,LeaveDate,LeaveId) values", TableName)
			fmt.Fprintf(w, "('%d','%s', '%d');\n",
				i, LDate.Format("2006-01-02"), LeaveId)
		}

	}
	w.Flush()

	//
}

func getTableHead2(tn string) string {
	return "DROP TABLE IF EXISTS " + tn + " CASCADE;\n create table " + tn + ` (
	id SERIAL PRIMARY KEY,
	employeeId Integer,
	LeaveDate date,
	LeaveId integer  
);`
}

// id, first_name,last_name,email,phone,job_title, branchId,  managerId
func getTableHead(tn string) string {
	return "DROP TABLE IF EXISTS " + tn + " CASCADE;\n create table " + tn + ` (
	id SERIAL PRIMARY KEY,
	first_name VARCHAR(50),
	last_name VARCHAR(50),
	email VARCHAR(80),
	phone VARCHAR(15),
	job_title VARCHAR(100),
	branchId integer,
	managerId integer 
);`
}
