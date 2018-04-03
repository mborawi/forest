package main

type day struct {
	Date  string `json:"date"`
	Count uint   `json:"count"`
}

type result struct {
	Year  int             `json:"year"`
	Title string          `json:"title"`
	Days  map[string]uint `json:"days"`
	Max   uint            `json:"max"`
	Min   uint            `json:"min"`
}
