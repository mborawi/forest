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
	f, err := os.Create("schema.sql")
	if err != nil {
		fmt.Println(err)
		return
	}
	w := bufio.NewWriter(f)
	TableName := "leaves"
	fmt.Fprintln(w, getTableHead2(TableName))
	TableName = "employees"
	fmt.Fprintln(w, getTableHead(TableName))
	w.Flush()
	f.Close()

	f, err = os.Create(TableName + ".csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	w = bufio.NewWriter(f)
	N := 10000
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
		fmt.Fprintf(w, "%d, %s, %s, %s, %s, %s, %d, %d\n",
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
			for j := 0; j < l; j++ {
				LDate := time.Date(y, time.Month(rand.Intn(12)+1), rand.Intn(28)+1, 0, 0, 0, 0, time.UTC)
				LeaveId := rand.Intn(20) + 1
				fmt.Fprintf(w, "%d, %d, %s, %d\n",
					c, i, LDate.Format("2006-01-02"), LeaveId)
				c += 1
			}
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

func getTableHead(tn string) string {
	return "DROP TABLE IF EXISTS " + tn + " CASCADE;\n create table " + tn + ` (
	id SERIAL PRIMARY KEY,
	first_name VARCHAR(50),
	last_name VARCHAR(50),
	email VARCHAR(80),
	phone VARCHAR(25),
	job_title VARCHAR(100),
	branchId integer,
	managerId integer 
);`
}
