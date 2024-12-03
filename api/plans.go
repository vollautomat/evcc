package api

type RepeatingPlanStruct struct {
	Days []int  `json:"days"` // 0-6 (Sunday-Saturday)
	Time     string `json:"time"`     // HH:MM
	Tz       string `json:"tz"`       // timezone in IANA format
	Soc      int    `json:"soc"`
	Active   bool   `json:"active"`
}
