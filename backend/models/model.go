package models

import (
	"time"
)

type Employee struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
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
	ID              uint `gorm:"primary_key"`
	CreatedAt       time.Time
	EmployeeID      uint `json:"-"`
	LeaveDate       time.Time
	LeaveNameID     uint
	LeaveTypeID     uint
	LeaveCategoryID uint
	Name            LeaveName
	Type            LeaveType
	Category        LeaveCategory
}

type LeaveRange struct {
	ID         uint `gorm:"primary_key"`
	CreatedAt  time.Time
	Start      time.Time
	Finish     time.Time
	EmployeeID uint
	Name       string
	Category   string
	Type       string
}

type LeaveName struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	Name      string `gorm:"size:20"`
}

type LeaveCategory struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	Name      string `gorm:"size:40"`
}

type LeaveType struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	Name      string `gorm:"size:40"`
}
