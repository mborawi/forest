package main

type day struct {
	Date  string `json:"date"`
	Count uint   `json:"count"`
}

type result struct {
	Year      int                 `json:"year"`
	Title     string              `json:"title"`
	FileTitle string              `json:"file_title"`
	Days      map[string]LeaveDay `json:"days"`
	Max       uint                `json:"max"`
	Min       uint                `json:"min"`
	PlCounts  []LeaveDay          `json:"pcounts"`
	UplCounts []LeaveDay          `json:"ucounts"`
	Dows      []DayCounts         `json:"dows"`
	Total     int                 `json:"total"`
}

type LeaveDay struct {
	Count     int    `json:"count"`
	NameID    uint   `json:"name_id"`
	CatId     uint   `json:"cat_id"`
	CatName   string `json:"cat,omitempty"`
	LeaveName string `json:"name,omitempty"`
	TypeId    uint   `json:"type_id"`
}

type YrLeaves struct {
	TotalCount uint
}

type DayCounts struct {
	DOW   uint   `json:"-"`
	Day   string `json:"day"`
	Count uint   `json:"count"`
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
}
