package models

import (
	"time"
)

type Employee struct {
	ID        uint `gorm:"primary_key"`
	FirstName string
	LastName  string
	Email     string
	Phone     string
	JobTitle  string
	ManagerID uint
	BranchID  uint
	Leaves    []Leave
}

type Leave struct {
	ID         uint `gorm:"primary_key" json:"-"`
	EmployeeID uint `json:"-"`
	LeaveID    uint
	LeaveDate  time.Time
}
