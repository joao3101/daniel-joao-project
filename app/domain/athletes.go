package domain

type Athlete struct {
	UUID     string `gorm:"primary_key" json:"uuid"`
	Age      int64  `json:"age"`
	Name     string `json:"name"`
	TeamUUID string `json:"team_uuid"`
}
