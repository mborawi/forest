package models

import (
	"time"
)

type Employee struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	FirstName string    `json:"-"`
	LastName  string    `json:"-"`
	FullName  string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	JobTitle  string    `json:"title"`
	ManagerID uint      `json:"manager_id"`
	BranchID  uint      `json:"-"`
	StartDate time.Time `json:"start_date"`
	Leaves    []Leave   `json:"-"`
}

type Leave struct {
	ID              uint          `gorm:"primary_key"`
	CreatedAt       time.Time     `json:"-"`
	EmployeeID      uint          `json:"-"`
	LeaveDate       time.Time     `json:"-"`
	LeaveNameID     uint          `json:"-"`
	LeaveTypeID     uint          `json:"-"`
	LeaveCategoryID uint          `json:"-"`
	Name            LeaveName     `json:"-"`
	Type            LeaveType     `json:"-"`
	Category        LeaveCategory `json:"-"`
}

type LeaveRange struct {
	ID         uint      `gorm:"primary_key"`
	CreatedAt  time.Time `json:"-"`
	Start      time.Time `json:"-"`
	Finish     time.Time `json:"-"`
	Hours      float64   `json:"-"`
	EmployeeID uint      `json:"-"`
	Name       string    `json:"-"`
	Category   string    `json:"-"`
	Type       string    `json:"-"`
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
