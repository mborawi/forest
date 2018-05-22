package main

type day struct {
	Date  string `json:"date"`
	Count uint   `json:"count"`
}

type result struct {
	Year      int             `json:"year"`
	Title     string          `json:"title"`
	FileTitle string          `json:"file_title"`
	Days      map[string]uint `json:"days"`
	Max       uint            `json:"max"`
	Min       uint            `json:"min"`
}
