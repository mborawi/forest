package main

import (
	"time"
)

type Result struct {
	Count uint      `json:"count"`
	Date  time.Time `json:"date"`
}
