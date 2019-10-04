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
	PMax      uint            `json:"pmax"`
	PMin      uint            `json:"pmin"`
	UDays     map[string]uint `json:"udays"`
	UMax      uint            `json:"umax"`
	UMin      uint            `json:"umin"`
	PTotal    uint            `json:"ptotal"`
	UTotal    uint            `json:"utotal"`
	Dows      []DayCounts     `json:"dows"`
}
