package main

import "time"

type Result struct {
	Date  time.Time `json:"date"`
	Count uint      `json:"count"`
}
