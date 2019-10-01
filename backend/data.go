package main

type DayCounts struct {
	DOW   uint   `json:"-"`
	Day   string `json:"day"`
	Count uint   `json:"count,omitempty"`
}

type team_result struct {
	Year      int             `json:"year"`
	Title     string          `json:"title"`
	FileTitle string          `json:"file_title"`
	PDays     map[string]uint `json:"pdays"`
	UDays     map[string]uint `json:"udays"`
	PTotal    uint            `json:"ptotal"`
	UTotal    uint            `json:"utotal"`
	Dows      []DayCounts     `json:"dows"`
}
